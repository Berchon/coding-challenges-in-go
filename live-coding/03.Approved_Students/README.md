# Alunos aprovados por media

### Descricao

Implementar uma funcao que recebe registros de notas por aluno e retorna os alunos aprovados. A solucao valida as notas, agrupa os registros por `StudentID`, soma as pontuacoes, calcula a media de cada aluno e inclui no resultado apenas quem atingiu media maior ou igual a 70.

### Requisitos

- Receber uma lista de registros contendo `StudentID`, `Subject` e `Score`.
- Validar que cada nota esteja entre 0 e 100.
- Retornar erro quando alguma nota for menor que 0 ou maior que 100.
- Agrupar as notas por aluno usando o `StudentID`.
- Calcular a media de cada aluno com base em todas as notas registradas para ele.
- Retornar apenas alunos com media maior ou igual a 70.
- Retornar uma lista vazia quando nao houver registros ou quando nenhum aluno for aprovado.
- Retornar, para cada aluno aprovado, o `StudentID` e a media calculada.
- Nao exigir ordenacao especifica no resultado, ja que a iteracao do mapa nao garante ordem.
