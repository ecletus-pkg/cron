package cron

import (
	"github.com/ecletus/ecletus"
	"github.com/ecletus/plug"
	"github.com/moisespsena-go/task"
	"github.com/robfig/cron"
)

type Plugin struct {
	Key  string
	cron *cron.Cron
}

func (p *Plugin) ProvideOptions() []string {
	return []string{p.Key}
}

func (p *Plugin) ProvidesOptions(options *plug.Options) {
	p.cron = cron.New()
	options.Set(p.Key, p.cron)
}

func (p *Plugin) Init(options *plug.Options) {
	agp := options.GetInterface(ecletus.ECLETUS).(*ecletus.Ecletus)
	_ = agp.AddTask(task.NewTask(func() (err error) {
		p.cron.Start()
		return nil
	}, func() {
		p.cron.Stop()
	}))
}
