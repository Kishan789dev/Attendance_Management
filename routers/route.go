package routers

import (
	"github.com/gorilla/mux"
	"github.com/kk/attendance_management/authentication/login"
	"github.com/kk/attendance_management/components/students"
	"github.com/kk/attendance_management/components/teachers"
)

type RouteImp struct {
	login    login.LoginRest
	students students.StudentRest
	teachers teachers.TeacherRest
}

func NewRoute(r2 *mux.Router, login login.LoginRest, students students.StudentRest, teachers teachers.TeacherRest) *RouteImp {
	return &RouteImp{
		login:    login,
		students: students,
		teachers: teachers,
	}
}

func (impl *RouteImp) InitialiseRouter(r2 *mux.Router) {

	r2.HandleFunc("/login", impl.login.Login).Methods("POST")

	r2.HandleFunc("/student", impl.students.AddStudent).Methods("POST")

	// // ********************Student attendance*****************

	r2.HandleFunc("/studentattendance/punchin", impl.students.StudentEntryPunchin).Methods("POST")
	r2.HandleFunc("/studentattendance/punchout", impl.students.StudentEntryPunchOut).Methods("POST")

	r2.HandleFunc("/studentattendance/student", impl.students.GetStudentattendance).Methods("POST")
	r2.HandleFunc("/classattendance", impl.teachers.GetClassattendance).Methods("POST")

	// 	// ************Teacher************************************

	r2.HandleFunc("/teacher", impl.teachers.AddTeacher).Methods("POST")

	// ********************Teacher attendance*****************

	r2.HandleFunc("/teacherattendance/punchin", impl.teachers.TeacherEntryPunchin).Methods("POST")
	r2.HandleFunc("/teacherattendance/punchout", impl.teachers.TeacherEntryPunchOut).Methods("POST")

	r2.HandleFunc("/teacherattendance", impl.teachers.GetTeacherattendance).Methods("POST")

}
