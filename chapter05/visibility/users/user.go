package users

type address struct {
	Region string
	street string
	nu     string
}

type User struct {
	ID int
	Name string
	birthday string
	addr address
}