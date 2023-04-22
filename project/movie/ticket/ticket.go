package ticket

import (
	"fmt"
)

func init() {
	fmt.Println("init ticket")
}

func BuyTicket(name string) {
	fmt.Printf("I bought a ticket for %s\n", name)

}
