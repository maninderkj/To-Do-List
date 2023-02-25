package controller

import (
	"encoding/json"
	"fmt"
	"golang/models"
	"net/http"
	"os"
	"strconv"
)

func HandlingTodos(w http.ResponseWriter, r *http.Request) {
	var Todos []models.ToDo

	InitializingToDos(&Todos)
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(Todos)
	}
	if r.Method == http.MethodPost {
		var anotherTodo models.ToDo
		json.NewDecoder(r.Body).Decode(&anotherTodo)

		Todos = append(Todos, anotherTodo)
		json.NewEncoder(w).Encode(Todos)

	}
	if r.Method == http.MethodDelete {
		query := r.URL.Query()
		fmt.Println(query)
	}
}

func HandlingHttpReq(w http.ResponseWriter, r *http.Request) {
	var Students []models.Student
	InitializingAndAddingToArray(&Students)

	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(Students)
	}
	if r.Method == http.MethodPost {
		//create another struct
		var anotherStudent models.Student
		//decode result from body
		json.NewDecoder(r.Body).Decode(&anotherStudent)
		Students = append(Students, anotherStudent)
		//send response
		json.NewEncoder(w).Encode(Students)
	}
	if r.Method == http.MethodDelete {
		query := r.URL.Query()
		//convert string to int
		id, err := strconv.Atoi(query.Get("id"))
		if err != nil {
			fmt.Print("error occured", err)
			os.Exit(1)

		}
		for index, student := range Students {
			if id == student.Id {
				Students = append(Students[:index], Students[index+1:]...)
			}
		}
		json.NewEncoder(w).Encode(Students)

	}
}
