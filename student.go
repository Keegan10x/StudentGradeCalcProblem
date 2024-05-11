package students

/* STUDENT LOGIC */
type Student struct {
	name  string
	marks []int
}

func NewStudent(inputName string, inputMarks []int) *Student {
	s := &Student{
		name:  inputName,
		marks: inputMarks,
	}
	return s
}

func (s *Student) TotalMark() int {
	accumulator := 0
	for _, mark := range s.marks {
		accumulator += mark
	}
	return accumulator
}

func (s *Student) AverageMark() int {
	total := s.TotalMark()
	return total / len(s.marks)
}

func (s *Student) FinalGrade(mark int) string {
	if mark >= 70 {
		return "distinction"
	} else if mark >= 55 && mark < 70 {
		return "merit"
	} else if mark >= 40 && mark < 55 {
		return "pass"
	} else {
		return "fail"
	}
}

func (s *Student) FullReport() *StudentReport {
	r := &StudentReport{
		name:    s.name,
		total:   s.TotalMark(),
		average: s.AverageMark(),
		grade:   s.FinalGrade(s.AverageMark()),
	}
	return r
}

type StudentReport struct {
	name    string
	total   int
	average int
	grade   string
}

/* COUNT ALL REPORTS */
func ClassReport(reports []*StudentReport) map[string]int {
	gradeTracker := make(map[string]int)
	for _, report := range reports {
		_, exists := gradeTracker[report.grade]
		if !exists {
			gradeTracker[report.grade] = 1
		} else {
			gradeTracker[report.grade]++
		}
	}
	return gradeTracker
}
