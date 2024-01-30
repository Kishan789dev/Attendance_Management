package bean

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Class   int    `json:"class"`
	Email   string `json:"email"`
}

type AttendanceStudent struct {
	StudentId              string `json:"studentid"`
	StudentAttendancedate  string `json:"studentattendancedate"`
	StudentAttendancemonth string `json:"studentattendancemonth"`
	StudentAttendanceyear  string `json:"studentattendanceyear"`

	StudentPresentStatus bool `json:"studentpresentstatus"`
	// StudentPunchIntime     bool `json:"studentintime"`
	// StudentPunchOuttime   bool `json:"studentouttime"`
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
