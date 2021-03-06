// Code generated by counterfeiter. DO NOT EDIT.
package lockfakes

import (
	"sync"

	"github.com/concourse/atc/db/lock"
)

type FakeLockDB struct {
	AcquireStub        func(id lock.LockID) (bool, error)
	acquireMutex       sync.RWMutex
	acquireArgsForCall []struct {
		id lock.LockID
	}
	acquireReturns struct {
		result1 bool
		result2 error
	}
	acquireReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ReleaseStub        func(id lock.LockID) (bool, error)
	releaseMutex       sync.RWMutex
	releaseArgsForCall []struct {
		id lock.LockID
	}
	releaseReturns struct {
		result1 bool
		result2 error
	}
	releaseReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLockDB) Acquire(id lock.LockID) (bool, error) {
	fake.acquireMutex.Lock()
	ret, specificReturn := fake.acquireReturnsOnCall[len(fake.acquireArgsForCall)]
	fake.acquireArgsForCall = append(fake.acquireArgsForCall, struct {
		id lock.LockID
	}{id})
	fake.recordInvocation("Acquire", []interface{}{id})
	fake.acquireMutex.Unlock()
	if fake.AcquireStub != nil {
		return fake.AcquireStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.acquireReturns.result1, fake.acquireReturns.result2
}

func (fake *FakeLockDB) AcquireCallCount() int {
	fake.acquireMutex.RLock()
	defer fake.acquireMutex.RUnlock()
	return len(fake.acquireArgsForCall)
}

func (fake *FakeLockDB) AcquireArgsForCall(i int) lock.LockID {
	fake.acquireMutex.RLock()
	defer fake.acquireMutex.RUnlock()
	return fake.acquireArgsForCall[i].id
}

func (fake *FakeLockDB) AcquireReturns(result1 bool, result2 error) {
	fake.AcquireStub = nil
	fake.acquireReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) AcquireReturnsOnCall(i int, result1 bool, result2 error) {
	fake.AcquireStub = nil
	if fake.acquireReturnsOnCall == nil {
		fake.acquireReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.acquireReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) Release(id lock.LockID) (bool, error) {
	fake.releaseMutex.Lock()
	ret, specificReturn := fake.releaseReturnsOnCall[len(fake.releaseArgsForCall)]
	fake.releaseArgsForCall = append(fake.releaseArgsForCall, struct {
		id lock.LockID
	}{id})
	fake.recordInvocation("Release", []interface{}{id})
	fake.releaseMutex.Unlock()
	if fake.ReleaseStub != nil {
		return fake.ReleaseStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.releaseReturns.result1, fake.releaseReturns.result2
}

func (fake *FakeLockDB) ReleaseCallCount() int {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return len(fake.releaseArgsForCall)
}

func (fake *FakeLockDB) ReleaseArgsForCall(i int) lock.LockID {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return fake.releaseArgsForCall[i].id
}

func (fake *FakeLockDB) ReleaseReturns(result1 bool, result2 error) {
	fake.ReleaseStub = nil
	fake.releaseReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) ReleaseReturnsOnCall(i int, result1 bool, result2 error) {
	fake.ReleaseStub = nil
	if fake.releaseReturnsOnCall == nil {
		fake.releaseReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.releaseReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acquireMutex.RLock()
	defer fake.acquireMutex.RUnlock()
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLockDB) recordInvocation(key string, args []interface{}) {
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

var _ lock.LockDB = new(FakeLockDB)
