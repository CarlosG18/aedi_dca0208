# Estrutura de dados abstrata: Fila de Prioridade

A fila de prioridades é um tipo abstrato de dados em que os elementos da fila possuem um grau de prioridade maior ou menor em relação aos outros elementos da fila. Um exemplo disso é dado no gerenciamento de rotinas e execuções que o processador precisa realizar para manter o sistema otimizado, pois existem aplicações que devem ser priorizadas em relação a outras.

# HEAP

heap é uma forma de implementar a estrutura de dados da fila de prioridades. com a heap binária, que se comporta como uma árvore binária, podemos implementar este tipo de dados. existem dois tipos de heap:

1. **MaxHeap**: em uma maxHeap os elementos "pais" são sempre maiores que seus elementos filhos de uma forma recursiva. vamos para um exemplo para entendermos melhor. analisamos a arvore binária abaixo:

<p align="center">
  <img src="https://github.com/CarlosG18/edi_dca0208/blob/main/conteudos/heap/maxheap.png" alt="MaxHeap">
</p>

nela podemos observar que o nó raiz `56` é maior que seus filhos `45` e `34`. se isso for válido para todos os filhos de `56` de modo recursivo, então estamos diante de uma **MaxHeap**.

2. **MinHeap**: de forma semelhante a MaxHeap, a MinHeap possui um valor do nó pai menor que os valores do seu nós filhos. vamos observar a MinHeap abaixo:

<p align="center">
  <img src="https://github.com/CarlosG18/edi_dca0208/blob/main/conteudos/heap/minheap.png" alt="MinHeap">
</p>

nela podemos ver que qualquer nó pai possui filhos com valores maiores. por exemplo, o nó `8` possui dois filhos `21` e `15` que são maiores que o nó pai `8`. diante disso temos uma MinHeap.

## Heap valida?

como sabemos se uma heap é válida? bom como já vimos anteriormente, temos dois tipos de heap, `MaxHeap` e `MinHeap`, caso algum nó filho seja contrário as definições de cada tipo de heap, esta heap não é considerada valida. já que uma heap é construida como uma árvore binária, caso um nó filho possua dois nós pais, que quebra a definição de árvore binária, também é considerado como uma heap não valida. outro fator é determinante para verificar se uma heap é valida, é a ordem em que os elementos são inseridos, ou seja, para que uma heap seja valida a inserção na arvore binaria deve acontecer da esquerda para a direita.

## forma de implementação

a Heap é implementada usando um `array` e uma variável para indicar quantos elementos foram inseridos neste array. a estrutura ficaria da seguinte maneira:

```go
type Heap struct{
  ele []int
  qtd_ele_inseridos int
}
```

***observação***: irei usar um vetor de inteiros, porém podemos usar qualquer outro tipo de dado.

## interface de uma Heap

A heap contem algumas funcionalidades básicas:
1. **Add**: toda heap possui o método de adicionar elementos. essa adição ocorre no final do vetor. após a adição é realizada uma flutuação do elemento inserido caso a posição dele não for a correta.
2. **Poll**: o método Poll é o equivalente ao desenfileirar, cujo será retornado o elemento que possui uma maior prioridade (MaxHeap) ou uma menor prioridade (MinHeap).
3. **Remove**: este método remove um elemento, independente de sua posição na heap. com isso, dependendo do local da remoção, devemos reorganizar a árvore para que ela continue respeitando as definições de `MaxHeap` ou `MinHeap`.

sua interface ficaria da seguinte maneira:

```go
type Heap interface{
  Add(ele int)
  Poll() int
  Remove(ele int)
}
```

## funcionamento da Heap

vamos agora expor como seria o funcionamento passo a passo de uma heap. vamos supor que já temos uma MaxHeap e que foram adicionados elementos aleatorios (respeitando sua regra) dada pela imagem abaixo:

<p align="center">
  <img src="https://github.com/CarlosG18/edi_dca0208/blob/main/conteudos/heap/exemplo.png" alt="exemplo_heap">
</p>

como podemos ver na imagem, cada elemento possui seu índice que conresponde ao vetor definido na estrutura da Heap. o vetor `ele` da nossa MaxHeap ficará assim:

```go 
ele = [73, 55, 64, 21, 44, 19]
```
### função Add()
agora vamos realizar a operação de `Add(72)` em nossa MaxHeap.

<p align="center">
  <img src="https://github.com/CarlosG18/edi_dca0208/blob/main/conteudos/heap/add(72).gif" alt="gif_add(72)" width="700">
</p>

na função `Add()` o elemento a ser inserido será colocado no final da arvore binaria. apos isso ele flutuará para sua posição correta, fazendo a comparação com o seu nó pai e se ele for maior a troca será realizada, caso contrário, a posição já está correta.

***dicas para a implementação da função Add()***:
- para obter o indicie do pai de um determinado nó, basta usar a seguinte formula:
```text
  index_nó_pai = (index - 1)/2
```

### função Poll()
agora vamos ver como a nossa MaxHeap se comporta ao executarmos a função `poll()`:

<p align="center">
  <img src="https://github.com/CarlosG18/edi_dca0208/blob/main/conteudos/heap/poll.gif" alt="gif_poll" width="700">
</p>

como observado no gif, reservamos o valor da raiz da nossa MaxHeap colocando em uma variavel auxiliar chamada `Topo`. apos isso, devemos colocar o ultimo elemento para o lugar da raiz. com isso realizado, precisamos verificar recursivamente qual o maior nó filho e realizar a troca. a logica inversa ocorre em uma MinHeap, ou seja, precisamos verificar o menor valor entre os nós filhos e realizar a troca. não havendo mais a possibilidade de troca (o elemento esta no seu local correto) retornamos a variavel `Topo`. 

***dicas para a implementação da função poll()***:
- para obter o indicie dos filhos de um determinado nó, basta obter o index do nó pai e os index dos nos filhos se darão pela aplicação da seguinte formula:
```text
  index_nó_esquerdo = 2 * index_pai + 1
  index_nó_direito = 2 * index_pai + 2
```