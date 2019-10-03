package dataBaseTypes

type Class struct {
	Students  []Student
	Id        int
	TeacherID int
	Name      string
}

func AddClass(classId int, className string, students []Student, teacherID int, db *[]Class) {
	realValue := *db

	*db = append(realValue, Class{students, classId, teacherID, className})

	//fmt.Println(realValue)
}
