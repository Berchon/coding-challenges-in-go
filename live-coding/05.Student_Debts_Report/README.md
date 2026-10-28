# Relatorio de dividas por aluno

### Descricao

Implementar uma funcao que gera um relatorio de dividas pendentes por aluno a partir de uma lista de transacoes. A solucao valida o status das transacoes, calcula o saldo por aluno considerando transacoes pendentes e pagas, ignora transacoes canceladas, filtra apenas saldos positivos e ordena o relatorio por maior divida.

### Requisitos

- Receber uma lista de transacoes contendo `StudentID`, `Amount` e `Status`.
- Aceitar apenas os status `PENDING`, `PAID` e `CANCELED`.
- Retornar erro quando uma transacao possuir status invalido.
- Considerar transacoes com status `PENDING` no calculo do saldo do aluno.
- Considerar transacoes com status `PAID` no calculo do saldo do aluno, permitindo valores negativos para reduzir a divida.
- Ignorar transacoes com status `CANCELED`.
- Agrupar os valores por `StudentID` para calcular o saldo final de cada aluno.
- Incluir no relatorio apenas alunos com saldo positivo.
- Excluir alunos com saldo zero ou negativo.
- Retornar uma lista vazia quando nao houver transacoes ou quando nenhum aluno tiver saldo positivo.
- Ordenar o relatorio por valor de divida em ordem decrescente.
- Em caso de empate no valor da divida, ordenar pelo menor `StudentID`.
