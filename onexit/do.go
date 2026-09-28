package onexit

func Do(job func()) {
	lock.Lock()
	defer lock.Unlock()

	jobs = append(jobs, job)
}
