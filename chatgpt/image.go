package chatgpt

import (
	"context"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func CreateImage(c *gogpt.Client, ctx context.Context, prompt string) (string, error) {
	req := gogpt.ImageRequest{
		Prompt:         prompt,
		Size:           gogpt.CreateImageSize512x512,
		ResponseFormat: gogpt.CreateImageResponseFormatURL,
	}
	resp, err := c.CreateImage(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Data[0].URL, nil
}
