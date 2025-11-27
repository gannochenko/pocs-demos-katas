package openai

import (
	"context"
	"worker/internal/domain"
	"worker/internal/interfaces"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/pkg/errors"
)

type OpenAIClient struct {
	client openai.Client
	config *domain.Config
}

func NewClient(config *domain.Config) interfaces.OpenAIClient {
	client := openai.NewClient(
		option.WithAPIKey(config.OpenAI.APIKey),
	)

	return &OpenAIClient{
		client: client,
		config: config,
	}
}

func (c *OpenAIClient) Complete(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	completion, err := c.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: openai.ChatModel(c.config.OpenAI.Model),
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
			openai.UserMessage(userPrompt),
		},
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to create chat completion")
	}

	if len(completion.Choices) == 0 {
		return "", errors.New("no completion choices returned")
	}

	return completion.Choices[0].Message.Content, nil
}
