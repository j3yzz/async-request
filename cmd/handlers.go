package main

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	Data any
}

func (app *Config) async(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := make(map[string]any)
	response["success"] = true

	rc, _ := app.Redis.Get(app.context, "request_counter").Result()
	response["old_rc"] = rc

	app.Redis.Incr(app.context, "request_counter")
	rc, _ = app.Redis.Get(app.context, "request_counter").Result()
	response["new_rc"] = rc

	p := Payload{Data: response}

	json.NewEncoder(w).Encode(p)
}
