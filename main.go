package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Expense struct {
	ID          string       `json:"id.omnitempty"`
	Description string       `json:"description.omnitempty"`
	TypeExpense string       `json:"typeexpense.omnitempty"`
	Date        *DateExpense `json:"date.omnitempty`
}

type DateExpense struct {
	Month string `json:"month.omnitempty"`
	Year  string `json:"year.omnitempty"`
}

var expenses []Expense

func GetExpenses(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(expenses)
}
func GetExpense(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range expenses {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
	json.NewEncoder(res).Encode(&Expense{})
}
func CreateExpense(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var expense Expense
	_ = json.NewDecoder(req.Body).Decode(&expense)
	expense.ID = params["id"]
	expenses = append(expenses, expense)
	json.NewEncoder(res).Encode(expense)
}

func DeleteExpense(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range expenses {
		if item.ID == params["id"] {
			expenses = append(expenses[:index], expenses[index+1:]...)
			break
		}
		json.NewEncoder(res).Encode(expenses)
	}
}

func main() {
	// seed
	expenses = append(expenses, Expense{ID: "1", Description: "Pizza", TypeExpense: "Out", Date: &DateExpense{Month: "Jun", Year: "2020"}})

	router := mux.NewRouter()
	router.HandleFunc("/expenses", GetExpenses).Methods("GET")
	router.HandleFunc("/expenses/{id}", GetExpense).Methods("GET")
	router.HandleFunc("/expenses/{id}", CreateExpense).Methods("POST")
	router.HandleFunc("/expenses/{id}", DeleteExpense).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
