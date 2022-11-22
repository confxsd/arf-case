package api

import (
	"bytes"
	mockdb "confxsd/arf-case/db/mock"
	db "confxsd/arf-case/db/sqlc"
	"confxsd/arf-case/token"
	"confxsd/arf-case/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateOfferAPI(t *testing.T) {
	user := randomUser(t)
	amount := float64(10)
	rate := float64(18.5)

	wallet := randomWallet(user.ID)

	wallet.Currency = util.USD
	toCurrency := util.TRY

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"from_currency": wallet.Currency,
				"to_currency":   toCurrency,
				"amount":        amount,
				"rate":          rate,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).Times(1).Return(wallet, nil)

				arg := db.CreateOfferParams{
					UserID:       user.ID,
					FromCurrency: wallet.Currency,
					ToCurrency:   toCurrency,
					Rate:         rate,
					Amount:       amount,
				}
				store.EXPECT().CreateOffer(gomock.Any(), gomock.Eq(arg)).Times(1)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchOffer(t, recorder.Body)
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

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/offers"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}

}

func requireBodyMatchOffer(t *testing.T, body *bytes.Buffer) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotOffer db.Offer
	err = json.Unmarshal(data, &gotOffer)

	require.NoError(t, err)
	require.Equal(t, gotOffer.Status, "active")
}
