package main

import "fmt"

type student struct {
	id int
}

type class struct {
	students  []student
	id        int
	teacherID int
}

type teacher struct {
	id      int
	classes []class
}

func main() {
	dataBaseStudent := make([]student, 0)
	dataBaseClass := make([]class, 0)
	dataBaseTeacher := make([]teacher, 0)
	//fmt.Println(dataBaseStudent)
	//fmt.Println(dataBaseClass)
	//fmt.Println(dataBaseTeacher)

	fmt.Println("adding to student")
	addStudent(29, &dataBaseStudent)
	fmt.Println("student added", dataBaseStudent)
	//deleteStudent(29, &dataBaseStudent)
	//fmt.Println("student deleted", dataBaseStudent)
	addClass(&dataBaseClass, dataBaseStudent, 1, 2)
	addTeacher(2, dataBaseClass, &dataBaseTeacher)
	fmt.Println(dataBaseTeacher)
	fmt.Println(dataBaseClass)
}

func addStudent(id int, db *[]student) {
	realValue := *db
	*db = append(realValue, student{id})
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

func addTeacher(teacherID int, classes []class, db *[]teacher) {
	realValue := *db
	*db = append(realValue, teacher{teacherID, classes})
	//fmt.Println(realValue)
}

func deleteTeacher(taggedforRemoval int, db *[]teacher) {
	realValue := *db
	tempSlice := make([]teacher, 0)
	for i := 0; i < len(realValue); i++ {

		if realValue[i].id == taggedforRemoval {
			continue
		}
		tempSlice = append(tempSlice, teacher{realValue[i].id, realValue[i].classes})
	}
	*db = tempSlice
}
func addClass(db *[]class, students []student, id int, teacherID int) {
	realValue := *db
	fmt.Println(realValue)
	*db = append(realValue, class{students, id, teacherID})

	//fmt.Println(realValue)
}
