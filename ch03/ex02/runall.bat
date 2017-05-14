pushd %~dp0
go run .\main.go > ex03_02_0.svg
go run .\main.go 1 > ex03_02_1.svg
go run .\main.go 2 > ex03_02_2.svg
go run .\main.go 3 > ex03_02_3.svg
popd