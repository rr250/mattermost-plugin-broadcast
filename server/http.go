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

	var item *Broadcast
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&item)
	if err != nil {
		p.API.LogError("Unable to decode JSON err=" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, reciept_id := range item.Usersid {

		channel, _ := p.API.GetDirectChannel(userID, reciept_id)
		postModel := &model.Post{
			UserId:    userID,
			ChannelId: channel.Id,
			Message:   item.Message,
		}
		_, err := p.API.CreatePost(postModel)

		if err != nil {
			p.API.LogError("Unable to Broadcast -- err=" + err.Error())

		}
	}
}
