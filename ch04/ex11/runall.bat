set OWNER=hrNakamura
set REPOSITRY=go_learn
pushd %~dp0
go run .\main.go -o %OWNER% -r %REPOSITRY% -c Read -n 1
go run .\main.go -o %OWNER% -r %REPOSITRY% -c Create -title "new issue" -body "test"
go run .\main.go -o %OWNER% -r %REPOSITRY% -c Close -n 2
go run .\main.go -o %OWNER% -r %REPOSITRY% -c Edit -n 4 -title "edit issue" -body "edited"
popd