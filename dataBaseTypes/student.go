package dataBaseTypes

type Student struct {
	Id   int
	Name string
}

//add student functino, needs a name and id value;
func AddStudent(id int, name string, db *[]Student) {
	realValue := *db
	*db = append(realValue, Student{id, name})
}

//Function used to delete students, students are selected for deletion by passsing
//an id.
func DeleteStudent(taggedforRemoval int, db *[]Student) {
	realValue := *db
	tempSlice := make([]Student, 0)
	for i := 0; i < len(realValue); i++ {

		if realValue[i].Id == taggedforRemoval {
			continue
		}
		tempSlice = append(tempSlice, Student{realValue[i].Id, realValue[i].Name})
	}
	*db = tempSlice
}

/*
func searchStudentDatabase(query string) results {
	var returnResults results

	//searches the student database for it
	for i := 0; i < len(dataBaseStudent); i++ {
		if dataBaseStudent[i].Name == query {
			returnResults.name = dataBaseStudent[i].Name
			returnResults.id = dataBaseStudent[i].Id
		}

	}
	for i := 0; i < len(dataBaseClass); i++ {
		for z := 0; z < len(dataBaseClass[i].Students); z++ {
			if dataBaseClass[i].Students[z].Name == query {

				returnResults.classes = append(returnResults.classes, dataBaseClass[i])

			}
		}
	}
	return returnResults
}
*/
