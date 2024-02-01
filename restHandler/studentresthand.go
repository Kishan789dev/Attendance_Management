package restHandler

import (
	"encoding/json"
	"fmt"
	"time"

	"log"
	"net/http"

	"github.com/go-pg/pg"
	bean "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

// ********************************STUDENT************************************************

// func GetStudents(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	db := dataBase.Connect()
// 	defer db.Close()
// 	var students []bean.Student
// 	if err := db.Model(&students).Select(); err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(students)

// }

// // func connect() {
// // 	panic("unimplemented")
// // }

// // func connect() {
// // 	panic("unimplemented")
// // }

// func GetStudent(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)

// 	db := dataBase.Connect()
// 	defer db.Close()

// 	// student_id := params["id"]
// 	student_id := params["id"]
// 	trr, err := strconv.Atoi(student_id)
// 	log.Println(err)

// 	students := &bean.Student{Sid: trr}

// 	if err := db.Model(students).WherePK().Select(); err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(students)

// }

// func AddStudent(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Print("hello2")
// 	w.Header().Set("Content-Type", "application/json")

// 	student := bean.Student{}
// 	_ = json.NewDecoder(r.Body).Decode(&student)

// 	db := dataBase.Connect()
// 	defer db.Close()
// 	// student.Id = uuid.New().String()
// 	if _, err := db.Model(&student).Insert(); err != nil {
// 		log.Println(err)
// 		// json.NewEncoder(w).Encode("error is line no 77")

// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(student)

// }

// func UpdateStudent(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	db := dataBase.Connect()
// 	defer db.Close()

// 	params := mux.Vars(r)

// 	student_id := params["id"]
// 	trr, err := strconv.Atoi(student_id)
// 	log.Println(err)
// 	students := &bean.Student{Sid: trr}

// 	_ = json.NewDecoder(r.Body).Decode(&students)
// 	yy, err := db.Model(students).WherePK().Set("name= ?,address=?,class=?,email=?", students.Name, students.Address, students.Class, students.Email).Update()
// 	log.Println(yy)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(students)

// }

// func DeleteStudent(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)

// 	db := dataBase.Connect()
// 	defer db.Close()

// 	student_id := params["id"]

// 	trr, err := strconv.Atoi(student_id)
// 	log.Println(err)

// 	students := &bean.Student{Sid: trr}
// 	result, err := db.Model(students).WherePK().Delete()

// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(result)

// }

// *****************************AttendanceStudent***********************************
// perform the first punchin in transaction

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

func GetClassattendance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get attendance request received")
	w.Header().Set("Content-Type", "application/json")
	db := dataBase.Connect()
	defer db.Close()
	// var student bean.Student

	var classtemp = &bean.Classtemp{}

	var classdata []bean.ClasstempRes
	json.NewDecoder(r.Body).Decode(&classtemp)

	// err := db.Model(&student).Where("class =?", classtemp.Class).Select()
	// err := db.Model(&classdata).
	// 	Column("students.class").
	// 	Column("student_attendances.date").Column("student_attendances.month").Column("student_attendances.year").
	// 	Join("INNER JOIN student_attendances on student_attendances.sid=students.sid").
	// 	Table("students").
	// 	Where("student_attendances.date=? AND student_attendances.month=? AND student_attendances.year=?", classtemp.Date, classtemp.Month, classtemp.Year).
	// 	Where("students.class =?", classtemp.Class).
	// 	Select()
	query := `select students.sid,students.name,students.class,student_attendances.date,student_attendances.month,student_attendances.year
			from students inner join student_attendances on students.sid=student_attendances.sid where student_attendances.date=? and
			student_attendances.month=? and student_attendances.year=? and students.class =?;`
	_, err := db.Query(&classdata, query, classtemp.Date, classtemp.Month, classtemp.Year, classtemp.Class)
	if err == pg.ErrNoRows {
		json.NewEncoder(w).Encode("no students are in this class ")
		return

	} else if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return

	} else {

		json.NewEncoder(w).Encode(classdata)
		return
	}

}

// func getClassattendance(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var studentattendace AttendanceStudent
// 	json.NewDecoder(r.Body).Decode(&studentattendace)
// 	json.NewEncoder(w).Encode(studentattendace)

// }
