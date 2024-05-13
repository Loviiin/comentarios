package produtos

// Produto representa um produto com seus atributos.
type Produto struct {
	ID          int     `json:"id"`    // ID do produto
	Name        string  `json:"name"`  // Nome do produto
	Price       float64 `json:"price"` // Preço do produto
	Description string  `json:"desc"`  // Descrição do produto
}
