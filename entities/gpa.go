package entities

type GPAGroup struct {
	MajorGPA     float64 `json:"majorGPA"`
	MajorScore   float64 `json:"majorScore"`
	OverallGPA   float64 `json:"overallGPA"`
	OverallScore float64 `json:"overallScore"`
	InitialGPA   float64 `json:"initialGPA"`
	InitialScore float64 `json:"initialScore"`
}

func NewGPAGroup() *GPAGroup {
	return &GPAGroup{0, 0, 0, 0, 0, 0}
}
