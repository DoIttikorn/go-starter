package double

import "testing"

type MockSearcher struct {
	phone         string
	methodsToCall map[string]bool
}

func (ms *MockSearcher) Search(people []*Person, firstName, lastName string) *Person {
	ms.methodsToCall["Search"] = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ms.phone,
	}
}

func (ms *MockSearcher) Create(people []*Person, firstName, lastName string) *Person {
	ms.methodsToCall["Create"] = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ms.phone,
	}
}

// คาดหวังว่า mock จะถูกเรียก method Search และ Create
func (ms *MockSearcher) ExpectToCall(method string) {
	if ms.methodsToCall == nil {
		ms.methodsToCall = make(map[string]bool)
	}

	ms.methodsToCall[method] = false
}

func (ms *MockSearcher) Verify(t *testing.T) {
	for method, called := range ms.methodsToCall {
		if !called {
			t.Errorf("Expected %s to be called", method)
		}
	}
}

func TestFindCallsSearchAndReturnPersonUsingMock(t *testing.T) {
	fakePhone := "555-555-5555"
	phoneBook := &PhoneBook{}
	mock := &MockSearcher{phone: fakePhone}
	mock.ExpectToCall("Search")

	phone, _ := phoneBook.Find(mock, "John", "Doe")

	if phone != fakePhone {
		t.Errorf("Expected %s, got %s", fakePhone, phone)
	}

	mock.Verify(t)
}
