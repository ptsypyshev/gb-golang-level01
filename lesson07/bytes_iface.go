package main

func main() {

}

//Заглушка, постараюсь реализовать в выходные, а то не укладываюсь в срок сдачи задания.

//Необходимо объявить свой тип, обернув в него тип []byte - (слайс байтов).
//
//Затем, необходимо реализовать на нем такие методы, чтобы он удовлетворял интферйесам io.Reader (из него можно читать байты) и io.Writer (а так же писать их туда).
//Затем, используя функции пакета io:
//
//С помощью io.WriteString записать в переменную вашего типа произвольную строку
//С помощью io.ReadAll( ) считать вашу строку обратно (вообще говоря, он возвращает слайс байт, но его легко привести к виду строки)