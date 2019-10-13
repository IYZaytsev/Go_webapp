package databasetypes

//Teacher used to store data about teacher in database
type Teacher struct {
	ID        int
	ClassList Classes
	Name      string
}

//Teachers is here to ensure proper parsing of JSON
type Teachers []Teacher
