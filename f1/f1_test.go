package f1

import (
	"context"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	if got, want := c.baseURL.String(), baseURL; got != want {
		t.Errorf("NewClient baseURL is %v, want %v", got, want)
	}

	c2 := NewClient(nil)
	if c.http == c2.http {
		t.Error("NewClient returned same http client, but they should be different")
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)

	url := baseURL + "/test"
	req, _ := c.newRequest(context.Background(), "GET", url, nil)

	if got, want := req.URL.String(), url; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", url, got, want)
	}

	if got := req.Body; got != nil {
		t.Errorf("NewRequest(%q) body is %v, want %v", url, got, nil)
	}

	if got, want := req.Header.Get("Accept"), "application/json"; got != want {
		t.Errorf("NewRequest(%q) Accept header is %v, want %v", url, got, want)
	}

	if got, want := req.Header.Get("Content-Type"), ""; got != want {
		t.Errorf("NewRequest(%q) Content-Type header is %v, want %v", url, got, want)
	}
}
