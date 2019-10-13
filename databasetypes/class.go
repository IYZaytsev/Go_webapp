package databasetypes

//Class used to store data about class
type Class struct {
	ID              int
	StudentList     Students
	AssignedTeacher Teacher
	Name            string
}

//Classes is here to ensure proper parsing of Json
type Classes []Class
