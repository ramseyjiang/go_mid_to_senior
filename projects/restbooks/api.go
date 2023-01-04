package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// var host = "http://localhost"
var port = "12345"
var connectionString = "root:12345678@tcp(127.0.0.1:3306)/go_web?charset=utf8&parseTime=True&loc=Local"

type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Mobile    int64  `json:"mobile,omitempty"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	apiRouter := router.PathPrefix("/api").Subrouter() // /api will give access to all the API endpoints

	// /api/health returns "health check passed"
	apiRouter.PathPrefix("/health").HandlerFunc(CheckHealth).Methods("GET")

	// /api/entries returns listing all the address book
	apiRouter.PathPrefix("/book/list").HandlerFunc(GetBooks).Methods("GET")

	// GET /api/book?id=1 returns the book with id 1.
	apiRouter.PathPrefix("/book/get").HandlerFunc(GetBookByID).Methods("GET")

	// POST /api/book creates an book
	apiRouter.PathPrefix("/book").HandlerFunc(CreateBook).Methods("POST")

	// PUT /api/book updates an book
	apiRouter.PathPrefix("/book").HandlerFunc(UpdateBook).Methods("PUT")

	// DELETE /api/book deletes an book
	apiRouter.PathPrefix("/book").HandlerFunc(DeleteBook).Methods("DELETE")

	// POST /api/upload-entries-CSV imports CSV into the database
	// apiRouter.PathPrefix("/upload-entries-csv").HandlerFunc(UploadEntriesThroughCSV).Methods("POST")

	// GET /api/download-entries-CSV exports CSV from the database
	// apiRouter.PathPrefix("/download-entries-csv").HandlerFunc(DownloadEntriesToCSV).Methods("GET")

	_ = http.ListenAndServe(":"+port, router)
}

func CheckHealth(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	_, err := fmt.Fprintf(writer, "health check passed")

	if err != nil {
		return
	}
}

// GetBooks : Get All Books
// URL : /entries
// Method: GET
// Output: JSON Encoded Entries object if found else JSON Encoded Exception.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			return
		}
	}(db)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not connect to the database")
		return
	}

	var books []User
	rows, err := db.Query("SELECT * from users;")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong.")
		return
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)

	for rows.Next() {
		var book User
		var id string
		var firstName sql.NullString
		var lastName sql.NullString
		var email sql.NullString
		var mobile sql.NullInt64

		_ = rows.Scan(&id, &firstName, &lastName, &email, &mobile)
		book.ID = id
		book.FirstName = firstName.String
		book.LastName = lastName.String
		book.Email = email.String
		book.Mobile = mobile.Int64
		books = append(books, book)
	}
	respondWithJSON(w, http.StatusOK, books)
}

// GetBookByID - Get Book By ID
// URL : /book?id=1
// Parameters: int id
// Method: GET
// Output: JSON Encoded Address Book object if found else JSON Encoded Exception.
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			return
		}
	}(db)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not connect to the database")
		return
	}

	id := r.URL.Query().Get("id")
	var firstName sql.NullString
	var lastName sql.NullString
	var email sql.NullString
	var mobile sql.NullInt64

	err = db.QueryRow("SELECT first_name, last_name, email, mobile from users where id=?", id).Scan(&firstName, &lastName, &email, &mobile)

	switch {
	case err == sql.ErrNoRows:
		respondWithError(w, http.StatusBadRequest, "No book found with the id="+id)
		return
	case err != nil:
		respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
		return
	default:
		var book User
		book.ID = id
		book.FirstName = firstName.String
		book.LastName = lastName.String
		book.Email = email.String
		book.Mobile = mobile.Int64
		respondWithJSON(w, http.StatusOK, book)
	}
}

// CreateBook - Create Book
// URL : /book
// Method: POST
// Body:
/*
 {
 "first_name":"John",
 "last_name":"Doe",
 "email":"john.doe@gmail.com",
 "mobile":"1234567890",
 }
*/
// Output: JSON Encoded Address Book Book object if created else JSON Encoded Exception.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not connect to the database")
		return
	}

	var book User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&book)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Decode met some problem occurred.")
		return
	}

	statement, err := db.Prepare("insert into users (id, first_name, last_name, email, mobile) values(?,?,?,?,?)")
	log.Println(err)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
		return
	}
	defer func(statement *sql.Stmt) {
		err1 := statement.Close()
		if err1 != nil {
			return
		}
	}(statement)

	res, err := statement.Exec(book.ID, book.FirstName, book.LastName, book.Email, book.Mobile)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "There was problem entering the book.")
		return
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 1 {
		id, _ := res.LastInsertId()
		book.ID = fmt.Sprint(id)
		respondWithJSON(w, http.StatusOK, book)
	}
}

// UpdateBook - Update Book
// URL : /book
// Method: PUT
// Body:
/*
 {
 "id":1,
 "first_name":"Krish",
 "last_name":"Bhanushali",
 "email":"krishsb2405@gmail.com",
 "mobile":"7798775575"
 }
*/
// Output: JSON Encoded Address Book Book object if updated else JSON Encoded Exception.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer func(db *sql.DB) {
		err1 := db.Close()
		if err1 != nil {
			return
		}
	}(db)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not connect to the database")
		return
	}

	var book User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&book)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
		return
	}

	statement, err := db.Prepare("update users set first_name=?, last_name=?, email=?, mobile=? where id=?")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
		return
	}

	defer func(statement *sql.Stmt) {
		err2 := statement.Close()
		if err2 != nil {
			return
		}
	}(statement)

	res, err1 := statement.Exec(book.FirstName, book.LastName, book.Email, book.Mobile, book.ID)
	if err1 != nil {
		respondWithError(w, http.StatusInternalServerError, "There was problem entering the book.")
		return
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 1 {
		respondWithJSON(w, http.StatusOK, book)
	}
}

// DeleteBook -  Delete Book By ID
// URL : /book?id=1
// Parameters: int id
// Method: DELETE
// Output: JSON Encoded Address Book object if found & deleted else JSON Encoded Exception.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer func(db *sql.DB) {
		err1 := db.Close()
		if err1 != nil {
			return
		}
	}(db)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not connect to the database")
		return
	}

	id := r.URL.Query().Get("id")
	var firstName sql.NullString
	var lastName sql.NullString
	var email sql.NullString
	var mobile sql.NullInt64
	err = db.QueryRow("SELECT first_name, last_name, email, mobile from users where id=?", id).Scan(&firstName, &lastName, &email, &mobile)

	switch {
	case err == sql.ErrNoRows:
		respondWithError(w, http.StatusBadRequest, "No book found with the id="+id)
		return
	case err != nil:
		respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
		return
	default:
		res, err1 := db.Exec("DELETE from users where id=?", id)
		if err1 != nil {
			respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
			return
		}

		count, err2 := res.RowsAffected()
		if err2 != nil {
			respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
			return
		}

		if count == 1 {
			var book User
			book.ID = id
			book.FirstName = firstName.String
			book.LastName = lastName.String
			book.Email = email.String
			book.Mobile = mobile.Int64

			respondWithJSON(w, http.StatusOK, book)

			return
		}
	}
}

/**
// UploadEntriesThroughCSV - Reads CSV, Parses the CSV and creates all the entries in the database
func UploadEntriesThroughCSV(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("csvFile")

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong while opening the CSV.")
		return
	}

	defer func(file multipart.File) {
		err3 := file.Close()
		if err3 != nil {
			return
		}
	}(file)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 5
	csvData, err := reader.ReadAll()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong while parsing the CSV.")
		return
	}

	var book User

	for _, row := range csvData {
		if row[1] != "first_name" {
			book.FirstName = row[1]
		}

		if row[2] != "last_name" {
			book.LastName = row[2]
		}

		if row[3] != "email" {
			book.EmailAddress = row[3]
		}

		if row[4] != "mobile" {
			book.PhoneNumber = row[4]
		}

		if book.FirstName != "" && book.LastName != "" && book.Email != "" && book.Mobile != "" {
			jsonString, err1 := json.Marshal(book)
			if err1 != nil {
				respondWithError(w, http.StatusBadRequest, "Something went wrong while parsing the CSV.")
				return
			}

			req, err2 := http.NewRequest("POST", host+":"+port+"/api/book", bytes.NewBuffer(jsonString))

			if err2 != nil {
				respondWithError(w, http.StatusBadRequest, "Something went wrong while requesting to the Creation endpoint.")
				return
			}

			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err3 := client.Do(req)

			if err3 != nil {
				respondWithError(w, http.StatusBadRequest, "Something went wrong while requesting to the Creation endpoint.")
				return
			}

			defer resp.Body.Close()

			if resp.Status == strconv.Itoa(http.StatusBadRequest) {
				respondWithError(w, http.StatusBadRequest, "Something went wrong while inserting.")
				return
			} else if resp.Status == strconv.Itoa(http.StatusInternalServerError) {
				respondWithError(w, http.StatusInternalServerError, "Something went wrong while inserting.")
				return
			}
		}
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"success": "Upload successful"})
}

// DownloadEntriesToCSV - GetAllEntries, creates a CSV and downloads the CSV.
func DownloadEntriesToCSV(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get(host + ":" + port + "/api/entries")

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Somehow host could not be reached.")
		return
	}

	defer response.Body.Close()

	var books []User

	data, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(data, &books)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to unmarshal data.")
		return
	}

	b := &bytes.Buffer{}
	t := time.Now().Unix()
	fileName := "address-book-" + strconv.Itoa(int(t)) + ".csv"
	writer := csv.NewWriter(b)
	heading := []string{"id", "first_name", "last_name", "email", "mobile"}
	_ = writer.Write(heading)

	for _, book := range books {
		var record []string
		_ = writer.Write(append(record, strconv.Itoa(book.ID), book.FirstName, book.LastName, book.EmailAddress, book.PhoneNumber))
	}

	writer.Flush()
	w.Header().Set("Content-Type", "text/csv") // setting the content type header to text/csv
	w.Header().Set("Content-Disposition", "attachment;filename="+fileName)
	w.WriteHeader(http.StatusOK)
	_, err1 := w.Write(b.Bytes())

	if err1 != nil {
		return
	}
}
*/

// RespondWithError is called on an error to return info regarding error
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Called for responses to encode and send json data
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// encode payload to json
	response, _ := json.Marshal(payload)

	// set headers and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)

	if err != nil {
		return
	}
}
