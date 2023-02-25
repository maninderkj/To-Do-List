package controller

import (
	"golang/models"
)

func InitializingAndAddingToArray(Students *[]models.Student) {
	student1 := models.Student{
		Name:  "shurty",
		Age:   20,
		Phone: 7896251778,
		Email: "shrutya@gmail.com",
		Id:    1,
	}
	student2 := models.Student{
		Name:  "shiny",
		Age:   21,
		Phone: 7896256678,
		Email: "shraaaaa@gmail.com",
		Id:    2,
	}
	student3 := models.Student{
		Name:  "ankiket",
		Age:   21,
		Phone: 7896256678,
		Email: "anki@gmail.com",
		Id:    3,
	}
	*Students = append(*Students, student1, student2, student3)
}

func InitializingToDos(Todos *[]models.ToDo) {
	todo := []models.ToDo{
		{Id: 1, TaskName: "have breakfast", DateTime: models.DateAndTime{Date: "9FEB", Time: "7:00 AM"}},
		{Id: 2, TaskName: "do prayer", DateTime: models.DateAndTime{Date: "9FEB", Time: "8:00 AM"}},
		{Id: 3, TaskName: "go college", DateTime: models.DateAndTime{Date: "9FEB", Time: "9:00 AM"}},
	}
	todo = append(todo, models.ToDo{Id: 4, TaskName: "do assignment", DateTime: models.DateAndTime{Date: "9FEB", Time: "12:00 AM"}})
	*Todos = todo

}
