package main

import (
	"fmt"
	"little_grade_calculator/calculator"
)

func main() {

	correct, wrong, total, grade := calculator.CalculateGrade("tftttttttttttttttttttttftffffffttttf")

	fmt.Printf("Correct answers: %d \n", correct)
	fmt.Printf("Wrong answers: %d \n", wrong)
	fmt.Printf("Total: %d \n", total)

	fmt.Println("------------------------------------")
	fmt.Printf("Your grade was: %d/%d \n\n", correct, total)
	fmt.Printf("         %.2f%% \n\n", grade)

	if grade >= 80 {
		fmt.Println("🎉🎉🎉 You did! 🎉🎉🎉")
	} else {
		fmt.Println("You failed, try again !")
	}
}
