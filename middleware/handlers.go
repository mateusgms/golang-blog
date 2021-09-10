package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-blog/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
type response struct {
	ID int64 `json: "id"`
	Message string `json: "message,omitempty"`
}

func createConnection() *sql.DB {
	err:= godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[Error] loading .env file")
	}

	db,err =: sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("[info] Successfully connected!")
	return db
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User

	err:= json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("[Error] Unable to decode the request body. %v", err)
	}

	insertID:= insertUser(user)

	res := response{
		ID: insertID,
		Message: "User created successfully"
	}

	json.NewEncoder(w).Encode(res)
}

func GetUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err:= strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("[Error] Unable to convert the string into int. %v", err)
	}

	json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")

	users, err := getAllUsers()
	if err != nil {
		log.Fatalf("[Error] Unable to get all users %v", err)
	}

	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("[Error] Unable to convert the string into int.  %v", err)
	}

	var user models.User

	err = json.NewDecoder(r.body).Decode(&user)
	if err != nil {
		log.Fatalf("[Error] Unable to decode the request body. %v", err)
	}

	updatedRows := updateUser(int64(id), user)

	msg := fmt.Sprintf("[Info] User updated successfully. Total rows/record affected %v", updatedRows)

	res:= response{
		ID: int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}
func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("[Error] Unable to convert the string into int. %v", err)
	}
	deletedRows := deleteUser(int64(id))

	msg := fmt.Sprintf("[Info] User updated successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID: int64(id),
		Message: msg
	}
	json.NewEnconder(w).Encode(res)
}

func insertUser(user models.User) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO users (name, email, password, role)`

	var id int64

	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Password, user.Role).Scan(&id)
	if err != nil {
		log.Fatalf("[Error] Unable to execute the query %v", err)
	}

	fmt.Printlf("[Info] Inserted a single record %v", id)

	return id
}
