package main


import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var mongoClient *mongo.Client
var err error
var database *mongo.Database
var answersCollection *mongo.Collection
var questionsCollection *mongo.Collection


func initMongo() bool {
	//setting connection string
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoConfig.username, mongoConfig.password,
		mongoConfig.host, mongoConfig.port)
	if mongoConfig.username == "" && mongoConfig.password == "" {
		connectionString = fmt.Sprintf("mongodb://%s:%s",
			mongoConfig.host, mongoConfig.port)
	} else if mongoConfig.username == "" || mongoConfig.password == "" {
		fmt.Println("Please provide MONGO_USER and MONGO_PASSWORD")
		return false
	}
	fmt.Println("Connection String : " + connectionString)

	//connecting with mongo db
	mongoClient, err = mongo.Connect(context.Background(), connectionString)
	if err != nil {
		fmt.Println("Mongo connection error occured!")
		fmt.Printf("Connection String : %s\n", connectionString)
		return false
	}

	//setting database and collection
	database = mongoClient.Database(mongoConfig.database)
	answersCollection = database.Collection("answers")
	questionsCollection = database.Collection("questions")

return true
}

// insert user answers to database
func insertAnswers(ans []answer_struct) bool {
	var intf []interface{}
	for a := range ans {
		intf = append(intf, ans[a])
	}
	_, err := answersCollection.InsertMany(context.Background(), intf)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// get questions from database
func getQuestions() []question_struct{
	var allQns []question_struct
	cursor, err := questionsCollection.Find(context.Background(), bson.NewDocument())
	if err != nil {
		fmt.Println(err)
	}

	for cursor.Next(context.Background()) {
		var result question_struct
		cursor.Decode(&result)
		allQns = append(allQns, result)
	}
	return allQns
}

func insertQuestions(qns []question_struct) bool {
	var intf []interface{}
	for q := range qns {
		intf = append(intf, qns[q])
	}
	_, err := questionsCollection.InsertMany(context.Background(), intf)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func getAnswers(id string) []answer_struct{
	var ans []answer_struct
	cursor, err := answersCollection.Find(context.Background(), bson.NewDocument(bson.EC.String("qid", id)))
	if err != nil {
		fmt.Println(err)
	}

	for cursor.Next(context.Background()) {
		var result answer_struct
		cursor.Decode(&result)
		ans = append(ans, result)
	}
	return ans
}

// func getAnswers() map[question_struct][]answer_struct{
// 	var m map[question_struct][]answer_struct
// 	m = make(map[question_struct]answer_struct)

// 	m[{"id":"001","blob":"jkl","timestamp":"890"}] = [
// 	{"id":"001","qid":"001","blob":"jkl","username":"yuii","timestamp":"890"},
// 	]

// 	return m
// }

func getAll() map[string][]answer_struct{
	var m map[string][]answer_struct
	m = make(map[string][]answer_struct)

	
	// m["john"] = append(x["john"], "id":"001","qid":"001","blob":"jkl","username":"yuii","timestamp":"890")
 //    // x["john"] = append(x["john"], "value1")

    cursor, err := answersCollection.Find(context.Background(), bson.NewDocument(bson.EC.String("qid", "001")))
	if err != nil {
		fmt.Println(err)
	}

	for cursor.Next(context.Background()) {
		var result answer_struct
		cursor.Decode(&result)
		// ans = append(ans, result)
		m["john"] = append(m["john"], result)
	}

	return m
}