package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	mockdb "serhatbxld/arf-case/db/mock"
	db "serhatbxld/arf-case/db/sqlc"
	"serhatbxld/arf-case/token"
	util "serhatbxld/arf-case/util"
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

func requireBodyMatchWallet(t *testing.T, body *bytes.Buffer, account db.Wallet) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotWallet db.Wallet
	err = json.Unmarshal(data, &gotWallet)
	require.NoError(t, err)
	require.Equal(t, account, gotWallet)
}
