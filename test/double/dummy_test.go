package double

import "testing"

type DummySearcher struct {
}

// สร้างขึ้นมาเพื่อเป็น dummy ให้ compile ผ่าน เหมือน mock แต่ไม่มีการทำงาน
func (ds DummySearcher) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{}
}

func TestFindsItShouldReturnsErrorWhenFirstNameOrLastEmpty(t *testing.T) {
	phoneBook := &PhoneBook{}
	want := ErrMissingArg

	dd := DummySearcher{}

	_, got := phoneBook.Find(dd, "", "")

	if got != want {
		t.Errorf("Want %s, got %s", want, got)
	}
}
