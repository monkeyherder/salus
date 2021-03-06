// This file was generated by counterfeiter
package adaptorsfakes

import (
	"sync"

	"github.com/monkeyherder/salus/checks"
	"github.com/monkeyherder/salus/checks/adaptors"
)

type FakeNotifier struct {
	BeforeCheckStub        func(checks.Check)
	beforeCheckMutex       sync.RWMutex
	beforeCheckArgsForCall []struct {
		arg1 checks.Check
	}
	AfterCheckStub        func(checks.Check)
	afterCheckMutex       sync.RWMutex
	afterCheckArgsForCall []struct {
		arg1 checks.Check
	}
	OnErrorStub        func(checks.Check, error)
	onErrorMutex       sync.RWMutex
	onErrorArgsForCall []struct {
		arg1 checks.Check
		arg2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNotifier) BeforeCheck(arg1 checks.Check) {
	fake.beforeCheckMutex.Lock()
	fake.beforeCheckArgsForCall = append(fake.beforeCheckArgsForCall, struct {
		arg1 checks.Check
	}{arg1})
	fake.recordInvocation("BeforeCheck", []interface{}{arg1})
	fake.beforeCheckMutex.Unlock()
	if fake.BeforeCheckStub != nil {
		fake.BeforeCheckStub(arg1)
	}
}

func (fake *FakeNotifier) BeforeCheckCallCount() int {
	fake.beforeCheckMutex.RLock()
	defer fake.beforeCheckMutex.RUnlock()
	return len(fake.beforeCheckArgsForCall)
}

func (fake *FakeNotifier) BeforeCheckArgsForCall(i int) checks.Check {
	fake.beforeCheckMutex.RLock()
	defer fake.beforeCheckMutex.RUnlock()
	return fake.beforeCheckArgsForCall[i].arg1
}

func (fake *FakeNotifier) AfterCheck(arg1 checks.Check) {
	fake.afterCheckMutex.Lock()
	fake.afterCheckArgsForCall = append(fake.afterCheckArgsForCall, struct {
		arg1 checks.Check
	}{arg1})
	fake.recordInvocation("AfterCheck", []interface{}{arg1})
	fake.afterCheckMutex.Unlock()
	if fake.AfterCheckStub != nil {
		fake.AfterCheckStub(arg1)
	}
}

func (fake *FakeNotifier) AfterCheckCallCount() int {
	fake.afterCheckMutex.RLock()
	defer fake.afterCheckMutex.RUnlock()
	return len(fake.afterCheckArgsForCall)
}

func (fake *FakeNotifier) AfterCheckArgsForCall(i int) checks.Check {
	fake.afterCheckMutex.RLock()
	defer fake.afterCheckMutex.RUnlock()
	return fake.afterCheckArgsForCall[i].arg1
}

func (fake *FakeNotifier) OnError(arg1 checks.Check, arg2 error) {
	fake.onErrorMutex.Lock()
	fake.onErrorArgsForCall = append(fake.onErrorArgsForCall, struct {
		arg1 checks.Check
		arg2 error
	}{arg1, arg2})
	fake.recordInvocation("OnError", []interface{}{arg1, arg2})
	fake.onErrorMutex.Unlock()
	if fake.OnErrorStub != nil {
		fake.OnErrorStub(arg1, arg2)
	}
}

func (fake *FakeNotifier) OnErrorCallCount() int {
	fake.onErrorMutex.RLock()
	defer fake.onErrorMutex.RUnlock()
	return len(fake.onErrorArgsForCall)
}

func (fake *FakeNotifier) OnErrorArgsForCall(i int) (checks.Check, error) {
	fake.onErrorMutex.RLock()
	defer fake.onErrorMutex.RUnlock()
	return fake.onErrorArgsForCall[i].arg1, fake.onErrorArgsForCall[i].arg2
}

func (fake *FakeNotifier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.beforeCheckMutex.RLock()
	defer fake.beforeCheckMutex.RUnlock()
	fake.afterCheckMutex.RLock()
	defer fake.afterCheckMutex.RUnlock()
	fake.onErrorMutex.RLock()
	defer fake.onErrorMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeNotifier) recordInvocation(key string, args []interface{}) {
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

var _ adaptors.Notifier = new(FakeNotifier)
