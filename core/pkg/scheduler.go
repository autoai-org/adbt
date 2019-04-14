package adbt

import "fmt"

func loadScheduler(config adbtConfig) Scheduler {
	sched := NewScheduler()
	for _, scheduler := range config.Jobs {
		job := func() {
			if scheduler.Database == "mongodb" {
				m := MongoDB{
					URI:  scheduler.URI,
					Name: scheduler.Name,
				}
				m.Backup()
			} else {
				_ = fmt.Errorf("Unsupported Database")
			}
		}
		sched.Schedule().Every().Day().At(scheduler.Time).Do(job, scheduler.Identifier)
	}
	go sched.Run()
	return sched
}