package main

import (
	"context"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const HOST_FILE = `C:\Windows\System32\drivers\etc\hosts`

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Query(query string) []string {
	return queryDomains(query)
}

func (a *App) GetBlockedDomains() []string {
	domains, err := getBlockedDomains()
	if err != nil {
		log.Print(err)
	}
	return domains
}

func (a *App) AddDomain(domain string) bool {
	addToHostsFile(domain)
	runtime.EventsEmit(a.ctx, "sync")
	return true
}

func (a *App) DeleteDomain(domain string) bool {
	removeFromHostsFile(domain)
	runtime.EventsEmit(a.ctx, "sync")
	return true
}
