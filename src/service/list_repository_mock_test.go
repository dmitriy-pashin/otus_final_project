package service

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i otus_final_project/src/repository.ListRepository -o ./src/service/list_repository_mock_test.go

import (
	"net"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ListRepositoryMock implements repository.ListRepository
type ListRepositoryMock struct {
	t minimock.Tester

	funcAdd          func(ip *net.IPNet) (b1 bool, err error)
	inspectFuncAdd   func(ip *net.IPNet)
	afterAddCounter  uint64
	beforeAddCounter uint64
	AddMock          mListRepositoryMockAdd

	funcDelete          func(ip *net.IPNet) (b1 bool, err error)
	inspectFuncDelete   func(ip *net.IPNet)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mListRepositoryMockDelete

	funcIsInList          func(ip net.IP) (b1 bool)
	inspectFuncIsInList   func(ip net.IP)
	afterIsInListCounter  uint64
	beforeIsInListCounter uint64
	IsInListMock          mListRepositoryMockIsInList
}

// NewListRepositoryMock returns a mock for repository.ListRepository
func NewListRepositoryMock(t minimock.Tester) *ListRepositoryMock {
	m := &ListRepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddMock = mListRepositoryMockAdd{mock: m}
	m.AddMock.callArgs = []*ListRepositoryMockAddParams{}

	m.DeleteMock = mListRepositoryMockDelete{mock: m}
	m.DeleteMock.callArgs = []*ListRepositoryMockDeleteParams{}

	m.IsInListMock = mListRepositoryMockIsInList{mock: m}
	m.IsInListMock.callArgs = []*ListRepositoryMockIsInListParams{}

	return m
}

type mListRepositoryMockAdd struct {
	mock               *ListRepositoryMock
	defaultExpectation *ListRepositoryMockAddExpectation
	expectations       []*ListRepositoryMockAddExpectation

	callArgs []*ListRepositoryMockAddParams
	mutex    sync.RWMutex
}

// ListRepositoryMockAddExpectation specifies expectation struct of the ListRepository.Add
type ListRepositoryMockAddExpectation struct {
	mock    *ListRepositoryMock
	params  *ListRepositoryMockAddParams
	results *ListRepositoryMockAddResults
	Counter uint64
}

// ListRepositoryMockAddParams contains parameters of the ListRepository.Add
type ListRepositoryMockAddParams struct {
	ip *net.IPNet
}

// ListRepositoryMockAddResults contains results of the ListRepository.Add
type ListRepositoryMockAddResults struct {
	b1  bool
	err error
}

// Expect sets up expected params for ListRepository.Add
func (mmAdd *mListRepositoryMockAdd) Expect(ip *net.IPNet) *mListRepositoryMockAdd {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("ListRepositoryMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &ListRepositoryMockAddExpectation{}
	}

	mmAdd.defaultExpectation.params = &ListRepositoryMockAddParams{ip}
	for _, e := range mmAdd.expectations {
		if minimock.Equal(e.params, mmAdd.defaultExpectation.params) {
			mmAdd.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAdd.defaultExpectation.params)
		}
	}

	return mmAdd
}

// Inspect accepts an inspector function that has same arguments as the ListRepository.Add
func (mmAdd *mListRepositoryMockAdd) Inspect(f func(ip *net.IPNet)) *mListRepositoryMockAdd {
	if mmAdd.mock.inspectFuncAdd != nil {
		mmAdd.mock.t.Fatalf("Inspect function is already set for ListRepositoryMock.Add")
	}

	mmAdd.mock.inspectFuncAdd = f

	return mmAdd
}

// Return sets up results that will be returned by ListRepository.Add
func (mmAdd *mListRepositoryMockAdd) Return(b1 bool, err error) *ListRepositoryMock {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("ListRepositoryMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &ListRepositoryMockAddExpectation{mock: mmAdd.mock}
	}
	mmAdd.defaultExpectation.results = &ListRepositoryMockAddResults{b1, err}
	return mmAdd.mock
}

//Set uses given function f to mock the ListRepository.Add method
func (mmAdd *mListRepositoryMockAdd) Set(f func(ip *net.IPNet) (b1 bool, err error)) *ListRepositoryMock {
	if mmAdd.defaultExpectation != nil {
		mmAdd.mock.t.Fatalf("Default expectation is already set for the ListRepository.Add method")
	}

	if len(mmAdd.expectations) > 0 {
		mmAdd.mock.t.Fatalf("Some expectations are already set for the ListRepository.Add method")
	}

	mmAdd.mock.funcAdd = f
	return mmAdd.mock
}

// When sets expectation for the ListRepository.Add which will trigger the result defined by the following
// Then helper
func (mmAdd *mListRepositoryMockAdd) When(ip *net.IPNet) *ListRepositoryMockAddExpectation {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("ListRepositoryMock.Add mock is already set by Set")
	}

	expectation := &ListRepositoryMockAddExpectation{
		mock:   mmAdd.mock,
		params: &ListRepositoryMockAddParams{ip},
	}
	mmAdd.expectations = append(mmAdd.expectations, expectation)
	return expectation
}

// Then sets up ListRepository.Add return parameters for the expectation previously defined by the When method
func (e *ListRepositoryMockAddExpectation) Then(b1 bool, err error) *ListRepositoryMock {
	e.results = &ListRepositoryMockAddResults{b1, err}
	return e.mock
}

// Add implements repository.ListRepository
func (mmAdd *ListRepositoryMock) Add(ip *net.IPNet) (b1 bool, err error) {
	mm_atomic.AddUint64(&mmAdd.beforeAddCounter, 1)
	defer mm_atomic.AddUint64(&mmAdd.afterAddCounter, 1)

	if mmAdd.inspectFuncAdd != nil {
		mmAdd.inspectFuncAdd(ip)
	}

	mm_params := &ListRepositoryMockAddParams{ip}

	// Record call args
	mmAdd.AddMock.mutex.Lock()
	mmAdd.AddMock.callArgs = append(mmAdd.AddMock.callArgs, mm_params)
	mmAdd.AddMock.mutex.Unlock()

	for _, e := range mmAdd.AddMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if mmAdd.AddMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAdd.AddMock.defaultExpectation.Counter, 1)
		mm_want := mmAdd.AddMock.defaultExpectation.params
		mm_got := ListRepositoryMockAddParams{ip}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAdd.t.Errorf("ListRepositoryMock.Add got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAdd.AddMock.defaultExpectation.results
		if mm_results == nil {
			mmAdd.t.Fatal("No results are set for the ListRepositoryMock.Add")
		}
		return (*mm_results).b1, (*mm_results).err
	}
	if mmAdd.funcAdd != nil {
		return mmAdd.funcAdd(ip)
	}
	mmAdd.t.Fatalf("Unexpected call to ListRepositoryMock.Add. %v", ip)
	return
}

// AddAfterCounter returns a count of finished ListRepositoryMock.Add invocations
func (mmAdd *ListRepositoryMock) AddAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAdd.afterAddCounter)
}

// AddBeforeCounter returns a count of ListRepositoryMock.Add invocations
func (mmAdd *ListRepositoryMock) AddBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAdd.beforeAddCounter)
}

// Calls returns a list of arguments used in each call to ListRepositoryMock.Add.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAdd *mListRepositoryMockAdd) Calls() []*ListRepositoryMockAddParams {
	mmAdd.mutex.RLock()

	argCopy := make([]*ListRepositoryMockAddParams, len(mmAdd.callArgs))
	copy(argCopy, mmAdd.callArgs)

	mmAdd.mutex.RUnlock()

	return argCopy
}

// MinimockAddDone returns true if the count of the Add invocations corresponds
// the number of defined expectations
func (m *ListRepositoryMock) MinimockAddDone() bool {
	for _, e := range m.AddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAdd != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddInspect logs each unmet expectation
func (m *ListRepositoryMock) MinimockAddInspect() {
	for _, e := range m.AddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ListRepositoryMock.Add with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		if m.AddMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ListRepositoryMock.Add")
		} else {
			m.t.Errorf("Expected call to ListRepositoryMock.Add with params: %#v", *m.AddMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAdd != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		m.t.Error("Expected call to ListRepositoryMock.Add")
	}
}

type mListRepositoryMockDelete struct {
	mock               *ListRepositoryMock
	defaultExpectation *ListRepositoryMockDeleteExpectation
	expectations       []*ListRepositoryMockDeleteExpectation

	callArgs []*ListRepositoryMockDeleteParams
	mutex    sync.RWMutex
}

// ListRepositoryMockDeleteExpectation specifies expectation struct of the ListRepository.Delete
type ListRepositoryMockDeleteExpectation struct {
	mock    *ListRepositoryMock
	params  *ListRepositoryMockDeleteParams
	results *ListRepositoryMockDeleteResults
	Counter uint64
}

// ListRepositoryMockDeleteParams contains parameters of the ListRepository.Delete
type ListRepositoryMockDeleteParams struct {
	ip *net.IPNet
}

// ListRepositoryMockDeleteResults contains results of the ListRepository.Delete
type ListRepositoryMockDeleteResults struct {
	b1  bool
	err error
}

// Expect sets up expected params for ListRepository.Delete
func (mmDelete *mListRepositoryMockDelete) Expect(ip *net.IPNet) *mListRepositoryMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ListRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ListRepositoryMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &ListRepositoryMockDeleteParams{ip}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the ListRepository.Delete
func (mmDelete *mListRepositoryMockDelete) Inspect(f func(ip *net.IPNet)) *mListRepositoryMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for ListRepositoryMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by ListRepository.Delete
func (mmDelete *mListRepositoryMockDelete) Return(b1 bool, err error) *ListRepositoryMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ListRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ListRepositoryMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &ListRepositoryMockDeleteResults{b1, err}
	return mmDelete.mock
}

//Set uses given function f to mock the ListRepository.Delete method
func (mmDelete *mListRepositoryMockDelete) Set(f func(ip *net.IPNet) (b1 bool, err error)) *ListRepositoryMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the ListRepository.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the ListRepository.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the ListRepository.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mListRepositoryMockDelete) When(ip *net.IPNet) *ListRepositoryMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ListRepositoryMock.Delete mock is already set by Set")
	}

	expectation := &ListRepositoryMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &ListRepositoryMockDeleteParams{ip},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up ListRepository.Delete return parameters for the expectation previously defined by the When method
func (e *ListRepositoryMockDeleteExpectation) Then(b1 bool, err error) *ListRepositoryMock {
	e.results = &ListRepositoryMockDeleteResults{b1, err}
	return e.mock
}

// Delete implements repository.ListRepository
func (mmDelete *ListRepositoryMock) Delete(ip *net.IPNet) (b1 bool, err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ip)
	}

	mm_params := &ListRepositoryMockDeleteParams{ip}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := ListRepositoryMockDeleteParams{ip}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("ListRepositoryMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the ListRepositoryMock.Delete")
		}
		return (*mm_results).b1, (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ip)
	}
	mmDelete.t.Fatalf("Unexpected call to ListRepositoryMock.Delete. %v", ip)
	return
}

// DeleteAfterCounter returns a count of finished ListRepositoryMock.Delete invocations
func (mmDelete *ListRepositoryMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of ListRepositoryMock.Delete invocations
func (mmDelete *ListRepositoryMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to ListRepositoryMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mListRepositoryMockDelete) Calls() []*ListRepositoryMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*ListRepositoryMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *ListRepositoryMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *ListRepositoryMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ListRepositoryMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ListRepositoryMock.Delete")
		} else {
			m.t.Errorf("Expected call to ListRepositoryMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to ListRepositoryMock.Delete")
	}
}

type mListRepositoryMockIsInList struct {
	mock               *ListRepositoryMock
	defaultExpectation *ListRepositoryMockIsInListExpectation
	expectations       []*ListRepositoryMockIsInListExpectation

	callArgs []*ListRepositoryMockIsInListParams
	mutex    sync.RWMutex
}

// ListRepositoryMockIsInListExpectation specifies expectation struct of the ListRepository.IsInList
type ListRepositoryMockIsInListExpectation struct {
	mock    *ListRepositoryMock
	params  *ListRepositoryMockIsInListParams
	results *ListRepositoryMockIsInListResults
	Counter uint64
}

// ListRepositoryMockIsInListParams contains parameters of the ListRepository.IsInList
type ListRepositoryMockIsInListParams struct {
	ip net.IP
}

// ListRepositoryMockIsInListResults contains results of the ListRepository.IsInList
type ListRepositoryMockIsInListResults struct {
	b1 bool
}

// Expect sets up expected params for ListRepository.IsInList
func (mmIsInList *mListRepositoryMockIsInList) Expect(ip net.IP) *mListRepositoryMockIsInList {
	if mmIsInList.mock.funcIsInList != nil {
		mmIsInList.mock.t.Fatalf("ListRepositoryMock.IsInList mock is already set by Set")
	}

	if mmIsInList.defaultExpectation == nil {
		mmIsInList.defaultExpectation = &ListRepositoryMockIsInListExpectation{}
	}

	mmIsInList.defaultExpectation.params = &ListRepositoryMockIsInListParams{ip}
	for _, e := range mmIsInList.expectations {
		if minimock.Equal(e.params, mmIsInList.defaultExpectation.params) {
			mmIsInList.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmIsInList.defaultExpectation.params)
		}
	}

	return mmIsInList
}

// Inspect accepts an inspector function that has same arguments as the ListRepository.IsInList
func (mmIsInList *mListRepositoryMockIsInList) Inspect(f func(ip net.IP)) *mListRepositoryMockIsInList {
	if mmIsInList.mock.inspectFuncIsInList != nil {
		mmIsInList.mock.t.Fatalf("Inspect function is already set for ListRepositoryMock.IsInList")
	}

	mmIsInList.mock.inspectFuncIsInList = f

	return mmIsInList
}

// Return sets up results that will be returned by ListRepository.IsInList
func (mmIsInList *mListRepositoryMockIsInList) Return(b1 bool) *ListRepositoryMock {
	if mmIsInList.mock.funcIsInList != nil {
		mmIsInList.mock.t.Fatalf("ListRepositoryMock.IsInList mock is already set by Set")
	}

	if mmIsInList.defaultExpectation == nil {
		mmIsInList.defaultExpectation = &ListRepositoryMockIsInListExpectation{mock: mmIsInList.mock}
	}
	mmIsInList.defaultExpectation.results = &ListRepositoryMockIsInListResults{b1}
	return mmIsInList.mock
}

//Set uses given function f to mock the ListRepository.IsInList method
func (mmIsInList *mListRepositoryMockIsInList) Set(f func(ip net.IP) (b1 bool)) *ListRepositoryMock {
	if mmIsInList.defaultExpectation != nil {
		mmIsInList.mock.t.Fatalf("Default expectation is already set for the ListRepository.IsInList method")
	}

	if len(mmIsInList.expectations) > 0 {
		mmIsInList.mock.t.Fatalf("Some expectations are already set for the ListRepository.IsInList method")
	}

	mmIsInList.mock.funcIsInList = f
	return mmIsInList.mock
}

// When sets expectation for the ListRepository.IsInList which will trigger the result defined by the following
// Then helper
func (mmIsInList *mListRepositoryMockIsInList) When(ip net.IP) *ListRepositoryMockIsInListExpectation {
	if mmIsInList.mock.funcIsInList != nil {
		mmIsInList.mock.t.Fatalf("ListRepositoryMock.IsInList mock is already set by Set")
	}

	expectation := &ListRepositoryMockIsInListExpectation{
		mock:   mmIsInList.mock,
		params: &ListRepositoryMockIsInListParams{ip},
	}
	mmIsInList.expectations = append(mmIsInList.expectations, expectation)
	return expectation
}

// Then sets up ListRepository.IsInList return parameters for the expectation previously defined by the When method
func (e *ListRepositoryMockIsInListExpectation) Then(b1 bool) *ListRepositoryMock {
	e.results = &ListRepositoryMockIsInListResults{b1}
	return e.mock
}

// IsInList implements repository.ListRepository
func (mmIsInList *ListRepositoryMock) IsInList(ip net.IP) (b1 bool) {
	mm_atomic.AddUint64(&mmIsInList.beforeIsInListCounter, 1)
	defer mm_atomic.AddUint64(&mmIsInList.afterIsInListCounter, 1)

	if mmIsInList.inspectFuncIsInList != nil {
		mmIsInList.inspectFuncIsInList(ip)
	}

	mm_params := &ListRepositoryMockIsInListParams{ip}

	// Record call args
	mmIsInList.IsInListMock.mutex.Lock()
	mmIsInList.IsInListMock.callArgs = append(mmIsInList.IsInListMock.callArgs, mm_params)
	mmIsInList.IsInListMock.mutex.Unlock()

	for _, e := range mmIsInList.IsInListMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1
		}
	}

	if mmIsInList.IsInListMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmIsInList.IsInListMock.defaultExpectation.Counter, 1)
		mm_want := mmIsInList.IsInListMock.defaultExpectation.params
		mm_got := ListRepositoryMockIsInListParams{ip}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmIsInList.t.Errorf("ListRepositoryMock.IsInList got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmIsInList.IsInListMock.defaultExpectation.results
		if mm_results == nil {
			mmIsInList.t.Fatal("No results are set for the ListRepositoryMock.IsInList")
		}
		return (*mm_results).b1
	}
	if mmIsInList.funcIsInList != nil {
		return mmIsInList.funcIsInList(ip)
	}
	mmIsInList.t.Fatalf("Unexpected call to ListRepositoryMock.IsInList. %v", ip)
	return
}

// IsInListAfterCounter returns a count of finished ListRepositoryMock.IsInList invocations
func (mmIsInList *ListRepositoryMock) IsInListAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIsInList.afterIsInListCounter)
}

// IsInListBeforeCounter returns a count of ListRepositoryMock.IsInList invocations
func (mmIsInList *ListRepositoryMock) IsInListBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIsInList.beforeIsInListCounter)
}

// Calls returns a list of arguments used in each call to ListRepositoryMock.IsInList.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmIsInList *mListRepositoryMockIsInList) Calls() []*ListRepositoryMockIsInListParams {
	mmIsInList.mutex.RLock()

	argCopy := make([]*ListRepositoryMockIsInListParams, len(mmIsInList.callArgs))
	copy(argCopy, mmIsInList.callArgs)

	mmIsInList.mutex.RUnlock()

	return argCopy
}

// MinimockIsInListDone returns true if the count of the IsInList invocations corresponds
// the number of defined expectations
func (m *ListRepositoryMock) MinimockIsInListDone() bool {
	for _, e := range m.IsInListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.IsInListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterIsInListCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIsInList != nil && mm_atomic.LoadUint64(&m.afterIsInListCounter) < 1 {
		return false
	}
	return true
}

// MinimockIsInListInspect logs each unmet expectation
func (m *ListRepositoryMock) MinimockIsInListInspect() {
	for _, e := range m.IsInListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ListRepositoryMock.IsInList with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.IsInListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterIsInListCounter) < 1 {
		if m.IsInListMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ListRepositoryMock.IsInList")
		} else {
			m.t.Errorf("Expected call to ListRepositoryMock.IsInList with params: %#v", *m.IsInListMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIsInList != nil && mm_atomic.LoadUint64(&m.afterIsInListCounter) < 1 {
		m.t.Error("Expected call to ListRepositoryMock.IsInList")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ListRepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAddInspect()

		m.MinimockDeleteInspect()

		m.MinimockIsInListInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ListRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ListRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddDone() &&
		m.MinimockDeleteDone() &&
		m.MinimockIsInListDone()
}
