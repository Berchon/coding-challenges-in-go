# Drawing Book

**Link do Desafio:** [https://www.hackerrank.com/challenges/drawing-book/problem?isFullScreen=true](https://www.hackerrank.com/challenges/drawing-book/problem?isFullScreen=true)

---

Enunciado:

Uma professora pede que os alunos abram seus livros em uma
determinada página. O aluno pode começar a virar as páginas
a partir do início ou do final do livro.

Ao abrir o livro pela primeira vez, a página 1 fica sempre
no lado direito.

Ao virar uma página, o aluno visualiza as duas páginas
seguintes. Todas as páginas, exceto a última, são impressas
nos dois lados. Dependendo da quantidade total de páginas,
a última página pode ser impressa somente na frente.

Dado um livro com n páginas e uma página desejada p,
determine o número mínimo de viradas necessárias para
chegar à página p, começando pelo início ou pelo final
do livro.

Descrição da função:

Complete a função pageCount.

## Parâmetros

- int32 n: quantidade total de páginas do livro.
- int32 p: número da página desejada.

## Retorno

- int32: número mínimo de viradas necessárias.

## Formato de entrada

- A primeira linha contém um inteiro n.
- A segunda linha contém um inteiro p.

## Restrições

- 1 <= n <= 10^5
- 1 <= p <= n

## Exemplo 0

Entrada:
6
2

Saída:
1

## Explicação

Começando pelo início, é necessária uma virada.
Começando pelo final, são necessárias duas viradas.
O resultado mínimo é 1.

## Exemplo 1

Entrada:
5
4

Saída:
0

## Explicação

A página 4 pode ser alcançada diretamente a partir
do final do livro, sem nenhuma virada.
