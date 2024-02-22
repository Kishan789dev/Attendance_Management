package teachers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// "time"

	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	"github.com/kk/attendance_management/authentication/getrole"
	token "github.com/kk/attendance_management/authentication/tokenvalidation"

	bean "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

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

func GetTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := token.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		json.NewEncoder(w).Encode("user is unauthorised")
		return

	}

	db := dataBase.Connect()
	defer db.Close()
	var teachers []bean.Teacher
	if err := db.Model(&teachers).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//
	json.NewEncoder(w).Encode(teachers)

}

func GetTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// log.Println("dsjflsdflsdjflksd")

	_, err := token.ValidateTokenAndGetEmail(w, r)
	// log.Println(email)

	if err != nil {
		json.NewEncoder(w).Encode("user is unauthorised")
		return
	}

	params := mux.Vars(r)

	db := dataBase.Connect()
	defer db.Close()

	// teacher_id := params["id"]
	teacher_id := params["id"]
	tid, err := strconv.Atoi(teacher_id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// log.Println(err)

	teachers := &bean.Teacher{Tid: tid}

	if err := db.Model(teachers).Where("tid=?", tid).Select(); err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("not authorised")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(teachers)

}

func AddTeacher(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	role, err := getrole.GetRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if role != 3 {
		json.NewEncoder(w).Encode("only principle can add teacher")
		return
	}

	var userdetails bean.Userdetails

	_ = json.NewDecoder(r.Body).Decode(&userdetails)

	teacher := bean.Teacher{Name: userdetails.Name, Address: userdetails.Address, Email: userdetails.Email}
	err = AddTeacherSvc(&teacher, &userdetails)

	if err != nil {

		// w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode("teacher added successfully")

}

func UpdateTeacher(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	role, err := getrole.GetRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	db := dataBase.Connect()
	defer db.Close()

	if role == 2 || role == 3 {

		params := mux.Vars(r)

		teacher_id := params["id"]
		trr, err := strconv.Atoi(teacher_id)
		log.Println(err)
		teachers := &bean.Teacher{Tid: trr}

		_ = json.NewDecoder(r.Body).Decode(&teachers)
		yy, err := db.Model(teachers).Where("tid=?", teacher_id).Set("name= ?,address=?,email=?", teachers.Name, teachers.Address, teachers.Email).Update()
		log.Println(yy)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(teachers)
	} else {
		json.NewEncoder(w).Encode("only student and principle  can update teacher")
		return
	}
}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	role, err := getrole.GetRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	db := dataBase.Connect()
	defer db.Close()
	if role == 3 {

		params := mux.Vars(r)

		db := dataBase.Connect()
		defer db.Close()

		teacher_id := params["id"]

		tid, err := strconv.Atoi(teacher_id)
		log.Println(err)

		teachers := &bean.Teacher{Tid: tid}

		var email string
		err = db.Model(teachers).Column("email").Where("tid=?", tid).Select(&email)
		if err != nil {
			if err == pg.ErrNoRows {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("user with this sid doesn't exist ")
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		result, err := db.Model(teachers).Where("tid=?", teacher_id).Delete()

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if result != nil {
			json.NewEncoder(w).Encode("data deleted from user table")

		}
		var usr bean.User
		res, err := db.Model(&usr).Where("email=?", email).Delete()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if res != nil {
			json.NewEncoder(w).Encode("data deleted from teacher table")
			return
		}

	} else {

		json.NewEncoder(w).Encode("only principle  can delete teacher")
		return
	}

}

// *****************TEACHER ATTENDANCE**************************

// func AddTeacherattendace(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var teacherattendace bean.TeacherAttendance
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

// func TeacherEntryPunchin(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	email, err := token.ValidateTokenAndGetEmail(w, r)
// 	fmt.Println(email)
// 	// fmt.Println("kjgdjkhgh")

// 	// role, err := getRole(w, r)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	db := dataBase.Connect()
// 	defer db.Close()
// 	var tid int
// 	var teacher bean.Teacher
// 	err = db.Model(&teacher).Column("tid").Where("email=?", email).Select(&tid)
// 	log.Println(tid)

// 	if err == nil {

// 		teacherattendance := bean.TeacherAttendance{Tid: tid}
// 		// err := db.Model(&teacherattendance).Where("id=? and date=?", teacherattendance.Sid, teacherattendance.Date).Select() // add date in where clause
// 		err := db.Model(&teacherattendance).Where("tid=? and date=? and month=? and year=? ", tid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise

// 		if err == pg.ErrNoRows {
// 			//  teacherattendace.PunchIntime=time.Now()
// 			log.Println(teacherattendance.Tid)

// 			teacherattendance.Date = time.Now().Day()
// 			teacherattendance.Month = int(time.Now().Month())
// 			teacherattendance.Year = time.Now().Year()
// 			fmt.Println(teacherattendance)

// 			_, err := db.Model(&teacherattendance).Insert()
// 			if err != nil {
// 				log.Println("166")
// 				log.Println(err)

// 				w.WriteHeader(http.StatusBadRequest)
// 				return
// 			}
// 			// log punch in

// 			punchin := &bean.TeacherLogPunchs{
// 				Aid:  teacherattendance.Aid,
// 				Time: time.Now().Add(time.Hour*5 + time.Minute*30),
// 				Type: 1,
// 			}
// 			_, err = db.Model(punchin).Insert()

// 			if err != nil {
// 				log.Println("182")

// 				log.Println(err)
// 				w.WriteHeader(http.StatusBadRequest)
// 				return
// 			}

// 		} else if err != nil {
// 			log.Println("190")

// 			log.Println(err)
// 			w.WriteHeader(http.StatusBadRequest)
// 			return

// 		} else {

// 			aid := teacherattendance.Aid

// 			punchtable := bean.TeacherLogPunchs{Aid: aid}

// 			pi_count, err := db.Model(&punchtable).Where("aid=? and type=?", aid, 1).Count()

// 			if err != nil {
// 				log.Println("205")

// 				log.Println(err)
// 				w.WriteHeader(http.StatusBadRequest)
// 				return
// 			}

// 			po_count, err := db.Model(&punchtable).Where("aid=? and type=?", aid, 2).Count()

// 			if err != nil {
// 				log.Println("215")

// 				log.Println(err)
// 				w.WriteHeader(http.StatusBadRequest)
// 				return
// 			}
// 			// json.NewEncoder(w).Encode(pi_count)
// 			// json.NewEncoder(w).Encode(po_count)

// 			if pi_count <= po_count {

// 				// punchtable.Time = time.Now()
// 				punchtable.Time = time.Now().Add(time.Hour*5 + time.Minute*30)

// 				punchtable.Type = 1
// 				_, err := db.Model(&punchtable).Insert()

// 				if err != nil {
// 					log.Println("216")

// 					log.Println(err)
// 					w.WriteHeader(http.StatusBadRequest)
// 					return
// 				}

// 			} else {

// 				json.NewEncoder(w).Encode("You have already punch in")
// 				return

// 			}

// 		}

// 		json.NewEncoder(w).Encode("punch in successful")
// 	} else {
// 		json.NewEncoder(w).Encode("you are not a teacher")
// 		return
// 	}

// }

// func TeacherEntryPunchout(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	email, err := token.ValidateTokenAndGetEmail(w, r)
// 	log.Println(email)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	db := dataBase.Connect()
// 	defer db.Close()
// 	var tid int
// 	var teacher bean.Teacher
// 	err = db.Model(&teacher).Column("tid").Where("email=?", email).Select(&tid)
// 	// log.Println(tid)

// 	if err == nil {

// 		teacherattendance := bean.TeacherAttendance{Tid: tid}

// 		err := db.Model(&teacherattendance).Where("tid=? and date=? and month=? and year=? ", tid, time.Now().Day(), int(time.Now().Month()), time.Now().Year()).Select() // add date in where claise
// 		if err == pg.ErrNoRows {
// 			json.NewEncoder(w).Encode(" no data found  so go for punch in first")
// 			return

// 		} else if err != nil {
// 			log.Println(err)
// 			w.WriteHeader(http.StatusBadRequest)
// 			return

// 		} else {

// 			aid := teacherattendance.Aid

// 			punchtable := bean.TeacherLogPunchs{Aid: aid}

// 			pi_count, err := db.Model(&punchtable).Where("aid=? and type=?", aid, 1).Count()

// 			if err != nil {
// 				log.Println(err)
// 				w.WriteHeader(http.StatusBadRequest)
// 				return
// 			}

// 			po_count, err := db.Model(&punchtable).Where("aid=? and type=?", aid, 2).Count()

// 			if err != nil {
// 				log.Println(err)
// 				w.WriteHeader(http.StatusBadRequest)
// 				return
// 			}

// 			if pi_count > po_count {

// 				// punchtable.Time = time.Now()
// 				punchtable.Time = time.Now().Add(time.Hour*5 + time.Minute*30)

// 				punchtable.Type = 2
// 				_, err := db.Model(&punchtable).Insert()

// 				if err != nil {
// 					log.Println(err)
// 					w.WriteHeader(http.StatusBadRequest)
// 					return
// 				}

// 			} else {

// 				json.NewEncoder(w).Encode("You have already punch out")
// 				return

// 			}

// 		}

// 		json.NewEncoder(w).Encode("punch out successful")
// 	} else {

// 		json.NewEncoder(w).Encode("you are not a teacher")
// 		return
// 	}
// }

func TeacherEntryPunchin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	email, err := token.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err, CustomErrTyp, tid := TeacherEntryPunchinSvc(email)
	log.Println("tidddd", tid)
	log.Println("errtypeeee", CustomErrTyp)

	if CustomErrTyp == 0 {
		json.NewEncoder(w).Encode("you are not a Teacher")
		return
	}

	var aid int
	if CustomErrTyp == 1 {

		err, aid = TeacherAttendanceWithPunchData(tid)
		fmt.Println("AIDDDDDD", aid)

		if err != nil {
			fmt.Println("kkkkk")

			w.WriteHeader(http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode("Success")
		return

	}
	// err, aid = TeacherAttendanceWithPunchData(tid )

	if err != nil {
		fmt.Println("jjjjjjjjj")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// err, aid = TeacherAttendanceWithPunchData(tid )
	log.Println("aid", tid)
	aid = tid
	err, str := TeacherPunchEntryInTable(aid)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(str)

}

func TeacherEntryPunchOut(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	email, err := token.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		log.Println("logout", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err, CustomErrTyp, _, aid := TeacherEntryPunchOutSvc(email)
	if CustomErrTyp == 0 {
		json.NewEncoder(w).Encode("you are not a Teacher")
		return
	}

	if err != nil && err != pg.ErrNoRows {
		log.Println("logout2222", err)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// var aid int
	if err == pg.ErrNoRows && CustomErrTyp == 1 {
		log.Println("logout3", err)

		json.NewEncoder(w).Encode(" no data found  so go for punch in first")

		return

	}
	log.Println("errrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr", aid)

	err, str := TeacherPunchOutEntryInTable(aid)
	if err != nil {
		// log.Println("logout777777", err)

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(str)

}

func GetTeacherattendance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, err := getrole.GetRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	email, err := token.ValidateTokenAndGetEmail(w, r)
	log.Println(email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var tid int

	if role == 1 {
		json.NewEncoder(w).Encode("u are a student so u cant get details")
		return

	}
	if role == 2 {

		err = GetTeacherattendanceSvcTidGetting(email, &tid)

		if err != nil {
			// making role as student so that it can't proceed  further
			role = 1
			//  errMsg:= make(map[string]string, 0)
			//  errMsg["error"] = err.Error()

			json.NewEncoder(w).Encode(err.Error())
			return
		}

	}

	if role == 2 || role == 3 {
		teacherattendance := &bean.TeacherAttendance{}

		json.NewDecoder(r.Body).Decode(&teacherattendance)

		if teacherattendance.Tid != 0 {
			tid = teacherattendance.Tid
		}

		err, teacherattendancedetail := GetTeacherAttendanceDetailsSvc(tid, teacherattendance.Month, teacherattendance.Year)

		if err != nil {
			fmt.Println("oooooo")

			// log.Println(err)
			errMsg := make(map[string]string, 0)
			errMsg["error"] = err.Error()

			json.NewEncoder(w).Encode(errMsg)
			// w.WriteHeader(http.StatusBadRequest)
			return

		} else {
			// fmt.Println("lllllllllll")
			// fmt.Println("ateendddddddd", teacherattendancedetail)

			json.NewEncoder(w).Encode(teacherattendancedetail)
			return
		}
	} else {
		json.NewEncoder(w).Encode("you are not a teacher")
		return
	}

}

func GetClassattendance(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	role, err := getrole.GetRole(w, r)
	// log.Println(email)
	if err != nil {
		// w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("you are not a teacher")

		return
	}
	if role == 1 || role == 3 {
		json.NewEncoder(w).Encode("you are not a teacher")
		return

	}

	// var tid int
	// var teacher bean.Teacher
	// err = db.Model(&teacher).Column("tid").Where("email=?", email).Select(&tid)
	// if err != nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	// var student bean.Student

	var classtemp = bean.Classtemp{}

	json.NewDecoder(r.Body).Decode(&classtemp)

	err, classdata := GetClassattendanceSvc(&classtemp)

	// err := db.Model(&student).Where("class =?", classtemp.Class).Select()

	// query := `select students.sid,students.name,students.class,student_attendances.date,student_attendances.month,student_attendances.year
	// 		from students inner join student_attendances on students.sid=student_attendances.sid where student_attendances.date=? and
	// 		student_attendances.month=? and student_attendances.year=? and students.class =?;`
	// _, err := db.Query(&classdata, query, classtemp.Date, classtemp.Month, classtemp.Year, classtemp.Class)

	// fmt.Println(err)
	if err != nil {
		errMsg := make(map[string]string, 0)
		errMsg["error"] = err.Error()

		json.NewEncoder(w).Encode(errMsg)
		return

	} else {
		// json.NewEncoder(w).Encode(len(classdata))
		if classdata != nil {
			json.NewEncoder(w).Encode(classdata)

		} else {
			errorMap := map[string]string{
				"error": "invalid inputs",
			}
			// return errorMap
			json.NewEncoder(w).Encode(errorMap)

		}
		return
	}

}
