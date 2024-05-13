package produtos

// ProdutoService é responsável por fornecer métodos para manipulação de produtos.
type ProdutoService struct {
	repo Repository // Repositório para acessar e manipular os dados dos produtos.
}

// NewProdutoService cria uma nova instância de ProdutoService com o repositório fornecido.
func NewProdutoService(repo Repository) *ProdutoService {
	return &ProdutoService{repo: repo}
}

// List retorna uma lista de todos os produtos.
func (s *ProdutoService) List() ([]*Produto, error) {
	return s.repo.List()
}

// Get retorna um produto com o ID fornecido.
func (s *ProdutoService) Get(id int) (*Produto, error) {
	return s.repo.GetByID(id)
}

// Create cria um novo produto.
func (s *ProdutoService) Create(produto Produto) (int64, error) {
	return s.repo.Create(&produto)
}

// Update atualiza as informações de um produto existente.
func (s *ProdutoService) Update(produto Produto) error {
	return s.repo.Update(produto)
}

// Delete exclui um produto com o ID fornecido.
func (s *ProdutoService) Delete(id int) error {
	return s.repo.Delete(id)
}
