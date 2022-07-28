package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"sync"
	"testing"
)

// add subtests, maybe 10 per type, for each step
// How do I keep breaking out the functionality to make it more testable?
// Revisit this tomorrow. I waited way too long.

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

func Test_healthCheck(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			healthCheck()
		})
	}
}

func Test_isBackendAlive(t *testing.T) {
	type args struct {
		u *url.URL
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBackendAlive(tt.args.u); got != tt.want {
				t.Errorf("isBackendAlive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lb(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb(tt.args.w, tt.args.r)
		})
	}
}

func TestGetRetryFromContext(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRetryFromContext(tt.args.r); got != tt.want {
				t.Errorf("GetRetryFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAttemptsFromContext(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAttemptsFromContext(tt.args.r); got != tt.want {
				t.Errorf("GetAttemptsFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServerPool_HealthCheck(t *testing.T) {
	type fields struct {
		backends []*Backend
		current  uint64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServerPool{
				backends: tt.fields.backends,
				current:  tt.fields.current,
			}
			s.HealthCheck()
		})
	}
}

func TestServerPool_GetNextPeer(t *testing.T) {
	type fields struct {
		backends []*Backend
		current  uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Backend
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServerPool{
				backends: tt.fields.backends,
				current:  tt.fields.current,
			}
			if got := s.GetNextPeer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerPool.GetNextPeer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServerPool_MarkBackendStatus(t *testing.T) {
	type fields struct {
		backends []*Backend
		current  uint64
	}
	type args struct {
		backendUrl *url.URL
		alive      bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServerPool{
				backends: tt.fields.backends,
				current:  tt.fields.current,
			}
			s.MarkBackendStatus(tt.args.backendUrl, tt.args.alive)
		})
	}
}

func TestServerPool_NextIndex(t *testing.T) {
	type fields struct {
		backends []*Backend
		current  uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServerPool{
				backends: tt.fields.backends,
				current:  tt.fields.current,
			}
			if got := s.NextIndex(); got != tt.want {
				t.Errorf("ServerPool.NextIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServerPool_AddBackend(t *testing.T) {
	type fields struct {
		backends []*Backend
		current  uint64
	}
	type args struct {
		backend *Backend
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServerPool{
				backends: tt.fields.backends,
				current:  tt.fields.current,
			}
			s.AddBackend(tt.args.backend)
		})
	}
}

func TestBackend_IsAlive(t *testing.T) {
	type fields struct {
		URL          *url.URL
		Alive        bool
		mux          sync.RWMutex
		ReverseProxy *httputil.ReverseProxy
	}
	tests := []struct {
		name      string
		fields    fields
		wantAlive bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Backend{
				URL:          tt.fields.URL,
				Alive:        tt.fields.Alive,
				mux:          tt.fields.mux,
				ReverseProxy: tt.fields.ReverseProxy,
			}
			if gotAlive := b.IsAlive(); gotAlive != tt.wantAlive {
				t.Errorf("Backend.IsAlive() = %v, want %v", gotAlive, tt.wantAlive)
			}
		})
	}
}

func TestBackend_SetAlive(t *testing.T) {
	type fields struct {
		URL          *url.URL
		Alive        bool
		mux          sync.RWMutex
		ReverseProxy *httputil.ReverseProxy
	}
	type args struct {
		alive bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Backend{
				URL:          tt.fields.URL,
				Alive:        tt.fields.Alive,
				mux:          tt.fields.mux,
				ReverseProxy: tt.fields.ReverseProxy,
			}
			b.SetAlive(tt.args.alive)
		})
	}
}
