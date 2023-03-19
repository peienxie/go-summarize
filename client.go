package main

import (
	"context"
	"errors"

	openai "github.com/sashabaranov/go-openai"
)

// Client interface represents an OpenAI client
type Client interface {
	SummarizeText(ctx context.Context, text string) (string, error)
}

// clientImpl struct implements the Client interface
type clientImpl struct {
	*openai.Client
}

// NewClient creates a new OpenAI client
func NewClient(apiKey string) Client {
	return clientImpl{openai.NewClient(apiKey)}
}

// Summarize generates a summary of the given text using the OpenAI API
func (c clientImpl) SummarizeText(ctx context.Context, inputText string) (string, error) {
	prompt := "Summarize the following text:\n" + inputText
	req := openai.CompletionRequest{
		Model:       openai.GPT3TextDavinci001,
		MaxTokens:   1024,
		Prompt:      prompt,
		Temperature: 0,
		Stream:      false,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	} else if len(resp.Choices) == 0 {
		return "", errors.New("OpenAI response has no choices")
	}
	return resp.Choices[0].Text, nil
}
