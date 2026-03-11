package grades

type StudentEntity struct {
	ID        int
	FirstName string
	LastName  string
}

type GradeEntity struct {
	ID        int
	StudentID int
	Title     string
	Type      GradeType
	Score     float32
}
