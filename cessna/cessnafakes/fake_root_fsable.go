// This file was generated by counterfeiter
package cessnafakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/cessna"
)

type FakeRootFSable struct {
	RootFSPathForStub        func(logger lager.Logger, worker cessna.Worker) (string, error)
	rootFSPathForMutex       sync.RWMutex
	rootFSPathForArgsForCall []struct {
		logger lager.Logger
		worker cessna.Worker
	}
	rootFSPathForReturns struct {
		result1 string
		result2 error
	}
	rootFSPathForReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRootFSable) RootFSPathFor(logger lager.Logger, worker cessna.Worker) (string, error) {
	fake.rootFSPathForMutex.Lock()
	ret, specificReturn := fake.rootFSPathForReturnsOnCall[len(fake.rootFSPathForArgsForCall)]
	fake.rootFSPathForArgsForCall = append(fake.rootFSPathForArgsForCall, struct {
		logger lager.Logger
		worker cessna.Worker
	}{logger, worker})
	fake.recordInvocation("RootFSPathFor", []interface{}{logger, worker})
	fake.rootFSPathForMutex.Unlock()
	if fake.RootFSPathForStub != nil {
		return fake.RootFSPathForStub(logger, worker)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.rootFSPathForReturns.result1, fake.rootFSPathForReturns.result2
}

func (fake *FakeRootFSable) RootFSPathForCallCount() int {
	fake.rootFSPathForMutex.RLock()
	defer fake.rootFSPathForMutex.RUnlock()
	return len(fake.rootFSPathForArgsForCall)
}

func (fake *FakeRootFSable) RootFSPathForArgsForCall(i int) (lager.Logger, cessna.Worker) {
	fake.rootFSPathForMutex.RLock()
	defer fake.rootFSPathForMutex.RUnlock()
	return fake.rootFSPathForArgsForCall[i].logger, fake.rootFSPathForArgsForCall[i].worker
}

func (fake *FakeRootFSable) RootFSPathForReturns(result1 string, result2 error) {
	fake.RootFSPathForStub = nil
	fake.rootFSPathForReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRootFSable) RootFSPathForReturnsOnCall(i int, result1 string, result2 error) {
	fake.RootFSPathForStub = nil
	if fake.rootFSPathForReturnsOnCall == nil {
		fake.rootFSPathForReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.rootFSPathForReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRootFSable) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.rootFSPathForMutex.RLock()
	defer fake.rootFSPathForMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRootFSable) recordInvocation(key string, args []interface{}) {
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

var _ cessna.RootFSable = new(FakeRootFSable)
