package mockdb

import (
	"context"
	"reflect"

	db "serhatbxld/arf-case/db/sqlc"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// GetUser mocks base method
func (m *MockStore) GetUser(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserByUsername mocks base method
func (m *MockStore) GetUserByUsername(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername
func (mr *MockStoreMockRecorder) GetUserByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// CreateSession mocks base method
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// GetSession mocks base method
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// CreateWallet mocks base method
func (m *MockStore) CreateWallet(arg0 context.Context, arg1 db.CreateWalletParams) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWallet indicates an expected call of CreateWallet
func (mr *MockStoreMockRecorder) CreateWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockStore)(nil).CreateWallet), arg0, arg1)
}

// AddWalletBalance mocks base method
func (m *MockStore) AddWalletBalance(arg0 context.Context, arg1 db.AddWalletBalanceParams) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWalletBalance", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddWalletBalance indicates an expected call of AddWalletBalance
func (mr *MockStoreMockRecorder) AddWalletBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWalletBalance", reflect.TypeOf((*MockStore)(nil).AddWalletBalance), arg0, arg1)
}

// DeleteWallet mocks base method
func (m *MockStore) DeleteWallet(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWallet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWallet indicates an expected call of DeleteWallet
func (mr *MockStoreMockRecorder) DeleteWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWallet", reflect.TypeOf((*MockStore)(nil).DeleteWallet), arg0, arg1)
}

// GetWallet mocks base method
func (m *MockStore) GetWallet(arg0 context.Context, arg1 int64) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWallet", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWallet indicates an expected call of GetWallet
func (mr *MockStoreMockRecorder) GetWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWallet", reflect.TypeOf((*MockStore)(nil).GetWallet), arg0, arg1)
}

// GetWalletByUserIdAndCurrency mocks base method
func (m *MockStore) GetWalletByUserIdAndCurrency(arg0 context.Context, arg1 db.GetWalletByUserIdAndCurrencyParams) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletByUserIdAndCurrency", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletByUserIdAndCurrency indicates an expected call of GetWalletByUserIdAndCurrency
func (mr *MockStoreMockRecorder) GetWalletByUserIdAndCurrency(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletByUserIdAndCurrency", reflect.TypeOf((*MockStore)(nil).GetWalletByIdAndCurrency), arg0, arg1)
}

// GetWalletForUpdate mocks base method
func (m *MockStore) GetWalletForUpdate(arg0 context.Context, arg1 int64) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletForUpdate indicates an expected call of GetWalletForUpdate
func (mr *MockStoreMockRecorder) GetWalletForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletForUpdate", reflect.TypeOf((*MockStore)(nil).GetWalletForUpdate), arg0, arg1)
}

// ListWallets mocks base method
func (m *MockStore) ListWallets(arg0 context.Context, arg1 db.ListWalletsParams) ([]db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWallets", arg0, arg1)
	ret0, _ := ret[0].([]db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWallets indicates an expected call of ListWallets
func (mr *MockStoreMockRecorder) ListWallets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWallets", reflect.TypeOf((*MockStore)(nil).ListWallets), arg0, arg1)
}

// UpdateWallet mocks base method
func (m *MockStore) UpdateWallet(arg0 context.Context, arg1 db.UpdateWalletParams) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWallet", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWallet indicates an expected call of UpdateWallet
func (mr *MockStoreMockRecorder) UpdateWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWallet", reflect.TypeOf((*MockStore)(nil).UpdateWallet), arg0, arg1)
}

// CreateOffer mocks base method
func (m *MockStore) CreateOffer(arg0 context.Context, arg1 db.CreateOfferParams) (db.Offer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOffer", arg0, arg1)
	ret0, _ := ret[0].(db.Offer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOffer indicates an expected call of CreateOffer
func (mr *MockStoreMockRecorder) CreateOffer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOffer", reflect.TypeOf((*MockStore)(nil).CreateOffer), arg0, arg1)
}

// GetOffer mocks base method
func (m *MockStore) GetOffer(arg0 context.Context, arg1 int64) (db.Offer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOffer", arg0, arg1)
	ret0, _ := ret[0].(db.Offer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOffer indicates an expected call of GetOffer
func (mr *MockStoreMockRecorder) GetOffer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffer", reflect.TypeOf((*MockStore)(nil).GetOffer), arg0, arg1)
}
