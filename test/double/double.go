package double

import "errors"

var (
	ErrMissingArg    = errors.New("FirstName and LastName are mandatory arguments")
	ErrNoPersonFound = errors.New("no person found")
)

type Queryer interface { // interface จะ map กับ struct ที่มี method ที่ตรงกัน
	Search(people []*Person, firstName string, lastName string) *Person
}

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}

type PhoneBook struct {
	People []*Person
}

func (p *PhoneBook) Find(query Queryer, firstName, lastName string) (string, error) {
	if firstName == "" || lastName == "" {
		return "", ErrMissingArg
	}

	person := query.Search(p.People, firstName, lastName)
	if person == nil {
		return "", ErrNoPersonFound
	}

	return person.Phone, nil
}
