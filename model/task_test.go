package model_test

import (
	"testing"
	"todo/model"
)

// go test -v model/*.go
func TestTask(t *testing.T) {
	task := model.Task{
		Title: "My Super Task",
	}
	if got, want := task.Title, "My Super Task"; got != want {
		t.Errorf("Title got %s; want %s", got, want)
	}
}
