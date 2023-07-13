# Árvores AVL

uma árvore AVL é uma árvore de busca binária que possui todos os seus nós balanceados. mas o que seria esse "balanceamento"? para garantirmos que as operações que envolvam a altura da árvore se mantenham sempre na complexidade de O(log(n)) precisamos garantir que todos os nós estejam balanceados e o que define se um nó esta balanceado é o `fator de balanço`.

## Fator de Balanço 

o fator de Balanço é definido como a diferença entre a altura da sub-árvore direita menos a altura da sub-árvore esquerda. com isso, temos as seguintes interpretações para o valor do fator de Balanço:

### Árvore AVL balanceada
<p align="center">
  <img src="https://github.com/CarlosG18/edi_dca0208/blob/main/conteudos/arvores/exemplo_avl.png" alt="exemple_avl" width="500">
</p>

nesta arvore árvore podemos observar valores para o fator de balanço que estão na faixa de `-1 <= FB <= 1`. uma arvore que possui todos os nós com fator de balanço nessa faixa é considerada uma árvore AVL balanceada. outro fato importante a levantar é que todos os nós folhas possuem o `fator de balanço` e a `altura` iguais a zero.

### Árvore AVL desbalanceada para a direita

<p align="center">
  <img src="https://github.com/CarlosG18/edi_dca0208/blob/main/conteudos/arvores/avl_direita.png" alt="avl_direita" width="500">
</p>

já nesse exemplo temos o nó `56` com o fator de balanço maior que o intervalo definido anteriormente como válido. como o valor do fator de Balanço deu positivo, temos o caso de desbalanceamento a direita.

### Árvore AVL desbalanceada para a esquerda

<p align="center">
  <img src="https://github.com/CarlosG18/edi_dca0208/blob/main/conteudos/arvores/avl_esquerda.png" alt="avl_esquerda" width="500">
</p>

nessa outra árvore avl, temos dois nós com o fator de balanço fora do intervalo correto, o nó `45` e o nó `56` possuem um desbalanceamento a esquerda devido o fator de balanço ser negativo.

bom já vimos que para uma árvore se encaixe em uma AVL ela deve está totalmente balanceada para garantir sua eficiência, mas como fazemos para corrigir uma arvore/sub-arvore que não está balanceada? a resposta está nas `rotações`.

## operação de rotação

- **rotação a esquerda**:

- **rotação a direita**:

## tipos de desbalanceamento e como corrigi-lás