package main

import (
	"runtime"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	goNum := runtime.NumCPU()
	printImage(goNum)
}

func BenchmarkMainHalf(b *testing.B) {
	goNum := runtime.NumCPU() / 2
	printImage(goNum)
}

func BenchmarkMainQuarter(b *testing.B) {
	goNum := runtime.NumCPU() / 4
	printImage(goNum)
}

func BenchmarkMainOne_Eighth(b *testing.B) {
	goNum := runtime.NumCPU() / 8
	printImage(goNum)
}

func BenchmarkMain2Times(b *testing.B) {
	goNum := runtime.NumCPU()
	printImage(goNum * 2)
}

func BenchmarkMain4Times(b *testing.B) {
	goNum := runtime.NumCPU()
	printImage(goNum * 4)
}

func BenchmarkMain8Times(b *testing.B) {
	goNum := runtime.NumCPU()
	printImage(goNum * 8)
}
