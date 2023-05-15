package channels

import (
	"fmt"
	"sync"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	time.Sleep(2 * time.Second)
	done <- true
	fmt.Println("This is not printed cause after writing to channel the RunHello function wont wait.")
}
func runHello() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function")
}

func worker(jobName string, infoChannel chan string) {
	time.Sleep(1 * time.Second)
	infoChannel <- "job " + jobName + " is done."
	close(infoChannel)
}

func sendWork(jobName string, wg *sync.WaitGroup) {
	//defer wg.Done()
	infoChannel := make(chan string)
	go worker(jobName, infoChannel)
	for {
		data, ok := <-infoChannel
		if ok == false {
			break
		}
		fmt.Println("channel data: ", data)
	}
}

func master() {

	wg := sync.WaitGroup{}
	wg.Add(3)

	go sendWork("job1", &wg)
	go sendWork("job2", &wg)
	go sendWork("job3", &wg)

	//wg.Wait()
	fmt.Println("All good")
}

func Run() {
	//runHello()
	master()
}
