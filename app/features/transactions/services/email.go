package services 

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/wanta-zulfikri/Event-Planning-App/config"
    "github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
	gomail "gopkg.in/mail.v2"
)


const (
	FrontEndURL = "http://goapp.com/INV-2323233232" 
	
    EmailHost = "smpt.gmail.com"
) 


type subjectBody struct {
	subject string 
	body   byte.Buffer
}  


func SendEmailPendingPayment(sdata config.SenderConfig, rdata transactions.Data) {
	log.Printf("[INFO] Sending email pending payment from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Tagihan %s", rdata.Invoice),
		body:    bytes.Buffer{},
	}

	t, err := getTemplate("pending.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		TWT     string
		IG      string
		FB      string
		URL     string
		Email   string
		Telpon  string
		Name    string
		Slogan  string
		Cusname string
		Pcode   string
		Invoice string
		Total   int
		Pmethod string
		Expire  string
	}{
		URL:     FrontEndURL + rdata.Invoice,
		TWT:     sdata.Twitter,
		FB:      sdata.Facebook,
		IG:      sdata.Instagram,
		Email:   sdata.Email,
		Telpon:  sdata.Phone,
		Cusname: rdata.Name,
		Invoice: rdata.Invoice,
		Slogan:  sdata.Slogan,
		Pmethod: rdata.PaymentMethod,
		Name:    sdata.Name,
		Total:   rdata.Total,
		Expire:  rdata.Expire,
		Pcode:   rdata.PaymentCode,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, rdata.Email, sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Email pending payment sent to %s", rdata.Email)

}
func SendEmailSuccessPayment(sdata config.SenderConfig, rdata transactions.Data) {
	log.Printf("[INFO] Sending email success payment from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Berhasil %s", rdata.Invoice),
		body:    bytes.Buffer{},
	}

	t, err := getTemplate("success.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		TWT     string
		IG      string
		FB      string
		URL     string
		Email   string
		Telpon  string
		Name    string
		Slogan  string
		Cusname string
		Invoice string
	}{
		URL:     FrontEndURL + rdata.Invoice,
		TWT:     sdata.Twitter,
		FB:      sdata.Facebook,
		IG:      sdata.Instagram,
		Email:   sdata.Email,
		Telpon:  sdata.Phone,
		Cusname: rdata.Name,
		Invoice: rdata.Invoice,
		Slogan:  sdata.Slogan,
		Name:    sdata.Name,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, rdata.Email, sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Email success payment sent to %s", rdata.Email)
}

func SendEmailCancelPayment(sdata config.SenderConfig, rdata transactions.Data) {
	log.Printf("[INFO] Sending email cancel payment from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Pembatalan %s", rdata.Invoice),
		body:    bytes.Buffer{},
	}

	t, err := getTemplate("cancel.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		TWT     string
		IG      string
		FB      string
		URL     string
		Email   string
		Telpon  string
		Name    string
		Slogan  string
		Cusname string
		Invoice string
	}{
		URL:     FrontEndURL + rdata.Invoice,
		TWT:     sdata.Twitter,
		FB:      sdata.Facebook,
		IG:      sdata.Instagram,
		Email:   sdata.Email,
		Telpon:  sdata.Phone,
		Cusname: rdata.Name,
		Invoice: rdata.Invoice,
		Slogan:  sdata.Slogan,
		Name:    sdata.Name,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, rdata.Email, sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Email cancel payment sent to %s", rdata.Email)

}
func SendEmailRefundPayment(sdata config.SenderConfig, rdata transactions.Data) {
	log.Printf("[INFO] Sending email refund payment from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Pengembalian %s", rdata.Invoice),
		body:    bytes.Buffer{},
	}

	t, err := getTemplate("refund.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		TWT     string
		IG      string
		FB      string
		URL     string
		Email   string
		Telpon  string
		Name    string
		Slogan  string
		Cusname string
		Invoice string
	}{
		URL:     FrontEndURL + rdata.Invoice,
		TWT:     sdata.Twitter,
		FB:      sdata.Facebook,
		IG:      sdata.Instagram,
		Email:   sdata.Email,
		Telpon:  sdata.Phone,
		Cusname: rdata.Name,
		Invoice: rdata.Invoice,
		Slogan:  sdata.Slogan,
		Name:    sdata.Name,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, rdata.Email, sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Email refund payment sent to %s", rdata.Email)

}

func getTemplate(htmlFile string) (t *template.Template, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fmt.Println("ini wd", wd)

	wd = wd + "/template/"

	t, err = template.ParseFiles(wd + htmlFile)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func sendEmail(emailSender, passSender, emailReceiver string, sb subjectBody) error {
	m := gomail.NewMessage()

	m.SetHeader("From", emailSender)
	m.SetHeader("To", emailReceiver)
	m.SetHeader("Subject", sb.subject)
	m.SetBody("text/html", string(sb.body.Bytes()))

	d := gomail.NewDialer(EmailHost, 587, emailSender, passSender)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}