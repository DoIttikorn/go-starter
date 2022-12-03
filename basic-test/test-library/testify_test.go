package testify_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}

func TestSomething(t *testing.T) {
	t.Run("not nil", func(t *testing.T) {
		pp := &Person{
			FirstName: "John",
		}

		if assert.NotNil(t, pp) {

			assert.Equal(t, "John", pp.FirstName)
		}
	})

	t.Run("nil", func(t *testing.T) {
		var p *Person

		assert.Nil(t, p)
	})

	t.Run("equal", func(t *testing.T) {
		want := 555
		got := 555

		assert.Equal(t, want, got)
	})

	t.Run("not equal", func(t *testing.T) {
		want := 555
		got := 666
		assert.NotEqual(t, want, got)
	})

}
