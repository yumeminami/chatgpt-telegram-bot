package chatgpt

import (
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func CreateEdit(output string, instruction string) (string, error) {
	fmt.Println("Edit")
	model := "text-davinci-edit-001"
	req := gogpt.EditsRequest{
		Model:       &model,
		Input:       output,
		Instruction: instruction,
		N:           1,
		Temperature: 0.7,
		TopP:        1,
	}
	resp, err := client.Edits(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Text, nil
}
