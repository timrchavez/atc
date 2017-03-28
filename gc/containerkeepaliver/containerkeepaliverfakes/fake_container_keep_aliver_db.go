// This file was generated by counterfeiter
package containerkeepaliverfakes

import (
	"sync"
	"time"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/gc/containerkeepaliver"
)

type FakeContainerKeepAliverDB struct {
	FindJobIDForBuildStub        func(buildID int) (int, bool, error)
	findJobIDForBuildMutex       sync.RWMutex
	findJobIDForBuildArgsForCall []struct {
		buildID int
	}
	findJobIDForBuildReturns struct {
		result1 int
		result2 bool
		result3 error
	}
	findJobIDForBuildReturnsOnCall map[int]struct {
		result1 int
		result2 bool
		result3 error
	}
	FindLatestSuccessfulBuildsPerJobStub        func() (map[int]int, error)
	findLatestSuccessfulBuildsPerJobMutex       sync.RWMutex
	findLatestSuccessfulBuildsPerJobArgsForCall []struct{}
	findLatestSuccessfulBuildsPerJobReturns     struct {
		result1 map[int]int
		result2 error
	}
	findLatestSuccessfulBuildsPerJobReturnsOnCall map[int]struct {
		result1 map[int]int
		result2 error
	}
	FindJobContainersFromUnsuccessfulBuildsStub        func() ([]db.SavedContainer, error)
	findJobContainersFromUnsuccessfulBuildsMutex       sync.RWMutex
	findJobContainersFromUnsuccessfulBuildsArgsForCall []struct{}
	findJobContainersFromUnsuccessfulBuildsReturns     struct {
		result1 []db.SavedContainer
		result2 error
	}
	findJobContainersFromUnsuccessfulBuildsReturnsOnCall map[int]struct {
		result1 []db.SavedContainer
		result2 error
	}
	UpdateExpiresAtOnContainerStub        func(handle string, ttl time.Duration) error
	updateExpiresAtOnContainerMutex       sync.RWMutex
	updateExpiresAtOnContainerArgsForCall []struct {
		handle string
		ttl    time.Duration
	}
	updateExpiresAtOnContainerReturns struct {
		result1 error
	}
	updateExpiresAtOnContainerReturnsOnCall map[int]struct {
		result1 error
	}
	GetAllPipelinesStub        func() ([]db.SavedPipeline, error)
	getAllPipelinesMutex       sync.RWMutex
	getAllPipelinesArgsForCall []struct{}
	getAllPipelinesReturns     struct {
		result1 []db.SavedPipeline
		result2 error
	}
	getAllPipelinesReturnsOnCall map[int]struct {
		result1 []db.SavedPipeline
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerKeepAliverDB) FindJobIDForBuild(buildID int) (int, bool, error) {
	fake.findJobIDForBuildMutex.Lock()
	ret, specificReturn := fake.findJobIDForBuildReturnsOnCall[len(fake.findJobIDForBuildArgsForCall)]
	fake.findJobIDForBuildArgsForCall = append(fake.findJobIDForBuildArgsForCall, struct {
		buildID int
	}{buildID})
	fake.recordInvocation("FindJobIDForBuild", []interface{}{buildID})
	fake.findJobIDForBuildMutex.Unlock()
	if fake.FindJobIDForBuildStub != nil {
		return fake.FindJobIDForBuildStub(buildID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findJobIDForBuildReturns.result1, fake.findJobIDForBuildReturns.result2, fake.findJobIDForBuildReturns.result3
}

func (fake *FakeContainerKeepAliverDB) FindJobIDForBuildCallCount() int {
	fake.findJobIDForBuildMutex.RLock()
	defer fake.findJobIDForBuildMutex.RUnlock()
	return len(fake.findJobIDForBuildArgsForCall)
}

func (fake *FakeContainerKeepAliverDB) FindJobIDForBuildArgsForCall(i int) int {
	fake.findJobIDForBuildMutex.RLock()
	defer fake.findJobIDForBuildMutex.RUnlock()
	return fake.findJobIDForBuildArgsForCall[i].buildID
}

func (fake *FakeContainerKeepAliverDB) FindJobIDForBuildReturns(result1 int, result2 bool, result3 error) {
	fake.FindJobIDForBuildStub = nil
	fake.findJobIDForBuildReturns = struct {
		result1 int
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeContainerKeepAliverDB) FindJobIDForBuildReturnsOnCall(i int, result1 int, result2 bool, result3 error) {
	fake.FindJobIDForBuildStub = nil
	if fake.findJobIDForBuildReturnsOnCall == nil {
		fake.findJobIDForBuildReturnsOnCall = make(map[int]struct {
			result1 int
			result2 bool
			result3 error
		})
	}
	fake.findJobIDForBuildReturnsOnCall[i] = struct {
		result1 int
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeContainerKeepAliverDB) FindLatestSuccessfulBuildsPerJob() (map[int]int, error) {
	fake.findLatestSuccessfulBuildsPerJobMutex.Lock()
	ret, specificReturn := fake.findLatestSuccessfulBuildsPerJobReturnsOnCall[len(fake.findLatestSuccessfulBuildsPerJobArgsForCall)]
	fake.findLatestSuccessfulBuildsPerJobArgsForCall = append(fake.findLatestSuccessfulBuildsPerJobArgsForCall, struct{}{})
	fake.recordInvocation("FindLatestSuccessfulBuildsPerJob", []interface{}{})
	fake.findLatestSuccessfulBuildsPerJobMutex.Unlock()
	if fake.FindLatestSuccessfulBuildsPerJobStub != nil {
		return fake.FindLatestSuccessfulBuildsPerJobStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findLatestSuccessfulBuildsPerJobReturns.result1, fake.findLatestSuccessfulBuildsPerJobReturns.result2
}

func (fake *FakeContainerKeepAliverDB) FindLatestSuccessfulBuildsPerJobCallCount() int {
	fake.findLatestSuccessfulBuildsPerJobMutex.RLock()
	defer fake.findLatestSuccessfulBuildsPerJobMutex.RUnlock()
	return len(fake.findLatestSuccessfulBuildsPerJobArgsForCall)
}

func (fake *FakeContainerKeepAliverDB) FindLatestSuccessfulBuildsPerJobReturns(result1 map[int]int, result2 error) {
	fake.FindLatestSuccessfulBuildsPerJobStub = nil
	fake.findLatestSuccessfulBuildsPerJobReturns = struct {
		result1 map[int]int
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerKeepAliverDB) FindLatestSuccessfulBuildsPerJobReturnsOnCall(i int, result1 map[int]int, result2 error) {
	fake.FindLatestSuccessfulBuildsPerJobStub = nil
	if fake.findLatestSuccessfulBuildsPerJobReturnsOnCall == nil {
		fake.findLatestSuccessfulBuildsPerJobReturnsOnCall = make(map[int]struct {
			result1 map[int]int
			result2 error
		})
	}
	fake.findLatestSuccessfulBuildsPerJobReturnsOnCall[i] = struct {
		result1 map[int]int
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerKeepAliverDB) FindJobContainersFromUnsuccessfulBuilds() ([]db.SavedContainer, error) {
	fake.findJobContainersFromUnsuccessfulBuildsMutex.Lock()
	ret, specificReturn := fake.findJobContainersFromUnsuccessfulBuildsReturnsOnCall[len(fake.findJobContainersFromUnsuccessfulBuildsArgsForCall)]
	fake.findJobContainersFromUnsuccessfulBuildsArgsForCall = append(fake.findJobContainersFromUnsuccessfulBuildsArgsForCall, struct{}{})
	fake.recordInvocation("FindJobContainersFromUnsuccessfulBuilds", []interface{}{})
	fake.findJobContainersFromUnsuccessfulBuildsMutex.Unlock()
	if fake.FindJobContainersFromUnsuccessfulBuildsStub != nil {
		return fake.FindJobContainersFromUnsuccessfulBuildsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findJobContainersFromUnsuccessfulBuildsReturns.result1, fake.findJobContainersFromUnsuccessfulBuildsReturns.result2
}

func (fake *FakeContainerKeepAliverDB) FindJobContainersFromUnsuccessfulBuildsCallCount() int {
	fake.findJobContainersFromUnsuccessfulBuildsMutex.RLock()
	defer fake.findJobContainersFromUnsuccessfulBuildsMutex.RUnlock()
	return len(fake.findJobContainersFromUnsuccessfulBuildsArgsForCall)
}

func (fake *FakeContainerKeepAliverDB) FindJobContainersFromUnsuccessfulBuildsReturns(result1 []db.SavedContainer, result2 error) {
	fake.FindJobContainersFromUnsuccessfulBuildsStub = nil
	fake.findJobContainersFromUnsuccessfulBuildsReturns = struct {
		result1 []db.SavedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerKeepAliverDB) FindJobContainersFromUnsuccessfulBuildsReturnsOnCall(i int, result1 []db.SavedContainer, result2 error) {
	fake.FindJobContainersFromUnsuccessfulBuildsStub = nil
	if fake.findJobContainersFromUnsuccessfulBuildsReturnsOnCall == nil {
		fake.findJobContainersFromUnsuccessfulBuildsReturnsOnCall = make(map[int]struct {
			result1 []db.SavedContainer
			result2 error
		})
	}
	fake.findJobContainersFromUnsuccessfulBuildsReturnsOnCall[i] = struct {
		result1 []db.SavedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerKeepAliverDB) UpdateExpiresAtOnContainer(handle string, ttl time.Duration) error {
	fake.updateExpiresAtOnContainerMutex.Lock()
	ret, specificReturn := fake.updateExpiresAtOnContainerReturnsOnCall[len(fake.updateExpiresAtOnContainerArgsForCall)]
	fake.updateExpiresAtOnContainerArgsForCall = append(fake.updateExpiresAtOnContainerArgsForCall, struct {
		handle string
		ttl    time.Duration
	}{handle, ttl})
	fake.recordInvocation("UpdateExpiresAtOnContainer", []interface{}{handle, ttl})
	fake.updateExpiresAtOnContainerMutex.Unlock()
	if fake.UpdateExpiresAtOnContainerStub != nil {
		return fake.UpdateExpiresAtOnContainerStub(handle, ttl)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateExpiresAtOnContainerReturns.result1
}

func (fake *FakeContainerKeepAliverDB) UpdateExpiresAtOnContainerCallCount() int {
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	return len(fake.updateExpiresAtOnContainerArgsForCall)
}

func (fake *FakeContainerKeepAliverDB) UpdateExpiresAtOnContainerArgsForCall(i int) (string, time.Duration) {
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	return fake.updateExpiresAtOnContainerArgsForCall[i].handle, fake.updateExpiresAtOnContainerArgsForCall[i].ttl
}

func (fake *FakeContainerKeepAliverDB) UpdateExpiresAtOnContainerReturns(result1 error) {
	fake.UpdateExpiresAtOnContainerStub = nil
	fake.updateExpiresAtOnContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerKeepAliverDB) UpdateExpiresAtOnContainerReturnsOnCall(i int, result1 error) {
	fake.UpdateExpiresAtOnContainerStub = nil
	if fake.updateExpiresAtOnContainerReturnsOnCall == nil {
		fake.updateExpiresAtOnContainerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateExpiresAtOnContainerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerKeepAliverDB) GetAllPipelines() ([]db.SavedPipeline, error) {
	fake.getAllPipelinesMutex.Lock()
	ret, specificReturn := fake.getAllPipelinesReturnsOnCall[len(fake.getAllPipelinesArgsForCall)]
	fake.getAllPipelinesArgsForCall = append(fake.getAllPipelinesArgsForCall, struct{}{})
	fake.recordInvocation("GetAllPipelines", []interface{}{})
	fake.getAllPipelinesMutex.Unlock()
	if fake.GetAllPipelinesStub != nil {
		return fake.GetAllPipelinesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getAllPipelinesReturns.result1, fake.getAllPipelinesReturns.result2
}

func (fake *FakeContainerKeepAliverDB) GetAllPipelinesCallCount() int {
	fake.getAllPipelinesMutex.RLock()
	defer fake.getAllPipelinesMutex.RUnlock()
	return len(fake.getAllPipelinesArgsForCall)
}

func (fake *FakeContainerKeepAliverDB) GetAllPipelinesReturns(result1 []db.SavedPipeline, result2 error) {
	fake.GetAllPipelinesStub = nil
	fake.getAllPipelinesReturns = struct {
		result1 []db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerKeepAliverDB) GetAllPipelinesReturnsOnCall(i int, result1 []db.SavedPipeline, result2 error) {
	fake.GetAllPipelinesStub = nil
	if fake.getAllPipelinesReturnsOnCall == nil {
		fake.getAllPipelinesReturnsOnCall = make(map[int]struct {
			result1 []db.SavedPipeline
			result2 error
		})
	}
	fake.getAllPipelinesReturnsOnCall[i] = struct {
		result1 []db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerKeepAliverDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findJobIDForBuildMutex.RLock()
	defer fake.findJobIDForBuildMutex.RUnlock()
	fake.findLatestSuccessfulBuildsPerJobMutex.RLock()
	defer fake.findLatestSuccessfulBuildsPerJobMutex.RUnlock()
	fake.findJobContainersFromUnsuccessfulBuildsMutex.RLock()
	defer fake.findJobContainersFromUnsuccessfulBuildsMutex.RUnlock()
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	fake.getAllPipelinesMutex.RLock()
	defer fake.getAllPipelinesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeContainerKeepAliverDB) recordInvocation(key string, args []interface{}) {
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

var _ containerkeepaliver.ContainerKeepAliverDB = new(FakeContainerKeepAliverDB)