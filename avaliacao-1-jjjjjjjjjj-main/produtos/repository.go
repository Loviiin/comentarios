package produtos

import "database/sql"

// Repository define a interface para operações de banco de dados relacionadas aos produtos.
type Repository interface {
	Create(product *Produto) (int64, error) // Criar um novo produto no banco de dados.
	GetByID(id int) (*Produto, error)       // Obter um produto pelo ID do banco de dados.
	List() ([]*Produto, error)              // Listar todos os produtos do banco de dados.
	Update(product *Produto) error          // Atualizar as informações de um produto no banco de dados.
	Delete(id int) error                    // Excluir um produto do banco de dados pelo ID.
}

// repository é uma implementação da interface Repository.
type repository struct {
	db *sql.DB // Conexão com o banco de dados.
}

// NewRepository cria uma nova instância de Repository com a conexão fornecida.
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// Create insere um novo produto no banco de dados e retorna o ID do produto inserido.
func (r *repository) Create(product *Produto) (int64, error) {
	result, err := r.db.Exec(`INSERT INTO products(name, price, description)
                              VALUES (?, ?, ?)`,
		product.Name, product.Price, product.Description)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetByID obtém um produto do banco de dados pelo ID fornecido.
func (r *repository) GetByID(id int) (*Produto, error) {
	row := r.db.QueryRow(`SELECT product_id, name, price, description
                          FROM products WHERE product_id = ?`, id)
	var product Produto
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Description)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// List retorna todos os produtos do banco de dados.
func (r *repository) List() ([]*Produto, error) {
	rows, err := r.db.Query(`
		SELECT product_id, name, price, description
		FROM products
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Produto
	for rows.Next() {
		var product Produto
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

// Update atualiza as informações de um produto no banco de dados.
func (r *repository) Update(product *Produto) error {
	_, err := r.db.Exec(`UPDATE products SET name = ?, price = ?, description = ? WHERE product_id = ?`,
		product.Name, product.Price, product.Description, product.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete exclui um produto do banco de dados pelo ID.
func (r *repository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM products WHERE product_id = ?`, id)
	if err != nil {
		return err
	}

	return nil
}
