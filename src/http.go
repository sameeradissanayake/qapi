package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetQuestion - get questions
func GetQuestion(w http.ResponseWriter, req *http.Request) {
	queryValues := req.URL.Query()

	// offset
	var offset int
	offset, err := strconv.Atoi(queryValues.Get("offset"))
	if err != nil {
		offset = 0
	}

	// limit
	var limit int
	limit, err = strconv.Atoi(queryValues.Get("limit"))
	if err != nil {
		limit = 10
	}

	allQuestions := getQuestions(offset, limit)
	json.NewEncoder(w).Encode(&allQuestions)
}

//SaveAnswer - save user answer
func SaveAnswer(w http.ResponseWriter, req *http.Request) {
	var allAnswers []answer_struct
	err = json.NewDecoder(req.Body).Decode(&allAnswers)
	if err != nil {
		fmt.Println(err)
		fmt.Println(req.Body)
	}
	if insertAnswers(allAnswers) {
		w.Write([]byte("{result: 'OK'}"))
		return
	}
	w.Write([]byte("{result: 'ERROR'}"))
}

func PostQuestions(w http.ResponseWriter, req *http.Request) {
	var allQuestions []question_struct
	err = json.NewDecoder(req.Body).Decode(&allQuestions)
	if err != nil {
		fmt.Println(err)
		fmt.Println(req.Body)
	}
	if insertQuestions(allQuestions) {
		w.Write([]byte("{result: 'OK'}"))
		return
	}
	w.Write([]byte("{result: 'ERROR'}"))
}

func GetAnswers(w http.ResponseWriter, req *http.Request) {
	queryValues := req.URL.Query()

	// offset
	var offset int
	offset, err := strconv.Atoi(queryValues.Get("offset"))
	if err != nil {
		offset = 0
	}

	// limit
	var limit int
	limit, err = strconv.Atoi(queryValues.Get("limit"))
	if err != nil {
		limit = 10
	}

	answers := getAnswers(queryValues.Get("id"), offset, limit)
	json.NewEncoder(w).Encode(&answers)
}

func GetAll(w http.ResponseWriter, req *http.Request) {

	queryValues := req.URL.Query()

	// offset
	var offset int
	offset, err := strconv.Atoi(queryValues.Get("offset"))
	if err != nil {
		offset = 0
	}

	// limit
	var limit int
	limit, err = strconv.Atoi(queryValues.Get("limit"))
	if err != nil {
		limit = 10
	}

	allQuestions := getAll(offset, limit)
	w.Write(allQuestions)
}

func initHttp() {
	router := mux.NewRouter()
	router.HandleFunc("/getquestions", GetQuestion).Methods("GET")
	router.HandleFunc("/postanswers", SaveAnswer).Methods("POST")
	router.HandleFunc("/postquestions", PostQuestions).Methods("POST")
	router.HandleFunc("/getanswers", GetAnswers).Methods("GET")
	router.HandleFunc("/getall", GetAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
