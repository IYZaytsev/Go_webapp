package main

import (
	"fmt"

	"./dataBaseTypes"
)

//Creating slices of stu
var dataBaseStudent []dataBaseTypes.Student
var dataBaseClass []dataBaseTypes.Class
var dataBaseTeacher []dataBaseTypes.Teacher

func main() {

	//fmt.Println(dataBaseStudent)
	//fmt.Println(dataBaseClass)
	//fmt.Println(dataBaseTeacher)

	fmt.Println("adding to student")
	dataBaseTypes.AddStudent(29, &dataBaseStudent)
	fmt.Println("student added", dataBaseStudent)
	//deleteStudent(29, &dataBaseStudent)
	//fmt.Println("student deleted", dataBaseStudent)
	dataBaseTypes.AddClass(&dataBaseClass, dataBaseStudent, 1, 2)
	dataBaseTypes.AddTeacher(2, dataBaseClass, &dataBaseTeacher)
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
				dataBaseTypes.DeleteStudent(response, &dataBaseStudent)
			}
			if entity == "teacher" {
				dataBaseTypes.DeleteTeacher(response, &dataBaseTeacher)
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
				dataBaseTypes.AddStudent(response, &dataBaseStudent)
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
