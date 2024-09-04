package aquagram

import (
	"encoding/json"
	"strconv"
)

type WebhookInfo struct {
	URL                          string       `json:"url"`
	HasCustomCertificate         bool         `json:"has_custom_certificate,omitempty"`
	PendingUpdatesCount          int          `json:"pending_update_count"`
	IpAddress                    string       `json:"ip_address,omitempty"`
	LastErrorDate                int64        `json:"last_error_date,omitempty"`
	LastErrorMessage             string       `json:"last_error_message,omitempty"`
	LastSynchronizationErrorDate int64        `json:"last_synchronization_error_date,omitempty"`
	MaxConnections               int          `json:"max_connections,omitempty"`
	AllowedUpdates               []UpdateType `json:"allowed_updates,omitempty"`
}

type SetWebhookParams struct {
	URL                string       `json:"url"`
	Certificate        *InputFile   `json:"certificate,omitempty"`
	IPAddress          string       `json:"ip_address,omitempty"`
	MaxConnections     int          `json:"max_connections,omitempty"` // 1-100, default: 40
	AllowedUpdates     []UpdateType `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool         `json:"drop_pending_updates,omitempty"`
	SecretToken        string       `json:"secret_token,omitempty"` // 1-256 characters
}

// https://core.telegram.org/bots/api#setwebhook
func (bot *Bot) SetWebhook(url string, params *SetWebhookParams) (bool, error) {
	var success bool

	if params == nil {
		params = new(SetWebhookParams)
	}

	params.URL = url

	if params.Certificate != nil {
		paramsMap, err := params.ToParams()
		if err != nil {
			return success, err
		}

		files := make(Files)
		files["certificate"] = params.Certificate

		data, err := bot.RawFile(bot.stopContext, "setWebhook", paramsMap, files)
		if err != nil {
			return success, err
		}

		result, err := ParseRawResult[bool](bot, data)
		if err != nil {
			return success, err
		}

		success = *result

		return success, nil
	}

	data, err := bot.Raw(bot.stopContext, "setWebhook", params)
	if err != nil {
		return success, err
	}

	result, err := ParseRawResult[bool](bot, data)
	if err != nil {
		return success, err
	}

	success = *result

	return success, nil
}

// https://core.telegram.org/bots/api#deletewebhook
func (bot *Bot) DeleteWebhook(dropPendingUpdates bool) (bool, error) {
	var success bool
	var params Params

	if dropPendingUpdates {
		params = make(Params)
		params["drop_pending_updates"] = TrueAsString
	}

	data, err := bot.Raw(bot.stopContext, "deleteWebhook", params)
	if err != nil {
		return success, err
	}

	result, err := ParseRawResult[bool](bot, data)
	if err != nil {
		return success, err
	}

	success = *result

	return success, nil
}

// https://core.telegram.org/bots/api#getwebhookinfo
func (bot *Bot) GetWebhookInfo() (*WebhookInfo, error) {
	data, err := bot.Raw(bot.stopContext, "getWebhookInfo", nil)
	if err != nil {
		return nil, err
	}

	return ParseRawResult[WebhookInfo](bot, data)
}

func (p *SetWebhookParams) ToParams() (Params, error) {
	params := make(Params)

	if p.IPAddress != EmptyString {
		params["ip_address"] = p.IPAddress
	}

	if p.MaxConnections != 0 {
		params["max_connections"] = strconv.Itoa(p.MaxConnections)
	}

	if p.AllowedUpdates != nil {
		data, err := json.Marshal(p.AllowedUpdates)
		if err != nil {
			return nil, err
		}

		params["allowed_updates"] = string(data)
	}

	if p.DropPendingUpdates {
		params["drop_pending_updates"] = TrueAsString
	}

	if p.SecretToken != EmptyString {
		params["secret_token"] = p.SecretToken
	}

	return params, nil
}
