/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: common/http_agent/http_agent_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIHttpAgent is a mock of IHttpAgent interface
type MockIHttpAgent struct {
	ctrl     *gomock.Controller
	recorder *MockIHttpAgentMockRecorder
}

// MockIHttpAgentMockRecorder is the mock recorder for MockIHttpAgent
type MockIHttpAgentMockRecorder struct {
	mock *MockIHttpAgent
}

// NewMockIHttpAgent creates a new mock instance
func NewMockIHttpAgent(ctrl *gomock.Controller) *MockIHttpAgent {
	mock := &MockIHttpAgent{ctrl: ctrl}
	mock.recorder = &MockIHttpAgentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIHttpAgent) EXPECT() *MockIHttpAgentMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockIHttpAgent) Get(path string, header http.Header, timeoutMs uint64, params map[string]string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", path, header, timeoutMs, params)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockIHttpAgentMockRecorder) Get(path, header, timeoutMs, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIHttpAgent)(nil).Get), path, header, timeoutMs, params)
}

// Post mocks base method
func (m *MockIHttpAgent) Post(path string, header http.Header, timeoutMs uint64, params map[string]string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", path, header, timeoutMs, params)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post
func (mr *MockIHttpAgentMockRecorder) Post(path, header, timeoutMs, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockIHttpAgent)(nil).Post), path, header, timeoutMs, params)
}

// Delete mocks base method
func (m *MockIHttpAgent) Delete(path string, header http.Header, timeoutMs uint64, params map[string]string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", path, header, timeoutMs, params)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockIHttpAgentMockRecorder) Delete(path, header, timeoutMs, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIHttpAgent)(nil).Delete), path, header, timeoutMs, params)
}

// Put mocks base method
func (m *MockIHttpAgent) Put(path string, header http.Header, timeoutMs uint64, params map[string]string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", path, header, timeoutMs, params)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put
func (mr *MockIHttpAgentMockRecorder) Put(path, header, timeoutMs, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockIHttpAgent)(nil).Put), path, header, timeoutMs, params)
}

// RequestOnlyResult mocks base method
func (m *MockIHttpAgent) RequestOnlyResult(method, path string, header http.Header, timeoutMs uint64, params map[string]string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestOnlyResult", method, path, header, timeoutMs, params)
	ret0, _ := ret[0].(string)
	return ret0
}

// RequestOnlyResult indicates an expected call of RequestOnlyResult
func (mr *MockIHttpAgentMockRecorder) RequestOnlyResult(method, path, header, timeoutMs, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestOnlyResult", reflect.TypeOf((*MockIHttpAgent)(nil).RequestOnlyResult), method, path, header, timeoutMs, params)
}

// Request mocks base method
func (m *MockIHttpAgent) Request(method, path string, header http.Header, timeoutMs uint64, params map[string]string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request", method, path, header, timeoutMs, params)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request
func (mr *MockIHttpAgentMockRecorder) Request(method, path, header, timeoutMs, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockIHttpAgent)(nil).Request), method, path, header, timeoutMs, params)
}
