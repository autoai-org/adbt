package adbt

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func something() {}

func TestScheduleNextRun(t *testing.T) {
	sched := NewScheduler()
	sched.activateTestMode()

	sched.Schedule().Every(10).Seconds().Do(something, "")
	s := sched.jobs[0].NextScheduledRun
	assert.Equal(t, s, time.Date(1, 1, 1, 1, 1, 10, 0, time.Local))

	sched.Schedule().Every(3).Minutes().Do(something, "")
	s = sched.jobs[1].NextScheduledRun
	assert.Equal(t, s, time.Date(1, 1, 1, 1, 4, 0, 0, time.Local))

	sched.Schedule().Every(4).Hours().Do(something, "")
	s = sched.jobs[2].NextScheduledRun
	assert.Equal(t, s, time.Date(1, 1, 1, 5, 1, 0, 0, time.Local))

	sched.Schedule().Every(2).Days().At("12:32").Do(something, "")
	s = sched.jobs[3].NextScheduledRun
	assert.Equal(t, s, time.Date(1, 1, 3, 12, 32, 0, 0, time.Local))

	sched.Schedule().Every(12).Weeks().Do(something, "")
	s = sched.jobs[4].NextScheduledRun
	assert.Equal(t, s, time.Date(1, 3, 26, 1, 1, 0, 0, time.Local))

	sched.Schedule().EverySingle().Second().Do(something, "") // EverySingle is "shorthand" for Every(1)
	s = sched.jobs[5].NextScheduledRun
	assert.Equal(t, s, time.Date(1, 1, 1, 1, 1, 1, 0, time.Local))

	sched.Schedule().EverySingle().Monday().At("9:10").Do(something, "")
	s = sched.jobs[6].NextScheduledRun
	assert.Equal(t, s, time.Date(1, 1, 8, 9, 10, 0, 0, time.Local))

	sched.Schedule().Every().Saturday().At("8:00").Do(something, "")
	s = sched.jobs[7].NextScheduledRun
	assert.Equal(t, s, time.Date(1, 1, 13, 8, 0, 0, 0, time.Local))

	sched.Schedule().Every().Tuesday().At("9:10").Do(something, "")
	s = sched.jobs[8].NextScheduledRun
	assert.Equal(t, s, time.Date(1, 1, 9, 9, 10, 0, 0, time.Local))

	sched.Schedule().Every().Wednesday().At("19:10").Do(something, "")
	s = sched.jobs[9].NextScheduledRun
	assert.Equal(t, s, time.Date(1, 1, 10, 19, 10, 0, 0, time.Local))

	sched.Schedule().Every().Thursday().At("19:10").Do(something, "")
	s = sched.jobs[10].NextScheduledRun
	assert.Equal(t, s, time.Date(1, 1, 11, 19, 10, 0, 0, time.Local))

	sched.Schedule().Every().Friday().At("19:10").Do(something, "")
	s = sched.jobs[11].NextScheduledRun
	assert.Equal(t, s, time.Date(1, 1, 12, 19, 10, 0, 0, time.Local))

	sched.Schedule().Every().Sunday().At("8:00").Do(something, "")
	s = sched.jobs[12].NextScheduledRun
	assert.Equal(t, s, time.Date(1, 1, 7, 8, 0, 0, 0, time.Local))
}

// Main test entrypoint
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
