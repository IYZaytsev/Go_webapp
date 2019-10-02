package dataBaseTypes

type Teacher struct {
	id      int
	classes []Class
	name    string
}

func AddTeacher(teacherID int, classes []Class, name string, db *[]Teacher) {
	realValue := *db
	*db = append(realValue, Teacher{teacherID, classes, name})
	//fmt.Println(realValue)
}

func DeleteTeacher(taggedforRemoval int, db *[]Teacher) {
	realValue := *db
	tempSlice := make([]Teacher, 0)
	for i := 0; i < len(realValue); i++ {

		if realValue[i].id == taggedforRemoval {
			continue
		}
		tempSlice = append(tempSlice, Teacher{realValue[i].id, realValue[i].classes, realValue[i].name})
	}
	*db = tempSlice
}
