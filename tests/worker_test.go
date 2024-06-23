package test18

import (
	"testing"

	test18 "test18.qstand.art/m/v2"
)

func TestWorker(t *testing.T) {
	limit := 1
	_, err := test18.Worker(limit)
	if err != nil {
		t.Fatalf(`worker(%d) produce an error %s`, limit, err)
	}
}
