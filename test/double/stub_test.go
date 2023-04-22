package double

import "testing"

type StubSearcher struct {
	phone string
}

// เป็นทำให้แน่ใจว่า ถ้าเราเรียก method Search จะได้ค่า phone ที่เรากำหนดไว้
func (ss StubSearcher) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

func TestReturnPerson(t *testing.T) {
	fakePhone := "1234567890"
	phoneBook := &PhoneBook{}

	ss := &StubSearcher{phone: fakePhone}
	phone, _ := phoneBook.Find(ss, "Jane", "Doe")

	if phone != fakePhone {
		t.Errorf("Want %s, got %s", fakePhone, phone)
	}

}
