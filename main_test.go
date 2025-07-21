package main

import (
	"os"
	"testing"

	"github.com/cryptowizard0/vmdocker_container/runtime"
	vmmSchema "github.com/hymatrix/hymx/vmm/schema"
	goarSchema "github.com/permadao/goar/schema"
)

// go test -tags=lua53 -v . -run TestOllama
func TestOllama(t *testing.T) {
	os.Setenv("RUNTIME_TYPE", "ollama")
	env := vmmSchema.Env{}
	runtime, err := runtime.New(
		env,
		"0x123",
		"0x84534",
		[]goarSchema.Tag{},
	)
	if err != nil {
		t.Fatalf("create runtime failed: %v", err)
	}

	msg := map[string]string{
		"Action":       "Chat",
		"Prompt":       "Hello, how are you?",
		"From":         "0x123",
		"Module":       "0x84534",
		"Id":           "0x131313",
		"Owner":        "0x123",
		"Block-Height": "100000",
		"Target":       "0x8454",
	}
	result, err := runtime.Apply("Chat", 1, msg)
	if err != nil {
		t.Fatalf("apply failed: %v", err)
	}
	t.Logf("apply ok! result: %s", result)
}

// go test -tags=lua53 -v . -run TestGoluaApply
func TestGoluaApply(t *testing.T) {
	os.Setenv("RUNTIME_TYPE", "golua")
	runtime, err := runtime.New(
		"0x8454",
		"0x123",
		"0x84534",
		"./ao/2.0.1",
		[]byte{},
		[]goarSchema.Tag{},
	)
	if err != nil {
		t.Fatalf("create runtime failed: %v", err)
	}

	msg := map[string]string{
		"Action":       "Info",
		"From":         "0x123",
		"Module":       "0x84534",
		"Id":           "0x131313",
		"Owner":        "0x123",
		"Block-Height": "100000",
		"Target":       "0x8454",
	}
	result, err := runtime.Apply("Info", 1, msg)
	if err != nil {
		t.Fatalf("apply failed: %v", err)
	}
	t.Logf("apply ok! result: %s", result)
}
