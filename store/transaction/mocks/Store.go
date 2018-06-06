// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/getamis/eth-indexer/model"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// Delete provides a mock function with given fields: from, to
func (_m *Store) Delete(from int64, to int64) error {
	ret := _m.Called(from, to)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, int64) error); ok {
		r0 = rf(from, to)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindTransaction provides a mock function with given fields: hash
func (_m *Store) FindTransaction(hash []byte) (*model.Transaction, error) {
	ret := _m.Called(hash)

	var r0 *model.Transaction
	if rf, ok := ret.Get(0).(func([]byte) *model.Transaction); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTransactionsByBlockHash provides a mock function with given fields: blockHash
func (_m *Store) FindTransactionsByBlockHash(blockHash []byte) ([]*model.Transaction, error) {
	ret := _m.Called(blockHash)

	var r0 []*model.Transaction
	if rf, ok := ret.Get(0).(func([]byte) []*model.Transaction); ok {
		r0 = rf(blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: data
func (_m *Store) Insert(data *model.Transaction) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Transaction) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}