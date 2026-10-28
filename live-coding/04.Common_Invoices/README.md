# Intersecao de notas fiscais

### Descricao

Implementar uma funcao que encontra as notas fiscais presentes tanto na lista do gateway quanto na lista interna. A solucao transforma a lista do gateway em um conjunto, percorre a lista interna e retorna cada nota comum uma unica vez.

### Requisitos

- Receber duas listas de inteiros: notas fiscais do gateway e notas fiscais internas.
- Retornar uma lista vazia quando qualquer uma das listas estiver vazia.
- Criar um conjunto a partir das notas fiscais do gateway.
- Percorrer as notas fiscais internas para identificar quais tambem existem no conjunto do gateway.
- Retornar apenas notas fiscais presentes nas duas listas.
- Evitar duplicatas no resultado, mesmo quando uma nota aparecer repetida em uma ou nas duas listas.
- Preservar a ordem em que as notas comuns sao encontradas na lista interna.
- Suportar valores positivos, negativos e zero como identificadores de nota fiscal.
