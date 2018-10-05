package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var err error
var answersCollection *mgo.Collection
var questionsCollection *mgo.Collection
var session mgo.Session

func initMongo() bool {
	fmt.Println(mongoConfig.database)
	//setting connection string
	connString := fmt.Sprintf("%s:%s", mongoConfig.host, mongoConfig.port)
	println(connString)
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{connString},
		Username: mongoConfig.username,
		Password: mongoConfig.password,
		Database: mongoConfig.database,
	})
	if err != nil {
		fmt.Println(err)
	}

	//setting database and collection
	answersCollection = session.DB(mongoConfig.database).C("answers")
	questionsCollection = session.DB(mongoConfig.database).C("questions")

	return true
}

// insert user answers to database
func insertAnswers(ans []answer_struct) bool {
	var intf []interface{}
	for a := range ans {
		intf = append(intf, ans[a])
	}

	err := answersCollection.Insert(intf...)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// get questions from database
func getQuestions(offset int, limit int) []question_struct {
	var allQns []question_struct
	query := questionsCollection.Find(bson.M{}).Skip(offset).Limit(limit)
	err := query.All(&allQns)
	if err != nil {
		fmt.Println(err)
		return []question_struct{}
	}
	if len(allQns) == 0 {
		return []question_struct{}
	}
	return allQns
}

func insertQuestions(qns []question_struct) bool {
	var intf []interface{}
	for a := range qns {
		intf = append(intf, qns[a])
	}

	err := questionsCollection.Insert(intf...)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func getAnswers(id string, offset int, limit int) []answer_struct {
	var ans []answer_struct
	queryInterface := make(map[string]string)
	queryInterface["qid"] = id
	query := answersCollection.Find(queryInterface).Skip(offset).Limit(limit)
	err := query.All(&ans)
	if err != nil {
		fmt.Println(err)
	}
	return ans
}

func getAll(offset int, limit int) []byte {

	values := []bson.M{}
	pipe := []bson.M{
		{"$lookup": bson.M{
			"from":         "answers",
			"localField":   "id",
			"foreignField": "qid",
			"as":           "answers",
		},
		},
	}
	pipe = append(pipe,
		bson.M{
			"$skip": offset,
		}, bson.M{
			"$limit": limit,
		})
	result := questionsCollection.Pipe(pipe)
	err := result.All(&values)
	if err != nil {
		fmt.Println(err)
	}

	jsString, err := bson.MarshalJSON(values)

	if err != nil {
		return []byte{}
	}
	return jsString

}

// getAll test function

// func testData() {
// 	qns := []question_struct{
// 		question_struct{
// 			"q1",
// 			"what is question 1?",
// 			"",
// 		},
// 		question_struct{
// 			"q2",
// 			"what is question 2?",
// 			"",
// 		},
// 		question_struct{
// 			"q3",
// 			"what is question 3?",
// 			"",
// 		},
// 		question_struct{
// 			"q4",
// 			"what is question 4?",
// 			"",
// 		},
// 	}
// 	insertQuestions(qns)

// 	ans := []answer_struct{
// 		answer_struct{
// 			"a0",
// 			"q1",
// 			"answer 0",
// 			"user 1",
// 			"",
// 		},
// 		answer_struct{
// 			"a1",
// 			"q3",
// 			"answer 1",
// 			"user 1",
// 			"",
// 		},
// 		answer_struct{
// 			"a2",
// 			"q4",
// 			"answer 2",
// 			"user 2",
// 			"",
// 		},
// 		answer_struct{
// 			"a3",
// 			"q3",
// 			"answer 3",
// 			"user 2",
// 			"",
// 		},
// 	}
// insertAnswers(ans)
// fmt.Println(string(getAll()))
// }
