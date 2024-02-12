package restHandler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"log"
	"net/http"

	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	services "github.com/kk/attendance_management/Services"
	auth "github.com/kk/attendance_management/authentication"
	bean "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	_, err := auth.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		json.NewEncoder(w).Encode("user is unauthorised")
		return

	}

	students, err := services.GetStudentsSvc()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(students)

}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := auth.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		json.NewEncoder(w).Encode("user is unauthorised")
		return
	}

	params := mux.Vars(r)
	student_id := params["id"]
	id, _ := strconv.Atoi(student_id)

	students, err := services.GetStudentSvc(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(students)

}

func AddStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	role, err := auth.GetRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Todo - make a function isPrincipal to check...
	if role != 3 {
		json.NewEncoder(w).Encode("only principle can add student")
		return
	}

	var userdetails bean.Userdetails
	err = json.NewDecoder(r.Body).Decode(&userdetails)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	student := bean.Student{
		Name:    userdetails.Name,
		Address: userdetails.Address,
		Class:   userdetails.Class,
		Email:   userdetails.Email,
	}

	_, err = services.AddStudentService(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = auth.AddUser(userdetails.Email, 1, userdetails.Password)
	if err != nil {
		json.NewEncoder(w).Encode("error in adding user to user table")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	role, err := auth.GetRole(w, r)
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

	role, err := auth.GetRole(w, r)
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
	email, err := auth.ValidateTokenAndGetEmail(w, r)

	// role, err := getRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	db := dataBase.Connect()
	defer db.Close()
	var sid int
	var student bean.Student
	err = db.Model(&student).Column("sid").Where("email=?", email).Select(&sid)

	if err == nil {

		studentattendance := bean.StudentAttendance{Sid: sid}

		// _ = json.NewDecoder(r.Body).Decode(&studentattendance)

		// err := db.Model(&studentattendance).Where("id=? and date=?", studentattendance.Sid, studentattendance.Date).Select() // add date in where clause
		err := db.Model(&studentattendance).Where("sid=? and date=? and month=? and year=? ", sid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise

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
			// log.Println("190")

			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			// json.NewEncoder(w).Encode("only student can puch in")
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
	email, err := auth.ValidateTokenAndGetEmail(w, r)

	// role, err := getRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	db := dataBase.Connect()
	defer db.Close()
	var sid int
	var student bean.Student
	err = db.Model(&student).Column("sid").Where("email=?", email).Select(&sid)
	// fmt.Println("email", email, "sid", sid)

	if err == nil {

		studentattendance := bean.StudentAttendance{Sid: sid}

		err := db.Model(&studentattendance).Where("sid=? and date=? and month=? and year=? ", sid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise
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

	email, err := auth.ValidateTokenAndGetEmail(w, r)

	// role, err := getRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	db := dataBase.Connect()
	defer db.Close()
	var sid int
	var student bean.Student
	err = db.Model(&student).Column("sid").Where("email=?", email).Select(&sid)
	fmt.Println("email", email, "sid", sid)

	if err == nil {

		studentattendance := &bean.StudentAttendance{Sid: sid}
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
			Where("student_attendances.sid=? AND student_attendances.month=? AND student_attendances.year=?", sid, studentattendance.Month, studentattendance.Year).
			Select()

		if err == pg.ErrNoRows {
			json.NewEncoder(w).Encode("no data found with this details ")
			return

		} else if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return

		} else {
			fmt.Println(studentattendance)
			if studentattendancedetail == nil {
				json.NewEncoder(w).Encode("student with this details doesn't exist")
				return

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
