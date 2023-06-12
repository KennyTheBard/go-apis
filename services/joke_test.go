package services

import "testing"

func TestGetJoke(t *testing.T) {
	joke := GetJoke()
	if len(joke.Delivery) == 0 || len(joke.Delivery) == 0 {
		t.Errorf("Missing joke!")
	}
}
