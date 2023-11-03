package main

import (
	"os"
	"time"

	clockface "learn-go-with-tests/maths/v3"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
