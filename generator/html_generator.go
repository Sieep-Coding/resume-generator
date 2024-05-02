package generator

import (
	"fmt"
	"os"
	"strings"

	"resume-generator/model"
)

// GenerateHTMLResume generates an HTML version of the resume
func GenerateHTMLResume(resume model.Resume) string {
	var sb strings.Builder

	sb.WriteString("<!DOCTYPE html>\n")
	sb.WriteString("<html>\n")
	sb.WriteString("<head>\n")
	sb.WriteString("  <title>Resume</title>\n")
	sb.WriteString("  <style>\n")
	sb.WriteString("    body { font-family: Arial, sans-serif; }\n")
	sb.WriteString("    h1, h2 { margin-bottom: 10px; }\n")
	sb.WriteString("    ul { margin-top: 5px; }\n")
	sb.WriteString("  </style>\n")
	sb.WriteString("</head>\n")
	sb.WriteString("<body>\n")

	sb.WriteString(fmt.Sprintf("  <h1>%s</h1>\n", resume.Name))
	sb.WriteString(fmt.Sprintf("  <p>Email: %s</p>\n", resume.Email))
	sb.WriteString(fmt.Sprintf("  <p>Phone: %s</p>\n", resume.Phone))

	sb.WriteString("  <h2>Summary</h2>\n")
	sb.WriteString(fmt.Sprintf("  <p>%s</p>\n", resume.Summary))

	sb.WriteString("  <h2>Education</h2>\n")
	sb.WriteString("  <ul>\n")
	for _, edu := range resume.Education {
		sb.WriteString(fmt.Sprintf("    <li>%s, %s (%d-%d)</li>\n", edu.Degree, edu.Institution, edu.StartYear, edu.EndYear))
	}
	sb.WriteString("  </ul>\n")

	sb.WriteString("  <h2>Experience</h2>\n")
	for _, exp := range resume.Experience {
		sb.WriteString("  <div>\n")
		sb.WriteString(fmt.Sprintf("    <h3>%s, %s</h3>\n", exp.Title, exp.Company))
		sb.WriteString(fmt.Sprintf("    <p>%s - %s</p>\n", exp.StartDate, exp.EndDate))
		sb.WriteString(fmt.Sprintf("    <p>%s</p>\n", exp.Description))
		sb.WriteString("  </div>\n")
	}

	sb.WriteString("  <h2>Skills</h2>\n")
	sb.WriteString("  <ul>\n")
	for _, skill := range resume.Skills {
		sb.WriteString(fmt.Sprintf("    <li>%s</li>\n", skill))
	}
	sb.WriteString("  </ul>\n")

	sb.WriteString("</body>\n")
	sb.WriteString("</html>\n")

	return sb.String()
}

// SaveResumeToFile saves the HTML resume to a file
func SaveResumeToFile(htmlResume string) error {
	file, err := os.Create("resume.html")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(htmlResume)
	if err != nil {
		return err
	}

	return nil
}
