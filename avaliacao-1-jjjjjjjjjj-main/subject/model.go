package subject

// Subject representa uma disciplina com seus atributos.
type Subject struct {
	Id       int64  `json:"id"`       // ID da disciplina
	Name     string `json:"name"`     // Nome da disciplina
	WordLoad int    `json:"workload"` // Carga horária da disciplina
}
