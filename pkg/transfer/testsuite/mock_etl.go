// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: transfer/etl (interfaces: Container,Field,Record,Schema)

// Package testsuite is a generated GoMock package.
package testsuite

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	etl "github.com/TencentBlueKing/bkmonitor-datalink/pkg/transfer/etl"
)

// MockContainer is a mock of Container interface.
type MockContainer struct {
	ctrl     *gomock.Controller
	recorder *MockContainerMockRecorder
}

// MockContainerMockRecorder is the mock recorder for MockContainer.
type MockContainerMockRecorder struct {
	mock *MockContainer
}

// NewMockContainer creates a new mock instance.
func NewMockContainer(ctrl *gomock.Controller) *MockContainer {
	mock := &MockContainer{ctrl: ctrl}
	mock.recorder = &MockContainerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContainer) EXPECT() *MockContainerMockRecorder {
	return m.recorder
}

// Copy mocks base method.
func (m *MockContainer) Copy() etl.Container {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy")
	ret0, _ := ret[0].(etl.Container)
	return ret0
}

// Copy indicates an expected call of Copy.
func (mr *MockContainerMockRecorder) Copy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockContainer)(nil).Copy))
}

// Del mocks base method.
func (m *MockContainer) Del(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockContainerMockRecorder) Del(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockContainer)(nil).Del), arg0)
}

// Get mocks base method.
func (m *MockContainer) Get(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockContainerMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockContainer)(nil).Get), arg0)
}

// Keys mocks base method.
func (m *MockContainer) Keys() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockContainerMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockContainer)(nil).Keys))
}

// Put mocks base method.
func (m *MockContainer) Put(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockContainerMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockContainer)(nil).Put), arg0, arg1)
}

// MockField is a mock of Field interface.
type MockField struct {
	ctrl     *gomock.Controller
	recorder *MockFieldMockRecorder
}

// MockFieldMockRecorder is the mock recorder for MockField.
type MockFieldMockRecorder struct {
	mock *MockField
}

// NewMockField creates a new mock instance.
func NewMockField(ctrl *gomock.Controller) *MockField {
	mock := &MockField{ctrl: ctrl}
	mock.recorder = &MockFieldMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockField) EXPECT() *MockFieldMockRecorder {
	return m.recorder
}

// DefaultValue mocks base method.
func (m *MockField) DefaultValue() (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultValue")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// DefaultValue indicates an expected call of DefaultValue.
func (mr *MockFieldMockRecorder) DefaultValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultValue", reflect.TypeOf((*MockField)(nil).DefaultValue))
}

// Name mocks base method.
func (m *MockField) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockFieldMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockField)(nil).Name))
}

// String mocks base method.
func (m *MockField) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockFieldMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockField)(nil).String))
}

// Transform mocks base method.
func (m *MockField) Transform(arg0, arg1 etl.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transform", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transform indicates an expected call of Transform.
func (mr *MockFieldMockRecorder) Transform(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transform", reflect.TypeOf((*MockField)(nil).Transform), arg0, arg1)
}

// MockRecord is a mock of Record interface.
type MockRecord struct {
	ctrl     *gomock.Controller
	recorder *MockRecordMockRecorder
}

// MockRecordMockRecorder is the mock recorder for MockRecord.
type MockRecordMockRecorder struct {
	mock *MockRecord
}

// NewMockRecord creates a new mock instance.
func NewMockRecord(ctrl *gomock.Controller) *MockRecord {
	mock := &MockRecord{ctrl: ctrl}
	mock.recorder = &MockRecordMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRecord) EXPECT() *MockRecordMockRecorder {
	return m.recorder
}

// Finish mocks base method.
func (m *MockRecord) Finish() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finish")
	ret0, _ := ret[0].(error)
	return ret0
}

// Finish indicates an expected call of Finish.
func (mr *MockRecordMockRecorder) Finish() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finish", reflect.TypeOf((*MockRecord)(nil).Finish))
}

// Name mocks base method.
func (m *MockRecord) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockRecordMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockRecord)(nil).Name))
}

// String mocks base method.
func (m *MockRecord) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockRecordMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockRecord)(nil).String))
}

// Transform mocks base method.
func (m *MockRecord) Transform(arg0, arg1 etl.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transform", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transform indicates an expected call of Transform.
func (mr *MockRecordMockRecorder) Transform(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transform", reflect.TypeOf((*MockRecord)(nil).Transform), arg0, arg1)
}

// MockSchema is a mock of Schema interface.
type MockSchema struct {
	ctrl     *gomock.Controller
	recorder *MockSchemaMockRecorder
}

// MockSchemaMockRecorder is the mock recorder for MockSchema.
type MockSchemaMockRecorder struct {
	mock *MockSchema
}

// NewMockSchema creates a new mock instance.
func NewMockSchema(ctrl *gomock.Controller) *MockSchema {
	mock := &MockSchema{ctrl: ctrl}
	mock.recorder = &MockSchemaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchema) EXPECT() *MockSchemaMockRecorder {
	return m.recorder
}

// String mocks base method.
func (m *MockSchema) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockSchemaMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockSchema)(nil).String))
}

// Transform mocks base method.
func (m *MockSchema) Transform(arg0, arg1 etl.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transform", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transform indicates an expected call of Transform.
func (mr *MockSchemaMockRecorder) Transform(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transform", reflect.TypeOf((*MockSchema)(nil).Transform), arg0, arg1)
}
