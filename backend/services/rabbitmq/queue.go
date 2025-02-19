package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

func SendMessageToQueue(message string) error {
	// Conectar ao RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5673/")
	if err != nil {
		return fmt.Errorf("erro ao conectar ao RabbitMQ: %w", err)
	}
	defer conn.Close()

	// Criar um canal de comunicação
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("erro ao criar canal: %w", err)
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
		return fmt.Errorf("erro ao declarar fila: %w", err)
	}

	// Enviar mensagem para a fila
	err = ch.Publish(
		"",     // Exchange
		q.Name, // Fila
		false,  // Não persistente
		false,  // Não aguardar confirmação
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return fmt.Errorf("erro ao enviar mensagem: %w", err)
	}
	fmt.Println("Mensagem enviada para a fila")
	return nil
}
