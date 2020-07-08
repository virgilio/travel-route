### Como executar a aplicação:

```shell
$ go get github.com/virgilio/travel-route
$ go install githube.com/virgilio/travel-route/cmd/travel-route
$ $GOPATH/bin/travel-route <path-to-input-file>

$ # using CURL to test the API:
$ curl -H "Content-Type: application/json"  "localhost:8080/bestRoute" --data '{"from":"GRU","to":"CDG"}'
$ curl -X POST -H "Content-Type: application/json"  "localhost:8080/addRoute" --data '{"From":"GRU","To":"SDU","Cost":10}'
$ curl -X POST -H "Content-Type: application/json"  "localhost:8080/addRoute" --data '{"From":"SDU","To":"CDG","Cost":20}'
$ curl -X GET -H "Content-Type: application/json"  "localhost:8080/addRoute" --data '{"From":"SDU","To":"ORL","Cost":20}' # Method not allowed
$ curl -H "Content-Type: application/json"  "localhost:8080/bestRoute" --data '{"from":"GRU","to":"CDG"}'

```
### Estrutura dos arquivos/pacotes:

  * `server` que contem a implementação da REST API, no caso, a applicação e um middleware que insere as informações de storage no requests
  * `shortestpath` que contem as estruturas de dados para calculo de caminhos minimos
  * `storage` que trata exclusivamente dos dados (há um sample input file nesse pacote)
  * `test` contem os testes de unidade

### Explique as decisões de design adotadas para a solução:

A aplicação possui um comando apenas que inicializa o serviço http e a linha de comando

Uma vez inicializado, ambas interfaces aguardam por requisições. Uma vez feita uma requisição, é lido o arquivo de voos disponiveis e montado o grafo

O vértice do grafo possui o nome da cidade e uma lista de partidas possíveis daquele aeroporto. Além disso, essa estrutura carrega seu custo mínimo ao ponto de partida e o caminho até lá (inicializados com -1 e vazio respectivamente). O Grafo é representado por um mapa de vértices

Com o ponto de partida, calculamos o caminho mais curto para todos os vértices, incluindo o destino usando algorítmo de Dijkstra de caminhos mínimos e retornamos, no final, o custo mínimo e o caminho da cidade de destino.

A construção do grafo e a função de menor caminho tem como receivers a lista de voos (carregada via modulo storage) e um vertive qualquer de partida respectivamente

### Descreva sua APÌ Rest de forma simplificada:

* POST /bestRoute `{"from": <str>, "to": <str>}`:
  * success: `{"Cost":<int>,"Cities": <[]str>}`
  * error: `{"error":<str>,"cause":<str>,"time":<str>}`
* POST /addRoute `{"from": <str>, "to": <str>, "cost": <int>}`
  * success: `{"From":<str>,"To":<str>,"Cost":<int>}`
  * error: `{"error":<str>,"cause":<str>,"time":<str>}`
