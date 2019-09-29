package dataBaseTypes

type Student struct {
	id int
}

func AddStudent(id int, db *[]Student) {
	realValue := *db
	*db = append(realValue, Student{id})
}

func DeleteStudent(taggedforRemoval int, db *[]Student) {
	realValue := *db
	tempSlice := make([]Student, 0)
	for i := 0; i < len(realValue); i++ {

		if realValue[i].id == taggedforRemoval {
			continue
		}
		tempSlice = append(tempSlice, Student{realValue[i].id})
	}
	*db = tempSlice
}
