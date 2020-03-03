// The MIT License (MIT)
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package domain is a generated GoMock package.
package domain

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	workflowservice "go.temporal.io/temporal-proto/workflowservice"
	reflect "reflect"
)

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// DeprecateDomain mocks base method
func (m *MockHandler) DeprecateDomain(ctx context.Context, deprecateRequest *workflowservice.DeprecateDomainRequest) (*workflowservice.DeprecateDomainResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeprecateDomain", ctx, deprecateRequest)
	ret0, _ := ret[0].(*workflowservice.DeprecateDomainResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeprecateDomain indicates an expected call of DeprecateDomain
func (mr *MockHandlerMockRecorder) DeprecateDomain(ctx, deprecateRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeprecateDomain", reflect.TypeOf((*MockHandler)(nil).DeprecateDomain), ctx, deprecateRequest)
}

// DescribeDomain mocks base method
func (m *MockHandler) DescribeDomain(ctx context.Context, describeRequest *workflowservice.DescribeDomainRequest) (*workflowservice.DescribeDomainResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDomain", ctx, describeRequest)
	ret0, _ := ret[0].(*workflowservice.DescribeDomainResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDomain indicates an expected call of DescribeDomain
func (mr *MockHandlerMockRecorder) DescribeDomain(ctx, describeRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDomain", reflect.TypeOf((*MockHandler)(nil).DescribeDomain), ctx, describeRequest)
}

// ListDomains mocks base method
func (m *MockHandler) ListDomains(ctx context.Context, listRequest *workflowservice.ListDomainsRequest) (*workflowservice.ListDomainsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDomains", ctx, listRequest)
	ret0, _ := ret[0].(*workflowservice.ListDomainsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDomains indicates an expected call of ListDomains
func (mr *MockHandlerMockRecorder) ListDomains(ctx, listRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDomains", reflect.TypeOf((*MockHandler)(nil).ListDomains), ctx, listRequest)
}

// RegisterDomain mocks base method
func (m *MockHandler) RegisterDomain(ctx context.Context, registerRequest *workflowservice.RegisterDomainRequest) (*workflowservice.RegisterDomainResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterDomain", ctx, registerRequest)
	ret0, _ := ret[0].(*workflowservice.RegisterDomainResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterDomain indicates an expected call of RegisterDomain
func (mr *MockHandlerMockRecorder) RegisterDomain(ctx, registerRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterDomain", reflect.TypeOf((*MockHandler)(nil).RegisterDomain), ctx, registerRequest)
}

// UpdateDomain mocks base method
func (m *MockHandler) UpdateDomain(ctx context.Context, updateRequest *workflowservice.UpdateDomainRequest) (*workflowservice.UpdateDomainResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDomain", ctx, updateRequest)
	ret0, _ := ret[0].(*workflowservice.UpdateDomainResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDomain indicates an expected call of UpdateDomain
func (mr *MockHandlerMockRecorder) UpdateDomain(ctx, updateRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDomain", reflect.TypeOf((*MockHandler)(nil).UpdateDomain), ctx, updateRequest)
}
