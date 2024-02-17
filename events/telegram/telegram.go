package telegram

import (
	"errors"
	"tg-bot/clients/telegram"
	"tg-bot/events"
	"tg-bot/lib/e"
)

type Processor struct {
	tg     *telegram.Client
	offset int
	//storage
}
type Meta struct {
	ChatId   int
	username string
}

var (
	errorUnknownEventType = errors.New("unknown event type")
	errorUnknownMetaType  = errors.New("unknown meta type")
)

func New(client *telegram.Client) *Processor {
	return &Processor{
		tg:     client,
		offset: 0,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("не могу получать события", err)
	}
	if len(updates) == 0 {
		return nil, nil
	}
	result := make([]events.Event, 0, len(updates))
	for _, u := range updates {
		result = append(result, event(u))
	}
	p.offset = updates[len(updates)-1].ID + 1
	return result, nil
}

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.ProcessMessage(event)
	default:
		return e.Wrap("cant process message", errorUnknownEventType)
	}
}

func (p *Processor) ProcessMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("cant process message", err)
	}

	if err := p.doCmd(event.Text, meta.ChatId, meta.username); err != nil {
		return e.Wrap("cant process message", err)
	}
	return nil

}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap("cant get meta", errorUnknownMetaType)
	}
	return res, nil
}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)

	result := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}

	if updType == events.Message {
		result.Meta = Meta{
			ChatId:   upd.Message.Chat.ID,
			username: upd.Message.From.Username,
		}
	}
	return result
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text
}

func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}
	return events.Message
}
