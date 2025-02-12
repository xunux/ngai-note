package main

import (
	"fmt"
	"time"
)

func main() {
	m := map[string]any{
		"age":     18,
		"hello":   "world",
		"now":     time.Now(),
		"now_fmt": time.Now().Format("2006-01-02 15:04:05"),
	}

	for k, v := range m {
		fmt.Printf("%s %v\n", k, v)
	}

}
