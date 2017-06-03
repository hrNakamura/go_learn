pushd %~dp0
go run .\main.go 1.23
go run .\main.go 12345.67
go run .\main.go -123.456789
popd
