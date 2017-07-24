package main

import (
	"runtime"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	goNum := runtime.NumCPU()
	printImage(goNum)
}

func BenchmarkMain2(b *testing.B) {
	goNum := runtime.NumCPU() / 2
	printImage(goNum)
}

func BenchmarkMain3(b *testing.B) {
	goNum := runtime.NumCPU() / 4
	printImage(goNum)
}

func BenchmarkMain4(b *testing.B) {
	goNum := runtime.NumCPU() / 8
	printImage(goNum)
}
