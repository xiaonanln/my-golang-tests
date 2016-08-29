package test

import (
	"testing"
	"sync"
)

var (
	lock sync.Mutex
)

func TestNothing(*testing.T) {
}

func deferFunc() {
	defer lock.Unlock()
	lock.Lock()
}

func noDeferFunc() {
	lock.Lock()
	lock.Unlock()
}

func justDeferFunc() {
	defer func(){} ()
}

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferFunc()
	}
}

func BenchmarkNoDefer(b *testing.B) {
	for i := 0; i< b.N; i++ {
		noDeferFunc()
	}
}

func BenchmarkJustDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		justDeferFunc()
	}
}
