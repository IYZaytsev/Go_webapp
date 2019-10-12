package databasetypes

//Student used to store information about the student
type Student struct {
	ID        int
	ClassList Classes
	Name      string
}
type Students []Student

//AddStudent needs a name and id value;
func AddStudent(id int, name string, db *[]Student) {
	realValue := *db
	*db = append(realValue, Student{id, nil, name})
}

//DeleteStudent Function used to delete students, students are selected for deletion by passsing
//an id.
func DeleteStudent(taggedforRemoval int, db *[]Student) {
	realValue := *db
	tempSlice := make([]Student, 0)
	for i := 0; i < len(realValue); i++ {

		if realValue[i].ID == taggedforRemoval {
			continue
		}
		tempSlice = append(tempSlice, Student{realValue[i].ID, nil, realValue[i].Name})
	}
	*db = tempSlice
}
