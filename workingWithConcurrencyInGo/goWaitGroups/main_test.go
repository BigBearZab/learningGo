package main

import (
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {

	wg.Add(1)

	go updateMessage("New message", &wg)

	wg.Wait()

	if !strings.Contains(msg, "New message") {
		t.Errorf("Expected 'New message', got %s", msg)
	}
}
