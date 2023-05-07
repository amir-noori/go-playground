package coroutines

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_printMessages(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	for i := 0; i < 50; i++ {
		PrintMessages()
	}
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	cosmosCount := strings.Count(output, "cosmos")
	worldCount := strings.Count(output, "world")
	universeCount := strings.Count(output, "universe")

	if cosmosCount != 50 && worldCount != 50 && universeCount != 50 {
		t.Fatalf("Print messages failed!")
	}
}
