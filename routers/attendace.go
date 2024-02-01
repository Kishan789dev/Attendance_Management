package routers

import (
	"github.com/gorilla/mux"
	rh "github.com/kk/attendance_management/restHandler"
)

func InitialiseRouter(r2 *mux.Router) {
	// r2.HandleFunc("/student/{id}", rh.GetStudent).Methods("GET")
	// r2.HandleFunc("/student", rh.GetStudents).Methods("GET")
	// r2.HandleFunc("/student", rh.AddStudent).Methods("POST")
	// r2.HandleFunc("/student/{id}", rh.UpdateStudent).Methods("PUT")
	// r2.HandleFunc("/student/{id}", rh.DeleteStudent).Methods("DELETE")

	// 	// ************Teacher************************************

	// ********************Student attendance*****************

	r2.HandleFunc("/studentattendance/punchin", rh.StudentEntryPunchin).Methods("POST")
	r2.HandleFunc("/studentattendance/punchout", rh.StudentEntryPunchOut).Methods("POST")

	// r.HandleFunc("/studentattendance/{id}/{month}/{year}", rh.GetStudentattendance).Methods("GET")
	r2.HandleFunc("/studentattendance", rh.GetClassattendance).Methods("GET")
	// r2.HandleFunc("/studentattendance/{class}/{date}/{month}/{year}", rh.GetClassattendance).Methods("GET")

	// log.Fatal(http.ListenAndServe(":5678", r2))
}

// 	r.HandleFunc("/techer/{id}", GetTeacher).Methods("GET")
// 	r.HandleFunc("/teacher", GetTeachers).Methods("GET")
// 	r.HandleFunc("/teacher", AddTeacher).Methods("POST")
// 	r.HandleFunc("/teacher{id}", UpdateTeacher).Methods("PUT")
// 	r.HandleFunc("/teacher{id}", DeleteTeacher).Methods("DELETE")

// 	// ********************Teacher attendance*****************

// 	r.HandleFunc("/teacherattendance", AddTeacherattendace).Methods("PUT")
// 	r.HandleFunc("/teacherattendace/{id}/{month}/{year}", GetTeacherattendace).Methods("GET")
// 	// principle
// 	// r.HandleFunc("/teacherattendace/{id}/{month}/{year}").Methods("GET")

// func createProducts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")

// 	// get connect
// 	db := connect()
// 	defer db.Close()

// 	// creating product instance
// 	product := &Product{
// 		ID: uuid.New().String(),
// 	}

// 	// decoding request

// 	_ = json.NewDecoder(r.Body).Decode(&product)

// 	// inserting into database
// 	_, err := db.Model(product).Insert()
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	// returning product

// 	json.NewEncoder(w).Encode(product)

// }

// func getProducts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	db := connect()
// 	defer db.Close()

// 	params := mux.Vars(r)

// 	product_id := params["id"]
// 	products := Product{ID: product_id}

// 	if err := db.Model(&products).Select(); err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 	}

// 	json.NewEncoder(w).Encode(products)

// }

// func getProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	db := connect()
// 	defer db.Close()

// 	params := mux.Vars(r)

// 	product_id := params["id"]

// 	products := Product{ID: product_id}
// 	if err := db.Model(&products).WherePK().Select(); err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(products)

// }

// func updateProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	db := connect()
// 	defer db.Close()

// 	params := mux.Vars(r)

// 	product_id := params["id"]

// 	products := Product{ID: product_id}

// 	_ = json.NewDecoder(r.Body).Decode(&products)
// 	_, err := db.Model(products).WherePK().Set("Name = ?,Quantity = ?", products.Name, products.Quantity).Update()
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(products)

// }

// func deleteProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")

// 	db := connect()
// 	defer db.Close()

// 	params := mux.Vars(r)

// 	product_id := params["id"]

// 	products := &Product{ID: product_id}

// 	result, err := db.Model(products).Where("id = ?", product_id).Delete()

// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return

// 	}
// 	json.NewEncoder(w).Encode(result)

// }
