package swarm

import (
	"github.com/deckarep/golang-set"
)

type JobContainer struct {
	reservoir mapset.Set
	queued    mapset.Set
	done      mapset.Set
}
type Job string

func (jc *JobContainer) AddJob(job Job) bool {
	is_new := job.IsNew()
	if !is_new {
		jc.reservoir.Add(job)
		return true
	}
	return false
}

// Returns the numbers of jobs in container
// that can be dispatched.
func (jc *JobContainer) ReservoirCount() int {
	return jc.reservoir.Cardinality()
}

// Broadcasts the job to peers, returns
// whether
func (job Job) IsNew() bool {
	return true
}
