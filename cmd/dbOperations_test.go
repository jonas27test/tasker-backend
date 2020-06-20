package main

import (
	"context"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	date = "01"
)

func Test_SetSleep(t *testing.T) {
	sleep := 8.0
	test := dbConn(t, "test")
	coll := Collection{Coll: test}
	coll.setSleep(date, sleep)

	d := coll.getDay(date)
	if d.Date != date || d.Sleep != sleep {
		t.Fatal()
	}

	cleanup(coll.Coll)
}

func Test_AddPList_Purpose(t *testing.T) {
	list := []Task{{Done: false, Name: "test"}}
	test := dbConn(t, "test")
	coll := Collection{Coll: test}
	coll.setPurposeList(date, list)

	d := coll.getDay(date)
	if d.Date != date || len(d.PurposeList) == 0 || d.PurposeList[0].Name != "test" {
		t.Fatal()
	}

	cleanup(coll.Coll)
}

func Test_AddPList_Pleasure(t *testing.T) {
	collection := "e2e"
	list := []Task{{Done: false, Name: "test"}}
	test := dbConn(t, collection)
	coll := Collection{Coll: test}
	coll.setPleasureList(date, list)

	d := coll.getDay(date)
	t.Error(d)
	if d.Date != date || len(d.PleasureList) == 0 || d.PleasureList[0].Name != "test" {
		t.Fatal()
	}

	cleanup(coll.Coll)
}
func Test_GetDay(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	day := createDay()
	test := dbConn(t, "test")
	coll := Collection{Coll: test}
	d := coll.getDay(day.Date)
	if day.Date != d.Date {
		t.Fatal()
	}
	t.Log(d)

	// cleanup(coll.c)
}

func createDay() Day {
	date := "2020-01-01"
	return Day{Date: date}
}

func dbConn(t *testing.T, collName string) *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://0.0.0.0"))
	if err != nil {
		t.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	return client.Database("test1").Collection(collName)
}

func cleanup(coll *mongo.Collection) {
	coll.Drop(context.Background())
}
