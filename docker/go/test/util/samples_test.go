package util

import (
	"sample-docker-gin/internal/util"
	"testing"
)

func TestSamples_Sample(t *testing.T) {
	actual := util.Sample(1, 2)
	expected := "Hi! from util.Sample! with 3"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
