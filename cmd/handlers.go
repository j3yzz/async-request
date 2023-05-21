package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Payload struct {
	Data any
}

func (app *Config) async(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	app.waitLock.Lock()

	response := make(map[string]any)
	response["success"] = true
	response["time"] = time.Now().UnixMilli()
	rc, _ := app.Redis.Get(app.context, "request_counter").Result()
	old_rc, _ := strconv.Atoi(rc)
	response["old_rc"] = old_rc
	app.Redis.Incr(app.context, "request_counter")
	rc, _ = app.Redis.Get(app.context, "request_counter").Result()
	new_rc, _ := strconv.Atoi(rc)
	response["new_rc"] = new_rc
	app.waitLock.Unlock()
	p := Payload{Data: response}

	json.NewEncoder(w).Encode(p)
}
