package main

import "sort"

type Student struct {
	ID     int
	Name   string
	Grades []float64
}

func (s Student) Average() float64 {
	total := 0.0
	for _, grade := range s.Grades {
		total += grade
	}
	avg := total / float64(len(s.Grades))
	return avg
}

func (s Student) Grade() string {
	avg := s.Average()
	switch {
	case avg > 9:
		return "A"
	case avg > 8:
		return "B"
	case avg > 7:
		return "C"
	case avg > 6:
		return "D"
	case avg > 5:
		return "E"
	default:
		return "F"
	}
}

type Registry struct {
	students []Student
	nextID   int
}

func (r *Registry) Add(name string, grades ...float64) int {
	student := Student{r.nextID, name, grades}
	r.students = append(r.students, student)
	r.nextID++
	return student.ID
}

func (r *Registry) FindByID(id int) (*Student, bool) {
	if id < 0 || id >= r.nextID {
		return nil, false
	}
	for i := 0; i < len(r.students); i++ {
		if id == r.students[i].ID {
			return &r.students[i], true
		}
	}
	return nil, false
}

func (r *Registry) TopN(n int) []Student {
	if len(r.students) < n {
		n = len(r.students)
	}
	sorted := make([]Student, len(r.students))
	copy(sorted, r.students)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Average() > sorted[j].Average()
	})
	return sorted[:n]
}
