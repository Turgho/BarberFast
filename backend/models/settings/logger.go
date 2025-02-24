package settings

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Função para configurar o arquivo de log
func SetupLogging() *os.File {
	// Define o diretório de logs
	logDir := "../logs"

	// Garante que o diretório de logs existe
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("Erro ao criar diretório de logs: %s", err)
	}

	// Cria o nome do arquivo de log com base na data
	logFileName := fmt.Sprintf("%s/app-%s.log", logDir, time.Now().Format("2006-01-02"))

	// Abre ou cria o arquivo de log
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo de log: %s", err)
	}

	// Configura o log para usar o arquivo como saída
	log.SetOutput(logFile)
	// Configura o formato do log
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Retorna o arquivo de log para fechá-lo depois
	return logFile
}
