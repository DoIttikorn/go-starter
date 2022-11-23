package ticket

import "testing"

// เป็นการ test โดยอาศัย Boundary values
func TestTicketPrice(t *testing.T) {
	t.Run("tiker $0 when age is 0", func(t *testing.T) {
		// Arrange
		want := 0.0
		const age uint = 0
		// Act
		result := Price(age)
		// Assert
		if result != want {
			t.Errorf("Price(0) = %f; want %f", result, want)
		}
	})

	t.Run("Ticket $0 when age is 3", func(t *testing.T) {
		// Arrange
		want := 0.0
		const age uint = 3
		// Act
		result := Price(age)
		// Assert
		if result != want {
			t.Errorf("Price(3) = %f; want %f", result, want)
		}
	})

	t.Run("Ticket $15 when age is 4", func(t *testing.T) {
		// Arrange
		want := 15.0
		const age uint = 4
		// Act
		result := Price(age)
		// Assert
		if result != want {
			t.Errorf("Price(4) = %f; want %f", result, want)
		}
	})

	t.Run("Ticket $15 when age is 15", func(t *testing.T) {
		// Arrange
		want := 15.0
		const age uint = 15
		// Act
		result := Price(age)
		// Assert
		if result != want {
			t.Errorf("Price(15) = %f; want %f", result, want)
		}
	})

	t.Run("Ticket $30 when age over 15", func(t *testing.T) {
		// Arrange
		want := 30.0
		const age uint = 16
		// Act
		result := Price(age)
		// Assert
		if result != want {
			t.Errorf("Price(16) = %f; want %f", result, want)
		}
	})

	t.Run("Ticket $30 when age over 30", func(t *testing.T) {
		// Arrange
		want := 30.0
		var age uint = 30
		// Act
		result := Price(age)
		// Assert
		if result != want {
			t.Errorf("Price(30) = %f; want %f", result, want)
		}
	})

	t.Run("Ticket $5 when age over 50", func(t *testing.T) {
		// Arrange
		want := 5.0
		var age uint = 51
		// Act
		result := Price(age)
		// Assert
		if result != want {
			t.Errorf("Price(51) = %f; want %f", result, want)
		}
	})
}
