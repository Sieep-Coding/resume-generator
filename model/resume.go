package model

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Resume represents the structure of a resume
type Resume struct {
	Name       string
	Email      string
	Phone      string
	Summary    string
	Education  []Education
	Experience []Experience
	Skills     []string
}

// Education represents an education entry in the resume
type Education struct {
	Degree      string
	Institution string
	StartYear   int
	EndYear     int
}

// Experience represents a work experience entry in the resume
type Experience struct {
	Title       string
	Company     string
	StartDate   string
	EndDate     string
	Description string
}

// GetResumeFromUser prompts the user for input and creates a Resume struct
func GetResumeFromUser() Resume {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name := readLine(reader)

	fmt.Print("Enter your email: ")
	email := readLine(reader)

	fmt.Print("Enter your phone number: ")
	phone := readLine(reader)

	fmt.Print("Enter your summary: ")
	summary := readLine(reader)

	fmt.Print("Enter the number of education entries: ")
	numEducation := readInt(reader)
	education := make([]Education, numEducation)
	for i := 0; i < numEducation; i++ {
		fmt.Printf("Education entry %d:\n", i+1)
		fmt.Print("  Degree: ")
		degree := readLine(reader)
		fmt.Print("  Institution: ")
		institution := readLine(reader)
		fmt.Print("  Start year: ")
		startYear := readInt(reader)
		fmt.Print("  End year: ")
		endYear := readInt(reader)
		education[i] = Education{
			Degree:      degree,
			Institution: institution,
			StartYear:   startYear,
			EndYear:     endYear,
		}
	}

	fmt.Print("Enter the number of experience entries: ")
	numExperience := readInt(reader)
	experience := make([]Experience, numExperience)
	for i := 0; i < numExperience; i++ {
		fmt.Printf("Experience entry %d:\n", i+1)
		fmt.Print("  Title: ")
		title := readLine(reader)
		fmt.Print("  Company: ")
		company := readLine(reader)
		fmt.Print("  Start date: ")
		startDate := readLine(reader)
		fmt.Print("  End date: ")
		endDate := readLine(reader)
		fmt.Print("  Description: ")
		description := readLine(reader)
		experience[i] = Experience{
			Title:       title,
			Company:     company,
			StartDate:   startDate,
			EndDate:     endDate,
			Description: description,
		}
	}

	fmt.Print("Enter your skills (comma-separated): ")
	skillsInput := readLine(reader)
	skills := strings.Split(skillsInput, ",")
	for i := range skills {
		skills[i] = strings.TrimSpace(skills[i])
	}

	return Resume{
		Name:       name,
		Email:      email,
		Phone:      phone,
		Summary:    summary,
		Education:  education,
		Experience: experience,
		Skills:     skills,
	}
}

// readLine reads a line of text from the reader
func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// readInt reads an integer from the reader
func readInt(reader *bufio.Reader) int {
	input := readLine(reader)
	num, _ := strconv.Atoi(input)
	return num
}
