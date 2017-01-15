package main

type verb struct {
	Name string
	Args []interface{}
	Body []interface{}
}

type variable struct {
	Key  string
	Path []interface{}
	Args []interface{}
}

type comment struct {
	Value string
}
