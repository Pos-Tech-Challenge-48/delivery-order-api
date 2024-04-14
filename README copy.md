# delivery- api

## O quê é?

Aplicativo que gerencia atividades de um serviço de pedidos em um restaurante. Desde a base de clientes, catálogo de produtos, pedidos e fila de preparo.


## Como executar


### Variáveis de ambiente

É necessário definir o arquivo `.env` com as váriaveis no padrão da `.env.example`, escolhendo as credenciais da forma que desejar. Um ponto de atenção: se estiver rodando o serviço do banco fora do docker compose e sem uma network configurada, é necessário definir o host do banco como `localhost`.


###  Executando com o Docker localmente

A melhor maneira de levantar a aplicação localmente é utilizando containers. Ao iniciar o projeto, basta executar o comando:

```
docker compose up --build
```

Isso fará com que o Docker faça a build das imagens e rode os serviços. 

Utilizamos para isso uma imagem de desenvolvimento, com foco em utilização "local" para que o desenvolvedor consiga rodar o projeto utilizando a imagem do Go correta, de dentro de um container que aponta para o sistema local de arquivos atrás de volumes.

A dockerfile usada para desenvolvimento é `Dockerfile.dev`. Ela orquestra os containers da aplicação e do banco de dados local, bem como usar a ferramenta `air` , que é uma dependência de desenvolvimento. O `air` é um hot reloader para o GO.

### Parar a aplicação

Para desligar os containers rodando basta rodar.

```
docker compose down
```


### Limpeza do ambiente

Caso necessite "limpar" o ambiente e excluir os volumes, basta executar.
Isso fará com que sua imagem fique zerada e limpará todos os dados do banco.

```
docker compose down -v --remove-orphans
```


### Docker produção

(Opcional)

Para subir um container standalone (sem o banco de dados) basta usar os comando do Makefile.

Para buildar uma imagem para produção, execute.
```
make build-prod
```

Para rodar uma imagem para produção localmente, execute. 

Lembre-se de ter uma instância do banco rodando e usar as variáveis corretamente para se conectar. Não automatizamos essa parte.

```
run-prod
```

Esse caminho não é recomendado para rodar localmente, já que a sua configuração é mais trabalhosa.

### Migrations

A aplicação conta com versionamento de migrações feita automaticamente. Utilizando o `docker-compose up --build`, a aplicação subirá junto com o banco e se responsabilizará por criar as tabelas, triggers e seeds.

As seeds serão dados "default" de `status` e `category`, que no começo iremos automaticamente disponibilizar como "estado inicial".

Para isso, utilizamos o pacote [golang-migrate](https://github.com/golang-migrate/migrate). 

(Opcional)

Para rodar as migrações diretamento no banco de forma manual é necessário instalar o pacote. Caso ainda não tenha instalado, você pode seguir a documentação ou executar o comando abaixo. Lembre-se que é necessário ter o `go` instalado em sua máquina.

```
make local-install
```

Para executar as migrations no banco, bastar estar com a instância do banco local rodando, e executar o comando abaixo.

```
make setup-dev
```

### Tests 

Para rodar os testes, basta executar:

```
go test -v ./... --coverprofile=c.out && go tool cover -html=c.out
```

### Update Swagger

Instale essas ferramentas: https://github.com/swaggo/gin-swagger

```
swag init -g ./cmd/api/main.go -o ./cmd/api/docs --parseDependency
```


### Kubernetes

A documentação do kubernetes pode ser acessada [aqui](https://github.com/Pos-Tech-Challenge-48/delivery-api/tree/main/kubernetes).