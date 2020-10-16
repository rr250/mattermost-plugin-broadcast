package main

type Broadcast struct {
	Message       string   `json:"message"`
	UserIdList    []string `json:"userIdList"`
	ChannelIdList []string `json:"channelIdList"`
}

type Team struct {
	TeamID string `json:"teamId"`
}
