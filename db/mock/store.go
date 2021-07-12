// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lightsOfTruth/dog-walker/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	db "github.com/lightsOfTruth/dog-walker/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddConversation mocks base method.
func (m *MockStore) AddConversation(arg0 context.Context, arg1 db.AddConversationParams) (db.Conversation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddConversation", arg0, arg1)
	ret0, _ := ret[0].(db.Conversation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddConversation indicates an expected call of AddConversation.
func (mr *MockStoreMockRecorder) AddConversation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddConversation", reflect.TypeOf((*MockStore)(nil).AddConversation), arg0, arg1)
}

// AddMessage mocks base method.
func (m *MockStore) AddMessage(arg0 context.Context, arg1 db.AddMessageParams) (db.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMessage", arg0, arg1)
	ret0, _ := ret[0].(db.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMessage indicates an expected call of AddMessage.
func (mr *MockStoreMockRecorder) AddMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMessage", reflect.TypeOf((*MockStore)(nil).AddMessage), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// GetMessageCommunication mocks base method.
func (m *MockStore) GetMessageCommunication(arg0 context.Context, arg1 uuid.UUID) ([]db.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageCommunication", arg0, arg1)
	ret0, _ := ret[0].([]db.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageCommunication indicates an expected call of GetMessageCommunication.
func (mr *MockStoreMockRecorder) GetMessageCommunication(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageCommunication", reflect.TypeOf((*MockStore)(nil).GetMessageCommunication), arg0, arg1)
}

// GetUserByEmail mocks base method.
func (m *MockStore) GetUserByEmail(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockStoreMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockStore)(nil).GetUserByEmail), arg0, arg1)
}

// GetUserById mocks base method.
func (m *MockStore) GetUserById(arg0 context.Context, arg1 uuid.UUID) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockStoreMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockStore)(nil).GetUserById), arg0, arg1)
}

// GetWalkers mocks base method.
func (m *MockStore) GetWalkers(arg0 context.Context) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalkers", arg0)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalkers indicates an expected call of GetWalkers.
func (mr *MockStoreMockRecorder) GetWalkers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalkers", reflect.TypeOf((*MockStore)(nil).GetWalkers), arg0)
}

// MessageTx mocks base method.
func (m *MockStore) MessageTx(arg0 context.Context, arg1 db.NewConversationMessageTxParams) (db.NewConversationMessageTxResultParams, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageTx", arg0, arg1)
	ret0, _ := ret[0].(db.NewConversationMessageTxResultParams)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MessageTx indicates an expected call of MessageTx.
func (mr *MockStoreMockRecorder) MessageTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageTx", reflect.TypeOf((*MockStore)(nil).MessageTx), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
