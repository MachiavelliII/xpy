package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/xmppo/go-xmpp"
)

func main() {
        // Define your credentials and recipient details
        username := "your_username@thesecure.biz"
        password := "your_password"
        recipient := "recipient@thesecure.biz"
        message := "Hello from Go xpy!"

        // Set up XMPP options with TLS enabled
        options := xmpp.Options{
                Host:          "xmpp_server.com:5222",  // XMPP server and port
                User:          username,
                Password:      password,
                NoTLS:         false,                   // Require TLS for secure connection
                StartTLS:      true,                    // Enable STARTTLS if supported by the server
                TLSConfig:     &tls.Config{InsecureSkipVerify: true},  // Adjust for local server testing; set to false in production
                Debug:         true,                    // Enable debugging output (Actually it's not necessary)
                Session:       true,
                Status:        "chat",
                StatusMessage: "xpy Hello !",
        }

        // Create an XMPP client
        client, err := options.NewClient()
        if err != nil {
                log.Fatalf("Failed to create XMPP client: %v", err)
        }

        // Send a message to the recipient
        _, err = client.Send(xmpp.Chat{
                Remote: recipient,
                Type:   "chat",
                Text:   message,
        })
        if err != nil {
                log.Fatalf("Failed to send message: %v", err)
        }
        fmt.Printf("Message sent to %s: %s\n", recipient, message)

        // Listen for incoming messages
        go func() {
                for {
                        chat, err := client.Recv()
                        if err != nil {
                                log.Printf("Failed to receive message: %v", err)
                                break
                        }
                        switch v := chat.(type) {
                        case xmpp.Chat:
                                fmt.Printf("Received message from %s: %s\n", v.Remote, v.Text)
                        case xmpp.Presence:
                                fmt.Printf("Presence from %s: %s\n", v.From, v.Show)
                        }
                }
        }()
		// API key: <2_@l?/9VoiZB56&6NPHPX#mK_gS$p&(EU'@egTc:ym/X)#/E_<yb
        // Keep the program running to receive messages
        time.Sleep(10 * time.Second)
}