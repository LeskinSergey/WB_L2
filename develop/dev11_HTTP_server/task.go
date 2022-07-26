package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	PORT = ":8080"
)

var events = NewM()

type Event struct {
	NameEvent string `json:"NameEvent"`
	Id        string `json:"Id"`
	Date      string `json:"Date"`
	time      time.Time
}

type DefenseMap struct {
	sync.RWMutex
	m map[int][]Event
}

func NewM() *DefenseMap {
	return &DefenseMap{
		m: make(map[int][]Event),
	}
}
func (d *DefenseMap) Add(key int, value Event) {
	d.Lock()
	defer d.Unlock()
	d.m[key] = append(d.m[key], value)
}
func (d *DefenseMap) Delete(key int, name string) {
	d.Lock()
	defer d.Unlock()
	for i, val := range d.m[key] {
		if val.NameEvent == name {
			d.m[key] = removeElementByIndex(d.m[key], i)
		}
	}
}
func (d *DefenseMap) Get(key int) []Event {
	d.Lock()
	defer d.Unlock()
	return d.m[key]
}
func (d *DefenseMap) Update(e Event, id int) error {
	d.Lock()
	defer d.Unlock()
	for i, val := range d.m[id] {
		if val.NameEvent == e.NameEvent {
			d.m[id][i].Date = e.Date
			return nil
		}
	}
	return fmt.Errorf("event not found")
}
func removeElementByIndex(a []Event, i int) []Event {
	return append(a[:i], a[i+1:]...)
}
func JSONError(text string) map[string]string {
	return map[string]string{"error": text}
}
func respondWithJSON(w http.ResponseWriter, code int, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal server error", 500)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(`"result" : "` + string(bytes) + `"`))
}

func parseBody(body io.ReadCloser, data *Event) error {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, data); err != nil {
		return err
	}
	return nil
}
func validate(event *Event) error {
	id, err := strconv.Atoi(event.Id)
	if err != nil {
		return err
	}
	if id <= 0 || event.NameEvent == "" {
		return fmt.Errorf("400")
	}
	event.time, err = time.Parse("2006-01-02", event.Date)
	if err != nil {
		return err
	}
	return nil
}
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := parseBody(r.Body, &event); err != nil {
		respondWithJSON(w, 400, JSONError("bad parameters"))
		return
	}
	if err := validate(&event); err != nil {
		respondWithJSON(w, 400, JSONError("bad validation"))
		return
	}
	id, err := strconv.Atoi(event.Id)
	if err != nil {
		respondWithJSON(w, 400, err.Error())
		return
	}
	if _, ok := events.m[id]; !ok {
		events.m[id] = make([]Event, 0)
	}
	events.Add(id, event)
	respondWithJSON(w, 200, event)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := parseBody(r.Body, &event); err != nil {
		respondWithJSON(w, 400, JSONError("bad parameters"))
		return
	}
	if err := validate(&event); err != nil {
		respondWithJSON(w, 400, JSONError("bad validation"))
		return
	}
	id, err := strconv.Atoi(event.Id)
	if err != nil {
		respondWithJSON(w, 400, err.Error())
		return
	}
	err = events.Update(event, id)
	if err != nil {
		respondWithJSON(w, 400, err.Error())
		return
	}
	respondWithJSON(w, 200, event)
}
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	if r.Method != http.MethodPost && r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := parseBody(r.Body, &event); err != nil {
		respondWithJSON(w, 400, JSONError("bad parameters"))
		return
	}
	if err := validate(&event); err != nil {
		respondWithJSON(w, 400, JSONError("bad validation"))
		return
	}
	id, err := strconv.Atoi(event.Id)
	if err != nil {
		respondWithJSON(w, 400, err.Error())
		return
	}
	events.Delete(id, event.NameEvent)
	respondWithJSON(w, 200, event)
}

func respondGETWithJson(w http.ResponseWriter, id int, typeHandle time.Duration) {
	goodEvents := make([]Event, 0)
	dayTime := time.Now().Add(time.Hour * 24 * typeHandle)
	for _, val := range events.m[id] {
		if val.time.Before(dayTime) {
			goodEvents = append(goodEvents, val)
		}
	}

	res := make([]byte, 0)
	for _, val := range goodEvents {
		tmp, _ := json.Marshal(val)
		res = append(res, tmp...)
	}
	w.Write([]byte(`"result" : "` + string(res) + `"`))
}
func EventsForDay(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	if _, ok := m["id"]; !ok {
		respondWithJSON(w, 400, JSONError("not found user id"))
		return
	}
	id, err := strconv.Atoi(m["id"][0])
	if err != nil {
		respondWithJSON(w, 400, JSONError(err.Error()))
		return
	}
	respondGETWithJson(w, id, time.Duration(1))
}
func EventsForWeek(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	if _, ok := m["id"]; !ok {
		respondWithJSON(w, 400, JSONError("not found user id"))
		return
	}
	id, err := strconv.Atoi(m["id"][0])
	if err != nil {
		respondWithJSON(w, 400, JSONError(err.Error()))
		return
	}
	respondGETWithJson(w, id, time.Duration(7))
}
func EventsForMonth(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	if _, ok := m["id"]; !ok {
		respondWithJSON(w, 400, JSONError("not found user id"))
		return
	}
	id, err := strconv.Atoi(m["id"][0])
	if err != nil {
		respondWithJSON(w, 400, JSONError(err.Error()))
		return
	}
	respondGETWithJson(w, id, time.Duration(30))
}
func MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println(
			"method", r.Method,
			"path", r.URL.EscapedPath(),
			"duration", time.Since(start),
		)
		next(w, r)
	}
}

func main() {
	//POST
	http.HandleFunc("/create_event", MiddlewareLogger(CreateEvent))
	http.HandleFunc("/update_event", MiddlewareLogger(UpdateEvent))
	http.HandleFunc("/delete_event", MiddlewareLogger(DeleteEvent))
	//GET
	http.HandleFunc("/events_for_day", MiddlewareLogger(EventsForDay))
	http.HandleFunc("/events_for_week", MiddlewareLogger(EventsForWeek))
	http.HandleFunc("/events_for_month", MiddlewareLogger(EventsForMonth))
	fmt.Println("Server listen on", PORT)
	log.Println(http.ListenAndServe(PORT, nil))
}
