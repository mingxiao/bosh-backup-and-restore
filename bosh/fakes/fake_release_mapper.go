// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/bosh"
)

type FakeReleaseMapper struct {
	NewReleaseMappingStub        func(manifest string, instanceNames []string) bosh.ReleaseMapping
	newReleaseMappingMutex       sync.RWMutex
	newReleaseMappingArgsForCall []struct {
		manifest      string
		instanceNames []string
	}
	newReleaseMappingReturns struct {
		result1 bosh.ReleaseMapping
	}
	newReleaseMappingReturnsOnCall map[int]struct {
		result1 bosh.ReleaseMapping
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReleaseMapper) NewReleaseMapping(manifest string, instanceNames []string) bosh.ReleaseMapping {
	var instanceNamesCopy []string
	if instanceNames != nil {
		instanceNamesCopy = make([]string, len(instanceNames))
		copy(instanceNamesCopy, instanceNames)
	}
	fake.newReleaseMappingMutex.Lock()
	ret, specificReturn := fake.newReleaseMappingReturnsOnCall[len(fake.newReleaseMappingArgsForCall)]
	fake.newReleaseMappingArgsForCall = append(fake.newReleaseMappingArgsForCall, struct {
		manifest      string
		instanceNames []string
	}{manifest, instanceNamesCopy})
	fake.recordInvocation("NewReleaseMapping", []interface{}{manifest, instanceNamesCopy})
	fake.newReleaseMappingMutex.Unlock()
	if fake.NewReleaseMappingStub != nil {
		return fake.NewReleaseMappingStub(manifest, instanceNames)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newReleaseMappingReturns.result1
}

func (fake *FakeReleaseMapper) NewReleaseMappingCallCount() int {
	fake.newReleaseMappingMutex.RLock()
	defer fake.newReleaseMappingMutex.RUnlock()
	return len(fake.newReleaseMappingArgsForCall)
}

func (fake *FakeReleaseMapper) NewReleaseMappingArgsForCall(i int) (string, []string) {
	fake.newReleaseMappingMutex.RLock()
	defer fake.newReleaseMappingMutex.RUnlock()
	return fake.newReleaseMappingArgsForCall[i].manifest, fake.newReleaseMappingArgsForCall[i].instanceNames
}

func (fake *FakeReleaseMapper) NewReleaseMappingReturns(result1 bosh.ReleaseMapping) {
	fake.NewReleaseMappingStub = nil
	fake.newReleaseMappingReturns = struct {
		result1 bosh.ReleaseMapping
	}{result1}
}

func (fake *FakeReleaseMapper) NewReleaseMappingReturnsOnCall(i int, result1 bosh.ReleaseMapping) {
	fake.NewReleaseMappingStub = nil
	if fake.newReleaseMappingReturnsOnCall == nil {
		fake.newReleaseMappingReturnsOnCall = make(map[int]struct {
			result1 bosh.ReleaseMapping
		})
	}
	fake.newReleaseMappingReturnsOnCall[i] = struct {
		result1 bosh.ReleaseMapping
	}{result1}
}

func (fake *FakeReleaseMapper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newReleaseMappingMutex.RLock()
	defer fake.newReleaseMappingMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReleaseMapper) recordInvocation(key string, args []interface{}) {
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

var _ bosh.ReleaseMapper = new(FakeReleaseMapper)
