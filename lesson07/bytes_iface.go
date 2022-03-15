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
	n = copy(p, m.content)
	m.offset += n
	return n, nil

}

func (m *myBytes) Write(p []byte) (n int, err error) {
	mSize := len(m.content)
	m.content = append(m.content, p...)
	return len(m.content) - mSize, nil
}

func main() {
	var tst myBytes
	_, err1 := io.WriteString(&tst, "Test")
	if err1 != nil {
		fmt.Printf("Something wrong! %v\n", err1)
		return
	}
	fmt.Printf("I've write to myBytes via io.WriteString, myBytes object is %v", tst)
	result, err2 := io.ReadAll(&tst)
	if err2 != nil {
		fmt.Printf("Something wrong! %v", err2)
		return
	}
	fmt.Printf("I've read from myBytes via io.ReadAll, result is %s\n", string(result))
}

//Заглушка, постараюсь реализовать в выходные, а то не укладываюсь в срок сдачи задания.

//Необходимо объявить свой тип, обернув в него тип []byte - (слайс байтов).
//
//Затем, необходимо реализовать на нем такие методы, чтобы он удовлетворял интферйесам io.Reader (из него можно читать байты) и io.Writer (а так же писать их туда).
//Затем, используя функции пакета io:
//
//С помощью io.WriteString записать в переменную вашего типа произвольную строку
//С помощью io.ReadAll( ) считать вашу строку обратно (вообще говоря, он возвращает слайс байт, но его легко привести к виду строки)
