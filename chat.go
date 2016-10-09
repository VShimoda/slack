package slack

import (
	"encoding/json"
	"errors"
	"net/url"
)

// ChatPostMessagesArgs is ChatPostMessages arguments
type ChatPostMessagesArgs struct {
	Token       string       `json:"token"`
	ChannelID   string       `json:"channel"` // Required
	Message     string       `json:"text"`    // Required (text)
	Parse       string       `json:"parse"`
	LinkNames   string       `json:"link_names"`
	Attachments []Attachment `json:"attachements"`
	UnfurlLinks string       `json:"unfurl_links"`
	UnfurlMedia string       `json:"unfurl_media"`
	Username    string       `json:"username"`
	AsUser      string       `json:"as_user"`
	IconURL     string       `json:"icon_url"`
	IconEmoji   string       `json:"icon_emoji"`
}

/*
{
    "attachments": [
        {
            "fallback": "Required plain-text summary of the attachment.",
            "color": "#36a64f",
            "pretext": "Optional text that appears above the attachment block",
            "author_name": "Bobby Tables",
            "author_link": "http://flickr.com/bobby/",
            "author_icon": "http://flickr.com/icons/bobby.jpg",
            "title": "Slack API Documentation",
            "title_link": "https://api.slack.com/",
            "text": "Optional text that appears within the attachment",
            "fields": [
                {
                    "title": "Priority",
                    "value": "High",
                    "short": false
                }
            ],
            "image_url": "http://my-website.com/path/to/image.jpg",
            "thumb_url": "http://example.com/path/to/thumb.png",
            "footer": "Slack API",
            "footer_icon": "https://platform.slack-edge.com/img/default_application_icon.png",
            "ts": 123456789
        }
    ]
}
*/

// Attachment is for PostMessage
type Attachment struct {
	Fallback   string  `json:"fallback"`
	Color      string  `json:"color"`
	Pretext    string  `json:"pretext"`
	AuthorName string  `json:"author_name"`
	AuthorLink string  `json:"author_link"`
	AuthorIcon string  `json:"author_icon"`
	Title      string  `json:"title"`
	TitleLink  string  `json:"title_link"`
	Text       string  `json:"text"`
	Fileds     []Field `json:"fields"`
	ImageURL   string  `json:"image_url"`
	ThumbURL   string  `json:"thumb_url"`
	Footer     string  `json:"footer"`
	FooterIcon string  `json:"footer_icon"`
	Ts         int     `json:"ts"`
}

// Field is a message table
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// ResponseChatPostMessage is a response of chatpostmessage
type ResponseChatPostMessage struct {
	Ok      bool                           `json:"ok"`
	Error   string                         `json:"error"`
	Channel string                         `json:"channel"`
	Ts      float64                        `json:"ts"`
	Message ResponseChatPostMessageMessage `json:"message"`
}

// ResponseChatPostMessageMessage is a message
type ResponseChatPostMessageMessage struct {
	Type  string  `json:"type"`
	User  string  `json:"user"`
	Text  string  `json:"text"`
	BotID string  `json:"bot_id"`
	Ts    float64 `json:"ts"`
}

// ChatPostMessages post messages to your channel
func (s Slack) ChatPostMessages(args *ChatPostMessagesArgs) error {
	endpoint := URLPrefix + "chat.postMessage"
	values := url.Values{}
	values.Set("token", s.Token)
	if args.ChannelID != "" {
		values.Set("channel", args.ChannelID)
	}
	if args.Message != "" {
		values.Set("text", args.Message)
	}
	if args.Parse != "" {
		values.Set("parse", args.Parse)
	}
	if args.LinkNames != "" {
		values.Set("link_names", args.LinkNames)
	}
	if args.Attachments != nil {
		for _, attachment := range args.Attachments {
			attach, err := json.Marshal(attachment)
			if err != nil {
				return err
			}
			values.Add("attachements", string(attach))
		}
	}
	if args.UnfurlLinks != "" {
		values.Add("unfurl_links", args.UnfurlLinks)
	}
	if args.UnfurlMedia != "" {
		values.Add("unfurl_media", args.UnfurlMedia)
	}
	if args.Username != "" {
		values.Add("username", args.Username)
	}
	if args.AsUser != "" {
		values.Add("as_user", args.AsUser)
	}
	if args.IconURL != "" {
		values.Add("icon_url", args.IconURL)
	}
	if args.IconEmoji != "" {
		values.Add("icon_emoji", args.IconEmoji)
	}

	body, err := requestPostForm(&s.Request, endpoint, &values)
	if err != nil {
		return err
	}

	respChatPostMessage := &ResponseChatPostMessage{}
	err = json.Unmarshal(body, &respChatPostMessage)
	if err != nil {
		return err
	}

	if !respChatPostMessage.Ok {
		return errors.New(respChatPostMessage.Error)
	}
	return nil
}
