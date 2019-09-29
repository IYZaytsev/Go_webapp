package dataBaseTypes

type Teacher struct {
	id      int
	classes []Class
}

func AddTeacher(teacherID int, classes []Class, db *[]Teacher) {
	realValue := *db
	*db = append(realValue, Teacher{teacherID, classes})
	//fmt.Println(realValue)
}

func DeleteTeacher(taggedforRemoval int, db *[]Teacher) {
	realValue := *db
	tempSlice := make([]Teacher, 0)
	for i := 0; i < len(realValue); i++ {

		if realValue[i].id == taggedforRemoval {
			continue
		}
		tempSlice = append(tempSlice, Teacher{realValue[i].id, realValue[i].classes})
	}
	*db = tempSlice
}
