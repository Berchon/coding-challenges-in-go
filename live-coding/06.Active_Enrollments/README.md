# Consolidacao de matriculas ativas

### Descricao

Implementar uma funcao que consolida eventos de matricula e retorna os cursos ativos de cada aluno. A solucao valida os eventos, mantem apenas o evento mais recente para cada par `StudentID` e `CourseCode`, considera ativos somente os cursos cujo evento mais recente esteja com status `ENROLLED` e ordena o resultado.

### Requisitos

- Receber uma lista de eventos contendo `StudentID`, `CourseCode`, `Status` e `Timestamp`.
- Validar que o `StudentID` seja um inteiro positivo.
- Validar que o `CourseCode` nao esteja vazio.
- Validar que o `Timestamp` seja um inteiro positivo.
- Aceitar apenas os status `ENROLLED`, `DROPPED` e `SUSPENDED`.
- Retornar erro quando qualquer evento possuir dados invalidos.
- Consolidar eventos usando a combinacao de `StudentID` e `CourseCode`.
- Para cada aluno e curso, considerar apenas o evento com maior `Timestamp`.
- Incluir no relatorio apenas cursos cujo evento mais recente tenha status `ENROLLED`.
- Excluir cursos cujo evento mais recente tenha status `DROPPED` ou `SUSPENDED`.
- Retornar uma lista vazia quando nao houver eventos ou quando nenhum curso estiver ativo.
- Ordenar os cursos ativos de cada aluno em ordem alfabetica.
- Ordenar o relatorio final por `StudentID` em ordem crescente.
