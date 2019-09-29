package dataBaseTypes

import "fmt"

type Class struct {
	students  []Student
	id        int
	teacherID int
}

func AddClass(db *[]Class, students []Student, id int, teacherID int) {
	realValue := *db
	fmt.Println(realValue)
	*db = append(realValue, Class{students, id, teacherID})

	//fmt.Println(realValue)
}
