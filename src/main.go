package main


type question_struct struct{
	Id 			string `json:"id,omitempty"`
	Blob		string `json:"blob,omitempty"`
	Timestamp	string `json:"timestamp,omitempty"`
}

type answer_struct struct{
	Id 			string	`json:"id,omitempty"`
	Qid 		string	`json:"qid,omitempty"`
	Blob 		string	`json:"blob,omitempty"`
	Username	string	`json:"username,omitempty"`
	Timestamp 	string	`json:"timestamp,omitempty"`
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