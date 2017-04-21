package main

import "testing"

func TestGetMessage(t *testing.T) {
	if GetMessage() != "" {
		t.Error("Where did that message come from")
	}
}
