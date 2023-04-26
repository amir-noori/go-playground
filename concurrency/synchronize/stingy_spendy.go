package synchronize

import (
	"sync"
	"time"
)

var lock = sync.Mutex{}
var money = 100

func stingy() {
	for i := 1; i < 1000; i++ {
		lock.Lock()
		money += 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("stingy done")
}

func spendy() {
	for i := 1; i < 1000; i++ {
		lock.Lock()
		money -= 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("spendy done")
}

func Stingy_Spendy() {

	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	println(money)
}
