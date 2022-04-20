package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//First Two Output to the User
	fmt.Println("The BMI Calculator with Golang")
	fmt.Println(strings.Repeat("=", 30))

	//Request for Data from User
	fmt.Print("Enter Your Weight(kg): ")
	weightInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter Your Height(m): ")
	heightInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	//Trimming Unwanted Strings gotten from input
	heightInput = strings.Trim(heightInput, "\r\n")
	weightInput = strings.Replace(weightInput, "\r\n", "", -1)

	//Convert input text to float64
	weight, err := strconv.ParseFloat(weightInput, 64)
	if err != nil {
		log.Fatal(err)
	}

	height, err := strconv.ParseFloat(heightInput, 64)
	if err != nil {
		log.Fatal(err)
	}

	//Calculating the BMI
	bmi := weight / (height * height)

	//Printing out final Value
	fmt.Printf("Your Body Mass Index(BMI) is: %.3f\n", bmi)

	//Final Prompt showing BMI Class
	if bmi < 16. {
		fmt.Println("You are Underweight(Severe Thinness)!!!")
	} else if bmi <= 16.9 {
		fmt.Println("You are Underweight(Moderate Thinness)!!")
	} else if bmi <= 18.4 {
		fmt.Println("You are Underweight(Mild Thinness)!")
	} else if bmi <= 24.9 {
		fmt.Println("You are in the Normal Range")
	} else if bmi <= 29.9 {
		fmt.Println("You are overweight(Pre-obese)!")
	} else if bmi <= 34.9 {
		fmt.Println("You are Obese(Class I)!!")
	} else if bmi <= 39.9 {
		fmt.Println("You are Obese(Class II!!")
	} else {
		fmt.Println("You are Obese(Class III))!")
	}
}
