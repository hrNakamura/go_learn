pushd %~dp0
go run .\main.go 12345 51234
go run .\main.go absde fdcab
go run .\main.go "日本語" "語日本"
popd
