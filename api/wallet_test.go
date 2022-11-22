package api

import (
	"bytes"
	mockdb "confxsd/arf-case/db/mock"
	db "confxsd/arf-case/db/sqlc"
	"confxsd/arf-case/token"
	util "confxsd/arf-case/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func randomWallet(UserID int64) db.Wallet {
	return db.Wallet{
		ID:       util.RandomInt(1, 1000),
		UserID:   UserID,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func TestCreateWalletAPI(t *testing.T) {
	user := randomUser(t)
	wallet := randomWallet(user.ID)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{{
		name: "OK",
		body: gin.H{
			"currency": wallet.Currency,
		},
		setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
		},
		buildStubs: func(store *mockdb.MockStore) {
			arg := db.CreateWalletParams{
				UserID:   wallet.UserID,
				Currency: wallet.Currency,
				Balance:  0,
			}

			store.EXPECT().
				CreateWallet(gomock.Any(), gomock.Eq(arg)).
				Times(1).
				Return(wallet, nil)
		},
		checkResponse: func(recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusOK, recorder.Code)
			requireBodyMatchWallet(t, recorder.Body, wallet)
		},
	}}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/wallets"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}

}

func requireBodyMatchWallet(t *testing.T, body *bytes.Buffer, wallet db.Wallet) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotWallet db.Wallet
	err = json.Unmarshal(data, &gotWallet)
	require.NoError(t, err)
	require.Equal(t, wallet, gotWallet)
}

func requireBodyMatchWallets(t *testing.T, body *bytes.Buffer, wallets []db.Wallet) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotWallets []db.Wallet
	err = json.Unmarshal(data, &gotWallets)
	require.NoError(t, err)
	require.Equal(t, wallets, gotWallets)
}

func TestListWalletsAPI(t *testing.T) {
	user := randomUser(t)

	n := 5
	wallets := make([]db.Wallet, n)
	for i := 0; i < n; i++ {
		wallets[i] = randomWallet(user.ID)
	}

	type Query struct {
		pageID   int
		pageSize int
	}

	testCases := []struct {
		name          string
		query         Query
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListWalletsParams{
					UserID: user.ID,
					Limit:  int32(n),
					Offset: 0,
				}

				store.EXPECT().
					ListWallets(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(wallets, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchWallets(t, recorder.Body, wallets)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/wallets"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Add query parameters to request URL
			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}

}
