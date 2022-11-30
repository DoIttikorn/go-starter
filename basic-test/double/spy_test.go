package double

import "testing"

type SpySearcher struct {
	phone           string
	searchWasCalled bool
	whatIsFirstName string
}

func (ss *SpySearcher) Search(people []*Person, firstName, lastName string) *Person {
	ss.searchWasCalled = true
	ss.whatIsFirstName = firstName
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

/*
SpySearcher จะเป็น struct ที่มี method Search ที่มี signature ตรงกับ method ใน interface Queryer ที่เราต้องการให้มัน implement

spy จะมีการ implement property เพิ่มเติมเพื่อให้เราสามารถทดสอบว่า spy นั้นถูกเรียกหรือไม่
spy จะมีการ implement property เพิ่มเติมเพื่อให้เราสามารถทดสอบว่าของที่รับมามีการเปลี่ยนแปลงหรือไม่
*/
func TestFindCallsSearchAndReturnsPerson(t *testing.T) {
	fakePhone := "555-555-5555"
	phoneBook := &PhoneBook{}

	spy := &SpySearcher{phone: fakePhone}

	phone, _ := phoneBook.Find(spy, "John", "Doe")

	if !spy.searchWasCalled {
		t.Error("Expected was not called")
	}

	if spy.whatIsFirstName != "John" {
		t.Errorf("Expected  was called with %s, want %s", spy.whatIsFirstName, "John")
	}

	if phone != fakePhone {
		t.Errorf("Expected %s, got %s", fakePhone, phone)
	}
}
