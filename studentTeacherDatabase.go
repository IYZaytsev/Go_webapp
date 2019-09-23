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

var dataBaseStudent []student
var dataBaseClass []class
var dataBaseTeacher []teacher

func main() {

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
	//fmt.Println(dataBaseTeacher)
	//fmt.Println(dataBaseClass)
	commandLinePrompt()
	fmt.Println(dataBaseStudent)
}

func commandLinePrompt() {

	exit := 0

	for exit == 0 {
		fmt.Println("What would you like to do ?(enter -1 to exit)")
		fmt.Println("1.Edit Students")
		fmt.Println("2.Display Students")
		fmt.Println("3.Edit classes")
		fmt.Println("4.Display classes")
		fmt.Println("5.Edit teachers")
		fmt.Println("6.Display teachers")

		var response int
		fmt.Scanln(&response)
		if response == -1 {
			break
		}
		//since display operations on happen on even numbers
		if response%2 == 0 {

		} else {
			promptEdit(response)
		}
	}

}
func promptEdit(userInput int) {

	if userInput == 1 {
		subPromptEdit("student")
	}

	if userInput == 3 {
		subPromptEdit("teacher")
	}
	if userInput == 5 {
		subPromptEdit("class")
	}
}

func subPromptEdit(entity string) {
	var exit int = 0
	var response int
	for exit == 0 {
		fmt.Println("1. remove ", entity)
		fmt.Println("2. add ", entity)
		fmt.Println("enter -1 to go back ")
		fmt.Scanln(&response)
		if response == -1 {
			return
		}
		if response == 1 {
			fmt.Println("Enter the Id of the student you want to delete (-1 to go back)")

			fmt.Scanln(&response)
			if response == -1 {
				break
			}
			if entity == "student" {
				deleteStudent(response, &dataBaseStudent)
			}
			if entity == "teacher" {
				deleteTeacher(response, &dataBaseTeacher)
			}
			if entity == "class" {
				//delete(response, &dataBaseTeacher)
			}
			fmt.Println(entity, " removed, enter another or go back (-1)")
		} else {
			fmt.Println("Enter the Id of the ", entity, "you want to add (-1 to go back)")

			fmt.Scanln(&response)
			if response == -1 {
				break
			}

			if entity == "student" {
				addStudent(response, &dataBaseStudent)
			}
			if entity == "teacher" {
				//addTeacher(response, &dataBaseTeacher)
			}
			if entity == "class" {
				//delete(response, &dataBaseTeacher)
			}
			fmt.Println(entity, " added, enter another or go back (-1)")
		}
	}

}
func promptView() {

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
