package test

import (
	"backen-oj/define"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

/**
发送邮件测试
 */

func TestSend(t *testing.T){
	e := email.NewEmail()
	e.From = "Get <1131927131@qq.com>"
	e.To = []string{"771002444@qq.com"}
	e.Subject = "邮箱发送测试"
	e.HTML = []byte("<h1>邮箱发送测试</h1>")
	err := e.SendWithTLS("smtp.qq.com:465",
		smtp.PlainAuth("", "1131927131@qq.com", define.MailPassword, "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		t.Fatal(err)
	}
}
