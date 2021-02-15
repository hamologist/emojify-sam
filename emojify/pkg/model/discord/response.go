package discord

type (
	Response struct {
		Type int          `json:"type"`
		Data ResponseData `json:"data"`
	}

	ResponseData struct {
		TTS             bool     `json:"tts"`
		Content         string   `json:"content"`
		Embeds          []string `json:"embeds"`
		AllowedMentions []string `json:"allowed_mentions"`
	}
)
