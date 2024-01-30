package restHandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	StudentId      int    `json:"studentid"`
	StudentName    string `json:"studentname"`
	StudentAddress string `json:"studentaddress"`
	StudentClass   int    `json:"studentclass"`
	StudentEmail   string `json:"studentemail"`
}

type AttendanceStudent struct {
	gorm.Model
	StudentId              int    `json:"studentid"`
	StudentAttendancedate  string `json:"studentattendancedate"`
	StudentAttendancemonth string `json:"studentattendancemonth"`
	StudentAttendanceyear  string `json:"studentattendanceyear"`

	StudentPresentStatus bool `json:"studentpresentstatus"`
	// StudentPunchIntime     bool `json:"studentintime"`
	// StudentPunchOuttime   bool `json:"studentouttime"`
}

// ********************************STUDENT************************************************

func InitialMigration() {

}

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var students []Student

	json.NewEncoder(w).Encode(students)

}

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var student Student

	json.NewEncoder(w).Encode(student)

}

func addStudentStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	json.NewEncoder(w).Encode(student)

}

func updateStudentStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var student Student

	json.NewDecoder(r.Body).Decode(&student)
	// save

	json.NewEncoder(w).Encode(student)

}

func deleteStudentStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var student Student
	// save

	json.NewEncoder(w).Encode("deleted successful")

}

// *****************************AttendanceStudent***********************************

func addStudentattendace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var studentattendace AttendanceStudent
	json.NewDecoder(r.Body).Decode(&studentattendace)
	json.NewEncoder(w).Encode(studentattendace)
}

func getStudentattendance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var studentattendace AttendanceStudent
	json.NewDecoder(r.Body).Decode(&studentattendace)
	json.NewEncoder(w).Encode(studentattendace)

}

func getClassattendance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var studentattendace AttendanceStudent
	json.NewDecoder(r.Body).Decode(&studentattendace)
	json.NewEncoder(w).Encode(studentattendace)

}
