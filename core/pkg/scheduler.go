package adbt

import "fmt"

type backupScheduler struct {
	Identifier  string `toml:"identifier" json:"identifier"`
	URI         string `toml:"uri" json:"uri"`
	Periodicity string `toml:"periodicity" json:"periodicity"`
	Time        string `toml:"time" json:"time"`
	Database    string `toml:"database" json:"database"`
	Name        string `toml:"name" json:"name"`
}

func (b *backupScheduler) run(sched Scheduler) {
	for _, job := range sched.jobs {
		if job.Identifier == b.Identifier {
			job.workFunc()
			fmt.Println("Finished")
		}
	}
}

func loadScheduler(config adbtConfig) Scheduler {
	sched := NewScheduler()
	for _, scheduler := range config.Jobs {
		job := func() {
			if scheduler.Database == "MongoDB" {
				m := MongoDB{
					URI:  scheduler.URI,
					Name: scheduler.Name,
				}
				m.Backup(scheduler.Identifier)
			} else {
				_ = fmt.Errorf("Unsupported Database")
			}
		}
		sched.Schedule().Every().Day().At(scheduler.Time).Do(job, scheduler.Identifier)
	}
	go sched.Run()
	return sched
}