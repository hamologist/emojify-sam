package discord

import (
	"emojify/pkg/model/discord"
	"github.com/hamologist/emojify-go/pkg/model"
	"github.com/hamologist/emojify-go/pkg/transformer"
)

func InteractionToResponse(interaction *discord.Interaction) (*discord.Response, error) {
	emojifyPayload := model.EmojifyPayload{
		Message: interaction.Data.Options[0].Value,
	}

	emojifyResponse, err := transformer.EmojifyTransformer(emojifyPayload)
	if err != nil {
		return nil, err
	}

	return &discord.Response{
		Type: 4,
		Data: discord.ResponseData{
			TTS:             false,
			Content:         emojifyResponse.Message,
			Embeds:          []string{},
			AllowedMentions: []string{},
		},
	}, nil
}
