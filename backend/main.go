package main

import (
	"fmt"
	"net/http"
	"os"
	"database/sql"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

func GetStoriesTitles() ([]string, error) {
	db, err := sql.Open("mysql", "root:secret@tcp(mysql:3306)/backend_dev")
	if err != nil {
		panic(err.Error())
	}

	table := "stories"
	query := fmt.Sprintf("SELECT title FROM %s;", table);

	res, err := db.Query(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to fetch stories' names from SQL database: %v\n", err)
		return []string{}, err
	}

	problems := []string{}
	var problem string
	for res.Next() {
		err = res.Scan(&problem)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning story: %v\n", err)
			return []string{}, err
		}
		problems = append(problems, problem)
	}

	return problems, nil
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	list, _ := GetStoriesTitles()

	json.NewEncoder(w).Encode(list)
}

func main() {
	fmt.Println("Server running on:", "localhost:8080")

	http.HandleFunc("/stories_titles", HandleRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Exiting")
}
