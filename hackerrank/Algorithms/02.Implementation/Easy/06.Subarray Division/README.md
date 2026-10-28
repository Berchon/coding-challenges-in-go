# Subarray Division (The Birthday Bar)

## Descrição do problema

Duas crianças, Lily e Ron, querem compartilhar uma barra de chocolate. Cada quadrado da barra contém um número inteiro.

Lily decide compartilhar um segmento contíguo da barra que satisfaça as seguintes condições:

* O comprimento do segmento deve ser igual ao mês de nascimento de Ron.
* A soma dos números presentes nos quadrados do segmento deve ser igual ao dia de nascimento de Ron.

Determine de quantas maneiras diferentes Lily pode dividir a barra de chocolate para atender a esses critérios.

## Descrição da função

Implemente a função `birthday` no editor.

**Parâmetros:**

* `s[n]`: números inteiros presentes nos quadrados da barra de chocolate.
* `d`: dia de nascimento de Ron.
* `m`: mês de nascimento de Ron.

## Retorno

* Um inteiro que representa o número de maneiras válidas de dividir a barra.

## Formato da entrada

1. A primeira linha contém um inteiro `n`, que representa a quantidade de quadrados da barra de chocolate.
2. A segunda linha contém `n` números inteiros separados por espaços, representando os valores dos quadrados.
3. A terceira linha contém dois números inteiros separados por espaço: `d`, o dia de nascimento de Ron, e `m`, o mês de nascimento de Ron.

## Exemplos

### Exemplo 0

**Entrada**

```text
5
1 2 1 3 2
3 2
```

**Saída**

```text
2
```

**Explicação**

Lily precisa encontrar segmentos de comprimento 2 cuja soma seja 3.

Existem dois segmentos que atendem aos critérios: `[1, 2]` e `[2, 1]`.

### Exemplo 1

**Entrada**

```text
6
1 1 1 1 1 1
3 2
```

**Saída**

```text
0
```

**Explicação**

Lily precisa selecionar quadrados consecutivos cujos valores somem 3, considerando segmentos de comprimento 2.

Não existe nenhum segmento que satisfaça essas condições. Portanto, o resultado é 0.

### Exemplo 2

**Entrada**

```text
1
4
4 1
```

**Saída**

```text
1
```

**Explicação**

Lily precisa selecionar um segmento de comprimento 1 cuja soma seja 4.

Como a barra contém apenas um quadrado, com valor 4, existe exatamente uma maneira válida de fazer a divisão.

## Link do desafio

https://www.hackerrank.com/challenges/the-birthday-bar/problem?isFullScreen=true
