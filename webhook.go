package main

import (
	"bytes"
	"encoding/json"
	"net/http"
    "os"
    "time"
)

type Embed struct {
	Title       string       `json:"title,omitempty"`
	Color       int          `json:"color,omitempty"`
	Fields      []EmbedField `json:"fields,omitempty"`
	Timestamp   string       `json:"timestamp,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type EmbedPayload struct {
	Embeds []Embed `json:"embeds"`
}

func sendAddition(res Reservation) error {
	embed := Embed{
		Title:       "新規掲載",
		Color:       5763719, // Green in decimal
		Fields: []EmbedField{
            {Name: "車", Value: res.Car, Inline: false},
			{Name: "出発店舗", Value: res.Departure, Inline: false},
			{Name: "到着地", Value: res.Arrival, Inline: false},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	body, _ := json.Marshal(EmbedPayload{Embeds: []Embed{embed}})
    webhookURL := os.Getenv("KATAMICHI_GO_WEBHOOK")
	resp, err := http.Post(webhookURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func sendDeletion(res Reservation) error {
	embed := Embed{
		Title:       "掲載終了",
		Color:       15548997, // Red in decimal
		Fields: []EmbedField{
			{Name: "車", Value: res.Car, Inline: false},
			{Name: "出発店舗", Value: res.Departure, Inline: false},
			{Name: "到着地", Value: res.Arrival, Inline: false},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	body, _ := json.Marshal(EmbedPayload{Embeds: []Embed{embed}})
    webhookURL := os.Getenv("KATAMICHI_GO_WEBHOOK")
	resp, err := http.Post(webhookURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
