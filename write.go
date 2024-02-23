package easyio

import (
	"bufio"
	"io"
	"log"
)

func WriteBytes(output io.Writer, input []byte) {
	buf := bufio.NewWriter(output)
	_, err := buf.Write(input)
	if err != nil {
		log.Println(err)
		return
	}

	err = buf.Flush()
	if err != nil {
		log.Println(err)
		return
	}
}

func WriteString(output io.Writer, input string) {
	WriteBytes(output, []byte(input))
}
