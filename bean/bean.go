package bean

import "time"

type Student struct {
	TableName struct{} `sql:"students" json:"-"`
	Sid       int      `json:"sid" sql:"sid"`
	Name      string   `json:"name" sql:"name"`
	Address   string   `json:"address" sql:"address"`
	Class     int      `json:"class" sql:"class"`
	Email     string   `json:"email" sql:"email"`
}

type StudentAttendance struct {
	TableName struct{} `sql:"student_attendances" json:"-"`
	Aid       int      `json:"aid" sql:"aid"`
	Sid       int      `json:"sid" sql:"sid"`
	Date      int      `json:"date" sql:"date"`
	Month     int      `json:"month" sql:"month"`
	Year      int      `json:"year" sql:"year"`
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

type Classtemp struct {
	// TableName struct{} `sql:"classtemp" json:"-"`
	Class int `json:"class" sql:"class"`
	Date  int `json:"date" sql:"date"`
	Month int `json:"month" sql:"month"`
	Year  int `json:"year" sql:"year"`
	// student Student
}

// type ClasstempRes struct {
// 	// TableName struct{} `sql:"classtemp" json:"-"`
// 	Sid   int    `json:"id" sql:"id"`
// 	Name  string `json:"name" sql:"name"`
// 	Class int    `json:"class" sql:"class"`
// 	Date  int    `json:"date" sql:"date"`
// 	Month int    `json:"month" sql:"month"`
// 	Year  int    `json:"year" sql:"year"`

// 	// student Student
// }

type ClasstempRes struct {
	Sid   int    `json:"sid" sql:"sid"`
	Name  string `json:"name" sql:"name"`
	Class int    `json:"class" sql:"class"`
	Date  int    `json:"date" sql:"date"`
	Month int    `json:"month" sql:"month"`
	Year  int    `json:"year" sql:"year"`
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
