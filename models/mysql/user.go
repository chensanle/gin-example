package mysql

type User struct {
	Uid      int
	Birthday int
}

func NewEmptyUser() *User {
	return &User{}
}

func (u *User) Create() error {
	return nil
}

func (u *User) Get() (*User, error) {
	return u, nil
}
