package models

type Student struct {
	Name  string
	Age   int
	Phone int
	Email string
	Id    int
}
type DateAndTime struct {
	Date string
	Time string
}

type ToDo struct {
	Id       int
	TaskName string
	DateTime DateAndTime
}
