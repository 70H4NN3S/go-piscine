package main

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
