package slack

import (
	"encoding/json"
	"errors"
	"net/url"
)

/*
{
    "ok": true,
    "channels": [
        {
            "id": "C024BE91L",
            "name": "fun",
            "created": 1360782804,
            "creator": "U024BE7LH",
            "is_archived": false,
            "is_member": false,
            "num_members": 6,
            "topic": {
                "value": "Fun times",
                "creator": "U024BE7LV",
                "last_set": 1369677212
            },
            "purpose": {
                "value": "This channel is for fun",
                "creator": "U024BE7LH",
                "last_set": 1360782804
            }
        },
        ....
    ]
}
*/

// ResponseChannelsList is struct of ChannelsList
type ResponseChannelsList struct {
	Ok       bool      `json:"ok"`
	ErrorStr string    `json:"error"`
	Channels []Channel `json:"channels"`
}

// Channel for Slack API Response
type Channel struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Created    int            `json:"created"`
	Creator    string         `json:"creator"`
	IsArchived bool           `json:"is_archived"`
	IsMember   bool           `json:"is_member"`
	NumMembers int            `json:"num_members"`
	Topic      ChannelTopic   `json:"topic"`
	Purpose    ChannelPurpose `json:"purpose"`
}

// ChannelTopic for Channel Topic
type ChannelTopic struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
}

// ChannelPurpose for Purpose
type ChannelPurpose struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
}

// ChannelsList show your slack channels list
// exclude_archived Optional, If its value 1, Don't return archived channels.
func (s Slack) ChannelsList(excludeArchived bool) ([]Channel, error) {
	endpoint := URLPrefix + "channels.list"
	values := url.Values{}
	values.Set("token", s.Token)
	if excludeArchived {
		values.Add("exclude_archived", "1")
	}

	body, err := requestPostForm(&s.Request, endpoint, &values)
	if err != nil {
		return []Channel{}, err
	}

	respChannel := &ResponseChannelsList{}
	err = json.Unmarshal(body, &respChannel)
	if err != nil {
		return []Channel{}, err
	}

	if !respChannel.Ok {
		return []Channel{}, errors.New(respChannel.ErrorStr)
	}
	return []Channel{}, err
}
