package bean

import "time"

type Student struct {
	Sid     int    `json:"sid"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Class   int    `json:"class"`
	Email   string `json:"email"`
}

type StudentAttendance struct {
	Aid   int `json:"aid"`
	Sid   int `json:"sid"`
	Date  int `json:"date"`
	Month int `json:"month"`
	Year  int `json:"year"`
	// Presentstatus bool      `json:"presentstatus"`
	// PunchIntime   time.Time `json:"punchintime"`
	// Puncouttime   time.Time `json:"punchouttime"`

	// StudentPunchIntime     bool `json:"studentintime"`
	// StudentPunchOuttime   bool `json:"studentouttime"`
}

type StudentLogPunchs struct {
	Lid  int       `json:"lid"`
	Aid  int       `json:"aid"`
	Time time.Time `json:"time"`
	Type int       `json:"type"`
}

// type Teacher struct {
// 	gorm.Model
// 	TeacherId      int    `json:"teacherid"`
// 	TeacherName    string `json:"teachername"`
// 	TeacherAddress string `json:"teacheraddress"`
// 	TeacherEmail   string `json:"teacheremail"`
// }

// type AttendanceTeacher struct {
// 	gorm.Model
// 	Id    int       `json:"teacherid"`
// 	Date  time.Time `json:"teacherattendancedate"`
// 	Month time.Time `json:"teacherattendancemonth"`
// 	Year  time.Time `json:"teacherattendanceyear"`

// 	PresentStatus bool `json:"teacherpresentstatus"`
// 	// TeacherPunchIntime    bool `json:"teacherintime"`
// 	// TeacherPunchOuttime   bool `json:"teacherouttime"`
// }
