package restHandler

import (
	"encoding/json"
	"strconv"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

// ********************************STUDENT************************************************

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dataBase.Connect()
	defer db.Close()
	var students []bean.Student
	if err := db.Model(&students).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(students)

}

// func connect() {
// 	panic("unimplemented")
// }

// func connect() {
// 	panic("unimplemented")
// }

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	db := dataBase.Connect()
	defer db.Close()

	// student_id := params["id"]
	student_id := params["id"]
	trr, err := strconv.Atoi(student_id)
	log.Println(err)

	students := &bean.Student{Id: trr}

	if err := db.Model(students).WherePK().Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(students)

}

func AddStudent(w http.ResponseWriter, r *http.Request) {
	// fmt.Print("hello2")
	w.Header().Set("Content-Type", "application/json")

	student := bean.Student{}
	_ = json.NewDecoder(r.Body).Decode(&student)
	db := dataBase.Connect()
	defer db.Close()
	// student.Id = uuid.New().String()
	if _, err := db.Model(&student).Insert(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(student)

}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	db := dataBase.Connect()
	defer db.Close()

	params := mux.Vars(r)

	student_id := params["id"]
	trr, err := strconv.Atoi(student_id)
	log.Println(err)
	students := &bean.Student{Id: trr}

	_ = json.NewDecoder(r.Body).Decode(&students)
	yy, err := db.Model(students).WherePK().Set("name= ?,address=?,class=?,email=?", students.Name, students.Address, students.Class, students.Email).Update()
	log.Println(yy)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(students)

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	db := dataBase.Connect()
	defer db.Close()

	student_id := params["id"]

	trr, err := strconv.Atoi(student_id)
	log.Println(err)

	students := &bean.Student{Id: trr}
	result, err := db.Model(students).WherePK().Delete()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)

}

// *****************************AttendanceStudent***********************************

// func addStudentattendace(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var studentattendace AttendanceStudent
// 	json.NewDecoder(r.Body).Decode(&studentattendace)
// 	json.NewEncoder(w).Encode(studentattendace)
// }

// func getStudentattendance(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var studentattendace AttendanceStudent
// 	json.NewDecoder(r.Body).Decode(&studentattendace)
// 	json.NewEncoder(w).Encode(studentattendace)

// }

// func getClassattendance(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var studentattendace AttendanceStudent
// 	json.NewDecoder(r.Body).Decode(&studentattendace)
// 	json.NewEncoder(w).Encode(studentattendace)

// }
