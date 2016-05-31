package main

import (
	"github.com/rcrowley/go-metrics"
	"fmt"
)

func main() {
	r := metrics.NewRegistry()
	for i := 1; i <= 5; i++ {
		c := metrics.NewCounter()

		conn := r.GetOrRegister("counter", c).(metrics.Counter)
		conn.Inc(1)
		fmt.Println(conn.Count())
	}
}
