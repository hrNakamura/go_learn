package main

func pipeline(num int) (in, out chan int) {
	out = make(chan int)
	start := out
	for i := 0; i < num; i++ {
		in = out
		out = make(chan int)
		go func(in, out chan int) {
			for v := range in {
				out <- v
			}
			close(out)
		}(in, out)
	}
	return start, out
}
