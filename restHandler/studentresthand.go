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

// ********************************STUDENT************************************************

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

// func connect() {
// 	panic("unimplemented")
// }

// func connect() {
// 	panic("unimplemented")
// }

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

	if err := db.Model(students).WherePK().Select(); err != nil {
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
	yy, err := db.Model(students).WherePK().Set("name= ?,address=?,class=?,email=?", students.Name, students.Address, students.Class, students.Email).Update()
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
	result, err := db.Model(students).WherePK().Delete()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)

}

// *****************************AttendanceStudent***********************************

func StudentEntryPunchin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	studentattendance := bean.StudentAttendance{}

	_ = json.NewDecoder(r.Body).Decode(&studentattendance)

	db := dataBase.Connect()

	defer db.Close()

	err := db.Model(&studentattendance).Where("id=?", studentattendance.Sid).Select()
	if err == pg.ErrNoRows {
		//  studentattendace.PunchIntime=time.Now()
		studentattendance.Date = time.Now().Day()
		studentattendance.Month = int(time.Now().Month())
		studentattendance.Date = time.Now().Year()

		_, err := db.Model(&studentattendance).Insert()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// log punch in

		punchin := &bean.StudentLogPunchs{
			Aid:  studentattendance.Aid,
			Time: time.Now(),
			Type: 1,
		}
		_, err = db.Model(punchin).Insert()

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

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

		if pi_count < po_count {

			punchtable.Time = time.Now()
			punchtable.Type = 1
			_, err := db.Model(&punchtable).Insert()

			if err != nil {
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

	err := db.Model(&studentattendance).Where("id=?", studentattendance.Sid).Select()
	if err == pg.ErrNoRows {
		json.NewEncoder(w).Encode(" no data found  so go for punch in first")

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

			punchtable.Time = time.Now()
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

// func getStudentattendance(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var studentattendace AttendanceStudent
// 	json.NewDecoder(r.Body).Decode(&studentattendace)
// 	json.NewEncoder(w).Encode(studentattendace)

// }

// func getClassattendance(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var studentattendace AttendanceStudent
// 	json.NewDecoder(r.Body).Decode(&studentattendace)
// 	json.NewEncoder(w).Encode(studentattendace)

// }
