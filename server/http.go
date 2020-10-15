package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func (p *Plugin) InitAPI() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/broadcast", p.broadcast).Methods("POST")
	return r
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
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

		for _, user := range channelUsers {
			userIDmap[user.Id] = exists
		}
	}
	for recieverID := range userIDmap {
		channel, err := p.API.GetDirectChannel(userID, recieverID)
		if err != nil {
			p.API.LogError("Unable to Broadcast -- err=" + err.Error())
		}
		postModel := &model.Post{
			UserId:    userID,
			ChannelId: channel.Id,
			Message:   broadcast.Message,
		}
		_, err = p.API.CreatePost(postModel)

		if err != nil {
			p.API.LogError("Unable to Broadcast -- err=" + err.Error())
		}

	}
}
