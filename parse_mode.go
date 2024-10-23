package aquagram

type ParseMode int8

const (
	ParseModeDefault ParseMode = iota
	ParseModeDisabled
	ParseModeHTML
	ParseModeMarkdown
	ParseModeMarkdownV2
)

func (mode ParseMode) String() string {
	switch mode {
	case ParseModeHTML:
		return "html"
	case ParseModeMarkdown:
		return "markdown"
	case ParseModeMarkdownV2:
		return "markdownv2"
	default:
		return ""
	}
}

func (mode ParseMode) MarshalJSON() ([]byte, error) {
	return []byte("\"" + mode.String() + "\""), nil
}

func (bot *Bot) ParseMode(mode ParseMode) string {
	if mode == ParseModeDefault {
		return bot.Config.DefaultParseMode.String()
	}

	return mode.String()
}
