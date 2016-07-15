// This file was generated by counterfeiter
package kawasakifakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/kawasaki"
)

type FakeConfigStore struct {
	SetStub        func(handle string, name string, value string)
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		handle string
		name   string
		value  string
	}
	GetStub        func(handle string, name string) (string, bool)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		handle string
		name   string
	}
	getReturns struct {
		result1 string
		result2 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfigStore) Set(handle string, name string, value string) {
	fake.setMutex.Lock()
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		handle string
		name   string
		value  string
	}{handle, name, value})
	fake.recordInvocation("Set", []interface{}{handle, name, value})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		fake.SetStub(handle, name, value)
	}
}

func (fake *FakeConfigStore) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *FakeConfigStore) SetArgsForCall(i int) (string, string, string) {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return fake.setArgsForCall[i].handle, fake.setArgsForCall[i].name, fake.setArgsForCall[i].value
}

func (fake *FakeConfigStore) Get(handle string, name string) (string, bool) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		handle string
		name   string
	}{handle, name})
	fake.recordInvocation("Get", []interface{}{handle, name})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(handle, name)
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *FakeConfigStore) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeConfigStore) GetArgsForCall(i int) (string, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].handle, fake.getArgsForCall[i].name
}

func (fake *FakeConfigStore) GetReturns(result1 string, result2 bool) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 string
		result2 bool
	}{result1, result2}
}

func (fake *FakeConfigStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConfigStore) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ kawasaki.ConfigStore = new(FakeConfigStore)