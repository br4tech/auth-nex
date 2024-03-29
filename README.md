# AuthNex

O AuthNex é um sistema de autenticação e autorização desenvolvido em Go para garantir segurança e controle de acesso eficientes em ambientes diversos. Com uma arquitetura flexível e modular, o AuthNex oferece uma solução completa para autenticação de usuários, controle granular de permissões e suporte multitenancy.

## Recursos Principais

- **Autenticação Segura:**
  - Baseada em JWT para autenticação robusta.
  - Gerenciamento flexível de credenciais de usuários.

- **Autorização Personalizada:**
  - Controle granular de permissões com base em RBAC ou políticas específicas.
  - Concessão e revogação dinâmica de permissões.

- **Multitenancy Integrado:**
  - Suporte completo para ambientes multitenants.
  - Associação de usuários a tenants para gerenciamento eficiente.

- **APIs Restful:**
  - APIs bem definidas para manipulação de usuários, permissões e configurações multitenants.
  - Comunicação segura via HTTPS.

- **Middleware de Segurança:**
  - Utilização de middlewares para autenticação e autorização em todas as requisições.
  - Proteção contra ataques comuns, como CSRF e injeção SQL.

- **Escalabilidade e Desempenho:**
  - Projetado para escalabilidade, permitindo o gerenciamento eficiente de grandes volumes de usuários e permissões.
  - Otimização de consultas de banco de dados para desempenho máximo.

## Como Usar

O AuthNex é fácil de implementar e se integra perfeitamente a diferentes tipos de aplicativos, desde sistemas web até APIs RESTful. Consulte nossa documentação abrangente e exemplos práticos para começar a utilizar o AuthNex e aprimorar a segurança e o controle de acesso em seu projeto.

Transforme a autenticação e autorização em uma experiência segura e flexível com o AuthNex em Go.

## Dev

database:
  ``` 
    docker run --name authnex -e POSTGRES_PASSWORD=123456 -d -p 5434:5432 postgres
    
    docker exec -it authnex bash

    psql -U postgres

    CREATE DATABASE authnexdb;
  ```


  Para executar a aplicacao devemos executar o comando abaixo:

  ```
   go run cmd/api/wire_gen.go cmd/api/main.go
   
  ```
   

![Diagrama em branco](https://github.com/br4tech/auth-nex/assets/26689902/13605cde-617b-46d6-a041-779d5a1bee2b)

