package main

import (
	"os"
	"time"

	clockface "learn-go-with-tests/maths/v4"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
