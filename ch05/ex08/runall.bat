pushd %~dp0
go run .\main.go https://golang.org/ lowframe
go run .\main.go https://golang.org/ nokey
popd
