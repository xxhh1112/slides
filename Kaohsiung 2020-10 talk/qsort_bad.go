package qsort

import "sync"
import "time"

func qsortBadInternal(input []int, wg *sync.WaitGroup) {
	// for demo the effect of go scheduler
	time.Sleep(1 * time.Nanosecond)

	defer wg.Done()

	// sentinal
	if len(input) <= 1 {
		return
	}

	pivotPos := qsortPartition(input)

	wg.Add(2)
	go qsortBadInternal(input[:pivotPos], wg)
	go qsortBadInternal(input[pivotPos+1:], wg)
}

func qsortBad(input []int) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go qsortBadInternal(input, &wg)
	wg.Wait()
}
