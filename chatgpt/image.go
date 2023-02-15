package chatgpt

import (
	gogpt "github.com/sashabaranov/go-gpt3"
)

func CreateImage(prompt string) (string, error) {
	req := gogpt.ImageRequest{
		Prompt:         prompt,
		Size:           gogpt.CreateImageSize512x512,
		ResponseFormat: gogpt.CreateImageResponseFormatURL,
	}
	resp, err := client.CreateImage(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Data[0].URL, nil
}

func CreateImageEdit(prompt string, image string) (string, error) {
	return "", nil
}
