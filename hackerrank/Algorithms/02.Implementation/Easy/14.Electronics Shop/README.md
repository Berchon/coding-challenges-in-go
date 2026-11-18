# Electronics Shop

**Link do Desafio:** [https://www.hackerrank.com/challenges/electronics-shop/problem?isFullScreen=true](https://www.hackerrank.com/challenges/electronics-shop/problem?isFullScreen=true)

---

# Descrição do problema:

Monica deseja comprar um teclado e um dispositivo USB (USB drive) em sua loja de eletrônicos favorita. A loja possui diversos modelos de cada produto. Monica deseja gastar o máximo possível dentro do seu orçamento.

Dadas as listas de preços dos teclados e dos dispositivos USB, além do orçamento disponível, determine e retorne o valor máximo que Monica poderá gastar na compra dos dois itens.

Caso não tenha dinheiro suficiente para comprar um teclado e um dispositivo USB, retorne `-1`.

**Exemplo:**

Considere os seguintes valores:

- `b = 60` (orçamento disponível).
- `keyboards = [40, 50, 60]` (preços dos teclados).
- `drives = [5, 8, 12]` (preços dos dispositivos USB).

Monica possui as seguintes combinações possíveis:

- Teclado de 40 + dispositivo USB de 12 = 52.
- Teclado de 50 + dispositivo USB de 8 = 58.

A segunda combinação é a mais cara dentro do orçamento disponível. Portanto, o resultado deve ser `58`.

Monica não pode comprar mais de dois itens, pois precisa adquirir exatamente um teclado e um dispositivo USB.

## Descrição da função

Complete a função `getMoneySpent` no editor disponibilizado pelo HackerRank.

**Parâmetros:**

- `int keyboards[n]`: array contendo os preços dos modelos de teclado disponíveis.
- `int drives[m]`: array contendo os preços dos dispositivos USB disponíveis.
- `int b`: orçamento disponível para a compra.

**Retorno:**

- `int`: o valor máximo que pode ser gasto na compra dos dois itens. Retorne `-1` caso não seja possível comprar ambos dentro do orçamento.

## Formato de entrada

- A primeira linha contém três números inteiros separados por espaço: `b`, `n` e `m`, representando, respectivamente, o orçamento disponível, a quantidade de modelos de teclado e a quantidade de modelos de dispositivos USB.
- A segunda linha contém `n` números inteiros separados por espaço, representando os preços de cada modelo de teclado.
- A terceira linha contém `m` números inteiros separados por espaço, representando os preços dos dispositivos USB.

## Restrições

- `1 ≤ n, m ≤ 1000`
- `1 ≤ b ≤ 10⁶`
- O preço de cada item está no intervalo inclusivo `[1, 10⁶]`.

## Exemplo de entrada 0

```text
10 2 3
3 1
5 2 8
```

## Exemplo de saída 0

```text
9
```

**Explicação 0:**

Compre o segundo modelo de teclado, que custa `1`, e o terceiro dispositivo USB, que custa `8`.

O custo total será `1 + 8 = 9`.

## Exemplo de entrada 1

```text
5 1 1
4
5
```

## Exemplo de saída 1

```text
-1
```

**Explicação 1:**

Não existe uma combinação que permita comprar um teclado e um dispositivo USB dentro do orçamento disponível, pois `4 + 5 = 9`, que ultrapassa o orçamento de `5`.

Portanto, retorne `-1`.