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

Esse projeto utiliza o viper para tratar as ENVS.

Configurar o arquivo `.env` para refletir seu cenario de banco de dados.Em nosso modelo estamos utilizando uma imagem docker do mysql.

- **Banco de dados:**
  MySQL: `docker run --name auth_nex -e MYSQL_ROOT_PASSWORD=admin -d -p 3306:3306 mysql --bind-address=0.0.0.0`

- **Wire:**
  O Wire é uma biblioteca em Go que simplifica a injeção de dependência e a criação de gráficos de objetos complexos. Com o Wire, é possível definir facilmente dependências entre diferentes componentes do seu aplicativo e gerar automaticamente o código necessário para construir esses componentes
  
  Execute na pasta CMD: `go run github.com/google/wire/cmd/wire`

- **Rodar aplicacao:**
   `go run cmd/server/main.go  cmd/server/wire_gen.go`

