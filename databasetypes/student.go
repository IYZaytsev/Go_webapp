package databasetypes

//Student used to store information about the student
type Student struct {
	ID        int
	ClassList Classes
	Name      string
}

//Students is here to ensure proper parsing of Json
type Students []Student
