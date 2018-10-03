package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

//GetQuestion - get questions
func GetQuestion(w http.ResponseWriter, req *http.Request) {
	allQuestions := getQuestions()
	json.NewEncoder(w).Encode(&allQuestions)
}

//SaveAnswer - save user answer
func SaveAnswer(w http.ResponseWriter, req *http.Request) {
	var allAnswers []answer_struct
	err= json.NewDecoder(req.Body).Decode(&allAnswers)
	if err != nil{
		fmt.Println(err)
		fmt.Println(req.Body)
	}
	if insertAnswers(allAnswers){
		w.Write([]byte("{result: 'OK'}"))
		return
	}
	w.Write([]byte("{result: 'ERROR'}"))
}

func PostQuestions(w http.ResponseWriter, req *http.Request) {
	var allQuestions []question_struct
	err= json.NewDecoder(req.Body).Decode(&allQuestions)
	if err != nil{
		fmt.Println(err)
		fmt.Println(req.Body)
	}
	if insertQuestions(allQuestions){
		w.Write([]byte("{result: 'OK'}"))
		return
	}
	w.Write([]byte("{result: 'ERROR'}"))
}

func GetAnswers(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	fmt.Println(params["id"])

	answers := getAnswers(params["id"])
	json.NewEncoder(w).Encode(&answers)
}

func GetAll(w http.ResponseWriter, req *http.Request) {
	allQuestions := getAll()
	json.NewEncoder(w).Encode(&allQuestions)
}

func initHttp() {
	router := mux.NewRouter()
	router.HandleFunc("/getquestions", GetQuestion).Methods("GET")
	router.HandleFunc("/postanswers", SaveAnswer).Methods("POST")
	router.HandleFunc("/postquestions", PostQuestions).Methods("POST")
	router.HandleFunc("/getanswers/{id}", GetAnswers).Methods("GET")
	router.HandleFunc("/getall", GetAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}