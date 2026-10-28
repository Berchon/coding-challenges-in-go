# Grading Students

**Link do Desafio:** [https://www.hackerrank.com/challenges/grading/problem?isFullScreen=true](https://www.hackerrank.com/challenges/grading/problem?isFullScreen=true)

---

DESCRIÇÃO DO PROBLEMA

A HackerLand University possui a seguinte política de notas:

- Todo estudante recebe uma nota no intervalo inclusivo de 0 a 100.
- Qualquer nota menor que 38 é considerada uma nota de reprovação.

Sam é professor na universidade e deseja arredondar as notas dos
estudantes de acordo com as seguintes regras:

- Se a diferença entre a nota e o próximo múltiplo de 5 for menor que 3,
  a nota deve ser arredondada para o próximo múltiplo de 5.
- Se a nota for menor que 38, nenhum arredondamento deve ser realizado,
  pois o resultado continuaria sendo uma nota de reprovação.

EXEMPLOS

- 84 deve ser arredondado para 85, pois 85 - 84 = 1, que é menor que 3.
- 57 não deve ser arredondado para 60, pois 60 - 57 = 3.
- 33 não deve ser arredondado, pois é menor que 38.

Dado o valor inicial da nota de cada estudante, implemente uma função
que realize o arredondamento de acordo com essas regras.

## FUNCTION DESCRIPTION

Complete a função gradingStudents com o seguinte parâmetro:

- grades: um array contendo as notas antes do arredondamento.

A função deve retornar:

- um array contendo as notas após o arredondamento.

## INPUT FORMAT

A primeira linha contém um único inteiro n, representando a quantidade
de estudantes.

Cada uma das n linhas seguintes contém um único inteiro grade,
representando a nota de um estudante.

## CONSTRAINTS

- 1 <= n <= 60
- 0 <= grade <= 100

## SAMPLE INPUT

4
73
67
38
33

## SAMPLE OUTPUT

75
67
40
33

EXPLICAÇÃO DO EXEMPLO

1. O estudante recebeu 73. O próximo múltiplo de 5 é 75.
   Como 75 - 73 = 2, a nota é arredondada para 75.

2. O estudante recebeu 67. O próximo múltiplo de 5 é 70.
   Como 70 - 67 = 3, a nota não é modificada.

3. O estudante recebeu 38. O próximo múltiplo de 5 é 40.
   Como 40 - 38 = 2, a nota é arredondada para 40.

4. O estudante recebeu 33. Como a nota é menor que 38,
   ela não é modificada.
