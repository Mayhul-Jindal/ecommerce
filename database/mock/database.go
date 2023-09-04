// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc (interfaces: Storer)

// Package mockDatabase is a generated GoMock package.
package mockDatabase

import (
	context "context"
	reflect "reflect"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStorer is a mock of Storer interface.
type MockStorer struct {
	ctrl     *gomock.Controller
	recorder *MockStorerMockRecorder
}

// MockStorerMockRecorder is the mock recorder for MockStorer.
type MockStorerMockRecorder struct {
	mock *MockStorer
}

// NewMockStorer creates a new mock instance.
func NewMockStorer(ctrl *gomock.Controller) *MockStorer {
	mock := &MockStorer{ctrl: ctrl}
	mock.recorder = &MockStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorer) EXPECT() *MockStorerMockRecorder {
	return m.recorder
}

// AddOrder mocks base method.
func (m *MockStorer) AddOrder(arg0 context.Context, arg1 database.AddOrderParams) (database.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrder", arg0, arg1)
	ret0, _ := ret[0].(database.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrder indicates an expected call of AddOrder.
func (mr *MockStorerMockRecorder) AddOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrder", reflect.TypeOf((*MockStorer)(nil).AddOrder), arg0, arg1)
}

// AddToCart mocks base method.
func (m *MockStorer) AddToCart(arg0 context.Context, arg1 database.AddToCartParams) (database.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToCart", arg0, arg1)
	ret0, _ := ret[0].(database.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddToCart indicates an expected call of AddToCart.
func (mr *MockStorerMockRecorder) AddToCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToCart", reflect.TypeOf((*MockStorer)(nil).AddToCart), arg0, arg1)
}

// CheckAdmin mocks base method.
func (m *MockStorer) CheckAdmin(arg0 context.Context, arg1 int64) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAdmin", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAdmin indicates an expected call of CheckAdmin.
func (mr *MockStorerMockRecorder) CheckAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAdmin", reflect.TypeOf((*MockStorer)(nil).CheckAdmin), arg0, arg1)
}

// CheckBookPurchased mocks base method.
func (m *MockStorer) CheckBookPurchased(arg0 context.Context, arg1 database.CheckBookPurchasedParams) (database.Purchase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBookPurchased", arg0, arg1)
	ret0, _ := ret[0].(database.Purchase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckBookPurchased indicates an expected call of CheckBookPurchased.
func (mr *MockStorerMockRecorder) CheckBookPurchased(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBookPurchased", reflect.TypeOf((*MockStorer)(nil).CheckBookPurchased), arg0, arg1)
}

// CheckEmailVerified mocks base method.
func (m *MockStorer) CheckEmailVerified(arg0 context.Context, arg1 int64) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmailVerified", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEmailVerified indicates an expected call of CheckEmailVerified.
func (mr *MockStorerMockRecorder) CheckEmailVerified(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmailVerified", reflect.TypeOf((*MockStorer)(nil).CheckEmailVerified), arg0, arg1)
}

// CreateBook mocks base method.
func (m *MockStorer) CreateBook(arg0 context.Context, arg1 database.CreateBookParams) (database.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", arg0, arg1)
	ret0, _ := ret[0].(database.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockStorerMockRecorder) CreateBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockStorer)(nil).CreateBook), arg0, arg1)
}

// CreatePurchase mocks base method.
func (m *MockStorer) CreatePurchase(arg0 context.Context, arg1 database.CreatePurchaseParams) (database.Purchase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePurchase", arg0, arg1)
	ret0, _ := ret[0].(database.Purchase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePurchase indicates an expected call of CreatePurchase.
func (mr *MockStorerMockRecorder) CreatePurchase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePurchase", reflect.TypeOf((*MockStorer)(nil).CreatePurchase), arg0, arg1)
}

// CreateReview mocks base method.
func (m *MockStorer) CreateReview(arg0 context.Context, arg1 database.CreateReviewParams) (database.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReview", arg0, arg1)
	ret0, _ := ret[0].(database.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReview indicates an expected call of CreateReview.
func (mr *MockStorerMockRecorder) CreateReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReview", reflect.TypeOf((*MockStorer)(nil).CreateReview), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStorer) CreateSession(arg0 context.Context, arg1 database.CreateSessionParams) (database.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(database.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStorerMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStorer)(nil).CreateSession), arg0, arg1)
}

// CreateTag mocks base method.
func (m *MockStorer) CreateTag(arg0 context.Context, arg1 database.CreateTagParams) (database.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTag", arg0, arg1)
	ret0, _ := ret[0].(database.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTag indicates an expected call of CreateTag.
func (mr *MockStorerMockRecorder) CreateTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockStorer)(nil).CreateTag), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStorer) CreateUser(arg0 context.Context, arg1 database.CreateUserParams) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStorerMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStorer)(nil).CreateUser), arg0, arg1)
}

// CreateUserTx mocks base method.
func (m *MockStorer) CreateUserTx(arg0 context.Context, arg1 database.CreateUserTxParams) (database.CreateUserTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserTx", arg0, arg1)
	ret0, _ := ret[0].(database.CreateUserTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserTx indicates an expected call of CreateUserTx.
func (mr *MockStorerMockRecorder) CreateUserTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserTx", reflect.TypeOf((*MockStorer)(nil).CreateUserTx), arg0, arg1)
}

// CreateVerifyEmail mocks base method.
func (m *MockStorer) CreateVerifyEmail(arg0 context.Context, arg1 database.CreateVerifyEmailParams) (database.VerifyEmail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVerifyEmail", arg0, arg1)
	ret0, _ := ret[0].(database.VerifyEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVerifyEmail indicates an expected call of CreateVerifyEmail.
func (mr *MockStorerMockRecorder) CreateVerifyEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVerifyEmail", reflect.TypeOf((*MockStorer)(nil).CreateVerifyEmail), arg0, arg1)
}

// DeleteBook mocks base method.
func (m *MockStorer) DeleteBook(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockStorerMockRecorder) DeleteBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockStorer)(nil).DeleteBook), arg0, arg1)
}

// DeleteCartItem mocks base method.
func (m *MockStorer) DeleteCartItem(arg0 context.Context, arg1 database.DeleteCartItemParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCartItem", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCartItem indicates an expected call of DeleteCartItem.
func (mr *MockStorerMockRecorder) DeleteCartItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCartItem", reflect.TypeOf((*MockStorer)(nil).DeleteCartItem), arg0, arg1)
}

// DeleteOrder mocks base method.
func (m *MockStorer) DeleteOrder(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockStorerMockRecorder) DeleteOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockStorer)(nil).DeleteOrder), arg0, arg1)
}

// DeletePurchase mocks base method.
func (m *MockStorer) DeletePurchase(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePurchase", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePurchase indicates an expected call of DeletePurchase.
func (mr *MockStorerMockRecorder) DeletePurchase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePurchase", reflect.TypeOf((*MockStorer)(nil).DeletePurchase), arg0, arg1)
}

// DeleteReview mocks base method.
func (m *MockStorer) DeleteReview(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReview", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReview indicates an expected call of DeleteReview.
func (mr *MockStorerMockRecorder) DeleteReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReview", reflect.TypeOf((*MockStorer)(nil).DeleteReview), arg0, arg1)
}

// DeleteSession mocks base method.
func (m *MockStorer) DeleteSession(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStorerMockRecorder) DeleteSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStorer)(nil).DeleteSession), arg0, arg1)
}

// DeleteTag mocks base method.
func (m *MockStorer) DeleteTag(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTag indicates an expected call of DeleteTag.
func (mr *MockStorerMockRecorder) DeleteTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTag", reflect.TypeOf((*MockStorer)(nil).DeleteTag), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStorer) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStorerMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStorer)(nil).DeleteUser), arg0, arg1)
}

// DeleteVerifyEmail mocks base method.
func (m *MockStorer) DeleteVerifyEmail(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVerifyEmail", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVerifyEmail indicates an expected call of DeleteVerifyEmail.
func (mr *MockStorerMockRecorder) DeleteVerifyEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVerifyEmail", reflect.TypeOf((*MockStorer)(nil).DeleteVerifyEmail), arg0, arg1)
}

// GetAllTags mocks base method.
func (m *MockStorer) GetAllTags(arg0 context.Context) ([]database.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTags", arg0)
	ret0, _ := ret[0].([]database.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTags indicates an expected call of GetAllTags.
func (mr *MockStorerMockRecorder) GetAllTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTags", reflect.TypeOf((*MockStorer)(nil).GetAllTags), arg0)
}

// GetBookById mocks base method.
func (m *MockStorer) GetBookById(arg0 context.Context, arg1 int64) (database.GetBookByIdRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookById", arg0, arg1)
	ret0, _ := ret[0].(database.GetBookByIdRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookById indicates an expected call of GetBookById.
func (mr *MockStorerMockRecorder) GetBookById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookById", reflect.TypeOf((*MockStorer)(nil).GetBookById), arg0, arg1)
}

// GetCartItemsByUserId mocks base method.
func (m *MockStorer) GetCartItemsByUserId(arg0 context.Context, arg1 int64) ([]database.GetCartItemsByUserIdRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartItemsByUserId", arg0, arg1)
	ret0, _ := ret[0].([]database.GetCartItemsByUserIdRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartItemsByUserId indicates an expected call of GetCartItemsByUserId.
func (mr *MockStorerMockRecorder) GetCartItemsByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartItemsByUserId", reflect.TypeOf((*MockStorer)(nil).GetCartItemsByUserId), arg0, arg1)
}

// GetHotSellingBooks mocks base method.
func (m *MockStorer) GetHotSellingBooks(arg0 context.Context, arg1 database.GetHotSellingBooksParams) ([]database.GetHotSellingBooksRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHotSellingBooks", arg0, arg1)
	ret0, _ := ret[0].([]database.GetHotSellingBooksRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHotSellingBooks indicates an expected call of GetHotSellingBooks.
func (mr *MockStorerMockRecorder) GetHotSellingBooks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHotSellingBooks", reflect.TypeOf((*MockStorer)(nil).GetHotSellingBooks), arg0, arg1)
}

// GetOrderById mocks base method.
func (m *MockStorer) GetOrderById(arg0 context.Context, arg1 database.GetOrderByIdParams) (database.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderById", arg0, arg1)
	ret0, _ := ret[0].(database.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderById indicates an expected call of GetOrderById.
func (mr *MockStorerMockRecorder) GetOrderById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderById", reflect.TypeOf((*MockStorer)(nil).GetOrderById), arg0, arg1)
}

// GetPurchasedBooks mocks base method.
func (m *MockStorer) GetPurchasedBooks(arg0 context.Context, arg1 int64) ([]database.GetPurchasedBooksRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPurchasedBooks", arg0, arg1)
	ret0, _ := ret[0].([]database.GetPurchasedBooksRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPurchasedBooks indicates an expected call of GetPurchasedBooks.
func (mr *MockStorerMockRecorder) GetPurchasedBooks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPurchasedBooks", reflect.TypeOf((*MockStorer)(nil).GetPurchasedBooks), arg0, arg1)
}

// GetReviewsByBookId mocks base method.
func (m *MockStorer) GetReviewsByBookId(arg0 context.Context, arg1 database.GetReviewsByBookIdParams) ([]database.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviewsByBookId", arg0, arg1)
	ret0, _ := ret[0].([]database.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReviewsByBookId indicates an expected call of GetReviewsByBookId.
func (mr *MockStorerMockRecorder) GetReviewsByBookId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviewsByBookId", reflect.TypeOf((*MockStorer)(nil).GetReviewsByBookId), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStorer) GetSession(arg0 context.Context, arg1 uuid.UUID) (database.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(database.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStorerMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStorer)(nil).GetSession), arg0, arg1)
}

// GetTotalCartAmountById mocks base method.
func (m *MockStorer) GetTotalCartAmountById(arg0 context.Context, arg1 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalCartAmountById", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalCartAmountById indicates an expected call of GetTotalCartAmountById.
func (mr *MockStorerMockRecorder) GetTotalCartAmountById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalCartAmountById", reflect.TypeOf((*MockStorer)(nil).GetTotalCartAmountById), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStorer) GetUser(arg0 context.Context, arg1 database.GetUserParams) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStorerMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStorer)(nil).GetUser), arg0, arg1)
}

// GetUserById mocks base method.
func (m *MockStorer) GetUserById(arg0 context.Context, arg1 int64) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockStorerMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockStorer)(nil).GetUserById), arg0, arg1)
}

// GetUserByUsername mocks base method.
func (m *MockStorer) GetUserByUsername(arg0 context.Context, arg1 string) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStorerMockRecorder) GetUserByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStorer)(nil).GetUserByUsername), arg0, arg1)
}

// GetUserRecommendations mocks base method.
func (m *MockStorer) GetUserRecommendations(arg0 context.Context, arg1 database.GetUserRecommendationsParams) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRecommendations", arg0, arg1)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRecommendations indicates an expected call of GetUserRecommendations.
func (mr *MockStorerMockRecorder) GetUserRecommendations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRecommendations", reflect.TypeOf((*MockStorer)(nil).GetUserRecommendations), arg0, arg1)
}

// SearchBooksV2 mocks base method.
func (m *MockStorer) SearchBooksV2(arg0 context.Context, arg1 database.SearchBooksV2Params) ([]database.SearchBooksV2Row, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchBooksV2", arg0, arg1)
	ret0, _ := ret[0].([]database.SearchBooksV2Row)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchBooksV2 indicates an expected call of SearchBooksV2.
func (mr *MockStorerMockRecorder) SearchBooksV2(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchBooksV2", reflect.TypeOf((*MockStorer)(nil).SearchBooksV2), arg0, arg1)
}

// UpdateBook mocks base method.
func (m *MockStorer) UpdateBook(arg0 context.Context, arg1 database.UpdateBookParams) (database.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", arg0, arg1)
	ret0, _ := ret[0].(database.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockStorerMockRecorder) UpdateBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockStorer)(nil).UpdateBook), arg0, arg1)
}

// UpdateOrder mocks base method.
func (m *MockStorer) UpdateOrder(arg0 context.Context, arg1 database.UpdateOrderParams) (database.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", arg0, arg1)
	ret0, _ := ret[0].(database.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockStorerMockRecorder) UpdateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockStorer)(nil).UpdateOrder), arg0, arg1)
}

// UpdateOrderTx mocks base method.
func (m *MockStorer) UpdateOrderTx(arg0 context.Context, arg1 database.UpdateOrderTxParams) (database.UpdateOrderTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderTx", arg0, arg1)
	ret0, _ := ret[0].(database.UpdateOrderTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrderTx indicates an expected call of UpdateOrderTx.
func (mr *MockStorerMockRecorder) UpdateOrderTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderTx", reflect.TypeOf((*MockStorer)(nil).UpdateOrderTx), arg0, arg1)
}

// UpdateReview mocks base method.
func (m *MockStorer) UpdateReview(arg0 context.Context, arg1 database.UpdateReviewParams) (database.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReview", arg0, arg1)
	ret0, _ := ret[0].(database.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateReview indicates an expected call of UpdateReview.
func (mr *MockStorerMockRecorder) UpdateReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReview", reflect.TypeOf((*MockStorer)(nil).UpdateReview), arg0, arg1)
}

// UpdateTag mocks base method.
func (m *MockStorer) UpdateTag(arg0 context.Context, arg1 database.UpdateTagParams) (database.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTag", arg0, arg1)
	ret0, _ := ret[0].(database.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTag indicates an expected call of UpdateTag.
func (mr *MockStorerMockRecorder) UpdateTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTag", reflect.TypeOf((*MockStorer)(nil).UpdateTag), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStorer) UpdateUser(arg0 context.Context, arg1 database.UpdateUserParams) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStorerMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStorer)(nil).UpdateUser), arg0, arg1)
}

// UpdateVerifyEmail mocks base method.
func (m *MockStorer) UpdateVerifyEmail(arg0 context.Context, arg1 database.UpdateVerifyEmailParams) (database.VerifyEmail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVerifyEmail", arg0, arg1)
	ret0, _ := ret[0].(database.VerifyEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVerifyEmail indicates an expected call of UpdateVerifyEmail.
func (mr *MockStorerMockRecorder) UpdateVerifyEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVerifyEmail", reflect.TypeOf((*MockStorer)(nil).UpdateVerifyEmail), arg0, arg1)
}

// VerifyEmailTx mocks base method.
func (m *MockStorer) VerifyEmailTx(arg0 context.Context, arg1 database.VerifyEmailTxParams) (database.VerifyEmailTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyEmailTx", arg0, arg1)
	ret0, _ := ret[0].(database.VerifyEmailTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyEmailTx indicates an expected call of VerifyEmailTx.
func (mr *MockStorerMockRecorder) VerifyEmailTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyEmailTx", reflect.TypeOf((*MockStorer)(nil).VerifyEmailTx), arg0, arg1)
}