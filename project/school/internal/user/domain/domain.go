package user

type (
	User struct {
		ID        int
		FirstName string
		LastName  string
		Email     string
		Password  string
	}

	UserInfo struct {
		ID        int
		FirstName string
		LastName  string
		Email     string
	}
)
