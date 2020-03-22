package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) //slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	//output:
	// Fibonacci(45) = 1134903170
	// 它需要大量的时间来执行，在此期间我们提供一个可见的提示，显示一个
	// 字符串"spinner"来指示程序依然在运行

}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
