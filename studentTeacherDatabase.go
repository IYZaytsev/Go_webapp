package main

import "fmt"

type student struct {
	id int
}

type class struct {
	students []student
}

type teacher struct {
	id      int
	classes []class
}

func main() {
	dataBaseStudent := make([]student, 0)
	//dataBaseClass := make([]class, 0)
	//dataBaseTeacher := make([]teacher, 0)
	//fmt.Println(dataBaseStudent)
	//fmt.Println(dataBaseClass)
	//fmt.Println(dataBaseTeacher)

	fmt.Println("adding to student")
	addStudent(29, &dataBaseStudent)
	fmt.Println("student added", dataBaseStudent)
	deleteStudent(29, &dataBaseStudent)
	fmt.Println("student deleted", dataBaseStudent)
}

func addStudent(id int, db *[]student) {
	realValue := *db
	fmt.Println(realValue)
	*db = append(realValue, student{id})

	//fmt.Println(realValue)
}
func deleteStudent(taggedforRemoval int, db *[]student) {
	realValue := *db
	tempSlice := make([]student, 0)
	for i := 0; i < len(realValue); i++ {

		if realValue[i].id == taggedforRemoval {
			continue
		}
		tempSlice = append(tempSlice, student{realValue[i].id})
	}
	*db = tempSlice
}
func addTeacher(id int, db *[]student) {
	realValue := *db
	fmt.Println(realValue)
	*db = append(realValue, student{id})

	//fmt.Println(realValue)
}

func addClass(id int, db *[]student) {
	realValue := *db
	fmt.Println(realValue)
	*db = append(realValue, student{id})

	//fmt.Println(realValue)
}
