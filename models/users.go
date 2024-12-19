package models

type User struct {
	ID       int    `"json":"id"`
	username string `"json":"username"`
	password string `"json":"password"`
}

type Crendential struct {
	username string `"json":"username"`
	password string `"json":"password"`
}
