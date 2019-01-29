package main

import (
	"fmt"

	"github.com/carabasdaniel/lesson1/operations"
	//	v "github.com/carabasdaniel/lesson1/validator"
)

func main() {
	//	var word1, word2 string
	//	var x int

	//	word1 = "test"
	//	word2 = "tST"
	//	//	fmt.Printf("Word 1=")
	//	fmt.Scanf("%s", &word1)
	//	fmt.Printf("Word 2=")
	//	fmt.Scanf("%s", &word2)
	//	fmt.Println(x)
	//	fmt.Println(funnelCompare(v.NewAsisvalidator(), word1, word2))

	//	fmt.Println(funnelCompare(v.NewLowerCaseValidator(), word1, word2))

	//	a := 123
	//	fmt.Printf("Initial value %d\n", a)
	//	changeValue(&a)
	//	fmt.Println("changed value ", a)
	parsedData, err := operations.ParseSampleJSON()
	if err != nil {
		panic(err)
	}

	fmt.Println(parsedData)

	fmt.Println(parsedData.Content.Items)

	//	var sampleMap map[string]string

	//	sampleMap = make(map[string]string)
	//	sampleMap["key1"] = "val1"
	//	sampleMap["key2"] = "val2"
	//	sampleMap["key3"] = "val3"

	//	for i := 0; i < 3; i++ {
	//		mapShow(sampleMap)
	//		fmt.Println("-----")
	//	}

	//	//	//check if map has key
	//	if value, ok := sampleMap["key5"]; ok {
	//		fmt.Println(value)
	//	} else {
	//		fmt.Println("No such key")
	//	}
}

func changeValue(number *int) {
	*number++
}

func mapShow(data map[string]string) {
	for key, value := range data {
		fmt.Printf("%s %s\n", key, value)
	}
}
