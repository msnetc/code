package listing25

import (
	"math/rand"
	"time"
	"sync"
	"fmt"
)

const (
	numberGorountines = 4
	taskLoad =10
)


var wg sync.WaitGroup

func init()  {
	rand.Seed(time.Now().Unix())
}
func  main()  {

	tasks := make(chan string, taskLoad)
	wg.Add(numberGorountines)

	for gr:=1; gr<=numberGorountines;gr++{
  		go workder(tasks, gr)
	}

	for post:=1; post <= taskLoad;post++{
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	close(tasks)
	wg.Wait();}

func workder (tasks chan string, worker int){
		defer  wg.Done();

	for{
		task, ok := <- tasks

		if !ok{
			fmt.Printf("workder：%d:shutting down\n", worker)
			return ;
		}

		// Display we are starting the work.
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// Randomly wait to simulate work time.
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// Display we finished the work.
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}

}
