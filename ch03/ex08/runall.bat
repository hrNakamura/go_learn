pushd %~dp0
go run .\main.go 128 > ex08_128.png
go run .\main.go 64 > ex08_64.png
go run .\main.go Float > ex08_Float.png
go run .\main.go Rat > ex08_Rat.png
popd