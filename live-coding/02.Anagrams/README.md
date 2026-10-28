# Verificador de anagramas

### Descricao

Implementar uma funcao que verifica se duas strings sao anagramas. A solucao converte as strings para slices de runes, valida se possuem o mesmo tamanho, ordena os caracteres de cada uma e compara os resultados ordenados.

### Requisitos

- Receber duas strings como entrada.
- Retornar `true` quando as duas strings tiverem exatamente os mesmos caracteres, apenas em ordem diferente ou igual.
- Retornar `false` quando as strings tiverem quantidades diferentes de caracteres.
- Considerar espacos como caracteres validos na comparacao.
- Considerar caracteres especiais como caracteres validos na comparacao.
- Suportar caracteres Unicode por meio de `rune`.
- Tratar a comparacao como case-sensitive, diferenciando maiusculas de minusculas.
- Considerar duas strings vazias como anagramas.
- Ordenar slices de runes em ordem crescente para apoiar a comparacao.
