package databasetypes

//Class used to store data about class
type Class struct {
	Students  []Student
	ID        int
	TeacherID int
	Name      string
}
type Classes []Class

//AddClass used to add class to database
func AddClass(classID int, className string, students []Student, teacherID int, db *[]Class) {
	realValue := *db

	*db = append(realValue, Class{students, classID, teacherID, className})

	//fmt.Println(realValue)
}
