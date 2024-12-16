package main

import "fmt"

type Animal struct {
	Type string
	Name string
	Age  int
}

func getAnimals() []Animal {
	return []Animal{
		{Type: "Корова", Name: "Боря", Age: 2},
		{Type: "Собака", Name: "Кузя", Age: 4},
		{Type: "Кошка", Name: "Маркиз", Age: 5},
	}
}

func preparePrint(animals []Animal) string {
	var result string

	for _, animal := range animals {
		result += fmt.Sprintf("Тип: %s, Имя: %s, Возраст: %d\n", animal.Type, animal.Name, animal.Age)
	}
	return result
}

func main() {
	animals := getAnimals()
	result := preparePrint(animals)
	fmt.Println(result)
}
