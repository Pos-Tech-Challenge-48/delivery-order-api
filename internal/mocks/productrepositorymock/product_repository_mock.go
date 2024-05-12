// Code generated by MockGen. DO NOT EDIT.
// Source: ./product.go
//
// Generated by this command:
//
//	mockgen -destination=./../../mocks/productrepositorymock/product_repository_mock.go -source=./product.go -package=productrepositorymock
//
// Package productrepositorymock is a generated GoMock package.
package productrepositorymock

import (
	context "context"
	reflect "reflect"

	entities "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockProductRepository) Add(ctx context.Context, product *entities.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockProductRepositoryMockRecorder) Add(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockProductRepository)(nil).Add), ctx, product)
}

// Delete mocks base method.
func (m *MockProductRepository) Delete(ctx context.Context, productID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductRepositoryMockRecorder) Delete(ctx, productID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductRepository)(nil).Delete), ctx, productID)
}

// DeleteImage mocks base method.
func (m *MockProductRepository) DeleteImage(ctx context.Context, productID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteImage", ctx, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteImage indicates an expected call of DeleteImage.
func (mr *MockProductRepositoryMockRecorder) DeleteImage(ctx, productID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImage", reflect.TypeOf((*MockProductRepository)(nil).DeleteImage), ctx, productID)
}

// GetAll mocks base method.
func (m *MockProductRepository) GetAll(ctx context.Context, params string) ([]entities.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, params)
	ret0, _ := ret[0].([]entities.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockProductRepositoryMockRecorder) GetAll(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockProductRepository)(nil).GetAll), ctx, params)
}

// GetByID mocks base method.
func (m *MockProductRepository) GetByID(ctx context.Context, ID string) (*entities.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, ID)
	ret0, _ := ret[0].(*entities.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockProductRepositoryMockRecorder) GetByID(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockProductRepository)(nil).GetByID), ctx, ID)
}

// GetCategoryID mocks base method.
func (m *MockProductRepository) GetCategoryID(ctx context.Context, categoryName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryID", ctx, categoryName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryID indicates an expected call of GetCategoryID.
func (mr *MockProductRepositoryMockRecorder) GetCategoryID(ctx, categoryName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryID", reflect.TypeOf((*MockProductRepository)(nil).GetCategoryID), ctx, categoryName)
}

// Update mocks base method.
func (m *MockProductRepository) Update(ctx context.Context, product *entities.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockProductRepositoryMockRecorder) Update(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductRepository)(nil).Update), ctx, product)
}

// UpdateImage mocks base method.
func (m *MockProductRepository) UpdateImage(ctx context.Context, productID, image string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateImage", ctx, productID, image)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateImage indicates an expected call of UpdateImage.
func (mr *MockProductRepositoryMockRecorder) UpdateImage(ctx, productID, image any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateImage", reflect.TypeOf((*MockProductRepository)(nil).UpdateImage), ctx, productID, image)
}
