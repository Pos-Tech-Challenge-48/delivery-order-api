# Kubernetes

Para trabalhar com o kubernetes, uma ferramenta indispensável é o `kubectl`, que é usando para se comunicar com api e realizar comandos no cluster. Para maiores informações sobre como instalá-lo siga o passo a passo do [link](https://kubernetes.io/docs/tasks/tools/).

Para executar o projeto utilizando-se de um cluster kubernetes escolhemos como ferramenta o `minikube` que permite executar as imagens do projeto localmente, sem necessidade de uma infra na nuvem.

Para fazer o download e instalação da ferramenta, siga o passo a passo do tutorial [no link](https://minikube.sigs.k8s.io/docs/start/).


- Para inicializar o cluster local:

```
minikube start
```

- Para verificar se o cluster foi inicializado corretamente:

```
kubectl get nodes -o wide
```

Nesse passo você já deve conseguir ver um node com o NAME "minikube" rodando com o STATUS Ready.


## Nosso projeto

Todas as configurações a respeito do kubernetes estão localizadas na pasta `/kubernetes` desse projeto. A divisão dos serviços foi feita em duas pastas principais diferentes:


### /db - Banco de Dados

Contém as configurações necessárias para subir o serviço `delivery-api`, que representa um banco de dados PostgreSQL na versão 16.1.

O racional para essa configuração foi baseada em estudos sobre qual o melhor kind para essa necessidade e encontramos a possibilidade de utilizar o `StatefulSet` como opção.

Principais vantagens do StatefulSet:
* gerencia os Pods, limites e a escalabilidade de acordo com suas configurações
* mantém uma identidade para cada pod, criando-o sempre com o mesmo nome
* consegue garantir a persistência dos dados valendo-se de PersistentVolume, PersistentVolumeClaim e a StorageClass. Nesse sentido, isso se faz muito importante pensando em um banco de dado, pois mesmo que os pods sejam "deletados" e "criados", os dados do banco se mantenham íntegros
* a StorageClass usada foi a `standard` do cluster

Temos os seguintes arquivos:
* `postgresql-db-configmap.yaml` - Configmap com os valores necessários para configurar um user no banco. Como é um ambiente de testes local e para aprendizado, não me preocupei com "secrets".

* `postgresql-db-service.yaml` - Service tipo LoadBalancer para redirecionar a carga externa para o pod do banco.

* `postgresql-db-statefulset.yaml` - Configuração do StatefulSet do banco, que irá subir apenas um Pod.

Para iniciar os serviços necessários do banco:
```
kubectl apply -f ./kubernetes/db/
```


### /api - Backend REST

Contém as configurações necessárias para subir o serviço da `delivery-api`, que é um Rest API criada em Golang, vinda da fase anterior.

A imagem do container contendo a `build` atualizada da app está no Docker HUB, nesse repositório `vitorsmap/delivery-api:v3`. Mais informações estão nesse [link](https://hub.docker.com/repository/docker/vitorsmap/delivery-api/general).

Temos os seguintes arquivos:

* `api-configmap.yaml` - configuração das variáveis de ambientes da api
* `api-deployment.yaml` - deployment que roda a imagem da api. 
    * Foram configuradas as rotas de livenesse e readiness probe  `/healthcheck`.
    * Limites dos recursos
    * ConfigMap
    * Replicas
* `api-service.yaml`
* `api-hpa.yaml` - Configuração de Horizontal Pod Autoscaler, podendo chegar à 5 pods. Baseado em utilização da CPU.
* `metrics.yaml` - Configuração de métricas para o Cluster


Para iniciar os serviços necessários da api:
```
kubectl apply -f ./kubernetes/api/
```


### /stressTest - Teste de stress 

Utilizamos o k6 como ferramenta para realizar os testes de stress.

Para rodar o stress test e validar o HPA:

```
k6 run ./kubernetes/stressTest/index.js 
```


## Desenho da Arquitetura

![Kubernetes Minikube](https://raw.githubusercontent.com/Pos-Tech-Challenge-48/delivery-api/main/kubernetes/images/fiapTech-kubernetes.drawio.png).
