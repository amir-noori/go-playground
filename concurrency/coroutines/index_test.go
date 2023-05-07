package coroutines

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_logInfo(t *testing.T) {

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	var wg sync.WaitGroup
	wg.Add(1)
	data := "Math"
	go logInfo(data, &wg)
	wg.Wait()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, data) {
		t.Fatalf("cannot find %s in output", data)
	}
}
