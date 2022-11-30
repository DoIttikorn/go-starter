package ticket

import "testing"

// เป็นการ test โดยอาศัย Boundary values
// การเขียน test แบบ table test
func TestTicketPrice(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want float64
	}{
		{name: "Free Ticket when age is 0", age: 0, want: 0.0},
		{name: "Free Ticket when age under 3", age: 3, want: 0.0},
		{name: "Ticket $15 when age at 4 year old", age: 4, want: 15.0},
		{name: "Ticket $15 when age is 15", age: 15, want: 15.0},
		{name: "Ticket $30 when age over 15", age: 16, want: 30.0},
		{name: "Ticket $30 when age is 50", age: 50, want: 30.0},
		{name: "Ticket $5 when age over 50", age: 51, want: 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Price(tt.age)

			if got != tt.want {
				t.Errorf("Price(%d) = %f, want %f", tt.age, got, tt.want)
			}

		})
	}
}
