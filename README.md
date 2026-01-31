# Echo (VATSIM Alias Trainer)

Uma plataforma de treinamento focada na memorização e agilidade de comandos de texto (alias) para controladores da rede VATSIM.

A ideia é simples: **Sem radar, apenas texto.** O sistema apresenta um cenário de voo e o usuário deve responder com o alias correto no tempo certo.
## Tech Stack
Projeto desenvolvido para aprendizado de arquitetura distribuída.

* **Ingestão de Dados:** Go
* **Core / Regras de Negócio:** Kotlin (Spring Boot)
* **Mensageria:** RabbitMQ
* **Banco de Dados:**
    * PostgreSQL (Dados Relacionais/Usuários)
    * MongoDB (Histórico/Cenários)
    * Redis (Cache de Sessão)
* **Infraestrutura:** Docker & Docker Compose