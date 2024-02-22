package students

import (
	"encoding/json"
	"fmt"

	// "log"
	"strconv"
	"strings"

	"log"
	"net/http"

	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	services "github.com/kk/attendance_management/Services"
	auth "github.com/kk/attendance_management/authentication"
	bean "github.com/kk/attendance_management/bean"
	"github.com/kk/attendance_management/dataBase"
)

func extractStaticPart(errorMessage string) string {
	// Find the index of ":"
	index := strings.Index(errorMessage, ":")
	if index != -1 {
		// Extract the static part before ":"
		return strings.TrimSpace(errorMessage[:index])
	}
	// Return the entire error message if ":" is not found
	return errorMessage
}

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

	if role != 3 {
		json.NewEncoder(w).Encode("only principle can add student")
		return
	}

	// fmt.Println("add student ")
	var userdetails bean.Userdetails
	err = json.NewDecoder(r.Body).Decode(&userdetails)
	fmt.Println(userdetails)
	if err != nil {
		// fmt.Println("err1 ")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	student := bean.Student{
		Name:    userdetails.Name,
		Address: userdetails.Address,
		Class:   userdetails.Class,
		Email:   userdetails.Email,
	}
	fmt.Println(student)

	err = services.AddStudentService(&student)
	if err != nil {
		// fmt.Println("err2 ")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = auth.AddUser(userdetails.Email, 1, userdetails.Password)
	if err != nil {
		// fmt.Println("err3 ")

		json.NewEncoder(w).Encode("error in adding user to user table")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// fmt.Println("success ")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(" Student added suceesfully")
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
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err, CustomErrTyp, sid := services.StudentEntryPunchinSvc(email)
	log.Println("sidddd", sid)
	log.Println("errtypeeee", CustomErrTyp)

	if CustomErrTyp == 0 {
		json.NewEncoder(w).Encode("you are not a student")
		return
	}

	var aid int
	if CustomErrTyp == 1 {

		err, aid = services.StudentAttendanceWithPunchData(sid)
		fmt.Println("AIDDDDDD", aid)

		if err != nil {
			fmt.Println("kkkkk")

			w.WriteHeader(http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode("Success")
		return
	}

	if err != nil {
		fmt.Println("jjjjjjjjj")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// err, aid = services.StudentAttendanceWithPunchData(sid)
	// log.Println("aiiiiioooooooooooood", aid)
	aid = sid
	// if CustomErrTyp == 2 {
	err, str := services.StudentPunchEntryInTable(aid)

	if err != nil {
		// json.NewEncoder(w).Encode(err)
		if len(str) > 0 {
			json.NewEncoder((w)).Encode(str)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	json.NewEncoder(w).Encode(str)
}

// }

func StudentEntryPunchOut(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	email, err := auth.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		log.Println("logout", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err, CustomErrTyp, _, aid := services.StudentEntryPunchOutSvc(email)
	if CustomErrTyp == 0 {
		json.NewEncoder(w).Encode("you are not a student")
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

	err, str := services.StudentPunchOutEntryInTable(aid)
	if err != nil {
		// log.Println("logout777777", err)

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(str)

}

func GetStudentattendance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	email, err := auth.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("error here")
		json.NewEncoder(w).Encode("kindly login")
		return
	}

	err, sid := services.GetStudentsAttendanceEmailSvc(email)

	if err != nil {
		json.NewEncoder(w).Encode("you are not a student")
		return
	}

	studentattendance := &bean.StudentAttendance{Sid: sid}
	json.NewDecoder(r.Body).Decode(&studentattendance)

	studentattendancedetail, err := services.FetchAttendanceFromDetailsSvc(studentattendance)

	if err != nil {
		// log.Println("yyyy1", err.Error())
		staticPart := extractStaticPart(err.Error())

		json.NewEncoder(w).Encode(staticPart)

		w.WriteHeader(http.StatusBadRequest)
		return

	} else {

		log.Println("yyyy2", err)

		if len(*studentattendancedetail) == 0 {

			mapp := map[string]string{
				"error": "student with this details doesn't exist",
			}
			json.NewEncoder(w).Encode(mapp)
			return
		}
		log.Println("yyyy3", studentattendancedetail)

		json.NewEncoder(w).Encode(studentattendancedetail)
		return
	}

}

// func extractStaticPart(s string) {
// 	panic("unimplemented")
// }

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// *************************************************************************************************************************************

// ****************TEACHER*********************
