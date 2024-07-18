//Вы собираете группу добровольцев и подвергаете их нагрузкам, имитирующим перегрузки при полете на корабле. После чего - измеряете пульс этих людей. В случае, если пульс кандидатов находится в диапазоне от 100 ударов в минуту до 140 (обе границы включительно) - вам подходит этот кандидат.
//
//Напишите программу, которая поможет вам решить проблему отбора кандидатов.
//
//Формат ввода
//Программа принимает на вход несколько вещественных чисел - значения пульса кандидатов (каждое с новой строки), в конце идет сигнал остановки - любое вещественное отрицательное число.
//
//Формат вывода
//Программа выводит следующую информацию - количество успешно отобранных кандидатов и на другой строке - минимальное и максимальное значение пульса отобранных кандидатов.
//Вывод
//92.5
//120.4
//139.2
//140.0
//99.9
//-1
//Вывод
//3
//120.4 140.0

package main

import (
	"fmt"
	"strings"
)

func main() {
	var onePass, secPass string

	for {
	outer:
		for {
			fmt.Print("Enter the password:")
			fmt.Scan(&onePass)
			if len(onePass) < 8 {
				fmt.Println("Слишком короткий пароль!")
				break outer
			} else if strings.Contains(onePass, "123") || strings.Contains(onePass, "qwe") {
				fmt.Println("Слишком простой пароль!")
				break outer
			}
			fmt.Print("Resume password:")
			fmt.Scan(&secPass)
			if onePass == secPass {
				fmt.Println("Ну наконец-то")
				return
			} else if onePass != secPass {
				fmt.Println("Введенные пароли различаются!")
			}
		}
	}
}
