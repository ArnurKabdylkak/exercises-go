package problem6

import (
	"sync"
)

func runTasks(init func()) {
	var wg sync.WaitGroup
	var once sync.Once

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			//TODO: modify so that load function gets called only once.
			once.Do(init)
		}()
	}
	wg.Wait()
}
