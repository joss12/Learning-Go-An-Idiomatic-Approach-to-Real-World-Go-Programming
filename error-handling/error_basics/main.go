package main


import (
	"fmt" 
"errors"
"os"
)


func calcReminderAndMod(numerator, denominator int)(int, int, error){
	if denominator == 0{
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main(){
	numerator := 20
	denominator := 3
	reminder, mod, err := calcReminderAndMod(numerator, denominator)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(reminder, mod)
}