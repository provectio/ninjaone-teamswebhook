package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
)

func sendTeams(webhookURL string, data RequestBody) {

	buffer := new(bytes.Buffer)

	err := tmpl.ExecuteTemplate(buffer, "base", data)
	if err != nil {
		log.Error("Error while executing template", "error", err)
		return
	}

	message := msftTeamsMessage{
		Title: data.Message,
		Text:  buffer.String(),
	}

	jsonObj, err := json.Marshal(message)
	if err != nil {
		log.Error("Failed serialize teams message", "error", err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(jsonObj))
	if err != nil {
		log.Error("Failed creating teams notification request", "error", err)
		return
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error("Failed sending teams message", "error", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		log.Error("Send request to teams failed, bad status code", "code", resp.StatusCode, "message", string(body))
	} else if string(body) != "1" {
		log.Error("Send request to teams failed", "message", string(body))
	}
}

type msftTeamsMessage struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
