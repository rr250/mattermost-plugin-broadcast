package main

import (
	"sync"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

type Plugin struct {
	plugin.MattermostPlugin
	router            *mux.Router
	configurationLock sync.RWMutex
	configuration     *configuration
}
