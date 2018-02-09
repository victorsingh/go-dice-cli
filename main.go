package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"reflect"
)

type Die struct {
	X numObject
}

type numObject struct {
	value int
	prob float64
	nextVal *numObject
}

func(v *Die) addNewValue(val int, prob float64) {
	v.X.nextVal = &numObject{val, prob, nil}
}

// Goal: Make die roll object oriented and functional

func main() {
	fmt.Println("Hey there, welcome to the die maker")
	
	buf := bufio.NewScanner(os.Stdin)
	bufExit := bufio.NewScanner(os.Stdin)

	for bufExit.Text() != "Exit" {
		fmt.Println("Enter the value here:")
		buf.Scan()
		desiredValue := buf.Text()

		fmt.Println("Now enter the probability of that value")
		buf.Scan()
		desiredProbability := buf.Text()

		convertDesireProb, err := strconv.ParseFloat(desiredProbability, 64)

		if err != nil {
			// handle error
			fmt.Print("You did not enter a proper value!")
			continue
		}

		if convertDesireProb >= 1 {
			fmt.Println("This cannot be a probability")
			break
		}

		fmt.Println(desiredValue, desiredProbability, reflect.TypeOf(convertDesireProb))
		v := Die{numObject{3, float64(2)/3, nil}}
		fmt.Println(v.X)

		fmt.Println("Enter Exit to end program (soon here will have an option to roll)")
		bufExit.Scan()
	}
	
}