package main

import (
	"fmt"
	"math/rand"
)

//find random item from slice
type Items struct {
	name string
	age int
}

var items = []Items{
	{name: "Noka", age: 30},
	{name: "Baco", age: 28},
	{name: "Gio", age: 25},
	{name: "Sandro", age: 30},
	{name: "Ana", age: 17},
}
func randomItem() {
	var random = rand.Intn(len(items))
    fmt.Printf("Random item: %v\n", items[random])
}

//Write a program that calculates the area of a rectangle. Prompt the user to enter the length and width, then calculate and display the area.

func calcArea(){
	var height int
	var width int
	fmt.Print("Enter height: ")
	fmt.Scan(&height)
	fmt.Print("Enter width: ")
	fmt.Scan(&width)
	fmt.Printf("Area of rectangle: %v\n", height * width)
}


//Write a program that converts temperature from Celsius to Fahrenheit. Prompt the user to enter the temperature in Celsius, then calculate and display the temperature in Fahrenheit.
func convertTemp(){
	 var celsius int
	 fmt.Print("Enter temperature in Celsius: ")
	 fmt.Scan(&celsius)
	 fmt.Printf("Temperature in Fahrenheit: %v\n", (celsius * 9/5) + 32)
}


//Write a program that checks if a number entered by the user is even or odd
 func checkEvenOdd(){
	var number uint
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	if number % 2 == 0 {
		fmt.Printf("%v is even\n", number)
	} else {
		fmt.Printf("%v is odd\n", number)
	}
 }

 //Write a program that checks if a year entered by the user is a leap year or not
 func isLeapYear() {
	var year uint
	fmt.Print("Enter a year: ")
	fmt.Scan(&year)
	if year % 4 == 0 {
		fmt.Printf("%v is a leap year\n", year)
	} else {
		fmt.Printf("%v is not a leap year\n", year)
	}
 }


 //Write a program that prints the first 10 natural numbers.
 func printNaturalNumbers(){
	 for i := 1; i <= 10; i++ {
		 fmt.Printf("%v\n", i)
	 }
 }

 //Write a program that calculates the factorial of a number entered by the user using a loop
func factorial(){
	var number int
	var result int = 1
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	for i := 1; i <= number; i++ {
		result *= i
	}
	fmt.Printf("Factorial of %v is %v\n", number, result)

}

//Write a program that calculates the sum of the first 100 numbers.
func sumOfFirst100Numbers(){
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("Sum of first 100 numbers: %v\n", sum)

}

//Write a function add that takes two integers as parameters and returns their sum
func add(a int, b int)int  {
	 sum := a + b
	fmt.Printf("Sum  is %v\n", sum)
	return sum

}
//Write a function isPrime that takes an integer as a parameter and returns true if it's a prime number, otherwise returns false.
func isPrime(number int) bool {
 var numbers = []int{2, 3, 5, 7}
	for _, i := range numbers{
		if number % i == 0 {
			return false
		}
	}
	return true
}


//Write a program that finds the maximum and minimum elements in an array of integer
func printArray(){
	var numbers = [5]int{1, 2, 3, 4, 5}
	max:= numbers[0]
	min:= numbers[0]
	for i := 0; i < len(numbers); i++ {
		if(numbers[i] > max){
			max = numbers[i]	
		}
		if(numbers[i] < min){
			min = numbers[i]	
		}
	}
	fmt.Printf("Max: %v\n", max)
	fmt.Printf("Min: %v\n", min)
}

//Declare a slice of strings and append new strings to it. Print the length and elements of the slice.
func appendStrings(){
	var names = []string{"Noka", "Baco", "Gio"}
	names = append(names, "Sandro", "Ana")
	fmt.Printf("Length: %v\n", len(names))
}

//Declare a map that stores the names of some countries and their capitals. Print each key-value pair.
func printMap(){
	countries := map[string]string{
		"Georgia":   "Tbilisi",
		"Germany":  "Berlin",
		"France":    "Paris",
		"Spain":     "Madrid",
	}
	for key, value := range countries {
		fmt.Printf("Country: %v, Capital: %v\n", key, value)
	}
}

//Write a program that counts the frequency of each character in a given string using a map.
 func countFrequency(){
	 var str = "Hello, World!"
	 var frequency = make(map[rune]int)
	 for _, char := range str {
		 frequency[char]++
	 }
	 fmt.Println(frequency)
 }

 //Define a struct type Person with fields name, age, and city. Create a few Person objects and print their details
type Person struct {
	name string
	age int
	city string

}
func printPerson(){
	person1 := Person{name: "Noka", age: 30, city: "Tbilisi"}
	person2 := Person{name: "Baco", age: 28, city: "Tbilisi"}
	fmt.Printf("Person1: %v\n", person1)
	fmt.Printf("Person2: %v\n", person2)

}

 func main() {
	//randomItem()
	// calcArea()
	//convertTemp()
	//checkEvenOdd()
	//isLeapYear()
	//printNaturalNumbers()
	// factorial()
	//sumOfFirst100Numbers()
	//add(10, 20)
	//fmt.Println(isPrime(7))
	//printArray()
	//appendStrings()
	//printMap()
	//printPerson()
}
