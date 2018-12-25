package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cryptix/wav"
	"github.com/google/uuid"
)

const subscriptionKey = "1ce031d210884246b0f0bc2a8783ba3e"
const locale = "en-US"
const format = "json"

func stt(client *http.Client) {
	req, _ := http.NewRequest("POST", "https://speech.platform.bing.com/speech/recognition/interactive/cognitiveservices/v1", nil)

	requestUUID := uuid.New()
	q := req.URL.Query()
	q.Add("language", locale)
	q.Add("locale", locale)
	q.Add("format", format)
	q.Add("requestid", requestUUID.String())
	req.URL.RawQuery = q.Encode()
	log.Printf("[debug] %s", req.URL.String())
}

func getKey(client *http.Client) {
	req, err := http.NewRequest("POST", "https://api.cognitive.microsoft.com/sts/v1.0/issueToken", nil)
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", "0")
	req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error %v", err)
	}
	log.Printf("%v", resp)
}

func main() {
	client := &http.Client{}

	fileInfo, _ := os.Stat("output.wav")
	wavFile, _ := os.Open("output.wav")
	wavReader, _ := wav.NewReader(wavFile, fileInfo.Size())

	stt(client)
	fmt.Println(wavReader)
}
