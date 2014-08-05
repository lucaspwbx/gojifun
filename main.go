package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

type Actor struct {
	Id   int
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var actors []Actor

func init() {
	actors = append(actors, Actor{1, "Lucas", 29})
	actors = append(actors, Actor{2, "Xiao", 30})
}

func NewActor(res http.ResponseWriter, req *http.Request) {
	var actor Actor
	err := json.NewDecoder(req.Body).Decode(&actor)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	actors = append(actors, actor)
	url := fmt.Sprintf("/actors/%d", len(actors))
	http.Redirect(res, req, url, http.StatusCreated)
}

func GetActors(res http.ResponseWriter, req *http.Request) {
	data, err := json.Marshal(actors)
	if err != nil {
		log.Println("Error marshalling JSON")
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(data)
}

func GetActor(c web.C, res http.ResponseWriter, req *http.Request) {
	var data []byte
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		log.Fatal("Error converting to integer")
	}
	for _, v := range actors {
		if v.Id == id {
			data, err = json.Marshal(v)
			res.Header().Set("Content-Type", "application/json")
			res.Write(data)
			break
		}
	}
}

func UpdateActor(c web.C, res http.ResponseWriter, req *http.Request) {
	var actor Actor
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		log.Fatal("Error converting to integer")
	}
	err := json.NewDecoder(req.Body).Decode(&actor)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	actors[i-1] = actor
}

func main() {
	goji.Get("/", GetActors)
	goji.Get("/actors/:id", GetActor)
	goji.Post("/actors", NewActor)
	goji.Put("/actors/:id", UpdateActor)
	goji.Serve()
}
