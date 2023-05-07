package coroutines

import (
	fmt "fmt"
	"sync"
)

func Run() {

	var wg sync.WaitGroup

	courses := []string{
		"Math",
		"Physics",
		"CS",
		"Logic",
		"Philosophy",
	}

	wg.Add(len(courses))
	for i, course := range courses {
		courseInfo := fmt.Sprintf("Course No %d is %s", i+1, course)
		go logInfo(courseInfo, &wg)
	}
	wg.Wait()

	wg.Add(1)
	logInfo("no more course!", &wg)
}

func logInfo(info string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(info)
}
