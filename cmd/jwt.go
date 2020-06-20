package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ID struct {
	ID string `json:"id"`
}

func (c *Controller) verifyJWT(jwt string) string {
	log.Println("verifyJWT")
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.AuthURL, nil)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("Authorization", jwt)
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	var id ID
	err = json.NewDecoder(resp.Body).Decode(&id)
	if err != nil {
		log.Panic(err)
	}
	return id.ID
}
