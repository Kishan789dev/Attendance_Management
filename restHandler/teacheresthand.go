package restHandler

// import (
// 	"encoding/json"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/mux"
// 	"gorm.io/gorm"
// )

// func InitialMigration() {

// }

// func GetTeachers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var teacher []Teacher

// 	json.NewEncoder(w).Encode(teacher)
// 	return
// }

// func getTeacher(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)

// 	var teacher Teacher

// 	json.NewEncoder(w).Encode(teacher)

// }

// func addTeacher(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var teacher Teacher
// 	json.NewDecoder(r.Body).Decode(&teacher)
// 	// id := teacher.TeacherId
// 	// newEmail := teacher.TeacherEmail
// 	json.NewEncoder(w).Encode(teacher)

// }

// func updateTeacher(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)

// 	var teacher Teacher

// 	json.NewDecoder(r.Body).Decode(&teacher)
// 	// save

// 	json.NewEncoder(w).Encode(teacher)

// }

// func deleteTeacher(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)

// 	var teacher Teacher
// 	// save

// 	json.NewEncoder(w).Encode("deleted successful")

// }

// // *****************TEACHER ATTENDANCE**************************

// func addTeacherattendace(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var teacherattendace AttendanceTeacher
// 	json.NewDecoder(r.Body).Decode(&teacherattendace)
// 	json.NewEncoder(w).Encode(teacherattendace)
// }

// func getTeacherattendace(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	id := params["id"]
// 	month := params["month"]
// 	year := params["year"]

// 	var teacherattendace AttendanceTeacher
// 	json.NewDecoder(r.Body).Decode(&teacherattendace)
// 	json.NewEncoder(w).Encode(teacherattendace)

// }
