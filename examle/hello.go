// package main

// import (
// 	"net/http"
// )

// // func main() {
// // 	fmt.Println("Hello World")
// // 	var x int
// // 	x = 2 * 4
// // 	fmt.Println(x)

// // 	y := 45
// // 	z := 3
// // 	sum := y + z
// // 	fmt.Println(sum)

// // 	a := [5]int{1, 2, 3, 4, 7}
// // 	a[2] = 2
// // 	b := []int{2, 3, 1, 6}
// // 	b = append(b, 13)
// // 	fmt.Println(b)
// // 	k := summ(2, 3)
// // 	fmt.Println(k)
// // }

// // func summ(x int, y int) int {
// // 	return x + y
// // }
// func main() {
// 	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("hello world"))
// 	})
// 	http.ListenAndServe(":8080", nil)
// }

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

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to meeting scheduler!")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/meetings", createSchedule).Methods("POST")
	router.HandleFunc("/meeting/{id}", getOneMeet).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
