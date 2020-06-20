package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Day struct {
	Date          string  `json:"date"`
	Sleep         float64 `json:"sleep"`
	WeightMorning float64 `json:"weightMorning"`
	WeightEvening float64 `json:"weightEvening"`
	PurposeList   []Task  `json:"purposeList"`
	PleasureList  []Task  `json:"pleasureList"`
}

type Task struct {
	Done bool   `json:"done"`
	Name string `json:"name"`
}

type Collection struct {
	Coll *mongo.Collection
}

func (coll *Collection) setSleep(date string, sleep float64) {
	log.Println("setSleep")
	filter := bson.D{{"date", coll.getDay(date).Date}}
	_, err := coll.Coll.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"sleep": sleep}})
	if err != nil {
		log.Println(err)
	}
}

func (coll *Collection) setWeigtMorning(date string, weightMorning float64) {
	log.Println("setWeightMorning")
	filter := bson.D{{"date", coll.getDay(date).Date}}
	_, err := coll.Coll.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"weightMorning": weightMorning}})
	if err != nil {
		log.Println(err)
	}
}
func (coll *Collection) setWeightEvening(date string, weightEvening float64) {
	log.Println("setWeightEvening")
	filter := bson.D{{"date", coll.getDay(date).Date}}
	_, err := coll.Coll.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"weightEvening": weightEvening}})
	if err != nil {
		log.Println(err)
	}
}

// sets the pleasureList for the the day
func (coll *Collection) setPleasureList(date string, pleasureList []Task) {
	log.Println("setPleasureList")
	filter := bson.D{{"date", coll.getDay(date).Date}}
	var err error
	_, err = coll.Coll.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"pleasureList": pleasureList}})
	if err != nil {
		log.Println(err)
	}
}

// sets the pleasureList for the the day
func (coll *Collection) setPurposeList(date string, purposeList []Task) {
	log.Println("setPurposeList")
	filter := bson.D{{"date", coll.getDay(date).Date}}
	var err error
	_, err = coll.Coll.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"purposeList": purposeList}})
	if err != nil {
		log.Println(err)
	}
}

// createDay returns either a newly created day or an already existing one
func (coll *Collection) getDay(date string) Day {
	log.Println("getDay")
	filter := bson.D{{"date", date}}
	var result Day
	err := coll.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		// Error if no document found
		d := Day{Date: date}
		_, err = coll.Coll.InsertOne(context.Background(), d)
		if err != nil {
			// can this be fatal and only shut down curr goroutine?
			log.Println(err)
		}
		if d.PleasureList == nil {
			d.PleasureList = []Task{{false, ""}}
		}
		if d.PurposeList == nil {
			d.PurposeList = []Task{{false, ""}}
		}
		return d
	}
	if result.PleasureList == nil {
		result.PleasureList = []Task{{false, ""}}
	}
	if result.PurposeList == nil {
		result.PurposeList = []Task{{false, ""}}
	}
	return result
}
