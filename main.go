package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"

    "github.com/gorilla/handlers"
)

// Função para lidar com o envio de email
func sendEmail(to, name, email, service, message string) error {
	from := "crrobg@gmail.com"        // O email de onde será enviado
	password := "grgi phia pyxc kzkw" // A senha do email

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Configuração da mensagem de email
	body := fmt.Sprintf("Nome: %s\nEmail: %s\nServiço: %s\nMensagem: %s", name, email, service, message)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: Solicitação de Orçamento\r\n\r\n" +
		body + "\r\n")

	// Configuração da autenticação SMTP
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Enviar o email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return err
	}

	return nil
}

// Função para processar os dados do formulário
func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	service := r.FormValue("service")
	message := r.FormValue("message")

	// Chama a função para enviar o email
	err := sendEmail("crrobg@gmail.com", name, email, service, message)
	if err != nil {
		http.Error(w, "Falha ao enviar o email", http.StatusInternalServerError)
		log.Printf("Erro ao enviar email: %v\n", err)
		return
	}

	// Resposta de sucesso
	fmt.Fprintf(w, "Email enviado com sucesso!")
}

func main() {
	http.HandleFunc("/send", handleForm)

	// Adiciona o middleware CORS
	loggedRouter := handlers.LoggingHandler(log.Writer(), http.DefaultServeMux)

	// Configura CORS
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Permite todas as origens
		handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	// Inicia o servidor com suporte a CORS
	fmt.Println("Servidor rodando em http://localhost:8080")
	if err := http.ListenAndServe(":8080", cors(loggedRouter)); err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
