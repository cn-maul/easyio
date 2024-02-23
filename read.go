package easyio

import (
	"bufio"
	"io"
	"log"
)

func Read(file io.Reader) []byte {
	r := bufio.NewReader(file)
	buf := make([]byte, 1024)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			log.Println(err)
			return nil
		}

		if n == 0 {
			break
		}
	}
	return buf
}
