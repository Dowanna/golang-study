package ex01

import (
	"bufio"
	"fmt"
)

type WordCounter int
type LineCounter int

func main() {
	p := []byte("世界　世界　世界")
	var c WordCounter
	c.Write(p)
	fmt.Printf("%v", c)
}

func (c *WordCounter) Write(p []byte) (int, error) {
	ploc := p
	for i := 0; i < len(ploc); {
		adv, token, _ := bufio.ScanWords(ploc, true)

		if token != nil {
			*c += 1
		}

		ploc = ploc[adv:]
	}
	return 0, nil
}
