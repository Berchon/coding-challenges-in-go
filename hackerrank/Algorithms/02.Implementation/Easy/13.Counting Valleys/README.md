# Counting Valleys

**Link do Desafio:** [https://www.hackerrank.com/challenges/counting-valleys/problem?isFullScreen=true](https://www.hackerrank.com/challenges/counting-valleys/problem?isFullScreen=true)

---

Descrição do problema:

Gary é um caminhante que registra cuidadosamente suas caminhadas.
Durante sua última caminhada, ele deu exatamente `steps` passos.
Para cada passo, registrou se foi uma subida (U) ou uma descida (D).

A caminhada sempre começa e termina no nível do mar. Cada passo
para cima ou para baixo representa uma mudança de uma unidade
na altitude.

Definições:

- Montanha: sequência de passos consecutivos acima do nível do mar,
  começando com uma subida a partir do nível do mar e terminando
  com uma descida de volta ao nível do mar.

- Vale: sequência de passos consecutivos abaixo do nível do mar,
  começando com uma descida a partir do nível do mar e terminando
  com uma subida de volta ao nível do mar.

Dada a sequência de subidas e descidas, determine quantos vales
foram percorridos durante a caminhada.

Descrição da função:

Complete a função countingValleys.

## Parâmetros

- steps (int32): quantidade de passos da caminhada.
- path (string): sequência de caracteres U e D que descreve
  os passos da caminhada.

## Retorno

- int32: quantidade de vales percorridos.

## Formato de entrada

- Primeira linha: inteiro steps, representando a quantidade de passos.
- Segunda linha: string path, contendo os passos da caminhada.

## Restrições

- 2 <= steps <= 10^6
- path[i] pertence ao conjunto {U, D}

## Exemplo de entrada

8
UDDDUDUU

## Exemplo de saída

1

## Explicação

A caminhada começa no nível do mar, desce abaixo dele, retorna ao
nível do mar e depois passa por uma montanha. O percurso contém
exatamente um vale.
