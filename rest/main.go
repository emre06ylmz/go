package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//User type
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// Existing code from above
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", returnAllUsers)
	myRouter.HandleFunc("/user", createNewUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{id}", returnSingleUser)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUsername := "root"
	dbPassword := "12345"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "sys"

	db, err := sql.Open(dbDriver, dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllUsers")
	var userList []User
	db := dbConn()

	selDB, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}
	var user = new(User)

	for selDB.Next() {
		var id int
		var username, email, password string
		err = selDB.Scan(&id, &username, &email, &password)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Username = username
		user.Email = email
		user.Password = password
		userList = append(userList, *user)
	}
	defer db.Close()

	json.NewEncoder(w).Encode(userList)
}

func returnSingleUser(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	vars := mux.Vars(r)
	nID := vars["Id"]

	selDB, err := db.Query("SELECT * FROM user WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}
	user := User{}
	for selDB.Next() {
		var id int
		var username, email, password string
		err = selDB.Scan(&id, &username, &email, &password)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Username = username
		user.Email = email
		user.Password = password
	}

	json.NewEncoder(w).Encode(user)
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new User struct
	// append this to our Users array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)

	db := dbConn()
	insForm, err := db.Prepare("INSERT INTO user(username, email, password) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(user.Username, user.Email, user.Password)
	defer db.Close()
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	vars := mux.Vars(r)
	nID, err := strconv.Atoi(vars["Id"])
	delForm, err := db.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(nID)
	log.Println("DELETE")
	defer db.Close()
}
