package rabbitmq

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/streadway/amqp"
	"gopkg.in/gomail.v2"
)

// sendEmail envia um e-mail usando Mailtrap
func sendEmail(content string) error {
	// Configura a mensagem de e-mail
	m := gomail.NewMessage()
	m.SetHeader("From", "youremail@example.com") // Substitua pelo seu e-mail
	m.SetHeader("To", "recipient@example.com")   // Substitua pelo destinatário real
	m.SetHeader("Subject", "Assunto do E-mail")
	m.SetBody("text/plain", content)

	// Configura o servidor SMTP do Mailtrap
	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_APP_PASSWORD")

	// Converter a porta de string para int
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Erro ao converter SMTP_PORT para int:", err)
		return err
	}

	d := gomail.NewDialer(host, port, user, password)

	// Envia o e-mail
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("erro ao enviar e-mail: %w", err)
	}

	fmt.Println("E-mail enviado!")
	return nil
}

// ConsumeMessages consome mensagens do RabbitMQ e envia e-mails
func ConsumeMessages() {
	// Conectar ao RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5673/") // Verifique a porta
	if err != nil {
		log.Fatalf("Erro ao conectar ao RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Criar um canal de comunicação
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Erro ao criar canal: %v", err)
	}
	defer ch.Close()

	// Declarar a fila
	q, err := ch.QueueDeclare(
		"emailQueue", // Nome da fila
		true,         // Durável
		false,        // Excluir quando não for mais necessário
		false,        // Exclusivo
		false,        // Aguardar a confirmação do servidor
		nil,          // Propriedades adicionais
	)
	if err != nil {
		log.Fatalf("Erro ao declarar fila: %v", err)
	}

	// Consumir mensagens da fila
	msgs, err := ch.Consume(
		q.Name, // Nome da fila
		"",     // Consumer
		true,   // Auto-acknowledge
		false,  // Exclusivo
		false,  // Não aguardar a confirmação do servidor
		false,  // Não pode ser consumido por outro consumidor
		nil,    // Propriedades adicionais
	)
	if err != nil {
		log.Fatalf("Erro ao consumir mensagens: %v", err)
	}

	// Processar as mensagens
	for msg := range msgs {
		// Verificar se a mensagem contém um corpo
		if len(msg.Body) == 0 {
			log.Println("Mensagem recebida sem conteúdo.")
			continue
		}

		// Tenta enviar o e-mail
		err := sendEmail(string(msg.Body))
		if err != nil {
			log.Printf("Erro ao enviar e-mail: %v", err)
		} else {
			log.Printf("E-mail enviado com sucesso: %s", string(msg.Body))
		}
	}
}
