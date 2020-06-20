package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	dbURL := flag.String("dburl", "mongodb://0.0.0.0:27017", "sets the urls where to connect to the db.")
	authURL := flag.String("authurl", "http://0.0.0.0:8000", "sets the auth url.")
	port := flag.String("p", ":8080", "sets the port.")
	flag.Parse()
	log.SetFlags(log.Lshortfile)
	log.Println("dburl is " + *dbURL)
	log.Println("authurl is " + *authURL)
	log.Println("port is " + *port)

	client, err := mongo.NewClient(options.Client().ApplyURI(*dbURL))
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo.tasker"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to mongoDB!")

	c := Controller{DBClient: client, AuthURL: *authURL + "/userid"}

	http.HandleFunc("/getDay", c.getDay)
	http.HandleFunc("/setPurposeList", c.setPurposeList)
	http.HandleFunc("/setPleasureList", c.setPleasureList)
	http.HandleFunc("/setWeightMorning", c.setWeigtMorning)
	http.HandleFunc("/setWeightEvening", c.setWeightEvening)
	http.HandleFunc("/setSleep", c.setSleep)
	http.HandleFunc("/healthz", healthz)
	log.Fatal(http.ListenAndServe(*port, nil))
}

/*

func m(){

	joe := client.Database("test").Collection("joe")

	var result Day

	// d := Day{Date: 1, MorningWeight: 0.1, EveningWeight: 0.1}

	// _, err = joe.InsertOne(context.Background(), d)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	filter := bson.D{{"date", 3}}

	err = joe.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println("hs")
		log.Fatal(err)
	}
	log.Println(result)

	var update []PObject

	copy(update, result.PleasureList)

	new := PObject{Task: "Do", Done: false}
	update = append(update, new)
	log.Println(result)
	result.PleasureList = update

	joe.UpdateOne(context.Background(), filter, bson.M{"$push": bson.M{"test.joe.pleasureList": new}})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("success")
	}

	fmt.Printf("Found a single document: %+v\n", result)

}
*/
