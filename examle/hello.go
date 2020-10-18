package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type meeting struct {
	ID                string    `json:"ID"`
	Title             string    `json:"Title"`
	Participants      int       `json:"Partcipants"`
	StartTime         time.Time `json:"StartTime"`
	EndTime           time.Time `json:"EndTime"`
	CreationTimestamp time.Time `json:"CreationTime"`
}

type part struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	RSVP  string `json:"RSVP"`
}

type allMeet []meeting
type allPart []part

var parts = allPart{
	{
		Name:  "Sayak",
		Email: "abc.sayak@gmail.com",
		RSVP:  "Yes",
	},
}

var meets = allMeet{
	{
		ID:                "1",
		Title:             "Learning GO",
		Participants:      10,
		StartTime:         time.Date(2020, 11, 14, 10, 45, 16, 0, time.UTC),
		EndTime:           time.Date(2020, 11, 15, 10, 45, 16, 0, time.UTC),
		CreationTimestamp: time.Date(2020, 8, 15, 10, 45, 16, 0, time.UTC),
	},
}

//creating a schedule
func createSchedule(w http.ResponseWriter, r *http.Request) {
	var newMeet meeting
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the meet title, participants, start_time and end_time only in order to update")
	}

	json.Unmarshal(reqBody, &newMeet)
	meets = append(meets, newMeet)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newMeet)
}

//getting one meet using an ID
func getOneMeet(w http.ResponseWriter, r *http.Request) {
	meetID := mux.Vars(r)["id"]

	for _, singleMeet := range meets {
		if singleMeet.ID == meetID {
			json.NewEncoder(w).Encode(singleMeet)
		}
	}
}

func getAllMeets(w http.ResponseWriter, r *http.Request) {
	meetStart := mux.Vars(r)["startTime"]
	meetEnd := mux.Vars(r)["endTime"]
	layout := "0001-01-01T00:00:00Z"
	tstart, err := time.Parse(layout, meetStart)
	tend, err := time.Parse(layout, meetEnd)
	if err != nil {
		fmt.Println(err)
	}
	for _, singleMeet := range meets {
		if singleMeet.StartTime == tstart && singleMeet.EndTime == tend {
			json.NewEncoder(w).Encode(singleMeet)
		}
	}
}

func allMeetPart(w http.ResponseWriter, r *http.Request) {

}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to meeting scheduler!")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/meetings", createSchedule).Methods("POST")
	router.HandleFunc("/meeting/{id}", getOneMeet).Methods("GET")
	router.HandleFunc("/meetings?start={startTime}&end={endTime}", getAllMeets).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// http://localhost:8080/meetings?meetings?start=2020-11-14T10:45:16Z&end=2020-11-15T10:45:16Z
