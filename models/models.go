package models

type Todo struct {
	Id        int
	Item      string
	Completed int
}
type View struct {
	Todos []Todo
}
