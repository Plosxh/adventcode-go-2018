package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main(){
	var bytesInputs []byte
	var frequencies []int
	duplicateFound := false
	bytesInputs, errRead := ioutil.ReadFile("./input.txt")

	if errRead != nil{
		fmt.Println(errRead)
	}
	inputs := strings.Split(string(bytesInputs),"\n")
	result := 0
	for duplicateFound != true {
		for i:=0; i<len(inputs); i++ {
			if strings.Contains(inputs[i],"-") {
				inputsInt, errConv :=strconv.Atoi(inputs[i][1:])
				if errConv != nil {
					fmt.Println(errConv)
				}
				result -= inputsInt
				if Contains(frequencies, result) {
					fmt.Print("found duplicate : ")
					fmt.Print(result)
					duplicateFound = true
					break
				}else {
					frequencies = append(frequencies,result)
				}
			} else {
				inputsInt, errConv :=strconv.Atoi(inputs[i][1:])
				if errConv != nil {
					fmt.Println(errConv)
				}
				result += inputsInt
				if Contains(frequencies, result) {
					fmt.Print("found duplicate : ")
					fmt.Print(result)
					duplicateFound = true
					break
				}else {
					frequencies = append(frequencies,result)
				}
			}

		}
	}
	fmt.Print("result : ")
	fmt.Print(result)
}

func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
