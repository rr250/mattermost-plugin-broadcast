package main

// OnActivate register the plugin command
func (p *Plugin) OnActivate() error {
	p.router = p.InitAPI()
	return nil
}
