package gpt

import (
	"GTG/config"
	"GTG/model"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)




func GenerateStreamWithGPT(prompt string, history *[]model.Message) {
	*history = append(*history, model.Message{Role: "user", Content: prompt})
	request := &model.RequestBody{
		Model: "gpt-3.5-turbo",
		Messages: *history,
		Stream: true,
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	BaseUrl, OpenaiKey, err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	req, err := http.NewRequest("POST", BaseUrl, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ OpenaiKey)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	var content string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Read Error: %s", err)
			return
		}
		if len(line) == 1 {
			// 忽略空行
			continue
		}

		var message model.ResponseBodyJSON
		jsonString := strings.TrimPrefix(line, "data: ")

		if strings.Contains(jsonString, "DONE") {
			break
		}

		err = json.Unmarshal([]byte(jsonString), &message)
		if err != nil {
			fmt.Printf("Parse Error: %s", err)
			break
		}

		for _, char := range message.Choices[0].Delta.Content {
			content += string(char)
			fmt.Print(string(char))
			time.Sleep(10 * time.Millisecond)
		}
	}

	//追加历史记录
	*history = append(*history, model.Message{Role: "assistant", Content: content})
}