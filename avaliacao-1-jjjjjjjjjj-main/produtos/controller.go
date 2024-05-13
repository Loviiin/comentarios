package produtos

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ProdutoController é responsável por lidar com as requisições relacionadas aos produtos e categorias.
type ProdutoController struct {
	produtoService  ProdutoService
	categoryService CategoryService
}

// NewProdutoController cria uma nova instância de ProdutoController com os serviços fornecidos.
func NewProdutoController(produtoService ProdutoService, categoryService CategoryService) *ProdutoController {
	return &ProdutoController{
		produtoService:  produtoService,
		categoryService: categoryService,
	}
}

// ListProdutos recupera todos os produtos e os envia como resposta.
func (c *ProdutoController) ListProdutos(w http.ResponseWriter, r *http.Request) {
	produtos, err := c.produtoService.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}

// GetProduto recupera um único produto com o ID fornecido e o envia como resposta.
func (c *ProdutoController) GetProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	produto, err := c.produtoService.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produto)
}

// CreateProduto decodifica o corpo da requisição em um novo produto e o cria no serviço.
func (c *ProdutoController) CreateProduto(w http.ResponseWriter, r *http.Request) {
	var produto Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	id, err := c.produtoService.Create(produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

// UpdateProduto decodifica o corpo da requisição em um produto existente e o atualiza no serviço.
func (c *ProdutoController) UpdateProduto(w http.ResponseWriter, r *http.Request) {
	var produto Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = c.produtoService.Update(produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteProduto remove o produto com o ID fornecido do serviço.
func (c *ProdutoController) DeleteProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = c.produtoService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// ListCategorias recupera todas as categorias e as envia como resposta.
func (c *ProdutoController) ListCategorias(w http.ResponseWriter, r *http.Request) {
	categorias, err := c.categoryService.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categorias)
}

// GetCategoria recupera uma única categoria com o ID fornecido e a envia como resposta.
func (c *ProdutoController) GetCategoria(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	categoria, err := c.categoryService.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categoria)
}

// CreateCategoria decodifica o corpo da requisição em uma nova categoria e a cria no serviço.
func (c *ProdutoController) CreateCategoria(w http.ResponseWriter, r *http.Request) {
	var categoria Categoria
	err := json.NewDecoder(r.Body).Decode(&categoria)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	id, err := c.categoryService.Create(categoria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

// UpdateCategoria decodifica o corpo da requisição em uma categoria existente e a atualiza no serviço.
func (c *ProdutoController) UpdateCategoria(w http.ResponseWriter, r *http.Request) {
	var categoria Categoria
	err := json.NewDecoder(r.Body).Decode(&categoria)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = c.categoryService.Update(categoria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteCategoria remove a categoria com o ID fornecido do serviço.
func (c *ProdutoController) DeleteCategoria(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	err = c.categoryService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
