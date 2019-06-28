package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	validator "github.com/gobuffalo/validate"
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

func convertStrNumToInt(strInt string) (int, error) {
	converted, err := strconv.ParseInt(strInt, 0, 0)

	if err != nil {
		return 0, errors.New("Cannot be converted to int")
	}

	return int(converted), nil
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
			log.Println("Connection error: ", err)
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

	err := encoder.Encode(map[string]interface{}{"httpCode": r.httpCode, "message": r.message, "data": r.data})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(map[string]interface{}{"httpCode": http.StatusInternalServerError, "message": "Internal server error", "data": nil})
		log.Println("Internal server error: ", err)
	}
}

// Persons persons table model
type Persons struct {
	ID          uint
	Name, Bio   string
	DateOfBirth int
	// going to save date of birth as unix timestamp
}

// PersonForm struct for person post requests
type PersonForm struct {
	Name, Bio, DateOfBirth string
}

// IsValid validates PersonForm struct
func (p *PersonForm) IsValid(errors *validator.Errors) {
	name := strings.TrimSpace(p.Name)
	bio := strings.TrimSpace(p.Bio)
	dateOfBirth := strings.TrimSpace(p.DateOfBirth)

	if len(name) < 5 {
		errors.Add("Name", "Name is too short")
	}

	if len(bio) < 5 {
		errors.Add("Bio", "Bio is too short")
	}

	if _, err := convertStrNumToInt(p.DateOfBirth); len(dateOfBirth) < 8 || err != nil {
		errors.Add("DateOfBirth", "Invalid birth of date provided")
	}
}

func handleGetAllPersons(w http.ResponseWriter, r *http.Request, conn Database) {
	db, _ := conn.connect()
	defer db.Close()

	var persons []Persons

	if db.Find(&persons); len(persons) == 0 {
		JSONResponse{http.StatusOK, "No record found", nil}.sendResponse(w)
		return
	}

	JSONResponse{http.StatusOK, "Success", persons}.sendResponse(w)
}

func handleGetPerson(w http.ResponseWriter, r *http.Request, conn Database) {
	db, _ := conn.connect()
	defer db.Close()

	var person Persons

	id, err := convertStrNumToInt(mux.Vars(r)["id"])
	if err != nil {
		JSONResponse{http.StatusNotFound, "Not found", nil}.sendResponse(w)
		return
	}

	if db.First(&person, id).RecordNotFound() {
		JSONResponse{http.StatusNotFound, "Not found", nil}.sendResponse(w)
		return
	}

	JSONResponse{http.StatusOK, "Success", person}.sendResponse(w)
}

func handleCreatePerson(w http.ResponseWriter, r *http.Request, conn Database) {
	db, _ := conn.connect()
	defer db.Close()

	formValues := PersonForm{
		Name:        r.FormValue("Name"),
		Bio:         r.FormValue("Bio"),
		DateOfBirth: r.FormValue("DateOfBirth"),
	}

	if err := validator.Validate(&formValues); err.Count() != 0 {
		JSONResponse{http.StatusBadRequest, "Bad request", err}.sendResponse(w)
		return
	}

	dateOfBirth, _ := convertStrNumToInt(formValues.Bio)
	newPerson := Persons{
		Name:        formValues.Name,
		Bio:         formValues.Bio,
		DateOfBirth: dateOfBirth,
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

	id, err := convertStrNumToInt(mux.Vars(r)["id"])
	if err != nil {
		JSONResponse{http.StatusNotFound, "Not found", nil}.sendResponse(w)
		return
	}

	if db.First(&person, id).RecordNotFound() {
		JSONResponse{http.StatusNotFound, "Not found", nil}.sendResponse(w)
		return
	}

	formValues := PersonForm{
		Name:        r.FormValue("Name"),
		Bio:         r.FormValue("Bio"),
		DateOfBirth: r.FormValue("DateOfBirth"),
	}

	if err := validator.Validate(&formValues); err.Count() != 0 {
		JSONResponse{http.StatusBadRequest, "Bad request", err}.sendResponse(w)
		return
	}

	person.Name = formValues.Name
	person.Bio = formValues.Bio
	dateOfBirth, _ := convertStrNumToInt(formValues.DateOfBirth)
	person.DateOfBirth = dateOfBirth

	if db.Save(&person).Error != nil {
		JSONResponse{http.StatusBadRequest, "Bad request", nil}.sendResponse(w)
		return
	}

	JSONResponse{http.StatusOK, "Success", person}.sendResponse(w)
}

func handleDeletePerson(w http.ResponseWriter, r *http.Request, conn Database) {
	db, _ := conn.connect()
	defer db.Close()

	var person Persons

	id, err := convertStrNumToInt(mux.Vars(r)["id"])
	if err != nil {
		JSONResponse{http.StatusNotFound, "Not found", nil}.sendResponse(w)
		return
	}

	result := db.First(&person, id)
	if result.RecordNotFound() {
		JSONResponse{http.StatusNotFound, "Not found", nil}.sendResponse(w)
		return
	}

	db.Delete(&person)

	JSONResponse{http.StatusOK, "Success", result.Value}.sendResponse(w)
}
