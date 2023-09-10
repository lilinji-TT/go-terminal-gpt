package model


type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}


type RequestBody struct {
	Model string `json:"model"`

	Messages []Message `json:"messages"`

	Stream bool `json:"stream"`
}

type ResponseBodyJSON struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Created int64 `json:"created"`
	Model  string `json:"model"`
	Choices []struct {
		Index        int    `json:"index"`
		Delta        struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"delta"`
		FinishReason *string `json:"finish_reason"`
	} `json:"choices"`
}


