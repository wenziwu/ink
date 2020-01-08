package trace

import (
	"bufio"
	"log"
	"os"
	"time"
)

var On bool = false

var logger *log.Logger
var start time.Time

func init() {
	buf := bufio.NewWriterSize(os.Stderr, 50000)
	logger = log.New(buf, "", 0)
	go func() {
		for range time.Tick(500 * time.Millisecond) {
			buf.Flush()
		}
	}()
}

func Start() {
	start = time.Now()
}

func Log(msg string, args ...interface{}) {
	if On {
		d := time.Since(start)
		// for relative times
		//start = time.Now()
		fms := float64(d) / float64(time.Millisecond)
		nargs := append([]interface{}{}, fms)
		nargs = append(nargs, args...)
		logger.Printf("%5.1fms "+msg, nargs...)
	}
}