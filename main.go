package main

import "fmt"

func main() {

	firstArr := NewArr(SizeArray())                // создаем новый массив величиной SizeArray
	secondArr := NewArr(SizeArray())               // создаем новый массив величиной SizeArray
	firstArr.Set()                                 // добавляем элементы в массив
	secondArr.Set()                                // добавляем элементы в массив
	fmt.Println(Intersection(firstArr, secondArr)) // выводим одинаковые элементы массивов

}

type IntArray struct {
	arr []string
}

func NewArr(size int) *IntArray {
	return &IntArray{
		arr: make([]string, size),
	}
}

var count int = 0                        // счетчик для вызова значений first, second
var value = [2]string{"first", "second"} // массив с элементами first, second

func SizeArray() int {
	var sizeArr int
	// делаем проверку на счетчик, если он 2 и более - сбрасываем счетчик на 0
	if count > 1 {
		count = 0
	}
	if count < 2 {
		// если счетчик 0, 1 - выводим значение из переменной value по индексу
		fmt.Printf("Enter %v array size: \n", value[count])
		fmt.Scanln(&sizeArr)
		count++
	}

	return sizeArr
}

func (in *IntArray) Set() {
	n := len(in.arr)
	// делаем проверку на счетчик, если он 2 и более - сбрасываем счетчик на 0
	if count > 1 {
		count = 0
	}
	if count < 2 {
		// если счетчик 0, 1 - выводим значение из переменной value по индексу
		fmt.Printf("Enter %v array size: \n", value[count])
		for i := 0; i < n; i++ {
			var value string
			fmt.Scanln(&value)
			in.arr[i] = value
		}
	}
	count++

}

// Intersection - выводит список одинаковых элементов массивов
func Intersection(first, second *IntArray) []string {
	newArr := make(map[int]string, 0)
	tmp := 0
	for i := range second.arr {
		for j := 0; j < len(first.arr); j++ {
			if second.arr[i] == first.arr[j] {
				newArr[tmp] = second.arr[i]
				tmp++
			}
		}
	}
	n := len(newArr)
	convertArr := make([]string, n)
	for i := 0; i < n; i++ {
		convertArr[i] = newArr[i]
	}
	return convertArr
}
