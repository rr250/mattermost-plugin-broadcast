package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func (p *Plugin) InitAPI() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/getallusersinteam", p.getAllUsersInTeam).Methods("POST")
	r.HandleFunc("/broadcast", p.broadcast).Methods("POST")
	return r
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}

func (p *Plugin) getAllUsersInTeam(w http.ResponseWriter, req *http.Request) {
	userID := req.Header.Get("Mattermost-User-ID")
	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}
	var team *Team
	err := json.NewDecoder(req.Body).Decode(&team)
	if err != nil {
		p.API.LogError("Unable to decode JSON err=" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userList, err1 := p.API.GetUsersInTeam(team.TeamID, 0, 100000)
	if err1 != nil {
		p.API.LogError("Unable to get users in team" + err1.Error())
	}
	p.API.LogInfo("", userList)
	json.NewEncoder(w).Encode(userList)
}

func (p *Plugin) broadcast(w http.ResponseWriter, req *http.Request) {
	userID := req.Header.Get("Mattermost-User-ID")
	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	var broadcast *Broadcast
	err := json.NewDecoder(req.Body).Decode(&broadcast)
	if err != nil {
		p.API.LogError("Unable to decode JSON err=" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userIDmap := make(map[string]struct{})
	var exists = struct{}{}
	for _, userID := range broadcast.UserIdList {
		userIDmap[userID] = exists
	}

	for _, channelID := range broadcast.ChannelIdList {

		channelStats, err := p.API.GetChannelStats(channelID)
		if err != nil {
			p.API.LogError("Unable to get channel stats" + err.Error())
		}
		channelUsers, err := p.API.GetUsersInChannel(channelID, "username", 0, int(int64(channelStats.MemberCount)+channelStats.GuestCount))
		if err != nil {
			p.API.LogError("Unable to get users in channel" + err.Error())
		}
		for _, user := range channelUsers {
			userIDmap[user.Id] = exists
		}
	}
	var broadcast1 Broadcast
	broadcast1.ID = model.NewId()
	userIDList := make([]string, 0, len(userIDmap))
	for userID := range userIDmap {
		userIDList = append(userIDList, userID)
	}
	broadcast1.UserIdList = userIDList
	p.AddBroadcast(broadcast1)
}

func (p *Plugin) sendBroadcast() {
	broadcastSummary, err1 := p.GetRecentBroadcast()
	if err1 != nil {
		p.API.LogError("Unable to Broadcast -- err=" + err1.(string))
	}
	broadcast, err2 := p.GetBroadcast(broadcastSummary.BroadcastID)
	if err2 != nil {
		p.API.LogError("Unable to Broadcast -- err=" + err2.(string))
	}
	for _, recieverID := range broadcast.UserIdList {
		channel, err := p.API.GetDirectChannel(broadcast.SenderUserID, recieverID)
		if err != nil {
			p.API.LogError("Unable to create direct channel -- err=" + err.Error())
		}
		for i := 0; i < 1000; i++ {
			postModel := &model.Post{
				UserId:    broadcast.SenderUserID,
				ChannelId: channel.Id,
				Message:   fmt.Sprintf("%d", i),
			}
			_, err3 := p.API.CreatePost(postModel)

			if err3 != nil {
				p.API.LogError("Unable to Broadcast -- err=" + err3.Error())
			}
		}
	}
}
