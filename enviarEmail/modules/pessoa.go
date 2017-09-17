package modules

//Pessoa Para identificar a pessoa que está enviando o email
type Pessoa struct {
	Nome    string `json:"nome"`
	Email   string `json:"email"`
	Produto Item   `json:"produto"`
}
