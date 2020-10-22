package main

import (
	"encoding/json"
	"fmt"
)

type Broadcast struct {
	ID            string
	Message       string   `json:"message"`
	UserIdList    []string `json:"userIdList"`
	ChannelIdList []string `json:"channelIdList"`
	SenderUserID  string
}

type Team struct {
	TeamID string `json:"teamId"`
}

type BroadcastList struct {
	BroadcastIDList []BroadcastSummary
}

type BroadcastSummary struct {
	BroadcastID string
}

func (p *Plugin) AddBroadcast(broadcast Broadcast) interface{} {
	broadcastJSON, err := json.Marshal(broadcast)
	if err != nil {
		p.API.LogError("failed to marshal broadcast %s", broadcast.ID)
		return fmt.Sprintf("failed to marshal broadcast %s", broadcast.ID)
	}
	err1 := p.API.KVSet("broadcast-"+broadcast.ID, broadcastJSON)
	if err1 != nil {
		p.API.LogError("failed KVSet %s", err1, broadcast)
		return fmt.Sprintf("failed KVSet %s", err1)
	}

	bytes, err2 := p.API.KVGet("broadcasts")
	if err2 != nil {
		p.API.LogError("failed KVGet %s", err)
		return fmt.Sprintf("failed KVGet %s", err)
	}
	broadcastSummary := BroadcastSummary{
		BroadcastID: broadcast.ID,
	}
	var broadcasts []BroadcastSummary
	if bytes != nil {
		if err = json.Unmarshal(bytes, &broadcasts); err != nil {
			return fmt.Sprintf("failed to unmarshal  %s", err)
		}
		broadcasts = append(broadcasts, broadcastSummary)
	} else {
		broadcasts = []BroadcastSummary{broadcastSummary}
	}
	broadcastsJSON, err := json.Marshal(broadcasts)
	if err != nil {
		p.API.LogError("failed to marshal broadcasts  %s", broadcasts)
		return fmt.Sprintf("failed to marshal broadcasts  %s", broadcasts)
	}
	err3 := p.API.KVSet("broadcasts", broadcastsJSON)
	if err3 != nil {
		p.API.LogError("failed KVSet", err3, broadcastsJSON)
		return fmt.Sprintf("failed KVSet %s", err3)
	}
	return nil
}

func (p *Plugin) GetRecentBroadcast() (BroadcastSummary, interface{}) {
	var broadcastSummary BroadcastSummary
	var broadcastSummaryList []BroadcastSummary
	bytes, err := p.API.KVGet("broadcasts")
	if err != nil {
		p.API.LogError("failed KVGet %s", err)
		return broadcastSummary, fmt.Sprintf("failed to unmarshal %s", err)
	}
	if bytes != nil {
		if err3 := json.Unmarshal(bytes, &broadcastSummaryList); err3 != nil {
			return broadcastSummary, fmt.Sprintf("failed to unmarshal %s", err3)
		}
		broadcastSummary := broadcastSummaryList[0]
		broadcastSummaryList = broadcastSummaryList[1:]
		broadcastsJSON, err := json.Marshal(broadcastSummaryList)
		if err != nil {
			p.API.LogError("failed to marshal broadcasts  %s", broadcastSummaryList)
			return broadcastSummary, fmt.Sprintf("failed to marshal broadcasts  %s", broadcastSummaryList)
		}
		err3 := p.API.KVSet("broadcasts", broadcastsJSON)
		if err3 != nil {
			p.API.LogError("failed KVSet", err3, broadcastsJSON)
			return broadcastSummary, fmt.Sprintf("failed KVSet %s", err3)
		}

	} else {
		return broadcastSummary, "No Broadcast found"
	}
	return broadcastSummaryList[0], nil
}

func (p *Plugin) GetBroadcast(broadcastID string) (Broadcast, interface{}) {
	var broadcast Broadcast
	bytes, err := p.API.KVGet("broadcast-" + broadcastID)
	if err != nil {
		p.API.LogError("failed KVGet %s", err)
		return broadcast, fmt.Sprintf("failed to unmarshal %s", err)
	}
	if bytes != nil {
		if err3 := json.Unmarshal(bytes, &broadcast); err3 != nil {
			return broadcast, fmt.Sprintf("failed to unmarshal %s", err3)
		}
	} else {
		return broadcast, "No Broadcast found"
	}
	return broadcast, nil
}
