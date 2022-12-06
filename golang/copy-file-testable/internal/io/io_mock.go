package io

import "github.com/stretchr/testify/mock"

type ReaderMock struct {
	mock.Mock
}

func NewReaderMock() *ReaderMock {
	return &ReaderMock{}
}

func (m *ReaderMock) Read(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Get(0).(int), args.Error(0)
}

func (m *ReaderMock) Close() (err error) {
	args := m.Called()
	return args.Error(0)
}
