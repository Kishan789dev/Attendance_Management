package routers

import (
	"github.com/gorilla/mux"
	rh "github.com/kk/attendance_management/restHandler"
)

func InitialiseRouter(r2 *mux.Router) {

	// r2.HandleFunc("/refresh", r/h.Refresh).Methods("GET")

	r2.HandleFunc("/login", rh.Login).Methods("POST")
	// r2.HandleFunc("/home", rh.Home).Methods("GET")
	r2.HandleFunc("/register", rh.Register).Methods("POST")

	r2.HandleFunc("/student/{id}", rh.GetStudent).Methods("GET")
	r2.HandleFunc("/students", rh.GetStudents).Methods("GET")
	// r2.HandleFunc("/student", rh.AddStudent).Methods("POST")
	r2.HandleFunc("/student/{id}", rh.UpdateStudent).Methods("PUT")
	r2.HandleFunc("/student/{id}", rh.DeleteStudent).Methods("DELETE")

	// // ********************Student attendance*****************

	r2.HandleFunc("/studentattendance/punchin", rh.StudentEntryPunchin).Methods("POST")
	r2.HandleFunc("/studentattendance/punchout", rh.StudentEntryPunchOut).Methods("POST")

	// r2.HandleFunc("/studentattendance/student", rh.GetStudentattendance).Methods("GET")
	r2.HandleFunc("/studentattendance", rh.GetClassattendance).Methods("GET")

	// r2.HandleFunc("/studentattendance/{class}/{date}/{month}/{year}", rh.GetClassattendance).Methods("GET")

	// log.Fatal(http.ListenAndServe(":5678", r2))q

	// 	// ************Teacher************************************

	r2.HandleFunc("/teacher/{id}", rh.GetTeacher).Methods("GET")
	r2.HandleFunc("/teachers", rh.GetTeachers).Methods("GET")
	// r2.HandleFunc("/teacher", rh.AddTeacher).Methods("POST")
	r2.HandleFunc("/teacher/{id}", rh.UpdateTeacher).Methods("PUT")
	r2.HandleFunc("/teacher/{id}", rh.DeleteTeacher).Methods("DELETE")

	// ********************Teacher attendance*****************

	r2.HandleFunc("/teacherattendance/punchin", rh.TeacherEntryPunchin).Methods("POST")
	r2.HandleFunc("/teacherattendance/punchout", rh.TeacherEntryPunchout).Methods("POST")

	r2.HandleFunc("/teacherattendance", rh.GetTeacherattendance).Methods("GET")

}
