package restHandler

import (
	"encoding/json"
	"strconv"
	"time"

	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	bean "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

// authorization

var jwtKey = []byte("secret_key")

// var users = map[string]string{
// 	"user1": "password1",
// 	"user2": "password2",
// }

// func Login(w http.ResponseWriter, r *http.Request) {

// 	var credentialStudent bean.Credentials
// 	// var credentialTeacher bean.Teacher
// 	err := json.NewDecoder(r.Body).Decode(&credentialStudent)
// 	// fmt.Println(credentialStudent)
// 	// err = json.NewDecoder(r.Body).Decode(&credentialTeacher)
// 	var getpassword bean.User

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	db := dataBase.Connect()

// 	defer db.Close()
// 	var mypass string

// 	err = db.Model(&getpassword).Column("password").Where("email=?", credentialStudent.Useremail).Select(&mypass)
// 	if err != nil {
// 		// log.Println("k0")
// 		if err == pg.ErrNoRows {

// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	expectedPassword := credentialStudent.Password
// 	actualPassword := mypass
// 	log.Println("k1")

// 	if expectedPassword != actualPassword {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	log.Println("k2")

// 	// expectedPassword, ok := users[credentials.Username]

// 	expirationTime := time.Now().Add(time.Minute * 5)

// 	claims := &bean.Claims{
// 		Useremail: credentialStudent.Useremail,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	http.SetCookie(w,
// 		&http.Cookie{
// 			Name:    "token",
// 			Value:   tokenString,
// 			Expires: expirationTime,
// 		})
// 	log.Println("kisahn")

// }

func ValidateTokenAndGetEmail(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			return "", err
		}
		w.WriteHeader(http.StatusBadRequest)
		return "", err
	}

	tokenStr := cookie.Value

	claims := &bean.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return "", err
		}
		w.WriteHeader(http.StatusBadRequest)
		return "", err
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return "", err
	}
	// check email expiration

	var newuser bean.User
	db := dataBase.Connect()
	defer db.Close()
	err = db.Model(&newuser).Where("email=?", claims.Useremail).Select()
	if err == pg.ErrNoRows {
		w.WriteHeader(http.StatusUnauthorized)
		return "", err
	}
	// if bean.Claims.ExpiresAt

	return claims.Useremail, err

}

// ********************************STUDENT************************************************

func GetStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	_, err := ValidateTokenAndGetEmail(w, r)
	if err != nil {
		json.NewEncoder(w).Encode("user is unauthorised")
		return

	}
	// var usr bean.User
	db := dataBase.Connect()
	defer db.Close()
	var students []bean.Student
	// var role int

	// err = db.Model(&usr).Column("role").Where("email=?", email).Select(&role)
	// if err != nil {
	// 	log.Println(err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// if role == 2 || role == 3  || role{
	if err := db.Model(&students).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(students)

	// } else {
	// 	json.NewEncoder(w).Encode("you are student you are not allowed to see the details of other students")
	// 	return

	// }
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := ValidateTokenAndGetEmail(w, r)

	if err != nil {
		json.NewEncoder(w).Encode("user is unauthorised")
		return
	}

	params := mux.Vars(r)

	db := dataBase.Connect()
	defer db.Close()

	// student_id := params["id"]
	student_id := params["id"]
	sid, err := strconv.Atoi(student_id)
	log.Println(err)

	students := &bean.Student{Sid: sid}

	if err := db.Model(students).Where("sid=?", sid).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(students)

}

func AddStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	role, err := getRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	db := dataBase.Connect()
	defer db.Close()
	if role == 3 {
		var userdetails bean.Userdetails

		_ = json.NewDecoder(r.Body).Decode(&userdetails)

		student := bean.Student{Name: userdetails.Name, Address: userdetails.Address, Class: userdetails.Class, Email: userdetails.Email}
		// _ = json.NewDecoder(r.Body).Decode(&student)

		// student.Id = uuid.New().String()

		if _, err := db.Model(&student).Insert(); err != nil {
			log.Println(err)
			// json.NewEncoder(w).Encode("error is line no 77")

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(student)
		// making an entry in user table

		AddUser(w, userdetails.Email, 1, userdetails.Password)

	} else {
		json.NewEncoder(w).Encode("only principle can add student")
		return

	}
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	role, err := getRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if role == 1 || role == 3 {

		db := dataBase.Connect()
		defer db.Close()

		params := mux.Vars(r)

		student_id := params["id"]
		trr, err := strconv.Atoi(student_id)
		log.Println(err)
		students := &bean.Student{Sid: trr}

		_ = json.NewDecoder(r.Body).Decode(&students)
		yy, err := db.Model(students).Where("sid=?", trr).Set("name= ?,address=?,class=?,email=?", students.Name, students.Address, students.Class, students.Email).Update()
		log.Println(yy)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(students)

	} else {
		json.NewEncoder(w).Encode("only student and principle  can update student")
		return

	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, err := getRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if role == 3 {

		params := mux.Vars(r)

		db := dataBase.Connect()
		defer db.Close()

		student_id := params["id"]

		sid, err := strconv.Atoi(student_id)
		log.Println(err)

		students := &bean.Student{Sid: sid}
		var email string
		err = db.Model(students).Column("email").Where("sid=?", sid).Select(&email)
		if err != nil {
			if err == pg.ErrNoRows {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("user with this sid doesn't exist ")
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var usr bean.User
		res, err := db.Model(&usr).Where("email=?", email).Delete()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if res != nil {
			json.NewEncoder(w).Encode("data deleted from user table")

		}

		result, err := db.Model(students).Where("sid=?", sid).Delete()

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if result != nil {
			json.NewEncoder(w).Encode("data deleted from student table")
			return
		}

		json.NewEncoder(w).Encode(result)

	} else {
		json.NewEncoder(w).Encode("only principle  can delete student")
		return
	}
}

// // *****************************AttendanceStudent***********************************
// // perform the first punchin in transaction

func StudentEntryPunchin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	role, err := getRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if role == 1 {

		studentattendance := bean.StudentAttendance{}

		_ = json.NewDecoder(r.Body).Decode(&studentattendance)

		db := dataBase.Connect()

		defer db.Close()

		// err := db.Model(&studentattendance).Where("id=? and date=?", studentattendance.Sid, studentattendance.Date).Select() // add date in where clause
		err := db.Model(&studentattendance).Where("sid=? and date=? and month=? and year=? ", studentattendance.Sid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise

		if err == pg.ErrNoRows {
			//  studentattendace.PunchIntime=time.Now()
			log.Println(studentattendance.Sid)

			studentattendance.Date = time.Now().Day()
			studentattendance.Month = int(time.Now().Month())
			studentattendance.Year = time.Now().Year()

			_, err := db.Model(&studentattendance).Insert()
			if err != nil {
				log.Println("166")
				log.Println(err)

				w.WriteHeader(http.StatusBadRequest)
				return
			}
			// log punch in

			punchin := &bean.StudentLogPunchs{
				Aid:  studentattendance.Aid,
				Time: time.Now().Add(time.Hour*5 + time.Minute*30),
				Type: 1,
			}
			_, err = db.Model(punchin).Insert()

			if err != nil {
				log.Println("182")

				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

		} else if err != nil {
			log.Println("190")

			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return

		} else {

			aid := studentattendance.Aid

			punchtable := bean.StudentLogPunchs{Aid: aid}

			pi_count, err := db.Model(&punchtable).Where("aid=? and type=?", aid, 1).Count()

			if err != nil {
				log.Println("205")

				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			po_count, err := db.Model(&punchtable).Where("aid=? and type=?", aid, 2).Count()

			if err != nil {
				log.Println("215")

				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			// json.NewEncoder(w).Encode(pi_count)
			// json.NewEncoder(w).Encode(po_count)

			if pi_count <= po_count {

				// punchtable.Time = time.Now()
				punchtable.Time = time.Now().Add(time.Hour*5 + time.Minute*30)

				punchtable.Type = 1
				_, err := db.Model(&punchtable).Insert()

				if err != nil {
					log.Println("216")

					log.Println(err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

			} else {

				json.NewEncoder(w).Encode("You have already punch in")
				return

			}

		}

		json.NewEncoder(w).Encode("punch in successful")
	} else {
		json.NewEncoder(w).Encode("you are not a student")
		return

	}
}

func StudentEntryPunchOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	role, err := getRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if role == 1 {

		studentattendance := bean.StudentAttendance{}

		_ = json.NewDecoder(r.Body).Decode(&studentattendance)

		db := dataBase.Connect()

		defer db.Close()

		err := db.Model(&studentattendance).Where("sid=? and date=? and month=? and year=? ", studentattendance.Sid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise
		if err == pg.ErrNoRows {
			json.NewEncoder(w).Encode(" no data found  so go for punch in first")
			return

		} else if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return

		} else {

			aid := studentattendance.Aid

			punchtable := bean.StudentLogPunchs{Aid: aid}

			pi_count, err := db.Model(&punchtable).Where("aid=? and type=?", aid, 1).Count()

			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			po_count, err := db.Model(&punchtable).Where("aid=? and type=?", aid, 2).Count()

			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if pi_count > po_count {

				// punchtable.Time = time.Now()
				punchtable.Time = time.Now().Add(time.Hour*5 + time.Minute*30)

				punchtable.Type = 2
				_, err := db.Model(&punchtable).Insert()

				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

			} else {

				json.NewEncoder(w).Encode("You have already punch out")
				return

			}

		}

		json.NewEncoder(w).Encode("punch out successful")
	} else {
		json.NewEncoder(w).Encode("you are not a student")
		return

	}
}

func GetStudentattendance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, err := getRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if role == 1 {

		db := dataBase.Connect()
		defer db.Close()
		var studentattendance bean.StudentAttendance
		json.NewDecoder(r.Body).Decode(&studentattendance)
		var studentattendancedetail []bean.StudentAttendancetemp
		err := db.Model(&studentattendancedetail).
			ColumnExpr(" DISTINCT student_attendances.date").
			Column("student_attendances.month").
			Column("student_attendances.year").
			Column("student_log_punchs.time").
			Column("student_log_punchs.type").
			Join("inner join student_attendances on student_attendances.aid=student_log_punchs.aid").
			Table("student_log_punchs").
			Where("student_attendances.sid=? AND student_attendances.month=? AND student_attendances.year=?", studentattendance.Sid, studentattendance.Month, studentattendance.Year).
			Select()

		if err == pg.ErrNoRows {
			json.NewEncoder(w).Encode("no data found with this details ")
			return

		} else if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return

		} else {
			// fmt.Println("Third case")
			if studentattendancedetail == nil {
				json.NewEncoder(w).Encode("student with this details doesn't exist")

			}
			json.NewEncoder(w).Encode(studentattendancedetail)
			return
		}

	} else {
		json.NewEncoder(w).Encode("you are not a student")
		return

	}

}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// *************************************************************************************************************************************

// ****************TEACHER*********************
