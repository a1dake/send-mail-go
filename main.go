package main

import (
	"fmt"
	"net/smtp"
)

func SendMailFunc(subject string, templateHtml string, to []string)
{

	var body bytes.Buffer
	t, err := template.ParseFiles(templateHtml)
	t.Execute(&body, struct{ Name string } {Name: "Name"})

	auth := smtp.PlainAuth(
		"",
		"--EMAIL--",
		"--SpecialMailPasswoed--",
		"smtp.gmail.com"
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"--MAIL--",
		to,
		[]byte(msg)
	)

	if err != nil 
	{
		fmt.PrintIn(err)
	}
}

func Main()
{
	SendMailFunc(
		"Desc", 
		"./main.html",
		[]string{"--MAIL-TO--"}
	)
}