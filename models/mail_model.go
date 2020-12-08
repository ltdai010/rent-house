package models

import (
	"github.com/astaxie/beego"
	"log"
	"net/smtp"
)

type Mail struct {
	To 		string `json:"to"`
	Subject string `json:"subject"`
	Msg     string `json:"msg"`
}



func (m *Mail) SendMail() error {
	to := "To: " + m.To +"\r\n"
	subject := "Subject: " + m.Subject  +"\r\n"
	message := []byte(to + subject + "\n" + m.Msg + "\r\n")
	err := smtp.SendMail(beego.AppConfig.String("gmail::host_port"), EmailAuth, EmailFrom, []string{"luongdai246@gmail.com"}, message)
	if err != nil {
		log.Println(err)
	}
	return err
}
