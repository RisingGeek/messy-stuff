package usecase

import (
	"sync"
	"time"

	"messy-stuff/utils"
)

/*
	Worker performs a simulated expensive task.
*/
func Worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second)
}

/*
	ProcessRoutines performs multiple simulated expensive tasks asynchronously
	It is used to do multiple tasks in parallel using WaitGroup which waits for all the goroutines to be returned
*/
func ProcessRoutines(workers int) (totalTime int64) {
	var wg sync.WaitGroup

	start := utils.GetCurrentTime()
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go Worker(i, &wg)
	}
	wg.Wait()

	end := utils.GetCurrentTime()
	totalTime = end - start
	return
}
