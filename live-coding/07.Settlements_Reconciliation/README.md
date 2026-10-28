# Conciliacao de liquidacoes por escola

### Descricao

Implementar uma funcao que reconcilia invoices com settlements e gera um resumo de arrecadacao por escola. A solucao valida invoices e settlements, monta um mapa de invoices unicas, considera apenas settlements confirmados que pagam exatamente o valor da invoice correspondente, contabiliza cada invoice no maximo uma vez e ordena o relatorio por total arrecadado.

### Requisitos

- Receber uma lista de invoices e uma lista de settlements.
- Validar que cada invoice tenha `InvoiceID` preenchido.
- Validar que cada invoice tenha `SchoolID` maior que zero.
- Validar que cada invoice tenha `Amount` maior que zero.
- Validar que cada invoice tenha `DueDate` preenchida.
- Retornar erro quando houver `InvoiceID` repetido na lista de invoices.
- Validar que cada settlement tenha `InvoiceID` preenchido.
- Validar que cada settlement tenha `PaidAmount` maior que zero.
- Validar que cada settlement tenha `PaidDate` preenchida.
- Aceitar apenas os status de settlement `CONFIRMED`, `REJECTED` e `REFUNDED`.
- Retornar erro quando alguma invoice ou settlement possuir dados invalidos.
- Ignorar settlements cujo `InvoiceID` nao exista na lista de invoices.
- Ignorar settlements com status diferente de `CONFIRMED`.
- Ignorar settlements cujo `PaidAmount` seja diferente do `Amount` da invoice correspondente.
- Contabilizar uma invoice liquidada apenas uma vez, usando o primeiro settlement valido encontrado na ordem da lista.
- Agrupar o resultado por `SchoolID`.
- Para cada escola, somar o total arrecadado em `TotalCollected`.
- Para cada escola, contar a quantidade de invoices liquidadas em `SettledInvoicesCount`.
- Retornar uma lista vazia quando nao houver dados conciliaveis.
- Ordenar o relatorio por `TotalCollected` em ordem decrescente.
- Em caso de empate no total arrecadado, ordenar por `SchoolID` em ordem crescente.
