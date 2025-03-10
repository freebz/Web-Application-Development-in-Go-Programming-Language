// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package auth

import (
	"context"
	"github.com/freebz/go_todo_app/entity"
	"sync"
)

// Ensure, that StoreMock does implement Store.
// If this is not the case, regenerate this file with moq.
var _ Store = &StoreMock{}

// StoreMock is a mock implementation of Store.
//
//	func TestSomethingThatUsesStore(t *testing.T) {
//
//		// make and configure a mocked Store
//		mockedStore := &StoreMock{
//			LoadFunc: func(ctx context.Context, key string) (entity.UserID, error) {
//				panic("mock out the Load method")
//			},
//			SaveFunc: func(ctx context.Context, key string, userID entity.UserID) error {
//				panic("mock out the Save method")
//			},
//		}
//
//		// use mockedStore in code that requires Store
//		// and then make assertions.
//
//	}
type StoreMock struct {
	// LoadFunc mocks the Load method.
	LoadFunc func(ctx context.Context, key string) (entity.UserID, error)

	// SaveFunc mocks the Save method.
	SaveFunc func(ctx context.Context, key string, userID entity.UserID) error

	// calls tracks calls to the methods.
	calls struct {
		// Load holds details about calls to the Load method.
		Load []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Key is the key argument value.
			Key string
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Key is the key argument value.
			Key string
			// UserID is the userID argument value.
			UserID entity.UserID
		}
	}
	lockLoad sync.RWMutex
	lockSave sync.RWMutex
}

// Load calls LoadFunc.
func (mock *StoreMock) Load(ctx context.Context, key string) (entity.UserID, error) {
	if mock.LoadFunc == nil {
		panic("StoreMock.LoadFunc: method is nil but Store.Load was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Key string
	}{
		Ctx: ctx,
		Key: key,
	}
	mock.lockLoad.Lock()
	mock.calls.Load = append(mock.calls.Load, callInfo)
	mock.lockLoad.Unlock()
	return mock.LoadFunc(ctx, key)
}

// LoadCalls gets all the calls that were made to Load.
// Check the length with:
//
//	len(mockedStore.LoadCalls())
func (mock *StoreMock) LoadCalls() []struct {
	Ctx context.Context
	Key string
} {
	var calls []struct {
		Ctx context.Context
		Key string
	}
	mock.lockLoad.RLock()
	calls = mock.calls.Load
	mock.lockLoad.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *StoreMock) Save(ctx context.Context, key string, userID entity.UserID) error {
	if mock.SaveFunc == nil {
		panic("StoreMock.SaveFunc: method is nil but Store.Save was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Key    string
		UserID entity.UserID
	}{
		Ctx:    ctx,
		Key:    key,
		UserID: userID,
	}
	mock.lockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	mock.lockSave.Unlock()
	return mock.SaveFunc(ctx, key, userID)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//
//	len(mockedStore.SaveCalls())
func (mock *StoreMock) SaveCalls() []struct {
	Ctx    context.Context
	Key    string
	UserID entity.UserID
} {
	var calls []struct {
		Ctx    context.Context
		Key    string
		UserID entity.UserID
	}
	mock.lockSave.RLock()
	calls = mock.calls.Save
	mock.lockSave.RUnlock()
	return calls
}
