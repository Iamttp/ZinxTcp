package main

import "fmt"

func main() {
	data := make([]byte, 512)
	data = append(data, 'h')
	data = append(data, 'e')
	data = append(data, 'l')
	for _, c := range data[:] {
		c--
	}
	fmt.Println(string(data))
}
