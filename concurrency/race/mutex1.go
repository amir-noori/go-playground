package race

import (
	"fmt"
	"sync"
)

// running the following code with -race option will warn you about race condition in updateMessage

var message1 string
var message2 string
var wg1 sync.WaitGroup
var wg2 sync.WaitGroup

func updateMessage(data string) {
	defer wg1.Done()
	message1 = data
}

func updateMessageWithMutex(data string, m *sync.Mutex) {
	defer wg2.Done()
	m.Lock()
	message1 = data
	m.Unlock()
}

func RunUpdateMessage() {
	message1 = "message 1"
	wg1.Add(2)
	go updateMessage("message 2")
	go updateMessage("message 3")
	wg1.Wait()

	message2 = "message 1"
	mutex := sync.Mutex{}
	wg2.Add(2)
	go updateMessageWithMutex("message 2", &mutex)
	go updateMessageWithMutex("message 3", &mutex)
	wg2.Wait()

	fmt.Println(message1)
}
