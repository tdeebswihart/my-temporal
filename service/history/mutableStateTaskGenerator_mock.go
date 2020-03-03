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
// Source: mutableStateTaskGenerator.go

// Package history is a generated GoMock package.
package history

import (
	gomock "github.com/golang/mock/gomock"
	shared "github.com/temporalio/temporal/.gen/go/shared"
	reflect "reflect"
	time "time"
)

// MockmutableStateTaskGenerator is a mock of mutableStateTaskGenerator interface
type MockmutableStateTaskGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockmutableStateTaskGeneratorMockRecorder
}

// MockmutableStateTaskGeneratorMockRecorder is the mock recorder for MockmutableStateTaskGenerator
type MockmutableStateTaskGeneratorMockRecorder struct {
	mock *MockmutableStateTaskGenerator
}

// NewMockmutableStateTaskGenerator creates a new mock instance
func NewMockmutableStateTaskGenerator(ctrl *gomock.Controller) *MockmutableStateTaskGenerator {
	mock := &MockmutableStateTaskGenerator{ctrl: ctrl}
	mock.recorder = &MockmutableStateTaskGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockmutableStateTaskGenerator) EXPECT() *MockmutableStateTaskGeneratorMockRecorder {
	return m.recorder
}

// generateWorkflowStartTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateWorkflowStartTasks(now time.Time, startEvent *shared.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateWorkflowStartTasks", now, startEvent)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateWorkflowStartTasks indicates an expected call of generateWorkflowStartTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateWorkflowStartTasks(now, startEvent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateWorkflowStartTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateWorkflowStartTasks), now, startEvent)
}

// generateWorkflowCloseTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateWorkflowCloseTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateWorkflowCloseTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateWorkflowCloseTasks indicates an expected call of generateWorkflowCloseTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateWorkflowCloseTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateWorkflowCloseTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateWorkflowCloseTasks), now)
}

// generateRecordWorkflowStartedTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateRecordWorkflowStartedTasks(now time.Time, startEvent *shared.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateRecordWorkflowStartedTasks", now, startEvent)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateRecordWorkflowStartedTasks indicates an expected call of generateRecordWorkflowStartedTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateRecordWorkflowStartedTasks(now, startEvent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateRecordWorkflowStartedTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateRecordWorkflowStartedTasks), now, startEvent)
}

// generateDelayedDecisionTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateDelayedDecisionTasks(now time.Time, startEvent *shared.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateDelayedDecisionTasks", now, startEvent)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateDelayedDecisionTasks indicates an expected call of generateDelayedDecisionTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateDelayedDecisionTasks(now, startEvent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateDelayedDecisionTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateDelayedDecisionTasks), now, startEvent)
}

// generateDecisionScheduleTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateDecisionScheduleTasks(now time.Time, decisionScheduleID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateDecisionScheduleTasks", now, decisionScheduleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateDecisionScheduleTasks indicates an expected call of generateDecisionScheduleTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateDecisionScheduleTasks(now, decisionScheduleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateDecisionScheduleTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateDecisionScheduleTasks), now, decisionScheduleID)
}

// generateDecisionStartTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateDecisionStartTasks(now time.Time, decisionScheduleID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateDecisionStartTasks", now, decisionScheduleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateDecisionStartTasks indicates an expected call of generateDecisionStartTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateDecisionStartTasks(now, decisionScheduleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateDecisionStartTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateDecisionStartTasks), now, decisionScheduleID)
}

// generateActivityTransferTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateActivityTransferTasks(now time.Time, event *shared.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateActivityTransferTasks", now, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateActivityTransferTasks indicates an expected call of generateActivityTransferTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateActivityTransferTasks(now, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateActivityTransferTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateActivityTransferTasks), now, event)
}

// generateActivityRetryTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateActivityRetryTasks(activityScheduleID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateActivityRetryTasks", activityScheduleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateActivityRetryTasks indicates an expected call of generateActivityRetryTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateActivityRetryTasks(activityScheduleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateActivityRetryTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateActivityRetryTasks), activityScheduleID)
}

// generateChildWorkflowTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateChildWorkflowTasks(now time.Time, event *shared.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateChildWorkflowTasks", now, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateChildWorkflowTasks indicates an expected call of generateChildWorkflowTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateChildWorkflowTasks(now, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateChildWorkflowTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateChildWorkflowTasks), now, event)
}

// generateRequestCancelExternalTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateRequestCancelExternalTasks(now time.Time, event *shared.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateRequestCancelExternalTasks", now, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateRequestCancelExternalTasks indicates an expected call of generateRequestCancelExternalTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateRequestCancelExternalTasks(now, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateRequestCancelExternalTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateRequestCancelExternalTasks), now, event)
}

// generateSignalExternalTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateSignalExternalTasks(now time.Time, event *shared.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateSignalExternalTasks", now, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateSignalExternalTasks indicates an expected call of generateSignalExternalTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateSignalExternalTasks(now, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateSignalExternalTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateSignalExternalTasks), now, event)
}

// generateWorkflowSearchAttrTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateWorkflowSearchAttrTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateWorkflowSearchAttrTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateWorkflowSearchAttrTasks indicates an expected call of generateWorkflowSearchAttrTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateWorkflowSearchAttrTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateWorkflowSearchAttrTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateWorkflowSearchAttrTasks), now)
}

// generateWorkflowResetTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateWorkflowResetTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateWorkflowResetTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateWorkflowResetTasks indicates an expected call of generateWorkflowResetTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateWorkflowResetTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateWorkflowResetTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateWorkflowResetTasks), now)
}

// generateActivityTimerTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateActivityTimerTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateActivityTimerTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateActivityTimerTasks indicates an expected call of generateActivityTimerTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateActivityTimerTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateActivityTimerTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateActivityTimerTasks), now)
}

// generateUserTimerTasks mocks base method
func (m *MockmutableStateTaskGenerator) generateUserTimerTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "generateUserTimerTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// generateUserTimerTasks indicates an expected call of generateUserTimerTasks
func (mr *MockmutableStateTaskGeneratorMockRecorder) generateUserTimerTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateUserTimerTasks", reflect.TypeOf((*MockmutableStateTaskGenerator)(nil).generateUserTimerTasks), now)
}
