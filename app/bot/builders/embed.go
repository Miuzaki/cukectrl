package builders

import "github.com/bwmarrin/discordgo"

// EmbedBuilder é uma estrutura para facilitar a construção de embeds de forma modular.
type EmbedBuilder struct {
	embed *discordgo.MessageEmbed
}

// NewEmbed cria um novo EmbedBuilder com um embed vazio.
func NewEmbed() *EmbedBuilder {
	return &EmbedBuilder{
		embed: &discordgo.MessageEmbed{},
	}
}

// SetTitle define o título do embed.
func (b *EmbedBuilder) SetTitle(title string) *EmbedBuilder {
	b.embed.Title = title
	return b
}

// SetDescription define a descrição do embed.
func (b *EmbedBuilder) SetDescription(description string) *EmbedBuilder {
	b.embed.Description = description
	return b
}

// SetColor define a cor do embed.
func (b *EmbedBuilder) SetColor(color int) *EmbedBuilder {
	b.embed.Color = color
	return b
}

// AddField adiciona um campo ao embed.
func (b *EmbedBuilder) AddField(name string, value string, inline bool) *EmbedBuilder {
	b.embed.Fields = append(b.embed.Fields, &discordgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	})
	return b
}

// SetFooter define o rodapé do embed.
func (b *EmbedBuilder) SetFooter(text string, iconURL string) *EmbedBuilder {
	b.embed.Footer = &discordgo.MessageEmbedFooter{
		Text:    text,
		IconURL: iconURL,
	}
	return b
}

// SetThumbnail define a thumbnail do embed.
func (b *EmbedBuilder) SetThumbnail(url string) *EmbedBuilder {
	b.embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL: url,
	}
	return b
}

func (b *EmbedBuilder) Build() *discordgo.MessageEmbed {
	return b.embed
}

func CreateEmbedResponse(embed *discordgo.MessageEmbed) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	}
}
