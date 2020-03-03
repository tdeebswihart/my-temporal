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
// Source: domainReplicationQueue.go

// Package persistence is a generated GoMock package.
package persistence

import (
	gomock "github.com/golang/mock/gomock"
	replication "github.com/temporalio/temporal/.gen/proto/replication"
	reflect "reflect"
)

// MockDomainReplicationQueue is a mock of DomainReplicationQueue interface
type MockDomainReplicationQueue struct {
	ctrl     *gomock.Controller
	recorder *MockDomainReplicationQueueMockRecorder
}

// MockDomainReplicationQueueMockRecorder is the mock recorder for MockDomainReplicationQueue
type MockDomainReplicationQueueMockRecorder struct {
	mock *MockDomainReplicationQueue
}

// NewMockDomainReplicationQueue creates a new mock instance
func NewMockDomainReplicationQueue(ctrl *gomock.Controller) *MockDomainReplicationQueue {
	mock := &MockDomainReplicationQueue{ctrl: ctrl}
	mock.recorder = &MockDomainReplicationQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDomainReplicationQueue) EXPECT() *MockDomainReplicationQueueMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockDomainReplicationQueue) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockDomainReplicationQueueMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockDomainReplicationQueue)(nil).Start))
}

// Stop mocks base method
func (m *MockDomainReplicationQueue) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockDomainReplicationQueueMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockDomainReplicationQueue)(nil).Stop))
}

// Publish mocks base method
func (m *MockDomainReplicationQueue) Publish(message interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockDomainReplicationQueueMockRecorder) Publish(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockDomainReplicationQueue)(nil).Publish), message)
}

// PublishToDLQ mocks base method
func (m *MockDomainReplicationQueue) PublishToDLQ(message interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishToDLQ", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishToDLQ indicates an expected call of PublishToDLQ
func (mr *MockDomainReplicationQueueMockRecorder) PublishToDLQ(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishToDLQ", reflect.TypeOf((*MockDomainReplicationQueue)(nil).PublishToDLQ), message)
}

// GetReplicationMessages mocks base method
func (m *MockDomainReplicationQueue) GetReplicationMessages(lastMessageID, maxCount int) ([]*replication.ReplicationTask, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicationMessages", lastMessageID, maxCount)
	ret0, _ := ret[0].([]*replication.ReplicationTask)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetReplicationMessages indicates an expected call of GetReplicationMessages
func (mr *MockDomainReplicationQueueMockRecorder) GetReplicationMessages(lastMessageID, maxCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicationMessages", reflect.TypeOf((*MockDomainReplicationQueue)(nil).GetReplicationMessages), lastMessageID, maxCount)
}

// UpdateAckLevel mocks base method
func (m *MockDomainReplicationQueue) UpdateAckLevel(lastProcessedMessageID int, clusterName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAckLevel", lastProcessedMessageID, clusterName)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAckLevel indicates an expected call of UpdateAckLevel
func (mr *MockDomainReplicationQueueMockRecorder) UpdateAckLevel(lastProcessedMessageID, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAckLevel", reflect.TypeOf((*MockDomainReplicationQueue)(nil).UpdateAckLevel), lastProcessedMessageID, clusterName)
}

// GetAckLevels mocks base method
func (m *MockDomainReplicationQueue) GetAckLevels() (map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAckLevels")
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAckLevels indicates an expected call of GetAckLevels
func (mr *MockDomainReplicationQueueMockRecorder) GetAckLevels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAckLevels", reflect.TypeOf((*MockDomainReplicationQueue)(nil).GetAckLevels))
}

// GetMessagesFromDLQ mocks base method
func (m *MockDomainReplicationQueue) GetMessagesFromDLQ(firstMessageID, lastMessageID, pageSize int, pageToken []byte) ([]*replication.ReplicationTask, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessagesFromDLQ", firstMessageID, lastMessageID, pageSize, pageToken)
	ret0, _ := ret[0].([]*replication.ReplicationTask)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMessagesFromDLQ indicates an expected call of GetMessagesFromDLQ
func (mr *MockDomainReplicationQueueMockRecorder) GetMessagesFromDLQ(firstMessageID, lastMessageID, pageSize, pageToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessagesFromDLQ", reflect.TypeOf((*MockDomainReplicationQueue)(nil).GetMessagesFromDLQ), firstMessageID, lastMessageID, pageSize, pageToken)
}

// UpdateDLQAckLevel mocks base method
func (m *MockDomainReplicationQueue) UpdateDLQAckLevel(lastProcessedMessageID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDLQAckLevel", lastProcessedMessageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDLQAckLevel indicates an expected call of UpdateDLQAckLevel
func (mr *MockDomainReplicationQueueMockRecorder) UpdateDLQAckLevel(lastProcessedMessageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDLQAckLevel", reflect.TypeOf((*MockDomainReplicationQueue)(nil).UpdateDLQAckLevel), lastProcessedMessageID)
}

// GetDLQAckLevel mocks base method
func (m *MockDomainReplicationQueue) GetDLQAckLevel() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDLQAckLevel")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDLQAckLevel indicates an expected call of GetDLQAckLevel
func (mr *MockDomainReplicationQueueMockRecorder) GetDLQAckLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDLQAckLevel", reflect.TypeOf((*MockDomainReplicationQueue)(nil).GetDLQAckLevel))
}

// RangeDeleteMessagesFromDLQ mocks base method
func (m *MockDomainReplicationQueue) RangeDeleteMessagesFromDLQ(firstMessageID, lastMessageID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RangeDeleteMessagesFromDLQ", firstMessageID, lastMessageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RangeDeleteMessagesFromDLQ indicates an expected call of RangeDeleteMessagesFromDLQ
func (mr *MockDomainReplicationQueueMockRecorder) RangeDeleteMessagesFromDLQ(firstMessageID, lastMessageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RangeDeleteMessagesFromDLQ", reflect.TypeOf((*MockDomainReplicationQueue)(nil).RangeDeleteMessagesFromDLQ), firstMessageID, lastMessageID)
}

// DeleteMessageFromDLQ mocks base method
func (m *MockDomainReplicationQueue) DeleteMessageFromDLQ(messageID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessageFromDLQ", messageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessageFromDLQ indicates an expected call of DeleteMessageFromDLQ
func (mr *MockDomainReplicationQueueMockRecorder) DeleteMessageFromDLQ(messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessageFromDLQ", reflect.TypeOf((*MockDomainReplicationQueue)(nil).DeleteMessageFromDLQ), messageID)
}
