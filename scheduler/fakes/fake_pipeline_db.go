// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/db/algorithm"
	"github.com/concourse/atc/scheduler"
)

type FakePipelineDB struct {
	CreateJobBuildStub        func(job string) (db.Build, error)
	createJobBuildMutex       sync.RWMutex
	createJobBuildArgsForCall []struct {
		job string
	}
	createJobBuildReturns struct {
		result1 db.Build
		result2 error
	}
	CreateJobBuildForCandidateInputsStub        func(job string) (db.Build, bool, error)
	createJobBuildForCandidateInputsMutex       sync.RWMutex
	createJobBuildForCandidateInputsArgsForCall []struct {
		job string
	}
	createJobBuildForCandidateInputsReturns struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	ScheduleBuildStub        func(buildID int, jobConfig atc.JobConfig) (bool, error)
	scheduleBuildMutex       sync.RWMutex
	scheduleBuildArgsForCall []struct {
		buildID   int
		jobConfig atc.JobConfig
	}
	scheduleBuildReturns struct {
		result1 bool
		result2 error
	}
	GetJobBuildForInputsStub        func(job string, inputs []db.BuildInput) (db.Build, error)
	getJobBuildForInputsMutex       sync.RWMutex
	getJobBuildForInputsArgsForCall []struct {
		job    string
		inputs []db.BuildInput
	}
	getJobBuildForInputsReturns struct {
		result1 db.Build
		result2 error
	}
	GetNextPendingBuildStub        func(job string) (db.Build, error)
	getNextPendingBuildMutex       sync.RWMutex
	getNextPendingBuildArgsForCall []struct {
		job string
	}
	getNextPendingBuildReturns struct {
		result1 db.Build
		result2 error
	}
	LoadVersionsDBStub        func() (algorithm.VersionsDB, error)
	loadVersionsDBMutex       sync.RWMutex
	loadVersionsDBArgsForCall []struct{}
	loadVersionsDBReturns struct {
		result1 algorithm.VersionsDB
		result2 error
	}
	GetLatestInputVersionsStub        func(versions algorithm.VersionsDB, job string, inputs []atc.JobInput) ([]db.BuildInput, error)
	getLatestInputVersionsMutex       sync.RWMutex
	getLatestInputVersionsArgsForCall []struct {
		versions algorithm.VersionsDB
		job      string
		inputs   []atc.JobInput
	}
	getLatestInputVersionsReturns struct {
		result1 []db.BuildInput
		result2 error
	}
	SaveResourceVersionsStub        func(atc.ResourceConfig, []atc.Version) error
	saveResourceVersionsMutex       sync.RWMutex
	saveResourceVersionsArgsForCall []struct {
		arg1 atc.ResourceConfig
		arg2 []atc.Version
	}
	saveResourceVersionsReturns struct {
		result1 error
	}
	UseInputsForBuildStub        func(buildID int, inputs []db.BuildInput) error
	useInputsForBuildMutex       sync.RWMutex
	useInputsForBuildArgsForCall []struct {
		buildID int
		inputs  []db.BuildInput
	}
	useInputsForBuildReturns struct {
		result1 error
	}
}

func (fake *FakePipelineDB) CreateJobBuild(job string) (db.Build, error) {
	fake.createJobBuildMutex.Lock()
	fake.createJobBuildArgsForCall = append(fake.createJobBuildArgsForCall, struct {
		job string
	}{job})
	fake.createJobBuildMutex.Unlock()
	if fake.CreateJobBuildStub != nil {
		return fake.CreateJobBuildStub(job)
	} else {
		return fake.createJobBuildReturns.result1, fake.createJobBuildReturns.result2
	}
}

func (fake *FakePipelineDB) CreateJobBuildCallCount() int {
	fake.createJobBuildMutex.RLock()
	defer fake.createJobBuildMutex.RUnlock()
	return len(fake.createJobBuildArgsForCall)
}

func (fake *FakePipelineDB) CreateJobBuildArgsForCall(i int) string {
	fake.createJobBuildMutex.RLock()
	defer fake.createJobBuildMutex.RUnlock()
	return fake.createJobBuildArgsForCall[i].job
}

func (fake *FakePipelineDB) CreateJobBuildReturns(result1 db.Build, result2 error) {
	fake.CreateJobBuildStub = nil
	fake.createJobBuildReturns = struct {
		result1 db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineDB) CreateJobBuildForCandidateInputs(job string) (db.Build, bool, error) {
	fake.createJobBuildForCandidateInputsMutex.Lock()
	fake.createJobBuildForCandidateInputsArgsForCall = append(fake.createJobBuildForCandidateInputsArgsForCall, struct {
		job string
	}{job})
	fake.createJobBuildForCandidateInputsMutex.Unlock()
	if fake.CreateJobBuildForCandidateInputsStub != nil {
		return fake.CreateJobBuildForCandidateInputsStub(job)
	} else {
		return fake.createJobBuildForCandidateInputsReturns.result1, fake.createJobBuildForCandidateInputsReturns.result2, fake.createJobBuildForCandidateInputsReturns.result3
	}
}

func (fake *FakePipelineDB) CreateJobBuildForCandidateInputsCallCount() int {
	fake.createJobBuildForCandidateInputsMutex.RLock()
	defer fake.createJobBuildForCandidateInputsMutex.RUnlock()
	return len(fake.createJobBuildForCandidateInputsArgsForCall)
}

func (fake *FakePipelineDB) CreateJobBuildForCandidateInputsArgsForCall(i int) string {
	fake.createJobBuildForCandidateInputsMutex.RLock()
	defer fake.createJobBuildForCandidateInputsMutex.RUnlock()
	return fake.createJobBuildForCandidateInputsArgsForCall[i].job
}

func (fake *FakePipelineDB) CreateJobBuildForCandidateInputsReturns(result1 db.Build, result2 bool, result3 error) {
	fake.CreateJobBuildForCandidateInputsStub = nil
	fake.createJobBuildForCandidateInputsReturns = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePipelineDB) ScheduleBuild(buildID int, jobConfig atc.JobConfig) (bool, error) {
	fake.scheduleBuildMutex.Lock()
	fake.scheduleBuildArgsForCall = append(fake.scheduleBuildArgsForCall, struct {
		buildID   int
		jobConfig atc.JobConfig
	}{buildID, jobConfig})
	fake.scheduleBuildMutex.Unlock()
	if fake.ScheduleBuildStub != nil {
		return fake.ScheduleBuildStub(buildID, jobConfig)
	} else {
		return fake.scheduleBuildReturns.result1, fake.scheduleBuildReturns.result2
	}
}

func (fake *FakePipelineDB) ScheduleBuildCallCount() int {
	fake.scheduleBuildMutex.RLock()
	defer fake.scheduleBuildMutex.RUnlock()
	return len(fake.scheduleBuildArgsForCall)
}

func (fake *FakePipelineDB) ScheduleBuildArgsForCall(i int) (int, atc.JobConfig) {
	fake.scheduleBuildMutex.RLock()
	defer fake.scheduleBuildMutex.RUnlock()
	return fake.scheduleBuildArgsForCall[i].buildID, fake.scheduleBuildArgsForCall[i].jobConfig
}

func (fake *FakePipelineDB) ScheduleBuildReturns(result1 bool, result2 error) {
	fake.ScheduleBuildStub = nil
	fake.scheduleBuildReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineDB) GetJobBuildForInputs(job string, inputs []db.BuildInput) (db.Build, error) {
	fake.getJobBuildForInputsMutex.Lock()
	fake.getJobBuildForInputsArgsForCall = append(fake.getJobBuildForInputsArgsForCall, struct {
		job    string
		inputs []db.BuildInput
	}{job, inputs})
	fake.getJobBuildForInputsMutex.Unlock()
	if fake.GetJobBuildForInputsStub != nil {
		return fake.GetJobBuildForInputsStub(job, inputs)
	} else {
		return fake.getJobBuildForInputsReturns.result1, fake.getJobBuildForInputsReturns.result2
	}
}

func (fake *FakePipelineDB) GetJobBuildForInputsCallCount() int {
	fake.getJobBuildForInputsMutex.RLock()
	defer fake.getJobBuildForInputsMutex.RUnlock()
	return len(fake.getJobBuildForInputsArgsForCall)
}

func (fake *FakePipelineDB) GetJobBuildForInputsArgsForCall(i int) (string, []db.BuildInput) {
	fake.getJobBuildForInputsMutex.RLock()
	defer fake.getJobBuildForInputsMutex.RUnlock()
	return fake.getJobBuildForInputsArgsForCall[i].job, fake.getJobBuildForInputsArgsForCall[i].inputs
}

func (fake *FakePipelineDB) GetJobBuildForInputsReturns(result1 db.Build, result2 error) {
	fake.GetJobBuildForInputsStub = nil
	fake.getJobBuildForInputsReturns = struct {
		result1 db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineDB) GetNextPendingBuild(job string) (db.Build, error) {
	fake.getNextPendingBuildMutex.Lock()
	fake.getNextPendingBuildArgsForCall = append(fake.getNextPendingBuildArgsForCall, struct {
		job string
	}{job})
	fake.getNextPendingBuildMutex.Unlock()
	if fake.GetNextPendingBuildStub != nil {
		return fake.GetNextPendingBuildStub(job)
	} else {
		return fake.getNextPendingBuildReturns.result1, fake.getNextPendingBuildReturns.result2
	}
}

func (fake *FakePipelineDB) GetNextPendingBuildCallCount() int {
	fake.getNextPendingBuildMutex.RLock()
	defer fake.getNextPendingBuildMutex.RUnlock()
	return len(fake.getNextPendingBuildArgsForCall)
}

func (fake *FakePipelineDB) GetNextPendingBuildArgsForCall(i int) string {
	fake.getNextPendingBuildMutex.RLock()
	defer fake.getNextPendingBuildMutex.RUnlock()
	return fake.getNextPendingBuildArgsForCall[i].job
}

func (fake *FakePipelineDB) GetNextPendingBuildReturns(result1 db.Build, result2 error) {
	fake.GetNextPendingBuildStub = nil
	fake.getNextPendingBuildReturns = struct {
		result1 db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineDB) LoadVersionsDB() (algorithm.VersionsDB, error) {
	fake.loadVersionsDBMutex.Lock()
	fake.loadVersionsDBArgsForCall = append(fake.loadVersionsDBArgsForCall, struct{}{})
	fake.loadVersionsDBMutex.Unlock()
	if fake.LoadVersionsDBStub != nil {
		return fake.LoadVersionsDBStub()
	} else {
		return fake.loadVersionsDBReturns.result1, fake.loadVersionsDBReturns.result2
	}
}

func (fake *FakePipelineDB) LoadVersionsDBCallCount() int {
	fake.loadVersionsDBMutex.RLock()
	defer fake.loadVersionsDBMutex.RUnlock()
	return len(fake.loadVersionsDBArgsForCall)
}

func (fake *FakePipelineDB) LoadVersionsDBReturns(result1 algorithm.VersionsDB, result2 error) {
	fake.LoadVersionsDBStub = nil
	fake.loadVersionsDBReturns = struct {
		result1 algorithm.VersionsDB
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineDB) GetLatestInputVersions(versions algorithm.VersionsDB, job string, inputs []atc.JobInput) ([]db.BuildInput, error) {
	fake.getLatestInputVersionsMutex.Lock()
	fake.getLatestInputVersionsArgsForCall = append(fake.getLatestInputVersionsArgsForCall, struct {
		versions algorithm.VersionsDB
		job      string
		inputs   []atc.JobInput
	}{versions, job, inputs})
	fake.getLatestInputVersionsMutex.Unlock()
	if fake.GetLatestInputVersionsStub != nil {
		return fake.GetLatestInputVersionsStub(versions, job, inputs)
	} else {
		return fake.getLatestInputVersionsReturns.result1, fake.getLatestInputVersionsReturns.result2
	}
}

func (fake *FakePipelineDB) GetLatestInputVersionsCallCount() int {
	fake.getLatestInputVersionsMutex.RLock()
	defer fake.getLatestInputVersionsMutex.RUnlock()
	return len(fake.getLatestInputVersionsArgsForCall)
}

func (fake *FakePipelineDB) GetLatestInputVersionsArgsForCall(i int) (algorithm.VersionsDB, string, []atc.JobInput) {
	fake.getLatestInputVersionsMutex.RLock()
	defer fake.getLatestInputVersionsMutex.RUnlock()
	return fake.getLatestInputVersionsArgsForCall[i].versions, fake.getLatestInputVersionsArgsForCall[i].job, fake.getLatestInputVersionsArgsForCall[i].inputs
}

func (fake *FakePipelineDB) GetLatestInputVersionsReturns(result1 []db.BuildInput, result2 error) {
	fake.GetLatestInputVersionsStub = nil
	fake.getLatestInputVersionsReturns = struct {
		result1 []db.BuildInput
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineDB) SaveResourceVersions(arg1 atc.ResourceConfig, arg2 []atc.Version) error {
	fake.saveResourceVersionsMutex.Lock()
	fake.saveResourceVersionsArgsForCall = append(fake.saveResourceVersionsArgsForCall, struct {
		arg1 atc.ResourceConfig
		arg2 []atc.Version
	}{arg1, arg2})
	fake.saveResourceVersionsMutex.Unlock()
	if fake.SaveResourceVersionsStub != nil {
		return fake.SaveResourceVersionsStub(arg1, arg2)
	} else {
		return fake.saveResourceVersionsReturns.result1
	}
}

func (fake *FakePipelineDB) SaveResourceVersionsCallCount() int {
	fake.saveResourceVersionsMutex.RLock()
	defer fake.saveResourceVersionsMutex.RUnlock()
	return len(fake.saveResourceVersionsArgsForCall)
}

func (fake *FakePipelineDB) SaveResourceVersionsArgsForCall(i int) (atc.ResourceConfig, []atc.Version) {
	fake.saveResourceVersionsMutex.RLock()
	defer fake.saveResourceVersionsMutex.RUnlock()
	return fake.saveResourceVersionsArgsForCall[i].arg1, fake.saveResourceVersionsArgsForCall[i].arg2
}

func (fake *FakePipelineDB) SaveResourceVersionsReturns(result1 error) {
	fake.SaveResourceVersionsStub = nil
	fake.saveResourceVersionsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePipelineDB) UseInputsForBuild(buildID int, inputs []db.BuildInput) error {
	fake.useInputsForBuildMutex.Lock()
	fake.useInputsForBuildArgsForCall = append(fake.useInputsForBuildArgsForCall, struct {
		buildID int
		inputs  []db.BuildInput
	}{buildID, inputs})
	fake.useInputsForBuildMutex.Unlock()
	if fake.UseInputsForBuildStub != nil {
		return fake.UseInputsForBuildStub(buildID, inputs)
	} else {
		return fake.useInputsForBuildReturns.result1
	}
}

func (fake *FakePipelineDB) UseInputsForBuildCallCount() int {
	fake.useInputsForBuildMutex.RLock()
	defer fake.useInputsForBuildMutex.RUnlock()
	return len(fake.useInputsForBuildArgsForCall)
}

func (fake *FakePipelineDB) UseInputsForBuildArgsForCall(i int) (int, []db.BuildInput) {
	fake.useInputsForBuildMutex.RLock()
	defer fake.useInputsForBuildMutex.RUnlock()
	return fake.useInputsForBuildArgsForCall[i].buildID, fake.useInputsForBuildArgsForCall[i].inputs
}

func (fake *FakePipelineDB) UseInputsForBuildReturns(result1 error) {
	fake.UseInputsForBuildStub = nil
	fake.useInputsForBuildReturns = struct {
		result1 error
	}{result1}
}

var _ scheduler.PipelineDB = new(FakePipelineDB)
