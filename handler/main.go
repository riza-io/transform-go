package main

import (
	"encoding/json"
	"log"
	"os"
)

type Request struct {
	Headers map[string]string `json:"headers"`
	Path    string            `json:"path"`
	Query   string            `json:"query"`
	Body    json.RawMessage   `json:"body"`
}

type ConsoleLine struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
}

type HandlerResponse struct {
	Request Request       `json:"request"`
	Lines   []ConsoleLine `json:"console"`
}

func main() {
	resp := HandlerResponse{
		Request: Request{
			Headers: map[string]string{
				"X-Test": "Go",
			},
			Body: []byte("{}"),
		},
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(&resp); err != nil {
		log.Fatal(err)
	}
}
