package modules

// Item são os itens encontrados pela promoção
type Item struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Preco string `json:"preco"`
	Link  string `json:"link"`
}
