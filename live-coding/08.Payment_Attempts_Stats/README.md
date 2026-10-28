# Estatísticas de Tentativas de Pagamento

### Descrição

Implementar uma função que recebe uma lista de tentativas de pagamento por escola e calcula as estatísticas de processamento. A solução valida cada tentativa, ignora status pendentes, contabiliza sucessos e falhas, calcula a taxa de falha de cada escola e retorna apenas as escolas qualificadas (pelo menos 3 transações processadas e taxa de falhas superior a 50%), ordenadas por maior taxa de falha.

### Requisitos

- Receber uma lista de `PaymentAttempt` contendo `AttemptID`, `SchoolID`, `Amount`, `Status` e `Timestamp`.
- Validar que `AttemptID` e `SchoolID` não estejam vazios.
- Validar que `Amount` seja maior que zero.
- Aceitar apenas status válidos (`SUCCESS`, `FAILED`, `PENDING`). Retornar erro caso contrário.
- Ignorar tentativas com status `PENDING` no cômputo das estatísticas.
- Para cada escola, totalizar tentativas processadas (`TotalProcessed = success + failed`) e falhas (`FailedAttempts`).
- Calcular a taxa de falha (`FailureRate = FailedAttempts / TotalProcessed`).
- Filtrar apenas escolas qualificadas: `TotalProcessed >= 3` e `FailureRate > 0.5`.
- Retornar uma lista de `SchoolStats` contendo `SchoolID`, `TotalProcessed`, `FailedAttempts` e `FailureRate`.
- Ordenar o resultado por `FailureRate` decrescente; em caso de empate, por `SchoolID` em ordem alfabética crescente.
- Retornar lista vazia caso não haja tentativas ou nenhuma escola se qualifique.
