package main


type question_struct struct{
	Id       string `json:"id,omitempty"`
	Question string `json:"question,omitempty"`
}

type answer_struct struct{
	Id 		string	`json:"id,omitempty"`
	Answer 	string	`json:"answer,omitempty"`
}

func main(){
	initMongo()
	initHttp()
/*
	var allAns []answer_struct
	ans1 := answer_struct{
		Id : "001",
		Answer : "answer 1",
	}
	allAns = append(allAns, ans1)
	allAns = append(allAns, ans1)
	allAns = append(allAns, ans1)
	insertAnswers(allAns)*/
}