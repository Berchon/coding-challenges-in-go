# Compare the Triplets

**Link do Desafio:** [https://www.hackerrank.com/challenges/compare-the-triplets/problem?isFullScreen=true](https://www.hackerrank.com/challenges/compare-the-triplets/problem?isFullScreen=true)

---

## Descrição

Alice e Bob criaram, cada um, um problema para o HackerRank.
Um avaliador atribui notas aos dois problemas em três categorias:
clareza do problema, originalidade e dificuldade.

As notas de Alice são representadas pelo triplet:
a = (a[0], a[1], a[2])

As notas de Bob são representadas pelo triplet:
b = (b[0], b[1], b[2])

A pontuação de comparação deve ser calculada comparando as notas
correspondentes de cada categoria:

- Se a[i] > b[i], Alice recebe 1 ponto.
- Se a[i] < b[i], Bob recebe 1 ponto.
- Se a[i] == b[i], ninguém recebe ponto.

A função deve retornar um array com duas posições:
- a primeira posição contém a pontuação de Alice;
- a segunda posição contém a pontuação de Bob.

## Exemplo

a = [1, 2, 3]
b = [3, 2, 1]

Comparações:
- a[0] < b[0] → Bob recebe 1 ponto.
- a[1] == b[1] → ninguém recebe ponto.
- a[2] > b[2] → Alice recebe 1 ponto.

Resultado:
[1, 1]

## Formato de entrada

A primeira linha contém três inteiros separados por espaço,
representando a[0], a[1] e a[2].

A segunda linha contém três inteiros separados por espaço,
representando b[0], b[1] e b[2].

## Restrições

1 <= a[i] <= 100
1 <= b[i] <= 100

## Exemplo de entrada

5 6 7
3 6 10

## Exemplo de saída

1 1
