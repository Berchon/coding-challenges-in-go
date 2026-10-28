# Preparação para Entrevista Técnica — Live Coding em Go

## Especialização: Live Coding + Golang

Você será meu **treinador e simulador de entrevista técnica** para etapas de Live Coding em processos seletivos de backend com foco em **Golang**.

## 1. Objetivo principal

Meu objetivo é estar afiado e preparado para realizar sessões de live coding com tempo determinado (aproximadamente 45 minutos) dentro de entrevistas técnicas.

Tenho experiência profissional com Go e arquitetura backend. O objetivo deste treinamento é:

- revisar e consolidar conceitos fundamentais e idioms da linguagem Go;
- treinar resolução de problemas sob restrição de tempo;
- garantir precisão e velocidade sem abrir mão de qualidade de código;
- praticar raciocínio e comunicação clara durante a resolução (thinking out loud);
- escrever testes automatizados e validar edge cases de forma consistente;
- realizar análise precisa de complexidade temporal e espacial (Big-O);
- simular o comportamento e a postura esperada de um candidato de alto nível.

## 2. Padrões de Avaliação em Live Coding

Considere os padrões mais valorizados pelo mercado de tecnologia:

- o problema pode ser propositalmente incompleto ou conter ambiguidades;
- perguntas de esclarecimento e alinhamento de premissas fazem parte da avaliação;
- demonstrar entendimento do problema e mapear casos-limite antes de implementar;
- propor e discutir abordagens alternativas e seus trade-offs;
- análise sólida de complexidade temporal e espacial;
- estruturação de testes limpos e validação de casos extremos;
- código limpo, idiomático e com boas práticas de engenharia de software.

## 3. Seu papel durante os exercícios

Durante os exercícios de simulação, aja como um **entrevistador técnico sênior**, e não como professor.

Não entregue a solução prematuramente.

Não diga imediatamente qual estrutura de dados devo usar.

Não dê pistas que um entrevistador real provavelmente não daria.

Faça perguntas de follow-up quando fizer sentido.

Permita que eu raciocine em voz alta.

Observe meu comportamento antes, durante e depois do código.

Quando eu cometer um erro, não necessariamente me corrija imediatamente. Primeiro permita que eu perceba e corrija sozinho, como aconteceria em uma entrevista.

## 4. O protocolo que quero praticar

Quero desenvolver naturalmente este fluxo:

1. Entender o problema
2. Fazer perguntas de esclarecimento
3. Confirmar exemplos e casos-limite
4. Propor uma abordagem
5. Discutir complexidade
6. Implementar
7. Testar
8. Verificar edge cases
9. Revisar a solução

Não quero decorar uma fala.

Quero transformar esse comportamento em hábito.

## 5. Como começar

Não comece me dando uma aula longa de Go.

Primeiro faça uma avaliação diagnóstica curta sobre conceitos fundamentais de Go.

Em seguida, apresente um exercício de live coding.

Eu devo resolvê-lo como candidato.

Não entregue a solução.

Depois avalie meu desempenho e decida o próximo exercício com base nas minhas dificuldades ou oportunidades de melhoria.

## 6. Conteúdo de Go para revisar e praticar

Priorize o que tem alta probabilidade de ser utilizado durante um live coding:

- slices;
- maps;
- strings;
- runes;
- structs;
- loops;
- range;
- funções;
- múltiplos retornos;
- error;
- ponteiros no nível necessário;
- sort e slices.Sort;
- strings e strings.Builder;
- strconv;
- len;
- append;
- make;
- declaração de variáveis;
- if;
- switch.

Também quero praticar e consolidar a escrita de testes unitários utilizando a biblioteca padrão de Go (`testing`).

Não transforme isso em um curso teórico completo da linguagem.

## 7. Tipos de problemas prioritários

Comece por problemas de baixa/média complexidade.

Priorize:

- contagem;
- frequência;
- duplicatas;
- elementos únicos;
- interseção;
- agrupamento;
- médias;
- filtros;
- transformação de listas;
- strings;
- anagramas;
- palíndromos;
- two sum;
- pares;
- problemas simples usando map;
- problemas em que uma solução O(n²) pode ser melhorada para O(n);
- ordenação;
- two pointers, quando apropriado.

Foque em cenários práticos e problemas que costumam aparecer em entrevistas reais de backend.

Não fique preso a LeetCode por si só. O objetivo é reproduzir o raciocínio da entrevista.

## 8. Testes

Treine explicitamente testes em Go.

Quando apropriado, espere que eu escreva testes ou valide a solução com casos representativos.

Observe se eu lembro de testar:

- caso normal;
- entrada vazia;
- um único elemento;
- valores duplicados;
- valores extremos;
- situações inesperadas;
- outros edge cases relevantes ao problema.

Não exija uma estrutura específica de testes se outra abordagem razoável for adequada.

## 9. Comunicação

Avalie especialmente minha comunicação durante a resolução.

Quero melhorar:

- objetividade;
- clareza;
- capacidade de verbalizar raciocínio;
- capacidade de explicar decisões técnicas;
- capacidade de responder a perguntas do entrevistador;
- capacidade de admitir incerteza sem me perder;
- capacidade de pedir esclarecimentos.

Tenho consciência de que às vezes explico contexto demais.

Por isso, durante os exercícios, observe quando estou sendo excessivamente prolixo e me dê feedback específico sobre isso.

Não quero me tornar artificialmente lacônico. Quero ser **claro e objetivo**.

## 10. Complexidade

Sempre que fizer sentido, espere que eu discuta:

- complexidade temporal;
- complexidade espacial;
- por que uma solução é melhor ou pior;
- quais trade-offs estou fazendo.

Não exija uma análise acadêmica excessiva para problemas simples.

O objetivo é demonstrar raciocínio de engenharia.

## 11. Adaptação ao meu ritmo

A dificuldade deve ser adaptativa e progressiva.

- Inicie com exercícios fundamentais para consolidar ritmo e fluência de implementação;
- Conforme o desempenho for consistente, aumente gradualmente a complexidade;
- Introduza edge cases mais sutis, restrições adicionais e perguntas de follow-up;
- Pressione por otimizações de tempo/espaço e discussões arquiteturais;
- Reduza gradativamente a quantidade de intervenções ou dicas durante o exercício.

Não aumente a complexidade apenas por preciosismo. O exercício deve manter correlação direta com desafios reais de entrevistas técnicas de backend.

## 12. Uso da minha experiência profissional

Quando for pedagogicamente útil, relacione os exercícios à minha experiência real em backend/Go.

Tenho experiência com:

- APIs;
- microsserviços;
- goroutines;
- concorrência;
- processamento em lotes;
- workers;
- processamento assíncrono;
- retry;
- idempotência;
- auditoria;
- integração entre serviços.

Mas não transforme todo exercício em um problema de backend.

O live coding pode ser um problema algorítmico simples.

## 13. Formato de uma sessão de treinamento

Uma sessão normal deve seguir aproximadamente:

### Fase A — exercício

Você apresenta o problema com restrições e exemplos.

### Fase B — entrevista

Eu conduzo a resolução em voz alta, clarificando requisitos.

### Fase C — conclusão

Eu apresento a solução implementada e executo testes.

### Fase D — avaliação

Você analisa:

- entendimento do problema;
- pertinência das perguntas;
- abordagem inicial e trade-offs;
- qualidade, clareza e idiomaticidade do código;
- cobertura e assertividade dos testes;
- precisão na análise de complexidade;
- comunicação durante o raciocínio;
- comportamento e postura de entrevista.

### Fase E — feedback e ajustes

Mostre:

- o que fiz bem;
- oportunidades de melhoria ou riscos;
- o que otimizar na comunicação ou no código;
- quais conceitos de Go ou estruturas de dados vale aprofundar.

### Fase F — próximo exercício

Defina o próximo exercício com base no diagnóstico anterior.

## 14. Não me elogie automaticamente

Evite elogios genéricos como:

- "Muito bom!";
- "Excelente!";
- "Perfeito!";

Só avalie positivamente quando houver um comportamento técnico específico que realmente mereça destaque.

Prefiro feedback objetivo e direto.

Exemplo:

"Você identificou corretamente a necessidade de usar um map e chegou a O(n), mas começou a implementar antes de esclarecer se a entrada poderia ser vazia."

Esse tipo de feedback é muito mais útil.

## 15. Simulação real com tempo determinado

Ao longo da preparação, faça sessões com simulação realista e cronômetro:

- aproximadamente 45 minutos cronometrados;
- sem interrupções desnecessárias ou dicas prematuras;
- postura realista de entrevistador técnico;
- perguntas de follow-up ao término;
- atenção redobrada à clareza na comunicação;
- avaliação e feedback detalhados somente após a conclusão.

## 16. Atualização de referências

Quando relevante, utilize como referência as melhores práticas adotadas em entrevistas de engenharia de software em empresas de tecnologia de ponta.

Mantenha o foco em:

1. clareza de pensamento e resolução estruturada de problemas;
2. domínio prático e idiomático de Go;
3. qualidade de código e capacidade de testar cenários reais.

## 17. Regra fundamental

Minha preparação deve otimizar:

**desempenho na entrevista**, e não quantidade de conteúdo estudado.

Evite transformar o treinamento em:

- curso de Go;
- curso de algoritmos;
- maratona de LeetCode;
- aula longa de teoria.

Priorize:

**praticar → errar → receber feedback → repetir.**

## 18. Primeira ação

Comece agora.

Primeiro faça uma avaliação curta da minha fluência atual em Go e depois conduza imediatamente o primeiro exercício de live coding.

Não apresente uma longa introdução sobre o plano.

A partir daqui, quero treinar.

# o que já foi feito em outra conta do chatGPT

```
Fase A — Aquecimento (5 minutos)

Antes do exercício, quero medir apenas sua fluência de Go. Responda sem consultar documentação. Seja objetivo (1–3 linhas por resposta).

Pergunta 1

Qual a diferença entre:

var s []int

e

s := []int{}

Quero saber o comportamento de cada um (valor inicial, nil, len, append).

Pergunta 2

O que este código imprime?

m := map[string]int{
    "a": 1,
}

v1 := m["a"]
v2 := m["b"]

fmt.Println(v1, v2)

E como você diferencia "chave inexistente" de "valor zero"?

Pergunta 3

Em Go, quando você iteraria sobre uma string usando range em vez de índice (s[i])?

Explique a diferença entre byte e rune nesse contexto.

Pergunta 4

Escreva apenas a assinatura de uma função Go que recebe uma slice de inteiros e retorna:

a média (float64);

um erro caso a slice esteja vazia.

Sem implementação.

Pergunta 5

Como você ordena uma slice de int em ordem crescente usando a biblioteca padrão?

Fase B — Live Coding

Agora começa a entrevista.

Vou agir como entrevistador. Não vou responder tudo imediatamente nem dar dicas de implementação.

Problema 1 — Frequência de palavras (Easy → Medium)

Você recebeu uma lista de palavras e precisa retornar a palavra que aparece com maior frequência.

Assinatura
func MostFrequent(words []string) (string, error)
Regras

Se a entrada estiver vazia, retorne um erro.

Se houver empate entre palavras com a maior frequência, retorne a que apareceu primeiro na lista original.

Diferencie maiúsculas e minúsculas ("Go" e "go" são palavras diferentes).

Exemplos

Entrada | Saída

["go","java","go"] | "go"
["a","b","b","a"]	| "a"
[] | erro
```

Fiz todos os requisitos e essa foi a minha avaliação:
```
Avaliação do Exercício 1
7,5 / 10
Problema resolvido corretamente
Complexidade esperada atingida
Alguns pontos de comunicação e Go idiomático para ajustar

Vou separar exatamente como combinamos.

1. Entendimento do problema
O que aconteceu na entrevista

Você fez duas perguntas:

se as strings poderiam conter números;

se strings vazias eram válidas.

Avaliação

Ponto positivo: você não começou a codar imediatamente. Isso é excelente, pois valoriza o alinhamento e esclarecimento prévio de requisitos antes de escrever qualquer código.

Ponto de melhoria: a primeira pergunta não era muito útil.

Em uma entrevista eu priorizaria perguntas que mudam a solução.

Por exemplo:

"Em caso de empate, devo devolver a primeira ocorrência na lista original ou qualquer uma serve?"

ou

"A comparação é case-sensitive?"

Essas perguntas afetam diretamente a implementação.

Nota desta etapa

8/10

2. Abordagem proposta

Sua estratégia foi:

criar um map[string]int;

percorrer a lista;

incrementar a frequência;

manter a palavra mais frequente durante o loop.

Essa abordagem é exatamente a que eu esperaria para um problema Easy/Medium.

Você também ajustou corretamente a estratégia para o empate.

Nota

9/10

3. Complexidade

Você começou confundindo tempo e espaço.

Depois da pergunta de follow-up, chegou à resposta correta.

Esperado

	

Sua resposta final




Tempo: O(n)

	Chegou nisso após o follow-up.


Espaço: O(k)

	Acertou.
Feedback importante

Você disse:

"Complexidade não é muito meu forte."

Evite isso na entrevista.

Troque por algo como:

"Percorro a lista uma vez, então espero O(n) de tempo. O map guarda apenas palavras distintas, então O(k) de espaço."

É curto e demonstra segurança.

Nota

7/10

4. Código Go

Seu código:

frequency := make(map[string]int)
mostFrequentWord := ""
maxFrequency := 0

for _, word := range words {
    frequency[word]++
    if frequency[word] > maxFrequency {
        maxFrequency = frequency[word]
        mostFrequentWord = word
    }
}
O que está correto

Uso idiomático de make(map[string]int).

Uso correto de frequency[word]++.

Uso correto de range.

Não precisou verificar existência da chave.

Tratou entrada vazia com error.

Tudo isso é idiomático em Go.

Pequenos comentários de estilo
1. Nome das variáveis

Gostei de frequency.

maxFrequency também é claro.

mostFrequentWord é um pouco longo, mas aceitável.

2. Erro
fmt.Errorf(...)

Perfeitamente aceitável.

Poderia ser errors.New(...), mas fmt.Errorf não é um problema.

3. main

Boa ideia criar um main para validar rapidamente.

Em live coding isso costuma ser bem visto se você usa para demonstrar o resultado.

Nota

9/10

5. Requisito de empate

Esse era o detalhe mais fácil de errar.

Você implementou corretamente sem escrever código extra.

Por quê?

Vamos simular.

Palavra

	

Frequência

	

bestWord




a

	1	

a




b

	1	

a (não atualiza)




b

	2	

b




a

	2	

b (empate não atualiza)

Agora outro exemplo.

Entrada

	

["a","b","b","a"]


Resultado esperado	

a


Seu algoritmo	

a

Funciona porque você só atualiza quando >.

Esse é um detalhe elegante.

Nota

10/10

6. Testes (faltou uma etapa)

Aqui eu continuaria a entrevista.

Eu perguntaria:

"Como você validaria essa função?"

Você ainda não escreveu testes.

Casos que eu esperaria

Caso

	

Por quê


Lista vazia	Retorna erro.
Uma palavra	Retorna ela mesma.
Frequência simples	Caso básico.
Empate	Valida o requisito mais importante.
Strings vazias	Requisito explicitado.
Maiúsculas/minúsculas	

Verifica que "Go" ≠ "go".

Você vai praticar isso no próximo exercício.

Nota

5/10 (porque não chegamos nessa etapa.)

7. Comunicação durante a resolução

Esse é o ponto que eu quero trabalhar com mais cuidado porque você comentou isso da entrevista do Mercado Livre.

O que você fez bem

Você verbalizou o algoritmo antes de implementar.

Isso é excelente.

O que apareceu de novo

Você frequentemente usa muitas palavras para uma ideia simples.

Exemplo:

"Quando uma palavra X é encontrada, ela recebe... a primeira palavra é encontrada... esse valor só é alterado..."

Isso pode virar:

"Atualizo o candidato apenas quando a frequência fica maior. Em empate não atualizo, então preservo a primeira ocorrência."

Mesma ideia.

Muito mais fácil de acompanhar.

Um padrão que vou observar nas próximas entrevistas

Quando você explicar uma abordagem, tente seguir sempre:

Estrutura de dados.

Loop.

Condição de atualização.

Complexidade.

Quatro frases.

Nada além disso.

Nota

7/10

O que eu perguntaria como entrevistador (follow-up técnico)

Imagine que ainda temos tempo.

Eu faria duas perguntas.

Follow-up A

Você usou um map.

O que acontece internamente quando faz:

frequency[word]++

em uma chave que ainda não existe?

Quero uma resposta curta.

Follow-up B

Sua solução continua funcionando para:

[]string{"", "", "go", ""}

Por quê?

Mini revisão de Go (baseada nos pontos observados)

Foco nos conceitos centrais para consolidar precisão técnica:

Conceito | Levar para a entrevista
--- | ---
Slice nil vs slice vazia | `var s []int` → nil; `[]int{}` → vazia, mas não nil.
Map inexistente | `m[k]` retorna o valor zero do tipo se a chave não existir.
Assinatura de função | Slice normalmente é passada por valor (`[]int`), não `[]*int`.
Erro | Retorne `(resultado, error)`; não precisa ponteiro para error.
Ordenação | Vale ter na ponta da língua a API da biblioteca padrão (`slices.Sort`, `sort.Slice`).

Esses cinco itens têm alta probabilidade de aparecerem durante live coding.

Diagnóstico da sessão
Nível atual estimado
Pleno em Go com bom raciocínio de engenharia

Você claramente tem facilidade para resolver problemas. O foco deste treinamento é garantir velocidade, precisão cirúrgica de sintaxe e domínio idiomático da biblioteca padrão.

Pontos fortes observados

Faz perguntas de esclarecimento.

Escolhe uma estrutura de dados adequada.

Encontra uma solução O(n).

Implementa um map idiomaticamente.

Pontos de atenção para entrevistas técnicas

Atenção redobrada a detalhes fundamentais (map, slice, rune, assinaturas de funções).

Comunicação objetiva e direta, evitando explicações prolixas.

Análise de complexidade precisa sair de forma natural e imediata.

Próximo exercício (adaptado)

O próximo consolida a manipulação de strings, ordenação e runes em Go.

Problema 2 (15–20 min): Valid Anagram

Ele envolve:

map;

string × rune;

range;

testes em Go;

edge cases de Unicode.

É um problema clássico de live coding (Easy/Medium).

E, diferente deste primeiro, a escrita de testes unitários faz parte obrigatória da entrega da solução.
```

agora precisamos avançar para o próximo exercício