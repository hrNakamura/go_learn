cd %~dp0
go test -bench=Benchmark
set GOARCH=386
go test -bench=BenchmarkWord
