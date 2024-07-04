package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Ultra-Smork/Ai_chat-bot/database"
	"github.com/go-chi/chi/v5"
)

type Request struct {
	Messages      []*Message `json:"messages"`
	System_prompt string     `json:"system_prompt,omitempty"`
	Temperature   float64    `json:"temperature"`
	Top_k         int64      `json:"top_k"`
	Top_p         float64    `json:"top_p"`
	Max_tokens    int64      `json:"max_tokens"`
	Web_access    bool       `json:"web_access"`
}

// check
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	Result      string `json:"result"`
	Status      bool   `json:"status"`
	Server_code int64  `json:"server_code"`
}

type Send_back struct {
	Question string `json:"request"`
	Answ     string `json:"response"`
}

// get message from frontend insert into database
func post_message(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	payload := "{\"messages\":[{\"role\":\"user\",\"content\":\"" + req.Messages[0].Content + "\"}],\"system_prompt\":\"\",\"temperature\":0.9,\"top_k\":5,\"top_p\":0.9,\"max_tokens\":256,\"web_access\":false}"
	var api_rngjesus Response
	api_rngjesus, err := get_response(payload)
	if err != nil {
		log.Fatal("PROBLEMS WITH API")
	}

	_, _, err = postgres.Save(req.Messages[0].Content, api_rngjesus.Result)
	if err != nil {
		fmt.Println("PROBLEMS WITH SAVE", err)
		return
	}
	ended := Send_back{
		Question: req.Messages[0].Content,
		Answ:     api_rngjesus.Result,
	}
	json.NewEncoder(w).Encode(ended)
}

// request to api with info from database

func get_response(prompt string) (Response, error) {

	url := "https://chatgpt-42.p.rapidapi.com/conversationgpt4-2"

	req, err := http.NewRequest("POST", url, strings.NewReader(prompt))
	if err != nil {
		return Response{}, err
	}

	req.Header.Add("x-rapidapi-key", "4454181f24mshfedd4d0e0d621b5p10a385jsn8d05d36ad4b9")
	req.Header.Add("x-rapidapi-host", "chatgpt-42.p.rapidapi.com")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Response{}, err
	}

	defer res.Body.Close()
	var middlething Response
	err = json.NewDecoder(res.Body).Decode(&middlething)
	if err != nil {
		return Response{}, err
	}

	return middlething, nil
}

func main() {

	err := postgres.Create()
	if err != nil {
		log.Fatal("PROBLEMS WITH DB", err)
	}

	router := chi.NewRouter()
	router.Post("/chat", post_message)

	log.Fatal(http.ListenAndServe(":8080", router))
}
