package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
	DBClient *mongo.Client
	AuthURL  string
}

func (c *Controller) getDay(w http.ResponseWriter, r *http.Request) {
	log.Println("getDay")
	user := c.verifyJWT(r.Header.Get("Authorization"))
	date := r.URL.Query()["date"][0]
	coll := Collection{Coll: c.collection(user)}
	day, err := json.Marshal(coll.getDay(date))
	if err != nil {
		log.Panicln(err)
	}
	log.Println(string(day))
	w.Write(day)
}

func (c *Controller) setSleep(w http.ResponseWriter, r *http.Request) {
	user := c.verifyJWT(r.Header.Get("Authorization"))
	date := r.URL.Query()["date"][0]
	sleep, err := strconv.ParseFloat(r.URL.Query()["value"][0], 64)
	if err != nil {
		log.Panicln(err)
		return
	}
	coll := Collection{Coll: c.collection(user)}
	coll.setSleep(date, sleep)
}

func (c *Controller) setWeigtMorning(w http.ResponseWriter, r *http.Request) {
	user := c.verifyJWT(r.Header.Get("Authorization"))
	date := r.URL.Query()["date"][0]
	weightMorning, err := strconv.ParseFloat(r.URL.Query()["value"][0], 64)
	log.Println(weightMorning)
	if err != nil {
		log.Panicln(err)
		return
	}
	coll := Collection{Coll: c.collection(user)}
	coll.setWeigtMorning(date, weightMorning)
}

func (c *Controller) setWeightEvening(w http.ResponseWriter, r *http.Request) {
	user := c.verifyJWT(r.Header.Get("Authorization"))
	date := r.URL.Query()["date"][0]
	weightEvening, err := strconv.ParseFloat(r.URL.Query()["value"][0], 64)
	if err != nil {
		log.Panicln(err)
		return
	}
	coll := Collection{Coll: c.collection(user)}
	coll.setWeightEvening(date, weightEvening)
}

func (c *Controller) setPleasureList(w http.ResponseWriter, r *http.Request) {
	user := c.verifyJWT(r.Header.Get("Authorization"))
	tasks := getTaskList(w, r)
	date := r.URL.Query()["date"][0]
	coll := Collection{Coll: c.collection(user)}
	coll.setPleasureList(date, tasks)
}

func (c *Controller) setPurposeList(w http.ResponseWriter, r *http.Request) {
	user := c.verifyJWT(r.Header.Get("Authorization"))
	tasks := getTaskList(w, r)
	date := r.URL.Query()["date"][0]
	coll := Collection{Coll: c.collection(user)}
	coll.setPurposeList(date, tasks)
}

func getTaskList(w http.ResponseWriter, r *http.Request) []Task {
	var t []Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return []Task{}
	}
	return t
}

func (c *Controller) collection(user string) *mongo.Collection {
	return c.DBClient.Database("tasker").Collection(user)
}
