# Fábrica de Dados de Teste

Este projeto fornece um conjunto de funções auxiliares para criar dados de teste em Go, facilitando a escrita de testes unitários e de integração em projetos que seguem a Clean Architecture e a Arquitetura Hexagonal.

## Abordagem:

- **Funções de Fábrica**: Cada struct possui uma função de fábrica correspondente (por exemplo, `NewCompanyFactory()`, `NewUserFactory()`, etc.) que retorna uma instância da struct preenchida com valores padrão úteis para testes.
- **Modularidade e Reutilização**: As fábricas são modulares e reutilizáveis, permitindo criar objetos complexos a partir de objetos mais simples.
- **Flexibilidade**: As fábricas podem ser customizadas com parâmetros ou opções para criar objetos com diferentes características.
- **Organização**: As fábricas são mantidas em uma pasta separada (`test/factories`) para manter o código de teste organizado e isolado do código da aplicação.

## Exemplo de Uso:

```go
package meu-pacote-de-teste

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/br4tech/auth-nex/internal/test/factories" 
)

func TestCreateCompanyUseCase_Execute(t *testing.T) {

    comapny := factories.NewCompanyFactory()

    // ... realiza alguma operação com a empresa

    // Faz asserções sobre o resultado da operação
    assert.Equal(t, resultadoEsperado, resultadoObtido)
} ```

## Fábricas Disponíveis:

- `NewAddressFactory()`
- `NewUserFactory()`
- `NewPartnerFactory()`
- `NewActivityFactory()`
- `NewCompanyFactory()`
- `NewPermissionFactory()`
- `NewTenantFactory()`


## Como Usar:

1. **Importe o pacote de fábricas**: `import "github.com/br4tech/auth-nex/internal/test/factories"`
2. **Chame as funções de fábrica**: `empresa := factories.NewTestCompany()`
3. **Utilize os objetos de teste**: Utilize os objetos criados pelas fábricas em seus testes.
4. **Adapte e Expanda**: Adapte as fábricas existentes e crie novas fábricas conforme necessário para atender às necessidades específicas do seu projeto.

## Observacao

1. Ao utilizar essa abordagem verifiquei que se nao esta sendo chamado a mesma factory dentro dela mesmo, caso isso ocorra sera gerado um erro de recursividade infinita.

