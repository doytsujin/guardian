// This file was generated by counterfeiter
package kawasakifakes

import (
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/kawasaki"
	"code.cloudfoundry.org/lager"
)

type FakeNetworker struct {
	CapacityStub        func() uint64
	capacityMutex       sync.RWMutex
	capacityArgsForCall []struct{}
	capacityReturns     struct {
		result1 uint64
	}
	NetworkStub        func(log lager.Logger, spec garden.ContainerSpec, pid int) error
	networkMutex       sync.RWMutex
	networkArgsForCall []struct {
		log  lager.Logger
		spec garden.ContainerSpec
		pid  int
	}
	networkReturns struct {
		result1 error
	}
	DestroyStub        func(log lager.Logger, handle string) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	destroyReturns struct {
		result1 error
	}
	NetInStub        func(log lager.Logger, handle string, externalPort, containerPort uint32) (uint32, uint32, error)
	netInMutex       sync.RWMutex
	netInArgsForCall []struct {
		log           lager.Logger
		handle        string
		externalPort  uint32
		containerPort uint32
	}
	netInReturns struct {
		result1 uint32
		result2 uint32
		result3 error
	}
	NetOutStub        func(log lager.Logger, handle string, rule garden.NetOutRule) error
	netOutMutex       sync.RWMutex
	netOutArgsForCall []struct {
		log    lager.Logger
		handle string
		rule   garden.NetOutRule
	}
	netOutReturns struct {
		result1 error
	}
	RestoreStub        func(log lager.Logger, handle string) error
	restoreMutex       sync.RWMutex
	restoreArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	restoreReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNetworker) Capacity() uint64 {
	fake.capacityMutex.Lock()
	fake.capacityArgsForCall = append(fake.capacityArgsForCall, struct{}{})
	fake.recordInvocation("Capacity", []interface{}{})
	fake.capacityMutex.Unlock()
	if fake.CapacityStub != nil {
		return fake.CapacityStub()
	} else {
		return fake.capacityReturns.result1
	}
}

func (fake *FakeNetworker) CapacityCallCount() int {
	fake.capacityMutex.RLock()
	defer fake.capacityMutex.RUnlock()
	return len(fake.capacityArgsForCall)
}

func (fake *FakeNetworker) CapacityReturns(result1 uint64) {
	fake.CapacityStub = nil
	fake.capacityReturns = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeNetworker) Network(log lager.Logger, spec garden.ContainerSpec, pid int) error {
	fake.networkMutex.Lock()
	fake.networkArgsForCall = append(fake.networkArgsForCall, struct {
		log  lager.Logger
		spec garden.ContainerSpec
		pid  int
	}{log, spec, pid})
	fake.recordInvocation("Network", []interface{}{log, spec, pid})
	fake.networkMutex.Unlock()
	if fake.NetworkStub != nil {
		return fake.NetworkStub(log, spec, pid)
	} else {
		return fake.networkReturns.result1
	}
}

func (fake *FakeNetworker) NetworkCallCount() int {
	fake.networkMutex.RLock()
	defer fake.networkMutex.RUnlock()
	return len(fake.networkArgsForCall)
}

func (fake *FakeNetworker) NetworkArgsForCall(i int) (lager.Logger, garden.ContainerSpec, int) {
	fake.networkMutex.RLock()
	defer fake.networkMutex.RUnlock()
	return fake.networkArgsForCall[i].log, fake.networkArgsForCall[i].spec, fake.networkArgsForCall[i].pid
}

func (fake *FakeNetworker) NetworkReturns(result1 error) {
	fake.NetworkStub = nil
	fake.networkReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworker) Destroy(log lager.Logger, handle string) error {
	fake.destroyMutex.Lock()
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("Destroy", []interface{}{log, handle})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(log, handle)
	} else {
		return fake.destroyReturns.result1
	}
}

func (fake *FakeNetworker) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeNetworker) DestroyArgsForCall(i int) (lager.Logger, string) {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return fake.destroyArgsForCall[i].log, fake.destroyArgsForCall[i].handle
}

func (fake *FakeNetworker) DestroyReturns(result1 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworker) NetIn(log lager.Logger, handle string, externalPort uint32, containerPort uint32) (uint32, uint32, error) {
	fake.netInMutex.Lock()
	fake.netInArgsForCall = append(fake.netInArgsForCall, struct {
		log           lager.Logger
		handle        string
		externalPort  uint32
		containerPort uint32
	}{log, handle, externalPort, containerPort})
	fake.recordInvocation("NetIn", []interface{}{log, handle, externalPort, containerPort})
	fake.netInMutex.Unlock()
	if fake.NetInStub != nil {
		return fake.NetInStub(log, handle, externalPort, containerPort)
	} else {
		return fake.netInReturns.result1, fake.netInReturns.result2, fake.netInReturns.result3
	}
}

func (fake *FakeNetworker) NetInCallCount() int {
	fake.netInMutex.RLock()
	defer fake.netInMutex.RUnlock()
	return len(fake.netInArgsForCall)
}

func (fake *FakeNetworker) NetInArgsForCall(i int) (lager.Logger, string, uint32, uint32) {
	fake.netInMutex.RLock()
	defer fake.netInMutex.RUnlock()
	return fake.netInArgsForCall[i].log, fake.netInArgsForCall[i].handle, fake.netInArgsForCall[i].externalPort, fake.netInArgsForCall[i].containerPort
}

func (fake *FakeNetworker) NetInReturns(result1 uint32, result2 uint32, result3 error) {
	fake.NetInStub = nil
	fake.netInReturns = struct {
		result1 uint32
		result2 uint32
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeNetworker) NetOut(log lager.Logger, handle string, rule garden.NetOutRule) error {
	fake.netOutMutex.Lock()
	fake.netOutArgsForCall = append(fake.netOutArgsForCall, struct {
		log    lager.Logger
		handle string
		rule   garden.NetOutRule
	}{log, handle, rule})
	fake.recordInvocation("NetOut", []interface{}{log, handle, rule})
	fake.netOutMutex.Unlock()
	if fake.NetOutStub != nil {
		return fake.NetOutStub(log, handle, rule)
	} else {
		return fake.netOutReturns.result1
	}
}

func (fake *FakeNetworker) NetOutCallCount() int {
	fake.netOutMutex.RLock()
	defer fake.netOutMutex.RUnlock()
	return len(fake.netOutArgsForCall)
}

func (fake *FakeNetworker) NetOutArgsForCall(i int) (lager.Logger, string, garden.NetOutRule) {
	fake.netOutMutex.RLock()
	defer fake.netOutMutex.RUnlock()
	return fake.netOutArgsForCall[i].log, fake.netOutArgsForCall[i].handle, fake.netOutArgsForCall[i].rule
}

func (fake *FakeNetworker) NetOutReturns(result1 error) {
	fake.NetOutStub = nil
	fake.netOutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworker) Restore(log lager.Logger, handle string) error {
	fake.restoreMutex.Lock()
	fake.restoreArgsForCall = append(fake.restoreArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("Restore", []interface{}{log, handle})
	fake.restoreMutex.Unlock()
	if fake.RestoreStub != nil {
		return fake.RestoreStub(log, handle)
	} else {
		return fake.restoreReturns.result1
	}
}

func (fake *FakeNetworker) RestoreCallCount() int {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return len(fake.restoreArgsForCall)
}

func (fake *FakeNetworker) RestoreArgsForCall(i int) (lager.Logger, string) {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return fake.restoreArgsForCall[i].log, fake.restoreArgsForCall[i].handle
}

func (fake *FakeNetworker) RestoreReturns(result1 error) {
	fake.RestoreStub = nil
	fake.restoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.capacityMutex.RLock()
	defer fake.capacityMutex.RUnlock()
	fake.networkMutex.RLock()
	defer fake.networkMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.netInMutex.RLock()
	defer fake.netInMutex.RUnlock()
	fake.netOutMutex.RLock()
	defer fake.netOutMutex.RUnlock()
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeNetworker) recordInvocation(key string, args []interface{}) {
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

var _ kawasaki.Networker = new(FakeNetworker)