# Sales by Match (Sock Merchant)

**Link do Desafio:** [https://www.hackerrank.com/challenges/sock-merchant/problem?isFullScreen=true](https://www.hackerrank.com/challenges/sock-merchant/problem?isFullScreen=true)

---

## Descrição

John trabalha em uma loja de roupas e possui uma grande
quantidade de meias que precisam ser separadas em pares
para venda.

Cada meia possui uma cor representada por um número inteiro.
Dado um array de inteiros representando as cores das meias,
determine quantos pares de meias da mesma cor podem ser
formados.

## Exemplo

n = 7
ar = [1, 2, 1, 2, 1, 3, 2]

É possível formar um par de meias da cor 1 e um par da
cor 2. Sobram três meias sem par, uma de cada cor.
Portanto, o resultado é 2.

Descrição da função:
Implemente a função sockMerchant para retornar a quantidade
de pares de meias que possuem cores correspondentes.

## Parâmetros

- n: quantidade de meias.
- ar: array contendo a cor de cada meia.

## Retorno

- int32: quantidade total de pares de meias.

## Formato de entrada

A primeira linha contém um inteiro n, representando a
quantidade de meias.
A segunda linha contém n inteiros separados por espaços,
representando as cores das meias.

## Restrições

- 1 <= n <= 100
- 1 <= ar[i] <= 100

## Exemplo de entrada

9
10 20 20 10 10 30 50 10 20

## Exemplo de saída

3

## Explicação

É possível formar três pares de meias da mesma cor.
