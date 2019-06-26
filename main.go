package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// initialize and declare routes
	router := mux.NewRouter()
	router.HandleFunc("/persons", withDbConnect(handleGetAllPersons)).Methods("GET")
	router.HandleFunc("/persons/{id}", withDbConnect(handleGetPerson)).Methods("GET")
	router.HandleFunc("/persons", withDbConnect(handleCreatePerson)).Methods("POST")
	router.HandleFunc("/persons/{id}", withDbConnect(handleUpdatePerson)).Methods("PUT")
	router.HandleFunc("/persons/{id}", withDbConnect(handleDeletePerson)).Methods("DELETE")

	startServer(8080, router)
}

func startServer(port int, router *mux.Router) {
	fmt.Println("Server is running")

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		log.Fatal("Starting server error: ", err)
	}
}

// Database database connection struct
type Database struct {
	user, pass, host, dbname string
}

func (db Database) connect() (*gorm.DB, error) {
	conn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", db.user, db.pass, db.host, db.dbname)

	database, err := gorm.Open("mysql", conn)
	if err != nil {
		log.Println("Database connection failed: ", err)
	}

	return database, err
}

func withDbConnect(fn func(http.ResponseWriter, *http.Request, Database)) http.HandlerFunc {
	conn := Database{"root", "", "localhost", "golang_backend_db"}

	return func(w http.ResponseWriter, r *http.Request) {
		testConn, err := conn.connect()
		defer testConn.Close()

		if err != nil {
			JSONResponse{http.StatusInternalServerError, "Internal server error", nil}.sendResponse(w)
			log.Println(err)
			return
		}

		fn(w, r, conn)
	}
}

// JSONResponse json response struct
type JSONResponse struct {
	httpCode int
	message  string
	data     interface{}
}

func (r JSONResponse) sendResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.httpCode)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(map[string]interface{}{"httpCode": r.httpCode, "message": r.message, "data": r.data}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(map[string]interface{}{"httpCode": http.StatusInternalServerError, "message": "Internal server error", "data": nil})
		log.Println("Internal server error: ", err)
	}
}

// Persons persons table model
type Persons struct {
	ID          uint
	Name, Bio   string
	DateOfBirth uint64
	// going to save date of birth as unix timestamp
}

func handleGetAllPersons(w http.ResponseWriter, r *http.Request, conn Database) {
	db, _ := conn.connect()
	defer db.Close()

	var persons []Persons

	if db.Find(&persons).RecordNotFound() {
		JSONResponse{http.StatusOK, "No record found", nil}.sendResponse(w)
		return
	}

	JSONResponse{http.StatusOK, "Success", persons}.sendResponse(w)
}

func handleGetPerson(w http.ResponseWriter, r *http.Request, conn Database) {
	db, _ := conn.connect()
	defer db.Close()

	var person Persons

	parameters := mux.Vars(r)
	if db.First(&person, parameters["id"]).RecordNotFound() {
		JSONResponse{http.StatusNotFound, "Not found", nil}.sendResponse(w)
		return
	}

	JSONResponse{http.StatusOK, "Success", person}.sendResponse(w)
}

func handleCreatePerson(w http.ResponseWriter, r *http.Request, conn Database) {
	db, _ := conn.connect()
	defer db.Close()

	name := r.FormValue("name")
	bio := r.FormValue("bio")
	dateOfBirth, _ := strconv.ParseUint(r.FormValue("dateOfBirth"), 0, 0)

	newPerson := Persons{Name: name, Bio: bio, DateOfBirth: dateOfBirth}

	if strings.TrimSpace(newPerson.Name) == "" || strings.TrimSpace(newPerson.Bio) == "" {
		JSONResponse{http.StatusBadRequest, "Bad request", nil}.sendResponse(w)
		return
	}

	db.NewRecord(newPerson)

	if dbs := db.Create(&newPerson); dbs.Error != nil {
		JSONResponse{http.StatusBadRequest, "Bad request", nil}.sendResponse(w)
		return
	}

	JSONResponse{http.StatusOK, "Success", newPerson}.sendResponse(w)
}

func handleUpdatePerson(w http.ResponseWriter, r *http.Request, conn Database) {
	db, _ := conn.connect()
	defer db.Close()

	var person Persons

	paramaters := mux.Vars(r)
	result := db.First(&person, paramaters["id"])
	if result.RecordNotFound() {
		JSONResponse{http.StatusNotFound, "Not found", nil}.sendResponse(w)
		return
	}

	person.Name = r.FormValue("name")
	person.Bio = r.FormValue("bio")
	dateOfBirth, _ := strconv.ParseUint(r.FormValue("dateOfBirth"), 0, 0)
	person.DateOfBirth = dateOfBirth

	if strings.TrimSpace(person.Name) == "" || strings.TrimSpace(person.Bio) == "" {
		JSONResponse{http.StatusBadRequest, "Bad request", nil}.sendResponse(w)
		return
	}

	if dbs := db.Save(&person); dbs.Error != nil {
		JSONResponse{http.StatusBadRequest, "Bad request", nil}.sendResponse(w)
		return
	}

	JSONResponse{http.StatusOK, "Success", person}.sendResponse(w)
}

func handleDeletePerson(w http.ResponseWriter, r *http.Request, conn Database) {
	db, _ := conn.connect()
	defer db.Close()

	var person Persons

	paramaters := mux.Vars(r)
	result := db.First(&person, paramaters["id"])
	if result.RecordNotFound() {
		JSONResponse{http.StatusNotFound, "Not found", nil}.sendResponse(w)
		return
	}

	db.Delete(&person)

	JSONResponse{http.StatusOK, "Success", result.Value}.sendResponse(w)
}
