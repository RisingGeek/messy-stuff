package usecase

import (
	"messy-stuff/domain"
	"messy-stuff/utils"
	"sync"
	"time"
)

// Worker performs a simulated expensive task.
func ChannelWorker(id int, channel chan domain.ChannelEl, wg *sync.WaitGroup) {
	defer wg.Done()

	start := utils.GetCurrentTime()
	time.Sleep(time.Second)
	end := utils.GetCurrentTime()
	channel <- domain.ChannelEl{ID: id, Time: end - start}
}

// Channels connect concurrent goroutines
func ProcessChannel(workers int) (data domain.WorkerInfo) {
	channel := make(chan domain.ChannelEl, workers)
	// WaitGroup is used to wait for the channel to send all values
	var wg sync.WaitGroup
	start := utils.GetCurrentTime()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go ChannelWorker(i, channel, &wg)
	}

	wg.Wait()
	close(channel)

	for el := range channel {
		data.Workers = append(data.Workers, el)
	}

	end := utils.GetCurrentTime()
	data.TotalTime = end - start
	return
}
