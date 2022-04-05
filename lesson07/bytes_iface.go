package main

import (
	"fmt"
	"io"
)

type myBytes struct {
	content []byte
	offset  int
}

func (m *myBytes) Read(p []byte) (n int, err error) {
	if m.offset >= len(m.content) {
		return 0, io.EOF
	}

	if len(m.content[m.offset:]) > len(p) {
		n = copy(p, m.content[m.offset:m.offset+len(p)])
	} else {
		n = copy(p, m.content[m.offset:])
	}

	m.offset += n
	fmt.Printf("Read %d bytes with %d offset of %d content.\n", n, m.offset, len(m.content)) // For debug only
	return n, nil

}

func (m *myBytes) Write(p []byte) (int, error) {
	m.content = append(m.content, p...)
	fmt.Printf("Write %d bytes\n", len(p)) // For debug only
	return len(p), nil
}

func main() {
	var tst myBytes
	_, err1 := io.WriteString(&tst, "Пустой интерфейс не имеет методов, следовательно, не накладывает никаких ограничений на принимаемое значение. Однако пустой интерфейс не рекомендуется использовать, так как он убирает любые проверки типа для хранимого значения   !   \nПустой интерфейс не имеет методов, следовательно, не накладывает никаких ограничений на принимаемое значение. Однако пустой интерфейс не рекомендуется использовать, так как он убирает любые проверки типа для хранимого значения")
	if err1 != nil {
		fmt.Printf("Something wrong! %v\n", err1)
		return
	}
	fmt.Printf("I've write to myBytes via io.WriteString, myBytes object is %v\n", tst)
	result, err2 := io.ReadAll(&tst)
	if err2 != nil {
		fmt.Printf("Something wrong! %v", err2)
		return
	}
	fmt.Printf("I've read from myBytes via io.ReadAll, result is %s\n", string(result))
}
