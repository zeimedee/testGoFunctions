package stringCharaters

import "fmt"

func ConvertStringToByteArray (input string){
	var series []byte = []byte(input)

	
	fmt.Println(series)
}

func FirstStringCharater (input string){
	var series []byte = []byte(input)

	firstCharater := string(series[0])
	
	fmt.Println(firstCharater)
}

func ReplaceCharater(input string, charaterReplacementCount int, newCharater string){
	var series []byte = []byte(input)
	var replaceCharater []byte = []byte(newCharater) 

	fmt.Printf(" %v %T\n", series[1], series[1])

	series[charaterReplacementCount - 1] = replaceCharater[0]

	fmt.Printf(" %v %T\n", series[charaterReplacementCount - 1], series[charaterReplacementCount - 1])

	fmt.Printf("New String : %v\n", string(series))
}