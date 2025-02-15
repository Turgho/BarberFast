package repositories

type Agendamento struct {
	DataInicio string `json:"data_inicio"`
	DataFim    string `json:"data_fim"`
	Status     string `json:"status"` // Ex: "confirmado", "cancelado"
}
