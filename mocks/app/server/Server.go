// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Server is an autogenerated mock type for the Server type
type Server struct {
	mock.Mock
}

// Run provides a mock function with given fields: ctx
func (_m *Server) Run(ctx context.Context) {
	_m.Called(ctx)
}
