package main

import (
	"fmt"
	"time"
)

// 28. Чтобудет выведено в результате выполнения это кода?

// Ответы:
// Ни чего, программазавершится раньше выводв текста на экран
// т.е. зависит от реализации компилятора и работы планировщика

func main() {
	for i:=0;i<10;i++ {
		go func() {
			fmt.Println(i)
		}()
		time.Sleep(time.Nanosecond)
	}
}

