package tasks

import (
	"encoding/base64"
	"encoding/json"
	"log"
)

type PayLoad struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func DecodeTask(msg string, task interface{}) (err error) {
	decodeMsg, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return
	}
	//msgBytes := []byte(decodeMsg)
	err = json.Unmarshal(decodeMsg, task)
	if err != nil{
		return
	}
	return
}

func SendEmail(base64PayLoad string) (bool, error){
	payload := PayLoad{}
	DecodeTask(base64PayLoad, &payload)

	from := "<your-google-email>@gmail.com"
	//pass := "<your-google-email-app-password>"
	to := payload.Email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: "+ payload.Subject +"\n\n" +
		payload.Body

	//err := smtp.SendMail("smtp.gmail.com:587",
	//	smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
	//	from, []string{to}, []byte(msg))

	//if err != nil {
	//	log.Printf("smtp error: %s", err)
	//	return false, nil
	//}
	log.Println(msg, "is sent to ", to, "from", from)

	return true, nil
}
