package aquagram

import (
	"context"
	"encoding/json"
	"fmt"
)

type PollingUpdater struct {
	Bot          *Bot
	LastUpdateID int
}

func NewPollingUpdater(bot *Bot) *PollingUpdater {
	updater := new(PollingUpdater)
	updater.Bot = bot

	return updater
}

func (updater *PollingUpdater) Start(ctx context.Context) {
	updater.LastUpdateID = -1

	for {
		select {
		case <-updater.Bot.stopContext.Done():
			break
		default:
		}

		params := GetUpdatesParams{}
		params.Offset = updater.LastUpdateID + 1

		updates, err := updater.Bot.GetUpdates(ctx, params)
		if err != nil {
			updater.Bot.Logger.Println(fmt.Errorf("%w: %w", ErrUpdaterError, err))
			continue
		}

		for _, update := range updates {
			updater.LastUpdateID = update.UpdateID
			updater.Bot.DispatchUpdate(update)
		}
	}
}

type Updates struct {
	Result []*Update `json:"result"`
}

type GetUpdatesParams struct {
	Offset         int      `json:"offset"`
	Limit          int      `json:"limit"`
	Timeout        int      `json:"timeout"`
	AllowedUpdates []string `json:"allowed_updates"`
}

func (bot *Bot) GetUpdates(ctx context.Context, params GetUpdatesParams) ([]*Update, error) {
	data, err := bot.Raw(ctx, "getUpdates", params)
	if err != nil {
		return nil, err
	}

	updates := new(Updates)

	if err := json.Unmarshal(data, updates); err != nil {
		return nil, err
	}

	return updates.Result, nil
}
