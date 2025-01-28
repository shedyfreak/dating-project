package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"

	"github.com/perso-proj/database"
)

func sendEmail(event database.Event, sub database.Subsriber) {
	// Sender data
	from := "customer@affinitys.ch"
	password := "Ch19980205?!"

	// Receiver data

	// SMTP server data
	smtpHost := "mail.infomaniak.com"
	smtpPort := "465" // Typically 587 for TLS or 465 for SSL

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Create a connection to the server using SSL
	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, &tls.Config{
		InsecureSkipVerify: false, // Optional, if you have issues with certificate verification
		ServerName:         smtpHost,
	})
	if err != nil {
		log.Println("Failed to connect: ", err)
		return
	}
	defer conn.Close()

	// Create an SMTP client from the SSL connection
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		log.Println("Failed to create SMTP client: ", err)
		return
	}

	// Authenticate with the server
	if err := client.Auth(auth); err != nil {
		log.Println("Failed to authenticate: ", err)
		return
	}

	// Set the sender and recipient
	if err := client.Mail(from); err != nil {
		log.Println("Failed to set sender: ", err)
		return
	}
	to := []string{sub.Email}
	for _, addr := range to {
		if err := client.Rcpt(addr); err != nil {
			log.Println("Failed to set recipient: ", err)
			return
		}
	}
	subject := "Subject: Welcome to Our Service!\n"
	mime := "MIME-Version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := subject + mime + emailTempl
	// Write the message to the server
	htmlMessage := []byte(message)
	writer, err := client.Data()
	if err != nil {
		log.Println("Failed to write data: ", err)
		return
	}
	_, err = writer.Write(htmlMessage)
	if err != nil {
		log.Println("Failed to write message: ", err)
		return
	}

	// Close the connection
	err = writer.Close()
	if err != nil {
		log.Println("Failed to close writer: ", err)
		return
	}

	// Quit the session
	client.Quit()

	fmt.Println("Email sent successfully!")
}

var emailTempl = `<!DOCTYPE html>
<html>
<head>
  <style>
    body {
      font-family: Arial, sans-serif;
      line-height: 1.6;
      background-color: #f9f9f9;
      margin: 0;
      padding: 0;
    }
    .email-container {
      max-width: 600px;
      margin: 20px auto;
      background: #ffffff;
      border: 1px solid #dddddd;
      border-radius: 8px;
      overflow: hidden;
    }
    .header {
      background-color: #4caf50;
      color: white;
      padding: 20px;
      text-align: center;
    }
    .content {
      padding: 20px;
      color: #333333;
    }
    .footer {
      background-color: #f1f1f1;
      text-align: center;
      padding: 10px;
      font-size: 12px;
      color: #666666;
    }
    a {
      color:rgb(250, 47, 247);
      text-decoration: none;
    }
  </style>
</head>
<body>
  <div class="email-container">
    <div class="header">
      <h1>Bienvenue!</h1>
    </div>
    <div class="content">
      <p>Cher Saverio,</p>
      <p>Merci d'être inscrit à notre evenement et à bientot TODO : ADD IMAGE;SUBSCRIBER NAME; EVENT NAME,DATE AND LOCATION</p>
      <p>Si t'as des questions <a href="mailto:Customer@affinitys.ch">contact</a>.</p>
      <p>Cordialement,<br> Affinitys Team</p>
    </div>
    <div class="footer">
      <p>&copy; 2025 Affinitys Sàrl. All rights reserved.</p>
    </div>
  </div>
</body>
</html>
`
