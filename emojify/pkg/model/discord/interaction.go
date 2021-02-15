package discord

type (
	Interaction struct {
		Type int             `json:"type"`
		Data InteractionData `json:"data"`
	}

	InteractionData struct {
		Options []OptionsData `json:"options"`
	}

	OptionsData struct {
		Name string `json:"name"`
		Value string `json:"value"`
	}
)
