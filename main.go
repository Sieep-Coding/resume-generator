package main

import (
	"fmt"

	"resume-generator/model"

	"resume-generator/generator"
)

func main() {
	// Prompt user for input
	resume := model.GetResumeFromUser()

	// Generate HTML resume
	htmlResume := generator.GenerateHTMLResume(resume)

	// Save HTML resume to file
	generator.SaveResumeToFile(htmlResume)

	fmt.Println("Resume generated successfully!")
}
