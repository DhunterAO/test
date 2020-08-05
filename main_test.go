package main

import "testing"

// test function for GenString, check the output
func TestGenString(t *testing.T) {
	if genStr := GenString(); genStr != "hello world" {
		t.Error("The generated string should be \"hello word\", got ", genStr)
	}
}
