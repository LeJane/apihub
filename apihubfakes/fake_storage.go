// This file was generated by counterfeiter
package apihubfakes

import (
	"sync"

	"github.com/apihub/apihub"
)

type FakeStorage struct {
	AddServiceStub        func(apihub.ServiceSpec) error
	addServiceMutex       sync.RWMutex
	addServiceArgsForCall []struct {
		arg1 apihub.ServiceSpec
	}
	addServiceReturns struct {
		result1 error
	}
	UpdateServiceStub        func(apihub.ServiceSpec) error
	updateServiceMutex       sync.RWMutex
	updateServiceArgsForCall []struct {
		arg1 apihub.ServiceSpec
	}
	updateServiceReturns struct {
		result1 error
	}
	FindServiceByHostStub        func(string) (apihub.ServiceSpec, error)
	findServiceByHostMutex       sync.RWMutex
	findServiceByHostArgsForCall []struct {
		arg1 string
	}
	findServiceByHostReturns struct {
		result1 apihub.ServiceSpec
		result2 error
	}
	ServicesStub        func() ([]apihub.ServiceSpec, error)
	servicesMutex       sync.RWMutex
	servicesArgsForCall []struct{}
	servicesReturns     struct {
		result1 []apihub.ServiceSpec
		result2 error
	}
	RemoveServiceStub        func(string) error
	removeServiceMutex       sync.RWMutex
	removeServiceArgsForCall []struct {
		arg1 string
	}
	removeServiceReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStorage) AddService(arg1 apihub.ServiceSpec) error {
	fake.addServiceMutex.Lock()
	fake.addServiceArgsForCall = append(fake.addServiceArgsForCall, struct {
		arg1 apihub.ServiceSpec
	}{arg1})
	fake.recordInvocation("AddService", []interface{}{arg1})
	fake.addServiceMutex.Unlock()
	if fake.AddServiceStub != nil {
		return fake.AddServiceStub(arg1)
	} else {
		return fake.addServiceReturns.result1
	}
}

func (fake *FakeStorage) AddServiceCallCount() int {
	fake.addServiceMutex.RLock()
	defer fake.addServiceMutex.RUnlock()
	return len(fake.addServiceArgsForCall)
}

func (fake *FakeStorage) AddServiceArgsForCall(i int) apihub.ServiceSpec {
	fake.addServiceMutex.RLock()
	defer fake.addServiceMutex.RUnlock()
	return fake.addServiceArgsForCall[i].arg1
}

func (fake *FakeStorage) AddServiceReturns(result1 error) {
	fake.AddServiceStub = nil
	fake.addServiceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) UpdateService(arg1 apihub.ServiceSpec) error {
	fake.updateServiceMutex.Lock()
	fake.updateServiceArgsForCall = append(fake.updateServiceArgsForCall, struct {
		arg1 apihub.ServiceSpec
	}{arg1})
	fake.recordInvocation("UpdateService", []interface{}{arg1})
	fake.updateServiceMutex.Unlock()
	if fake.UpdateServiceStub != nil {
		return fake.UpdateServiceStub(arg1)
	} else {
		return fake.updateServiceReturns.result1
	}
}

func (fake *FakeStorage) UpdateServiceCallCount() int {
	fake.updateServiceMutex.RLock()
	defer fake.updateServiceMutex.RUnlock()
	return len(fake.updateServiceArgsForCall)
}

func (fake *FakeStorage) UpdateServiceArgsForCall(i int) apihub.ServiceSpec {
	fake.updateServiceMutex.RLock()
	defer fake.updateServiceMutex.RUnlock()
	return fake.updateServiceArgsForCall[i].arg1
}

func (fake *FakeStorage) UpdateServiceReturns(result1 error) {
	fake.UpdateServiceStub = nil
	fake.updateServiceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) FindServiceByHost(arg1 string) (apihub.ServiceSpec, error) {
	fake.findServiceByHostMutex.Lock()
	fake.findServiceByHostArgsForCall = append(fake.findServiceByHostArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindServiceByHost", []interface{}{arg1})
	fake.findServiceByHostMutex.Unlock()
	if fake.FindServiceByHostStub != nil {
		return fake.FindServiceByHostStub(arg1)
	} else {
		return fake.findServiceByHostReturns.result1, fake.findServiceByHostReturns.result2
	}
}

func (fake *FakeStorage) FindServiceByHostCallCount() int {
	fake.findServiceByHostMutex.RLock()
	defer fake.findServiceByHostMutex.RUnlock()
	return len(fake.findServiceByHostArgsForCall)
}

func (fake *FakeStorage) FindServiceByHostArgsForCall(i int) string {
	fake.findServiceByHostMutex.RLock()
	defer fake.findServiceByHostMutex.RUnlock()
	return fake.findServiceByHostArgsForCall[i].arg1
}

func (fake *FakeStorage) FindServiceByHostReturns(result1 apihub.ServiceSpec, result2 error) {
	fake.FindServiceByHostStub = nil
	fake.findServiceByHostReturns = struct {
		result1 apihub.ServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) Services() ([]apihub.ServiceSpec, error) {
	fake.servicesMutex.Lock()
	fake.servicesArgsForCall = append(fake.servicesArgsForCall, struct{}{})
	fake.recordInvocation("Services", []interface{}{})
	fake.servicesMutex.Unlock()
	if fake.ServicesStub != nil {
		return fake.ServicesStub()
	} else {
		return fake.servicesReturns.result1, fake.servicesReturns.result2
	}
}

func (fake *FakeStorage) ServicesCallCount() int {
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	return len(fake.servicesArgsForCall)
}

func (fake *FakeStorage) ServicesReturns(result1 []apihub.ServiceSpec, result2 error) {
	fake.ServicesStub = nil
	fake.servicesReturns = struct {
		result1 []apihub.ServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) RemoveService(arg1 string) error {
	fake.removeServiceMutex.Lock()
	fake.removeServiceArgsForCall = append(fake.removeServiceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RemoveService", []interface{}{arg1})
	fake.removeServiceMutex.Unlock()
	if fake.RemoveServiceStub != nil {
		return fake.RemoveServiceStub(arg1)
	} else {
		return fake.removeServiceReturns.result1
	}
}

func (fake *FakeStorage) RemoveServiceCallCount() int {
	fake.removeServiceMutex.RLock()
	defer fake.removeServiceMutex.RUnlock()
	return len(fake.removeServiceArgsForCall)
}

func (fake *FakeStorage) RemoveServiceArgsForCall(i int) string {
	fake.removeServiceMutex.RLock()
	defer fake.removeServiceMutex.RUnlock()
	return fake.removeServiceArgsForCall[i].arg1
}

func (fake *FakeStorage) RemoveServiceReturns(result1 error) {
	fake.RemoveServiceStub = nil
	fake.removeServiceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addServiceMutex.RLock()
	defer fake.addServiceMutex.RUnlock()
	fake.updateServiceMutex.RLock()
	defer fake.updateServiceMutex.RUnlock()
	fake.findServiceByHostMutex.RLock()
	defer fake.findServiceByHostMutex.RUnlock()
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	fake.removeServiceMutex.RLock()
	defer fake.removeServiceMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStorage) recordInvocation(key string, args []interface{}) {
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

var _ apihub.Storage = new(FakeStorage)