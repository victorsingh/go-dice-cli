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

	fmt.Println("Enter the value here:")
	buf.Scan()
	desiredValue := buf.Text()

	fmt.Println("Now enter the probability of that value")
	buf.Scan()
	desiredProbability := buf.Text()

	convertDesireProb, err := strconv.Atoi(desiredProbability)

	if err != nil {
		// handle error
	 }

	fmt.Println(desiredValue, desiredProbability, reflect.TypeOf(convertDesireProb))
	v := Die{numObject{3, float64(2)/3, nil}}
	fmt.Println(v.X)
}