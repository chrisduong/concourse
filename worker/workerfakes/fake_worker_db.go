// This file was generated by counterfeiter
package workerfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/db/lock"
	"github.com/concourse/atc/worker"
)

type FakeWorkerDB struct {
	PutTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterStub        func(container db.Container, maxLifetime time.Duration) error
	putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterMutex       sync.RWMutex
	putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterArgsForCall []struct {
		container   db.Container
		maxLifetime time.Duration
	}
	putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterReturns struct {
		result1 error
	}
	putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterReturnsOnCall map[int]struct {
		result1 error
	}
	GetContainerStub        func(string) (db.SavedContainer, bool, error)
	getContainerMutex       sync.RWMutex
	getContainerArgsForCall []struct {
		arg1 string
	}
	getContainerReturns struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}
	getContainerReturnsOnCall map[int]struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}
	FindContainerByIdentifierStub        func(db.ContainerIdentifier) (db.SavedContainer, bool, error)
	findContainerByIdentifierMutex       sync.RWMutex
	findContainerByIdentifierArgsForCall []struct {
		arg1 db.ContainerIdentifier
	}
	findContainerByIdentifierReturns struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}
	findContainerByIdentifierReturnsOnCall map[int]struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}
	ReapContainerStub        func(handle string) error
	reapContainerMutex       sync.RWMutex
	reapContainerArgsForCall []struct {
		handle string
	}
	reapContainerReturns struct {
		result1 error
	}
	reapContainerReturnsOnCall map[int]struct {
		result1 error
	}
	GetPipelineByIDStub        func(pipelineID int) (db.SavedPipeline, error)
	getPipelineByIDMutex       sync.RWMutex
	getPipelineByIDArgsForCall []struct {
		pipelineID int
	}
	getPipelineByIDReturns struct {
		result1 db.SavedPipeline
		result2 error
	}
	getPipelineByIDReturnsOnCall map[int]struct {
		result1 db.SavedPipeline
		result2 error
	}
	ReapVolumeStub        func(handle string) error
	reapVolumeMutex       sync.RWMutex
	reapVolumeArgsForCall []struct {
		handle string
	}
	reapVolumeReturns struct {
		result1 error
	}
	reapVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	AcquireVolumeCreatingLockStub        func(lager.Logger, int) (lock.Lock, bool, error)
	acquireVolumeCreatingLockMutex       sync.RWMutex
	acquireVolumeCreatingLockArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
	}
	acquireVolumeCreatingLockReturns struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	acquireVolumeCreatingLockReturnsOnCall map[int]struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	AcquireContainerCreatingLockStub        func(lager.Logger, int) (lock.Lock, bool, error)
	acquireContainerCreatingLockMutex       sync.RWMutex
	acquireContainerCreatingLockArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
	}
	acquireContainerCreatingLockReturns struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	acquireContainerCreatingLockReturnsOnCall map[int]struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorkerDB) PutTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLater(container db.Container, maxLifetime time.Duration) error {
	fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterMutex.Lock()
	ret, specificReturn := fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterReturnsOnCall[len(fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterArgsForCall)]
	fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterArgsForCall = append(fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterArgsForCall, struct {
		container   db.Container
		maxLifetime time.Duration
	}{container, maxLifetime})
	fake.recordInvocation("PutTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLater", []interface{}{container, maxLifetime})
	fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterMutex.Unlock()
	if fake.PutTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterStub != nil {
		return fake.PutTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterStub(container, maxLifetime)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterReturns.result1
}

func (fake *FakeWorkerDB) PutTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterCallCount() int {
	fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterMutex.RLock()
	defer fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterMutex.RUnlock()
	return len(fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterArgsForCall)
}

func (fake *FakeWorkerDB) PutTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterArgsForCall(i int) (db.Container, time.Duration) {
	fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterMutex.RLock()
	defer fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterMutex.RUnlock()
	return fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterArgsForCall[i].container, fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterArgsForCall[i].maxLifetime
}

func (fake *FakeWorkerDB) PutTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterReturns(result1 error) {
	fake.PutTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterStub = nil
	fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) PutTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterReturnsOnCall(i int, result1 error) {
	fake.PutTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterStub = nil
	if fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterReturnsOnCall == nil {
		fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) GetContainer(arg1 string) (db.SavedContainer, bool, error) {
	fake.getContainerMutex.Lock()
	ret, specificReturn := fake.getContainerReturnsOnCall[len(fake.getContainerArgsForCall)]
	fake.getContainerArgsForCall = append(fake.getContainerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetContainer", []interface{}{arg1})
	fake.getContainerMutex.Unlock()
	if fake.GetContainerStub != nil {
		return fake.GetContainerStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getContainerReturns.result1, fake.getContainerReturns.result2, fake.getContainerReturns.result3
}

func (fake *FakeWorkerDB) GetContainerCallCount() int {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return len(fake.getContainerArgsForCall)
}

func (fake *FakeWorkerDB) GetContainerArgsForCall(i int) string {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return fake.getContainerArgsForCall[i].arg1
}

func (fake *FakeWorkerDB) GetContainerReturns(result1 db.SavedContainer, result2 bool, result3 error) {
	fake.GetContainerStub = nil
	fake.getContainerReturns = struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) GetContainerReturnsOnCall(i int, result1 db.SavedContainer, result2 bool, result3 error) {
	fake.GetContainerStub = nil
	if fake.getContainerReturnsOnCall == nil {
		fake.getContainerReturnsOnCall = make(map[int]struct {
			result1 db.SavedContainer
			result2 bool
			result3 error
		})
	}
	fake.getContainerReturnsOnCall[i] = struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) FindContainerByIdentifier(arg1 db.ContainerIdentifier) (db.SavedContainer, bool, error) {
	fake.findContainerByIdentifierMutex.Lock()
	ret, specificReturn := fake.findContainerByIdentifierReturnsOnCall[len(fake.findContainerByIdentifierArgsForCall)]
	fake.findContainerByIdentifierArgsForCall = append(fake.findContainerByIdentifierArgsForCall, struct {
		arg1 db.ContainerIdentifier
	}{arg1})
	fake.recordInvocation("FindContainerByIdentifier", []interface{}{arg1})
	fake.findContainerByIdentifierMutex.Unlock()
	if fake.FindContainerByIdentifierStub != nil {
		return fake.FindContainerByIdentifierStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findContainerByIdentifierReturns.result1, fake.findContainerByIdentifierReturns.result2, fake.findContainerByIdentifierReturns.result3
}

func (fake *FakeWorkerDB) FindContainerByIdentifierCallCount() int {
	fake.findContainerByIdentifierMutex.RLock()
	defer fake.findContainerByIdentifierMutex.RUnlock()
	return len(fake.findContainerByIdentifierArgsForCall)
}

func (fake *FakeWorkerDB) FindContainerByIdentifierArgsForCall(i int) db.ContainerIdentifier {
	fake.findContainerByIdentifierMutex.RLock()
	defer fake.findContainerByIdentifierMutex.RUnlock()
	return fake.findContainerByIdentifierArgsForCall[i].arg1
}

func (fake *FakeWorkerDB) FindContainerByIdentifierReturns(result1 db.SavedContainer, result2 bool, result3 error) {
	fake.FindContainerByIdentifierStub = nil
	fake.findContainerByIdentifierReturns = struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) FindContainerByIdentifierReturnsOnCall(i int, result1 db.SavedContainer, result2 bool, result3 error) {
	fake.FindContainerByIdentifierStub = nil
	if fake.findContainerByIdentifierReturnsOnCall == nil {
		fake.findContainerByIdentifierReturnsOnCall = make(map[int]struct {
			result1 db.SavedContainer
			result2 bool
			result3 error
		})
	}
	fake.findContainerByIdentifierReturnsOnCall[i] = struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) ReapContainer(handle string) error {
	fake.reapContainerMutex.Lock()
	ret, specificReturn := fake.reapContainerReturnsOnCall[len(fake.reapContainerArgsForCall)]
	fake.reapContainerArgsForCall = append(fake.reapContainerArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("ReapContainer", []interface{}{handle})
	fake.reapContainerMutex.Unlock()
	if fake.ReapContainerStub != nil {
		return fake.ReapContainerStub(handle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.reapContainerReturns.result1
}

func (fake *FakeWorkerDB) ReapContainerCallCount() int {
	fake.reapContainerMutex.RLock()
	defer fake.reapContainerMutex.RUnlock()
	return len(fake.reapContainerArgsForCall)
}

func (fake *FakeWorkerDB) ReapContainerArgsForCall(i int) string {
	fake.reapContainerMutex.RLock()
	defer fake.reapContainerMutex.RUnlock()
	return fake.reapContainerArgsForCall[i].handle
}

func (fake *FakeWorkerDB) ReapContainerReturns(result1 error) {
	fake.ReapContainerStub = nil
	fake.reapContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) ReapContainerReturnsOnCall(i int, result1 error) {
	fake.ReapContainerStub = nil
	if fake.reapContainerReturnsOnCall == nil {
		fake.reapContainerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.reapContainerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) GetPipelineByID(pipelineID int) (db.SavedPipeline, error) {
	fake.getPipelineByIDMutex.Lock()
	ret, specificReturn := fake.getPipelineByIDReturnsOnCall[len(fake.getPipelineByIDArgsForCall)]
	fake.getPipelineByIDArgsForCall = append(fake.getPipelineByIDArgsForCall, struct {
		pipelineID int
	}{pipelineID})
	fake.recordInvocation("GetPipelineByID", []interface{}{pipelineID})
	fake.getPipelineByIDMutex.Unlock()
	if fake.GetPipelineByIDStub != nil {
		return fake.GetPipelineByIDStub(pipelineID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getPipelineByIDReturns.result1, fake.getPipelineByIDReturns.result2
}

func (fake *FakeWorkerDB) GetPipelineByIDCallCount() int {
	fake.getPipelineByIDMutex.RLock()
	defer fake.getPipelineByIDMutex.RUnlock()
	return len(fake.getPipelineByIDArgsForCall)
}

func (fake *FakeWorkerDB) GetPipelineByIDArgsForCall(i int) int {
	fake.getPipelineByIDMutex.RLock()
	defer fake.getPipelineByIDMutex.RUnlock()
	return fake.getPipelineByIDArgsForCall[i].pipelineID
}

func (fake *FakeWorkerDB) GetPipelineByIDReturns(result1 db.SavedPipeline, result2 error) {
	fake.GetPipelineByIDStub = nil
	fake.getPipelineByIDReturns = struct {
		result1 db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerDB) GetPipelineByIDReturnsOnCall(i int, result1 db.SavedPipeline, result2 error) {
	fake.GetPipelineByIDStub = nil
	if fake.getPipelineByIDReturnsOnCall == nil {
		fake.getPipelineByIDReturnsOnCall = make(map[int]struct {
			result1 db.SavedPipeline
			result2 error
		})
	}
	fake.getPipelineByIDReturnsOnCall[i] = struct {
		result1 db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerDB) ReapVolume(handle string) error {
	fake.reapVolumeMutex.Lock()
	ret, specificReturn := fake.reapVolumeReturnsOnCall[len(fake.reapVolumeArgsForCall)]
	fake.reapVolumeArgsForCall = append(fake.reapVolumeArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("ReapVolume", []interface{}{handle})
	fake.reapVolumeMutex.Unlock()
	if fake.ReapVolumeStub != nil {
		return fake.ReapVolumeStub(handle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.reapVolumeReturns.result1
}

func (fake *FakeWorkerDB) ReapVolumeCallCount() int {
	fake.reapVolumeMutex.RLock()
	defer fake.reapVolumeMutex.RUnlock()
	return len(fake.reapVolumeArgsForCall)
}

func (fake *FakeWorkerDB) ReapVolumeArgsForCall(i int) string {
	fake.reapVolumeMutex.RLock()
	defer fake.reapVolumeMutex.RUnlock()
	return fake.reapVolumeArgsForCall[i].handle
}

func (fake *FakeWorkerDB) ReapVolumeReturns(result1 error) {
	fake.ReapVolumeStub = nil
	fake.reapVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) ReapVolumeReturnsOnCall(i int, result1 error) {
	fake.ReapVolumeStub = nil
	if fake.reapVolumeReturnsOnCall == nil {
		fake.reapVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.reapVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) AcquireVolumeCreatingLock(arg1 lager.Logger, arg2 int) (lock.Lock, bool, error) {
	fake.acquireVolumeCreatingLockMutex.Lock()
	ret, specificReturn := fake.acquireVolumeCreatingLockReturnsOnCall[len(fake.acquireVolumeCreatingLockArgsForCall)]
	fake.acquireVolumeCreatingLockArgsForCall = append(fake.acquireVolumeCreatingLockArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("AcquireVolumeCreatingLock", []interface{}{arg1, arg2})
	fake.acquireVolumeCreatingLockMutex.Unlock()
	if fake.AcquireVolumeCreatingLockStub != nil {
		return fake.AcquireVolumeCreatingLockStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.acquireVolumeCreatingLockReturns.result1, fake.acquireVolumeCreatingLockReturns.result2, fake.acquireVolumeCreatingLockReturns.result3
}

func (fake *FakeWorkerDB) AcquireVolumeCreatingLockCallCount() int {
	fake.acquireVolumeCreatingLockMutex.RLock()
	defer fake.acquireVolumeCreatingLockMutex.RUnlock()
	return len(fake.acquireVolumeCreatingLockArgsForCall)
}

func (fake *FakeWorkerDB) AcquireVolumeCreatingLockArgsForCall(i int) (lager.Logger, int) {
	fake.acquireVolumeCreatingLockMutex.RLock()
	defer fake.acquireVolumeCreatingLockMutex.RUnlock()
	return fake.acquireVolumeCreatingLockArgsForCall[i].arg1, fake.acquireVolumeCreatingLockArgsForCall[i].arg2
}

func (fake *FakeWorkerDB) AcquireVolumeCreatingLockReturns(result1 lock.Lock, result2 bool, result3 error) {
	fake.AcquireVolumeCreatingLockStub = nil
	fake.acquireVolumeCreatingLockReturns = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) AcquireVolumeCreatingLockReturnsOnCall(i int, result1 lock.Lock, result2 bool, result3 error) {
	fake.AcquireVolumeCreatingLockStub = nil
	if fake.acquireVolumeCreatingLockReturnsOnCall == nil {
		fake.acquireVolumeCreatingLockReturnsOnCall = make(map[int]struct {
			result1 lock.Lock
			result2 bool
			result3 error
		})
	}
	fake.acquireVolumeCreatingLockReturnsOnCall[i] = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) AcquireContainerCreatingLock(arg1 lager.Logger, arg2 int) (lock.Lock, bool, error) {
	fake.acquireContainerCreatingLockMutex.Lock()
	ret, specificReturn := fake.acquireContainerCreatingLockReturnsOnCall[len(fake.acquireContainerCreatingLockArgsForCall)]
	fake.acquireContainerCreatingLockArgsForCall = append(fake.acquireContainerCreatingLockArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("AcquireContainerCreatingLock", []interface{}{arg1, arg2})
	fake.acquireContainerCreatingLockMutex.Unlock()
	if fake.AcquireContainerCreatingLockStub != nil {
		return fake.AcquireContainerCreatingLockStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.acquireContainerCreatingLockReturns.result1, fake.acquireContainerCreatingLockReturns.result2, fake.acquireContainerCreatingLockReturns.result3
}

func (fake *FakeWorkerDB) AcquireContainerCreatingLockCallCount() int {
	fake.acquireContainerCreatingLockMutex.RLock()
	defer fake.acquireContainerCreatingLockMutex.RUnlock()
	return len(fake.acquireContainerCreatingLockArgsForCall)
}

func (fake *FakeWorkerDB) AcquireContainerCreatingLockArgsForCall(i int) (lager.Logger, int) {
	fake.acquireContainerCreatingLockMutex.RLock()
	defer fake.acquireContainerCreatingLockMutex.RUnlock()
	return fake.acquireContainerCreatingLockArgsForCall[i].arg1, fake.acquireContainerCreatingLockArgsForCall[i].arg2
}

func (fake *FakeWorkerDB) AcquireContainerCreatingLockReturns(result1 lock.Lock, result2 bool, result3 error) {
	fake.AcquireContainerCreatingLockStub = nil
	fake.acquireContainerCreatingLockReturns = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) AcquireContainerCreatingLockReturnsOnCall(i int, result1 lock.Lock, result2 bool, result3 error) {
	fake.AcquireContainerCreatingLockStub = nil
	if fake.acquireContainerCreatingLockReturnsOnCall == nil {
		fake.acquireContainerCreatingLockReturnsOnCall = make(map[int]struct {
			result1 lock.Lock
			result2 bool
			result3 error
		})
	}
	fake.acquireContainerCreatingLockReturnsOnCall[i] = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterMutex.RLock()
	defer fake.putTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLaterMutex.RUnlock()
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	fake.findContainerByIdentifierMutex.RLock()
	defer fake.findContainerByIdentifierMutex.RUnlock()
	fake.reapContainerMutex.RLock()
	defer fake.reapContainerMutex.RUnlock()
	fake.getPipelineByIDMutex.RLock()
	defer fake.getPipelineByIDMutex.RUnlock()
	fake.reapVolumeMutex.RLock()
	defer fake.reapVolumeMutex.RUnlock()
	fake.acquireVolumeCreatingLockMutex.RLock()
	defer fake.acquireVolumeCreatingLockMutex.RUnlock()
	fake.acquireContainerCreatingLockMutex.RLock()
	defer fake.acquireContainerCreatingLockMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeWorkerDB) recordInvocation(key string, args []interface{}) {
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

var _ worker.WorkerDB = new(FakeWorkerDB)
