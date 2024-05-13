package subject

// Service é responsável por fornecer métodos para manipulação de disciplinas.
type Service struct {
	repo *Repository // Repositório para acessar e manipular os dados das disciplinas.
}

// NewService cria uma nova instância de Service com o repositório fornecido.
func NewService(repository *Repository) *Service {
	return &Service{repo: repository}
}

// GetByStudentID retorna todas as disciplinas associadas a um determinado ID de estudante.
func (s *Service) GetByStudentID(studentID int) ([]Subject, error) {
	// Chama o método GetByStudentID do repositório para recuperar as disciplinas associadas ao ID do estudante.
	return s.repo.GetByStudentID(studentID)
}
