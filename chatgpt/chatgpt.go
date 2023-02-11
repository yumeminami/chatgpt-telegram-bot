package chatgpt

import (
	"bytes"
	"chatgpt-telegram-bot/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Creates a completion for the provided prompt and parameters
func Completions(prompt string) (*string, error) {
	if prompt == "" {
		panic("prompt is empty")
	}

	openai_api_key := config.GetConfig().Openai_API_Key
	if openai_api_key == "" {
		panic("openai api key is empty")
	}

	completion_request_body := GetCompletionsRequest(prompt)
	completion_request_data, err := json.Marshal(completion_request_body)
	if err != nil {
		panic(err)
	}
	completion_request, err := http.NewRequest("POST", completion_url, bytes.NewBuffer(completion_request_data))
	if err != nil {
		panic(err)
	}
	completion_request.Header.Set("Content-Type", "application/json")
	completion_request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", openai_api_key))

	client := &http.Client{}
	completion_response, err := client.Do(completion_request)
	if err != nil {
		panic(err)
	}

	defer completion_response.Body.Close()
	if completion_response.StatusCode != 200 {
		panic(fmt.Sprintf("completion request failed with status code %d", completion_response.StatusCode))
	}

	completion_response_data, err := ioutil.ReadAll(completion_response.Body)
	if err != nil {
		panic(err)
	}

	var completion_response_body CompletionResponse
	err = json.Unmarshal(completion_response_data, &completion_response_body)
	if err != nil {
		panic(err)
	}

	// get the text from the completion response
	if len(completion_response_body.Choices) == 0 {
		panic("completion response body choices is empty")
	}
	text := completion_response_body.Choices[0].Text
	return &text, nil
}

// Creates a new edit for the provided input, instruction, and parameters.
// Its meaning is you can edit the previous input to regenerate a new output.
func Edit(msg string) (*string, error) {
	return nil, nil
}
