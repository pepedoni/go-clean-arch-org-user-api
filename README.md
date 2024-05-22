# go-clean-arch-org-user-api

O projeto foi implementado seguindo a arquitetura Clean Architeture. 
Desse modo, o dominio da aplicação foi separado na pasta domain e 
as ferramentas e tecnologias utilizadas foram agrupadas na pasta
infra de modo que a logica do negocio não estão acoplada as ferramentas.


## Framework web
Foi usado o framework web gin e as rotas foram configuradas no arquivo 
`infra/api/routes.go`. Nesse arquivo também as dependencias foram criadas 
e injetadas conforme necessário.


## Cache
As rotas rotas GET `/api/organizations/:orgId`, GET `/api/organizations` e GET `/api/users/:userId`
foram cachedas utilizando o redis. Sendo que as rotas de PUT e POST invalidam o cache 
seguindo regras especificias para cada caso. Esse processo foi feito por meio do arquivo `/infra/middlewares/redis_middleware`.
Dessa forma, antes de chamar a rota é verificado se ela está armazenada no cache e retorna caso esteja. Caso contrario
ele espera a execução e caso seja executada com sucesso é armazenada no cache para futuras requisições. Além disso, 
recebe um parametro ttl que pode ser customizado para cada caso.

## Autenticação
Foi criado o arquivo `infra/middlewares/auth_middleware.go` que intercepta as rotas e verifica o token setado no 
header Authorization. Caso não esteja presente ou seja inválido um erro 401 é retornado. 

Além disso, uma rota /login foi criada e os e-mails validos nesse caso estão armazenados em uma variavel.

Dados para login:
{
    "email": "teste@teste.com",
    "password": "123456"
}

## Containerização
Foi utilizado o docker-compose para conteinerização. No caso do postgres e do redis foi utilizada uma imagem padrão
e para aplicação foi criada uma imagem por meio do Dockerfile (para faciltiar o desenvolvimento com live reload).