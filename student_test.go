package students_test

import (
	"fmt"
	"main/device-api/internal/students"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStudent(t *testing.T) {
	StudentNames := []string{"joe", "bob", "dave"} // 1D STUDENT ARRAY

	StudentMarks := [][]int{ //	2D STUDENT-MARK ARRAY
		[]int{77, 83, 86}, // joe
		[]int{77, 83, 86}, // bob
		[]int{40, 15, 25}, // dave
	}

	assert.True(t, len(StudentNames) == len(StudentMarks))

	reports := make([]*students.StudentReport, 0)
	for idx := range StudentNames {

		currentStudent := students.NewStudent(StudentNames[idx], StudentMarks[idx])

		currentReport := currentStudent.FullReport()
		reports = append(reports, currentReport)
	}

	result := students.ClassReport(reports)
	for key, value := range result {
		fmt.Println("logging grade occurances for class")
		fmt.Println(key, value)
		fmt.Println(" ")
	}
}
