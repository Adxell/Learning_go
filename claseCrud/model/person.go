package model 

type Community struct {
	Name string
}

type Communities []Community

type Person struct {
	Name string
	Age uint8

	Communities Communities
}
type Persons []Person