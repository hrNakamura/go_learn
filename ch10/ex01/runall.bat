cd %~dp0
.\mandelbrot.exe | go run .\main.go > mandelbrot.jpg
.\mandelbrot.exe | go run .\main.go -f gif > mandelbrot.gif
.\mandelbrot.exe | go run .\main.go -f png > mandelbrot.png
