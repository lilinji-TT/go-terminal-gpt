package gpt

import (
	"GTG/config"
	"GTG/model"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func GenerateStreamWithGPT(prompt string, history *[]model.Message, modelName string) {
	*history = append(*history, model.Message{Role: "user", Content: prompt})

	request := &model.RequestBody{
		Model:    modelName,
		Messages: *history,
		Stream:   true,
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	baseURL, openaiKey, err := config.ReadConfig()
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openaiKey)
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		var response model.ErrorResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Println("Error:", err)
		}
		log.Println(response.Error.Message)
		return
	}

	reader := bufio.NewReader(resp.Body)
	var content string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Read Error: %s", err)
			return
		}
		if len(line) == 1 {
			continue
		}

		jsonString := strings.TrimPrefix(line, "data: ")
		if strings.Contains(jsonString, "DONE") {
			break
		}

		var message model.ResponseBodyJSON
		err = json.Unmarshal([]byte(jsonString), &message)
		if err != nil {
			log.Printf("Parse Error: %s", err)
			break
		}

		for _, char := range message.Choices[0].Delta.Content {
			content += string(char)
			fmt.Print(string(char))
			time.Sleep(10 * time.Millisecond)
		}
	}

	*history = append(*history, model.Message{Role: "assistant", Content: content})
}
