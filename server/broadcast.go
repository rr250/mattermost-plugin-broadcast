package main

type Broadcast struct {
	Message       string   `json:"message"`
	UserIdList    []string `json:"userIdList"`
	ChannelIdList []string `json:"channelIdList"`
}
