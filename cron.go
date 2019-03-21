package cron

import (
	"github.com/ecletus/ecletus"
	"github.com/ecletus/plug"
	"github.com/moisespsena-go/task"
	"github.com/robfig/cron"
)

type Plugin struct {
	Key string
}

func (p *Plugin) ProvideOptions() []string {
	return []string{p.Key}
}

func (p *Plugin) Init(options *plug.Options) {
	c := cron.New()
	options.Set(p.Key, c)
	agp := options.GetInterface(ecletus.AGHAPE).(*ecletus.Ecletus)
	_ = agp.AddTask(task.NewTask(func() (err error) {
		c.Start()
		return nil
	}, c.Stop))
}
