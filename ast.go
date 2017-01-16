package main

type verb struct {
	Name string
	Args []interface{}
	Body []*verb
}

type variable struct {
	Key  string
	Path []interface{}
	Args []interface{}
}
