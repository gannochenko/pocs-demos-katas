package slack

import (
	"context"
	"regexp"
	"worker/internal/domain"
	"worker/internal/interfaces"

	"github.com/pkg/errors"
	"github.com/slack-go/slack"
)

type SlackClient struct {
	client *slack.Client
	config *domain.Config
}

func NewClient(config *domain.Config) interfaces.SlackClient {
	client := slack.New(config.Slack.Token)

	return &SlackClient{
		client: client,
		config: config,
	}
}

// ConvertMarkdownToSlack converts Markdown formatting to Slack's mrkdwn format
// - # Header → *Header*
// - ## Header → *Header*
// - ### Header → *Header*
// - **bold** → *bold*
// - [text](url) → <url|text>
// - Emojis :emoji: remain unchanged
func ConvertMarkdownToSlack(markdown string) string {
	result := markdown

	// Convert headers: # Header, ## Header, ### Header → *Header*
	headerRegex := regexp.MustCompile(`(?m)^#{1,6}\s+(.+)$`)
	result = headerRegex.ReplaceAllString(result, `*$1*`)

	// Convert bold: **text** → *text*
	boldRegex := regexp.MustCompile(`\*\*([^*]+)\*\*`)
	result = boldRegex.ReplaceAllString(result, `*$1*`)

	// Convert links: [text](url) → <url|text>
	linkRegex := regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	result = linkRegex.ReplaceAllString(result, `<$2|$1>`)

	return result
}

/**
* https://api.slack.com/apps
* Slack formatting (mrkdwn):
* - Bold: *text*
* - Italic: _text_
* - Strikethrough: ~text~
* - Code: `text`
* - Link: <url|text>
*/
func (c *SlackClient) SendMessage(ctx context.Context, channelID, message string) error {
	// Convert Markdown to Slack formatting
	slackMessage := ConvertMarkdownToSlack(message)

	_, _, err := c.client.PostMessageContext(
		ctx,
		channelID,
		slack.MsgOptionText(slackMessage, false),
		slack.MsgOptionDisableLinkUnfurl(),
		slack.MsgOptionDisableMediaUnfurl(),
	)
	if err != nil {
		return errors.Wrap(err, "failed to send slack message")
	}

	return nil
}
