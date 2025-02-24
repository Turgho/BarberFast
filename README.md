# 🏆 BarberFast

**BarberFast** é um sistema de gestão para barbearias, desenvolvido para otimizar o agendamento de clientes, gerenciamento de profissionais e controle financeiro. O objetivo é proporcionar uma experiência eficiente para barbearias e seus clientes.

## 🚀 Tecnologias Utilizadas

### Backend
- **Golang** - API REST para manipulação de dados
- **Gorm** - ORM para interação com MySQL
- **MySQL** - Banco de dados relacional
- **RabbitMQ** - Mensageria para comunicação assíncrona
- **JWT** - Autenticação segura
- **Swagger** - Documentação da API

### Frontend
- **Next.js** - Framework React para construção da interface
- **Tailwind CSS** - Estilização moderna e responsiva

## 📦 Instalação

### 1️⃣ Pré-requisitos
Certifique-se de ter instalado:
- [Go](https://go.dev/dl/)
- [Node.js](https://nodejs.org/)
- [MySQL](https://www.mysql.com/)
- [RabbitMQ](https://www.rabbitmq.com/)
  
> Container do RabbitMQ precisa estar rodando de fundo no docker!

### 2️⃣ Configuração do Banco de Dados
1. Crie um banco de dados no MySQL:
   ```sql
   CREATE DATABASE barberfast_db;
   ```
   
2. Crie um usuário para acessar o Banco de Dados:
   ```sql
   CREATE USER 'novo_usuario'@'localhost' IDENTIFIED BY 'minhasenha';
   GRANT ALL PRIVILEGES ON barberfast_db.* TO 'meu_usuario'@'localhost' WITH GRANT OPTION;
   FLUSH PRIVILEGES;
   ```
  
4. Configure seu arquivo `.env` na raiz do projeto:
   ```env
   DB_USER=meu_usuario
   DB_PASSWORD=minha_senha
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=barberfast_db

   RABBITMQ_URL="amqp://guest:guest@localhost:5673/"
   SMTP_APP_PASSWORD="sua_senha_app"
   SMTP_HOST="smtp.gmail.com"
   SMTP_PORT=587
   SMTP_USER="seu_gmail"
   ```

   > **Atenção:** A senha do aplicativo deve ser gerada diretamente no Google para autenticação SMTP.
   > 
   > Para mais informações, acesse: [Link para geração de senha de aplicativo](https://support.google.com/accounts/answer/185833?hl=pt-BR)

### 3️⃣ Executando o Backend
```sh
cd BarberFast/backend/cmd
go run main.go
```

### 4️⃣ Executando o Frontend
```sh
cd BarberFast/frontend
npm install
npm run dev
```

### Acessando o Site
```sh
http://localhost:3000
```

### Acessando a documentação
```sh
http://localhost:5050/swagger/index.html
```

## 📌 Recursos
- ✅ Cadastro e gerenciamento de clientes e profissionais 
- ✅ Agendamento online de horários 
- ✅ Controle financeiro da barbearia 
- ✅ Interface amigável e responsiva 
- ✅ Autenticação segura com JWT 
- ✅ Comunicação assíncrona com RabbitMQ 
- ✅ Documentação da API com Swagger

## 📄 Contribuição
Se deseja contribuir, siga os passos:
1. Faça um **fork** do repositório
2. Crie uma **branch** (`git checkout -b feature/minha-feature`)
3. Faça o **commit** das suas alterações (`git commit -m 'Minha feature'`)
4. Faça o **push** para a branch (`git push origin feature/minha-feature`)
5. Abra um **Pull Request**

## 📜 Licença
Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE.txt) para mais detalhes.

---

Atenciosamente, Turgho.

