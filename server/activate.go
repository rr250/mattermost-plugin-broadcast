package main

import (
	"time"

	"github.com/mattermost/mattermost-plugin-api/cluster"
	"github.com/pkg/errors"
)

// OnActivate register the plugin command
func (p *Plugin) OnActivate() error {
	p.router = p.InitAPI()
	job, cronErr := cluster.Schedule(
		p.API,
		"BackgroundJob",
		cluster.MakeWaitForRoundedInterval(1*time.Minute),
		p.sendBroadcast,
	)
	if cronErr != nil {
		return errors.Wrap(cronErr, "failed to schedule background job")
	}
	p.backgroundJob = job
	return nil
}
