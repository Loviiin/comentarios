package main

import (
	"avaliacao-1/produtos"
	"avaliacao-1/subject"
	"avaliacao-1/user"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	if err := createServer(); err != nil {
		log.Panic(err)
	}
}

func connectDB() *sql.DB {
	// Implemente a conexão com o banco de dados aqui.
}

func createServer() error {
	db := connectDB()

	subjectRepository := subject.NewRepository(db)
	subjectService := subject.NewService(subjectRepository)

	produtoRepository := produtos.NewRepository(db)
	produtoService := produtos.NewProdutoService(produtoRepository)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// Inicialização do roteador usando gorilla/mux.
	router := mux.NewRouter()

	// Rota /products que permite criar um novo produto (método POST), listar todos os produtos (método GET) e buscar um produto pelo ID (método GET).
	router.HandleFunc("/products", produtos.NewProdutoController(produtoService).CreateProduto).Methods("POST")
	router.HandleFunc("/products", produtos.NewProdutoController(produtoService).ListProdutos).Methods("GET")
	router.HandleFunc("/products/{id}", produtos.NewProdutoController(produtoService).GetProduto).Methods("GET")

	// Rota /products/{product_id} que permite atualizar os dados de um produto existente (método PUT) e excluir um produto (método DELETE).
	router.HandleFunc("/products/{id}", produtos.NewProdutoController(produtoService).UpdateProduto).Methods("PUT")
	router.HandleFunc("/products/{id}", produtos.NewProdutoController(produtoService).DeleteProduto).Methods("DELETE")

	// Rota /categories que permite criar uma nova categoria (método POST), listar todas as categorias (método GET) e buscar uma categoria pelo ID (método GET).
	// Rota /categories/{category_id} que permite atualizar os dados de uma categoria existente (método PUT) e excluir uma categoria (método DELETE).
	// Rota /products/{product_id}/categories que permite adicionar uma categoria a um produto (método POST) e remover uma categoria de um produto (método DELETE).
	// Rota /categories/{category_id}/products que permite listar todos os produtos de uma determinada categoria (método GET).

	// Implemente as rotas restantes aqui.

	// Configuração do CORS.
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://example.com"}, // Domínios permitidos.
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	// Adicionando middlewares.
	handler := corsHandler.Handler(router)
	handler = addPoweredByHeader(handler)
	handler = addLogging(handler)

	// Iniciar o servidor HTTP.
	fmt.Println("Servidor rodando na porta 8080...")
	return http.ListenAndServe(":8080", handler)
}

// Middleware para adicionar um cabeçalho X-Powered-By a todas as respostas da API.
func addPoweredByHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Powered-By", "GoAPI")
		next.ServeHTTP(w, r)
	})
}

// Middleware para logging.
func addLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requisição recebida: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
