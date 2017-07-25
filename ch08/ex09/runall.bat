@setlocal enabledelayedexpansion
cd %~dp0
go run ./main.go -v %GOPATH% %GOROOT%
