package restHandler

import (
	"encoding/json"
	"strconv"
	"time"

	"log"
	"net/http"

	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	bean "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

// // ********************************STUDENT************************************************

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dataBase.Connect()
	defer db.Close()
	var students []bean.Student
	if err := db.Model(&students).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(students)

}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	db := dataBase.Connect()
	defer db.Close()

	// student_id := params["id"]
	student_id := params["id"]
	trr, err := strconv.Atoi(student_id)
	log.Println(err)

	students := &bean.Student{Sid: trr}

	if err := db.Model(students).Where("sid=?", trr).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(students)

}

func AddStudent(w http.ResponseWriter, r *http.Request) {
	// fmt.Print("hello2")
	w.Header().Set("Content-Type", "application/json")

	student := bean.Student{}
	_ = json.NewDecoder(r.Body).Decode(&student)

	db := dataBase.Connect()
	defer db.Close()
	// student.Id = uuid.New().String()
	if _, err := db.Model(&student).Insert(); err != nil {
		log.Println(err)
		// json.NewEncoder(w).Encode("error is line no 77")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(student)

}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

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

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	db := dataBase.Connect()
	defer db.Close()

	student_id := params["id"]

	trr, err := strconv.Atoi(student_id)
	log.Println(err)

	students := &bean.Student{Sid: trr}
	result, err := db.Model(students).Where("sid=?", trr).Delete()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if result != nil {
		json.NewEncoder(w).Encode("data deleted successfully")
		return
	}

	json.NewEncoder(w).Encode(result)

}

// // *****************************AttendanceStudent***********************************
// // perform the first punchin in transaction

func StudentEntryPunchin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
}

func StudentEntryPunchOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
}

func GetStudentattendance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// *************************************************************************************************************************************

// ****************TEACHER*********************
