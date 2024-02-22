package routers

import (
	"github.com/gorilla/mux"
	"github.com/kk/attendance_management/authentication/login"
	"github.com/kk/attendance_management/components/students"
	"github.com/kk/attendance_management/components/teachers"
)

func InitialiseRouter(r2 *mux.Router) {

	// r2.HandleFunc("/refresh", r/h.Refresh).Methods("GET")

	r2.HandleFunc("/login", login.Login).Methods("POST")
	// r2.HandleFunc("/home", rh.Home).Methods("GET")
	// r2.HandleFunc("/register", login.Register).Methods("POST")

	r2.HandleFunc("/student/{id}", students.GetStudent).Methods("GET")
	r2.HandleFunc("/students", students.GetStudents).Methods("GET")
	r2.HandleFunc("/student", students.AddStudent).Methods("POST")
	r2.HandleFunc("/student/{id}", students.UpdateStudent).Methods("PUT")
	r2.HandleFunc("/student/{id}", students.DeleteStudent).Methods("DELETE")

	// // ********************Student attendance*****************

	r2.HandleFunc("/studentattendance/punchin", students.StudentEntryPunchin).Methods("POST")
	r2.HandleFunc("/studentattendance/punchout", students.StudentEntryPunchOut).Methods("POST")

	r2.HandleFunc("/studentattendance/student", students.GetStudentattendance).Methods("POST")
	r2.HandleFunc("/classattendance", teachers.GetClassattendance).Methods("POST")

	// r2.HandleFunc("/studentattendance/{class}/{date}/{month}/{year}", rh.GetClassattendance).Methods("GET")

	// log.Fatal(http.ListenAndServe(":5678", r2))q

	// 	// ************Teacher************************************

	r2.HandleFunc("/teacher/{id}", teachers.GetTeacher).Methods("GET")
	r2.HandleFunc("/teachers", teachers.GetTeachers).Methods("GET")
	r2.HandleFunc("/teacher", teachers.AddTeacher).Methods("POST")
	r2.HandleFunc("/teacher/{id}", teachers.UpdateTeacher).Methods("PUT")
	r2.HandleFunc("/teacher/{id}", teachers.DeleteTeacher).Methods("DELETE")

	// ********************Teacher attendance*****************

	r2.HandleFunc("/teacherattendance/punchin", teachers.TeacherEntryPunchin).Methods("POST")
	r2.HandleFunc("/teacherattendance/punchout", teachers.TeacherEntryPunchOut).Methods("POST")

	r2.HandleFunc("/teacherattendance", teachers.GetTeacherattendance).Methods("POST")

}
