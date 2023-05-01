/*

Пояснение к коду: В консоль при запуске кода нужно подавать только одно число.
ВАЖНО: число должно быть целое. Чтобы завершить ввод чисел необходимо ввести
число 0.

*/


package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    // открытие файла для записи
    file, err := os.Create("numbers.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // запись последовательности чисел в файл
    var n,score int
    score = 0
    for {
        	fmt.Print("Введите число (0 - конец): ")
        	fmt.Scan(&n)
        	if n == 0 {
        		break
        	}
        	fmt.Fprintln(file, n)
        	score += 1
    }

    // открытие файла для чтения
    file, err = os.Open("numbers.txt")
    if err != nil {
    	fmt.Println(err)
    	return
    }
    defer file.Close() // Закрытие файла после завершения работы
    
    scanner := bufio.NewScanner(file)

    var numbers []int
    for scanner.Scan() {
        number, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println("Вы написали букву", err)
            return
        }
        numbers = append(numbers, number)
    }

    fmt.Println("Введенные числа:", numbers)
    
    // Поиск первого положительного числа и последнего отрицательного числа в последовательности
    var positive int
	var positivePosition int
	var negative int
	var negativePosition int

	for i, v := range numbers {
		if v > 0 && positive == 0 {
			positive = v
			positivePosition = i
		}
		if v < 0 {
			negative = v
			negativePosition = i
		}
	}
	fmt.Printf("Первое положительное число [%d] в позиции [%d]\n", positive, positivePosition)
	fmt.Printf("Последнее отрицательное число [%d] в позиции [%d]\n ", negative, negativePosition)
}
