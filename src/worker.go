package test18

// worker.go
// Wee are need to limit executing ability to spread on N cpus.
// worker pattern help us.

// struct Task {
// 	id int
// }

import (
	"fmt"
)

func Worker(limit int) (int, error) {
	fmt.Printf("Start worker with %d limit\n", limit)
	return 0, nil
}
