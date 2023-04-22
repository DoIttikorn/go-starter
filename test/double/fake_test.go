package double

import "testing"

type FakeSearcher struct {
}

func (fs FakeSearcher) Search(people []*Person, firstName, lastName string) *Person {

	if len(people) == 0 {

		return nil
	}

	return people[0]
}

func TestFindCallsSearchAndReturnsEmptyStringForNoPerson(t *testing.T) {
	phoneBook := &PhoneBook{}

	phone, _ := phoneBook.Find(FakeSearcher{}, "John", "Doe")

	if phone != "" {
		t.Errorf("Expected empty string, got %s", phone)
	}

}
