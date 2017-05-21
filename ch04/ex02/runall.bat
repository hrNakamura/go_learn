pushd %~dp0
go run .\main.go
go run .\main.go -sha 512
go run .\main.go -sha 384
popd