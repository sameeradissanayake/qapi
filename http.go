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

func initHttp() {
	router := mux.NewRouter()
	router.HandleFunc("/questions", GetQuestion).Methods("GET")
	router.HandleFunc("/answers", SaveAnswer).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}