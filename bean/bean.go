package bean

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

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
}

type StudentLogPunchs struct {
	Lid  int       `json:"lid"`
	Aid  int       `json:"aid"`
	Time time.Time `json:"time"`
	Type int       `json:"type"`
}

type StudentAttendancetemp struct {
	TableName struct{} `sql:"student_log_punchs" json:"-"`

	Date  int       `json:"date" sql:"date"`
	Month int       `json:"month" sql:"month"`
	Year  int       `json:"year" sql:"year"`
	Time  time.Time `json:"time"  sql:"time"`
	Type  int       `json:"type"   sql:"type"`
}

// **************TEACHER**************

type Teacher struct {
	TableName struct{} `sql:"teachers" json:"-"`
	Tid       int      `json:"tid" sql:"tid"`
	Name      string   `json:"name" sql:"name"`
	Address   string   `json:"address" sql:"address"`
	Email     string   `json:"email" sql:"email"`
}

type TeacherAttendance struct {
	TableName struct{} `sql:"teacher_attendances" json:"-"`
	Aid       int      `json:"aid" sql:"aid"`
	Tid       int      `json:"tid" sql:"tid"`
	Date      int      `json:"date" sql:"date"`
	Month     int      `json:"month" sql:"month"`
	Year      int      `json:"year" sql:"year"`
}

type TeacherLogPunchs struct {
	Lid  int       `json:"lid"`
	Aid  int       `json:"aid"`
	Time time.Time `json:"time"`
	Type int       `json:"type"`
}

type TeacherLogPunchstemp struct {
	// TableName struct{} `sql:"teacher_log_punchs" json:"-"`
	// Lid int `json:"-"`
	// Aid int `json:"-"`

	Date  int       `json:"-" sql:"date"`
	Month int       `json:"-" sql:"month"`
	Year  int       `json:"-" sql:"year"`
	Time  time.Time `json:"time"  sql:"time"`
	Type  int       `json:"type"   sql:"type"`
}

type TeacherAttendancetemp struct {
	TableName struct{} `sql:"teacher_log_punchs" json:"-"`

	Date  int       `json:"-" sql:"date"`
	Month int       `json:"-" sql:"month"`
	Year  int       `json:"-" sql:"year"`
	Time  time.Time `json:"time"  sql:"time"`
	Type  int       `json:"type"   sql:"type"`
}

// ****CLASS*********

type Classtemp struct {
	Class int `json:"class" sql:"class"`
	Date  int `json:"date" sql:"date"`
	Month int `json:"month" sql:"month"`
	Year  int `json:"year" sql:"year"`
}
type ClasstempRes struct {
	TableName struct{} `sql:"students" json:"-"`

	Sid   int    `json:"-" sql:"sid"`
	Name  string `json:"name" sql:"name"`
	Class int    `json:"-" sql:"class"`
	Date  int    `json:"-" sql:"date"`
	Month int    `json:"-" sql:"month"`
	Year  int    `json:"-" sql:"year"`
}

type User struct {
	Uid      int    `json:"uid" sql:"uid"`
	Email    string `json:"email" sql:"email"`
	Role     int    `json:"role" sql:"role"`
	Password string `jso:"password" sql:"password"`
}

type Userdetails struct {
	TableName struct{} `sql:"users" json:"-"`
	Name      string   `json:"name" sql:"name"`
	Address   string   `json:"address" sql:"address"`
	Class     int      `json:"class" sql:"class"`
	Email     string   `json:"email" sql:"email"`
	// Role      int      `json:"role" sql:"role"`
	Password string `jso:"password" sql:"password"`
}

type Credentials struct {
	Useremail string `json:"useremail"`
	Password  string `json:"password"`
}

type Claims struct {
	Useremail string `json:"useremail"`
	jwt.StandardClaims
}
