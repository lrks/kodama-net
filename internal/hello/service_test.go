package hello_test

import (
	"context"
	"testing"

	"github.com/lrks/kodama-net/internal/hello"
)

func TestService_Greet(t *testing.T) {
	svc := hello.NewService()
	ctx := context.Background()

	got, err := svc.Greet(ctx, "Go")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "Hello, Go!"
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}
