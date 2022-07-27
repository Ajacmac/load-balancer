package main

import (
	"testing"
)

// func TestAddBackend(t *testing.T) {
// 	original := serverPool

// 	serverPool.AddBackend(&Backend{
// 		URL:          serverUrl,
// 		Alive:        true,
// 		ReverseProxy: proxy,
// 	})
// 	log.Printf("Configured server: %s\n", serverUrl)
// }

// func TestSetupBackends(t *testing.T) {
// 	// add subtests, maybe 10 per type, for each step
// 	// How do I keep breaking out the functionality to make it more testable?
// }

func Test_setupBackends(t *testing.T) {
	type args struct {
		serverList string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupBackends(tt.args.serverList)
		})
	}
}
