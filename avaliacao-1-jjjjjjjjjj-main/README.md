[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/ac6fRa9I)
# Avaliacao 1

Crie uma API RESTful que possua as seguintes funcionalidades:

1. Rota `/products` que permita criar um novo produto (método POST), listar todos os produtos (método GET) e buscar um produto pelo ID (método GET).
2. Rota `/products/{product_id}` que permita atualizar os dados de um produto existente (método PUT) e excluir um produto (método DELETE).
3. Rota `/categories` que permita criar uma nova categoria (método POST), listar todas as categorias (método GET) e buscar uma categoria pelo ID (método GET).
4. Rota `/categories/{category_id}` que permita atualizar os dados de uma categoria existente (método PUT) e excluir uma categoria (método DELETE).
5. Rota `/products/{product_id}/categories` que permita adicionar uma categoria a um produto (método POST) e remover uma categoria de um produto (método DELETE).
6. Rota `/categories/{category_id}/products` que permita listar todos os produtos de uma determinada categoria (método GET).
7. Implemente o CORS para limitar o acesso da API somente a domínios específicos.

Além das funcionalidades acima, a sua API deve ter as seguintes características:

* Utilizar middlewares para adicionar um cabeçalho `X-Powered-By` a todas as respostas da API, contendo o valor `GoAPI`.
* Utilizar HTTP caching para melhorar o desempenho da API, evitando que recursos sejam carregados repetidamente.
* Utilizar logging para registrar as ações realizadas na API.

Você será avaliado pela sua capacidade de:

* Criar rotas no Golang;
* Implementar os métodos HTTP necessários para cada rota;
* Implementar o CORS para garantir a segurança da API;
* Utilizar middlewares para adicionar um cabeçalho `X-Powered-By` a todas as respostas da API;
* Utilizar HTTP caching para melhorar o desempenho da API;
* Utilizar logging para registrar as ações realizadas na API.
* Escrever código limpo e bem organizado.

Você pode utilizar o MySQL para realizar a persistência dos dados.
