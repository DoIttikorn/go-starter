package sum

import "testing"

func TestSum(t *testing.T) {

	t.Run("should return 3 when input 1 and input 2", func(t *testing.T) {
		// Arrange
		want := 3
		// Act
		result := Sum(1, 2)
		// Assert
		if result != 3 {
			t.Errorf("Sum(1, 2) = %d; want %d", result, want)
		}
	})

	t.Run("should return 1 when 1 and 0", func(t *testing.T) {
		// Arrange
		want := 1
		// Act
		result := Sum(1, 0)
		// Assert
		if result != 1 {
			t.Errorf("Sum(1, 0) = %d; want %d", result, want)
		}
	})

	t.Run("should return -2 when -1 and -1", func(t *testing.T) {
		// Arrange
		want := -2
		// Act
		result := Sum(-1, -1)
		// Assert
		if result != -2 {
			t.Errorf("Sum(-1, -1) = %d; want %d", result, want)
		}
	})

}
