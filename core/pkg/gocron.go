// gocron enables simple and intuitive scheduling in Go.
//
// Examples:
//		sched.Schedule().Every(10).Seconds().Do(something)
//		sched.Schedule().Every(3).Minutes().Do(something)
//		sched.Schedule().Every(4).Hours().Do(something)
//		sched.Schedule().Every(2).Days().At("12:32").Do(something)
//		sched.Schedule().Every(12).Weeks().Do(something)
//		sched.Schedule().Every(1).Monday().Do(something)
//		sched.Schedule().Every(1).Saturday().At("8:00").Do(something)
package adbt

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// TimeUnit is an numeration used for handling
// time units internally.
type TimeUnit int

const (
	none = iota
	second
	minute
	hour
	day
	week
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

var timeNow = func() time.Time {
	return time.Now()
}

// Job struct handles all the data required to
// schedule and run jobs.
type Job struct {
	Identifier string `json:"identifier"`
	scheduler  *Scheduler
	unit       TimeUnit
	Frequency  int  `json:"freq"`
	UseAt      bool `json:"usedAt"`
	AtHour     int  `json:"atHour"`
	AtMinute   int  `json:"atMinute"`
	workFunc   func()

	NextScheduledRun time.Time `json:"nextRun"`
}

// Every is a method that fills the given Job struct with the given frequency
func (j *Job) Every(frequencies ...int) *Job {
	l := len(frequencies)

	switch l {
	case 0:
		j.Frequency = 1
	case 1:
		if frequencies[0] <= 0 {
			panic("Every expects frequency to be greater than of equal to 1")
		}
		j.Frequency = frequencies[0]
	default:
		panic("Every expects 0 or 1 arguments")
	}

	return j
}

// At method fills the given Job struct atHour and atMinute fields
// with the provided information
func (j *Job) At(t string) *Job {
	j.UseAt = true
	j.AtHour, _ = strconv.Atoi(strings.Split(t, ":")[0])
	j.AtMinute, _ = strconv.Atoi(strings.Split(t, ":")[1])
	return j
}

// Do method fills the given job struct with the function pointer
// to the job (user provided task) itself.
func (j *Job) Do(function func(), identifier string) string {
	j.Identifier = identifier
	j.workFunc = function
	j.scheduleNextRun()
	j.scheduler.jobs = append(j.scheduler.jobs, *j)
	return j.Identifier
}

func (j *Job) due() bool {
	return timeNow().After(j.NextScheduledRun)
}

// Generally, At() can only be used then unit is day or WEEKDAY
func (j *Job) isAtUsedIncorrectly() bool {
	return j.UseAt == true && (j.unit == second || j.unit == minute ||
		j.unit == hour || j.unit == week)
}

// Returns false when job unit is Day or any of the weekdays, vice versa.
// Used for scheduling when job frequency is 1, because day and WEEKDAY
// can be used with At() function which requires different scheduling approach.
func (j *Job) unitNotDayOrWEEKDAY() bool {
	return j.unit == second || j.unit == minute ||
		j.unit == hour || j.unit == week
}

// Returns false when job unit is or any of the weekdays, vice versa.
// Used for scheduling when job frequency is > 1, because we need to
// manually check for unit since we can't schedule WEEKDAYS with
// frequency > 1 .
func (j *Job) unitNotWEEKDAY() bool {
	return j.unit == second || j.unit == minute ||
		j.unit == hour || j.unit == day ||
		j.unit == week
}

func (j *Job) scheduleNextRun() {
	// If Every(frequency) == 1, unit can be anything .
	// At() can be used only with day and WEEKDAY
	if j.Frequency == 1 {

		// Panic if usage of "At()" is incorrect
		if j.isAtUsedIncorrectly() {
			panic(
				`Cannot schedule Every(1) with At()
				 when unit is not day or WEEKDAY`,
			) // TODO: Turn this into err
		}

		// Handle everything except day and WEEKDAY -- these guys don't use At()
		if j.unitNotDayOrWEEKDAY() {
			if j.NextScheduledRun == (time.Time{}) {
				j.NextScheduledRun = timeNow()
			}

			switch j.unit {
			case second:
				j.NextScheduledRun = j.NextScheduledRun.Add(1 * time.Second)
			case minute:
				j.NextScheduledRun = j.NextScheduledRun.Add(1 * time.Minute)
			case hour:
				j.NextScheduledRun = j.NextScheduledRun.Add(1 * time.Hour)
			case week:
				// 168 hours in a week
				j.NextScheduledRun = j.NextScheduledRun.Add(168 * time.Hour)
			}
		} else {
			// Handle Day and WEEKDAY  --  these guys use At()
			switch j.unit {
			case day:
				if j.NextScheduledRun == (time.Time{}) {
					now := timeNow()
					lastMidnight := time.Date(
						now.Year(),
						now.Month(),
						now.Day(),
						0, 0, 0, 0,
						time.Local,
					)
					if j.UseAt == true {
						j.NextScheduledRun = lastMidnight.Add(
							time.Duration(j.AtHour)*time.Hour +
								time.Duration(j.AtMinute)*time.Minute,
						)
					} else {
						j.NextScheduledRun = lastMidnight
					}
				}
				j.NextScheduledRun = j.NextScheduledRun.Add(24 * time.Hour)

			case monday:
				j.scheduleWeekday(time.Monday)
			case tuesday:
				j.scheduleWeekday(time.Tuesday)
			case wednesday:
				j.scheduleWeekday(time.Wednesday)
			case thursday:
				j.scheduleWeekday(time.Thursday)
			case friday:
				j.scheduleWeekday(time.Friday)
			case saturday:
				j.scheduleWeekday(time.Saturday)
			case sunday:
				j.scheduleWeekday(time.Sunday)
			}

		}

		fmt.Println("Scheduled for ", j.NextScheduledRun)

	} else {
		// If Every(frequency) > 1, unit has to be either
		// second, minute, hour, day, week - not a WEEKDAY .
		// At() can be used only with day

		// Panic if usage of "At()" is incorrect
		if j.isAtUsedIncorrectly() {
			panic("Cannot schedule Every(>1) with At() when unit is not day")
			// TODO: Turn this into err
		}

		// Unlike when frequency = 1, here unit can't be anyhing.
		// We have to check that it isn't a WEEKDAY
		if j.unitNotWEEKDAY() {

			// Handle everything except day -- these guys don't use At()
			if j.unit != day {
				if j.NextScheduledRun == (time.Time{}) {
					j.NextScheduledRun = timeNow()
				}

				switch j.unit {
				case second:
					j.NextScheduledRun = j.NextScheduledRun.Add(
						time.Duration(j.Frequency) * time.Second,
					)
				case minute:
					j.NextScheduledRun = j.NextScheduledRun.Add(
						time.Duration(j.Frequency) * time.Minute,
					)
				case hour:
					j.NextScheduledRun = j.NextScheduledRun.Add(
						time.Duration(j.Frequency) * time.Hour,
					)
				case week:
					j.NextScheduledRun = j.NextScheduledRun.Add(
						time.Duration(j.Frequency*168) * time.Hour,
					) // 168 hours in a week

				}
			} else {
				// Handle Day  --  these guy uses At()
				if j.NextScheduledRun == (time.Time{}) {
					now := timeNow()
					lastMidnight := time.Date(
						now.Year(),
						now.Month(),
						now.Day(),
						0, 0, 0, 0,
						time.Local,
					)
					if j.UseAt == true {
						j.NextScheduledRun = lastMidnight.Add(
							time.Duration(j.AtHour)*time.Hour +
								time.Duration(j.AtMinute)*time.Minute,
						)
					} else {
						j.NextScheduledRun = lastMidnight
					}
				}
				j.NextScheduledRun = j.NextScheduledRun.Add(
					time.Duration(j.Frequency*24) * time.Hour,
				)
			}

		} else {
			panic("Cannot schedule Every(>1) when unit is WEEKDAY")
			// TODO: Turn this into err
		}

		fmt.Println("Scheduled for ", j.NextScheduledRun)
		// TODO: Turn this into a log

	}
	return
}

func (j *Job) scheduleWeekday(dayOfWeek time.Weekday) {
	if j.NextScheduledRun == (time.Time{}) {
		now := timeNow()
		lastWeekdayMidnight := time.Date(
			now.Year(),
			now.Month(),
			now.Day()-int(now.Weekday()-dayOfWeek),
			0, 0, 0, 0,
			time.Local)
		if j.UseAt == true {
			j.NextScheduledRun = lastWeekdayMidnight.Add(
				time.Duration(j.AtHour)*time.Hour +
					time.Duration(j.AtMinute)*time.Minute,
			)
		} else {
			j.NextScheduledRun = lastWeekdayMidnight
		}
	}
	j.NextScheduledRun = j.NextScheduledRun.Add(7 * 24 * time.Hour)
}

// Second method fills the given job struct with seconds
func (j *Job) Second() *Job {
	j.unit = second
	return j
}

// Seconds method fills the given job struct with seconds
func (j *Job) Seconds() *Job {
	j.unit = second
	return j
}

// Minute method fills the given job struct with minutes
func (j *Job) Minute() *Job {
	j.unit = minute
	return j
}

// Minutes method fills the given job struct with minutes
func (j *Job) Minutes() *Job {
	j.unit = minute
	return j
}

// Hour method fills the given job struct with hours
func (j *Job) Hour() *Job {
	j.unit = hour
	return j
}

// Hours method fills the given job struct with hours
func (j *Job) Hours() *Job {
	j.unit = hour
	return j
}

// Day method fills the given job struct with days
func (j *Job) Day() *Job {
	j.unit = day
	return j
}

// Days method fills the given job struct with days
func (j *Job) Days() *Job {
	j.unit = day
	return j
}

// Week method fills the given job struct with weeks
func (j *Job) Week() *Job {
	j.unit = week
	return j
}

// Weeks method fills the given job struct with weeks
func (j *Job) Weeks() *Job {
	j.unit = week
	return j
}

// Monday method fills the given job struct with monday
func (j *Job) Monday() *Job {
	j.unit = monday
	return j
}

// Tuesday method fills the given job struct with tuesday
func (j *Job) Tuesday() *Job {
	j.unit = tuesday
	return j
}

// Wednesday method fills the given job struct with wednesday
func (j *Job) Wednesday() *Job {
	j.unit = wednesday
	return j
}

// Thursday method fills the given job struct with thursday
func (j *Job) Thursday() *Job {
	j.unit = thursday
	return j
}

// Friday method fills the given job struct with friday
func (j *Job) Friday() *Job {
	j.unit = friday
	return j
}

// Saturday method fills the given job struct with saturday
func (j *Job) Saturday() *Job {
	j.unit = saturday
	return j
}

// Sunday method fills the given job struct with sunday
func (j *Job) Sunday() *Job {
	j.unit = sunday
	return j
}

// Scheduler type is used to store a group of jobs (Job structs)
type Scheduler struct {
	identifier string
	jobs       []Job
	quit       chan bool
}

// NewScheduler creates and returns a new Scheduler
func NewScheduler() Scheduler {
	return Scheduler{
		identifier: uuid.New().String(),
		jobs:       make([]Job, 0),
	}
}

// activateTestMode method sets the timeNow func for testing,
// by setting the current time to a fixed value
func (s *Scheduler) activateTestMode() {
	timeNow = func() time.Time {
		return time.Date(1, 1, 1, 1, 1, 0, 0, time.Local)
	}
}

// Run method on the Scheduler type runs the scheduler.
// This is a blocking method, and should be run as a goroutine.
func (s *Scheduler) Run() {
	for {
		select {
		case <-s.quit:
			log.Println("Stopping old scheduler...")
			return
		default:
			for jobIdx := range s.jobs {
				job := &s.jobs[jobIdx]
				if job.due() {
					job.scheduleNextRun()
					go job.workFunc()
				}
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func (s *Scheduler) Stop() {
	s.quit <- true
}

// Schedule method on the Scheduler creates a new Job
// and prepares is for "filling"
func (s *Scheduler) Schedule() *Job {
	newJob := Job{
		Identifier:       uuid.New().String(),
		scheduler:        s,
		unit:             none,
		Frequency:        1,
		UseAt:            false,
		AtHour:           0,
		AtMinute:         0,
		workFunc:         nil,
		NextScheduledRun: time.Time{}, // zero value
	}
	return &newJob
}
