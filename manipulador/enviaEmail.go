package manipulador

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"

	"github.com/brunobog/enviarEmail/modules"
)

// EnviaEmail para enviar email com os dados
func EnviaEmail(w http.ResponseWriter, r *http.Request) {

	url := r.URL.String()
	log.Println(url)

	// Lendo a respsta
	corpo, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println("[EnviarEmail] Erro ao ler o r.body", err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	// log.Println(string(corpo))
	// UnMarshall
	var pessoa modules.Pessoa
	if err := json.Unmarshal(corpo, &pessoa); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	if err := r.Body.Close(); err != nil {
		log.Println("[EnviarEmail] Erro ao fechar oo Body")
		http.Error(w, err.Error(), 400)
		panic(err)
	}

	log.Println(pessoa)
	sendPromo(pessoa)
	return
}

func sendPromo(pessoa modules.Pessoa) {
	from := "anymail@gmail.com"
	pass := "Not Easy ;)"
	to := pessoa.Email
	body := "Encontramos seu produto " + pessoa.Produto.Nome
	body += "\nCom o preço: " + pessoa.Produto.Preco
	body += "\n No link: " + pessoa.Produto.Link

	msg := "From: " + "Seu muambator preferido@nada.com" + "\n" +
		"To: " + to + "\n" +
		"Subject: Encontramos o Produto\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("[EnviarEmail] smtp error: %s", err)
		return
	}
}
