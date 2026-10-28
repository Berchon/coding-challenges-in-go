# Between Two Sets

**Link do Desafio:** [https://www.hackerrank.com/challenges/between-two-sets/problem?isFullScreen=true](https://www.hackerrank.com/challenges/between-two-sets/problem?isFullScreen=true)

---

## Descrição

São fornecidos dois arrays de inteiros, a e b.

Precisamos encontrar quantos inteiros existem entre os dois arrays, seguindo
duas condições:

1. Todos os elementos do array a devem ser fatores (divisores) do número
   considerado.

2. O número considerado deve ser fator (divisor) de todos os elementos
   do array b.

Um número só deve ser contado quando satisfizer simultaneamente as duas
condições.

## Exemplo

a = [2, 4]
b = [16, 32, 96]

Os números que satisfazem as duas condições são:

4, 8 e 16

Portanto, o resultado é:

3

## Explicação

Para 4:
- 2 divide 4
- 4 divide 4
- 4 divide 16, 32 e 96

Para 8:
- 2 divide 8
- 4 divide 8
- 8 divide 16, 32 e 96

Para 16:
- 2 divide 16
- 4 divide 16
- 16 divide 16, 32 e 96

Logo, existem 3 números entre os dois conjuntos.

Entrada:

A primeira linha contém a quantidade de elementos dos arrays a e b.
A segunda linha contém os elementos do array a.
A terceira linha contém os elementos do array b.

Saída:

Retornar a quantidade de inteiros que satisfazem as duas condições.

## Exemplo de entrada

2 3
2 4
16 32 96

## Exemplo de saída

3
