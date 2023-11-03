package main

import (
	"os"
	"time"

	"learn-go-with-tests/maths/v5/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
