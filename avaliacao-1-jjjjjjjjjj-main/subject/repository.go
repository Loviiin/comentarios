package subject

import "database/sql"

// Repository gerencia as operações de banco de dados relacionadas às disciplinas.
type Repository struct {
	db *sql.DB // Conexão com o banco de dados.
}

// NewRepository cria uma nova instância de Repository com a conexão fornecida.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// GetByStudentID retorna todas as disciplinas associadas a um determinado ID de estudante.
func (r *Repository) GetByStudentID(studentID int) ([]Subject, error) {
	// Consulta SQL para selecionar disciplinas associadas a um determinado ID de estudante.
	rows, err := r.db.Query(`
		SELECT su.id,
		       su.name,
		       su.workload
		FROM subjects su
		         JOIN students_subjects ss ON su.id = ss.subject_id
		WHERE ss.student_id = ?`, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []Subject
	// Iterar sobre cada linha do resultado da consulta.
	for rows.Next() {
		var subject Subject
		// Extrair os valores da linha e atribuí-los à estrutura Subject.
		err := rows.Scan(&subject.Id, &subject.Name, &subject.WordLoad)
		if err != nil {
			return nil, err
		}
		// Adicionar a disciplina à lista de disciplinas.
		subjects = append(subjects, subject)
	}

	return subjects, nil
}
