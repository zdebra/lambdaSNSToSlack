package main

// SlackPayload defines slack's expected message
//
// Sending message
// For a simple message, your JSON payload could contain a text property at minimum. This is the text that will be posted to the channel.
//
// Adding links
// To create a link in your text, enclose the URL in <> angle brackets.
// For example: payload={"text": "<https://slack.com>"} will post a clickable link to https://slack.com.
//
// Customized Appearence
// You can customize the name and icon of your Incoming Webhook in the Integration Settings section below.
// However, you can override the displayed name by sending "username": "new-bot-name" in your JSON payload.
// You can also override the bot icon either with "icon_url": "https://slack.com/img/icons/app-57.png"
// or "icon_emoji": ":ghost:".
//
// Channel Override
// Incoming webhooks have a default channel, but it can be overridden in your JSON payload.
// A public channel can be specified with "channel": "#other-channel", and a Direct Message with "channel": "@username".
type SlackPayload struct {
	Text      string `json:"text"` // To create a link in your text, enclose the URL in <> angle brackets
	Username  string `json:"username,omitempty"`
	IconURL   string `json:"icon_url,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
	Channel   string `json:"channel,omitempty"`
}
