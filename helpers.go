package tg

import (
	"encoding/json"
	"net/url"
)

// Convert struct to url values map
// TODO: temp implementation
func structToValues(obj interface{}) (url.Values, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	rawMap := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &rawMap); err != nil {
		return nil, err
	}

	values := url.Values{}
	for key := range rawMap {
		values.Set(key, string(rawMap[key]))
	}

	return values, nil
}

type sendMessageParams struct {
	SendMessageOptions
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type sendPhotoParams struct {
	ChatID int64  `json:"chat_id"`
	Photo  string `json:"photo,omitempty"`
	SendPhotoOptions
}

func newSendPhotoParams(chatID int64, photo string, options *SendPhotoOptions) *sendPhotoParams {
	params := &sendPhotoParams{
		ChatID: chatID,
		Photo:  photo,
	}

	if options != nil {
		params.SendPhotoOptions = *options
	}

	return params
}

type sendAudioParams struct {
	ChatID int64  `json:"chat_id"`
	Audio  string `json:"audio,omitempty"`
	SendAudioOptions
}

func newSendAudioParams(chatID int64, audio string, options *SendAudioOptions) *sendAudioParams {
	params := &sendAudioParams{
		ChatID: chatID,
		Audio:  audio,
	}

	if options != nil {
		params.SendAudioOptions = *options
	}

	return params
}

type sendDocumentParams struct {
	ChatID   int64  `json:"chat_id"`
	Document string `json:"document,omitempty"`
	SendDocumentOptions
}

func newSendDocumentParams(chatID int64, document string, options *SendDocumentOptions) *sendDocumentParams {
	params := &sendDocumentParams{
		ChatID:   chatID,
		Document: document,
	}

	if options != nil {
		params.SendDocumentOptions = *options
	}

	return params
}

type sendStickerParams struct {
	ChatID  int64  `json:"chat_id"`
	Sticker string `json:"sticker,omitempty"`
	SendStickerOptions
}

func newSendStickerParams(chatID int64, sticker string, options *SendStickerOptions) *sendStickerParams {
	params := &sendStickerParams{
		ChatID:  chatID,
		Sticker: sticker,
	}

	if options != nil {
		params.SendStickerOptions = *options
	}

	return params
}

type sendVideoParams struct {
	ChatID int64  `json:"chat_id"`
	Video  string `json:"video,omitempty"`
	SendVideoOptions
}

func newSendVideoParams(chatID int64, video string, options *SendVideoOptions) *sendVideoParams {
	params := &sendVideoParams{
		ChatID: chatID,
		Video:  video,
	}

	if options != nil {
		params.SendVideoOptions = *options
	}

	return params
}

type sendVoiceParams struct {
	ChatID int64  `json:"chat_id"`
	Voice  string `json:"voice,omitempty"`
	SendVoiceOptions
}

func newSendVoiceParams(chatID int64, voice string, options *SendVoiceOptions) *sendVoiceParams {
	params := &sendVoiceParams{
		ChatID: chatID,
		Voice:  voice,
	}

	if options != nil {
		params.SendVoiceOptions = *options
	}

	return params
}

type sendLocationParams struct {
	ChatID    int64   `json:"chat_id"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	SendLocationOptions
}

func newSendLocationParams(chatID int64, latitude, longitude float64, options *SendLocationOptions) *sendLocationParams {
	params := &sendLocationParams{
		ChatID:    chatID,
		Latitude:  latitude,
		Longitude: longitude,
	}

	if options != nil {
		params.SendLocationOptions = *options
	}

	return params
}

type sendVenueParams struct {
	ChatID    int64   `json:"chat_id"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Title     string  `json:"title,omitempty"`
	Address   string  `json:"address,omitempty"`
	SendVenueOptions
}

func newSendVenueParams(chatID int64, latitude, longitude float64, title, address string, options *SendVenueOptions) *sendVenueParams {
	params := &sendVenueParams{
		ChatID:    chatID,
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}

	if options != nil {
		params.SendVenueOptions = *options
	}

	return params
}

type sendContactParams struct {
	ChatID      int64  `json:"chat_id"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	SendContactOptions
}

func newSendContactParams(chatID int64, phoneNumber, firstName, lastName string, options *SendContactOptions) *sendContactParams {
	params := &sendContactParams{
		ChatID:      chatID,
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
		LastName:    lastName,
	}

	if options != nil {
		params.SendContactOptions = *options
	}

	return params
}

type sendGameParams struct {
	ChatID        int64  `json:"chat_id"`
	GameShortName string `json:"game_short_name"`
	SendGameOptions
}

type setGameScoreParams struct {
	UserID int64 `json:"chat_id"`
	Score  int   `json:"score"`
	SetGameScoreOptions
}

type getGameHighScoresParams struct {
	UserID int64 `json:"chat_id"`
	GetGameHighScoresOptions
}

type answerCallbackQueryParams struct {
	CallbackQueryID string `json:"callback_query_id"`
	AnswerCallbackQueryOptions
}

func newAnswerCallbackQueryParams(callbackQueryID string, options *AnswerCallbackQueryOptions) *answerCallbackQueryParams {
	params := &answerCallbackQueryParams{
		CallbackQueryID: callbackQueryID,
	}

	if options != nil {
		params.AnswerCallbackQueryOptions = *options
	}

	return params
}

type editMessageTextParams struct {
	ChatID          int64  `json:"chat_id,omitempty"`
	MessageID       int64  `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
	Text            string `json:"text"`
	EditMessageTextOptions
}

type editMessageCationParams struct {
	ChatID          int64  `json:"chat_id,omitempty"`
	MessageID       int64  `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
	EditMessageCationOptions
}

type editMessageReplyMarkupParams struct {
	ChatID          int64       `json:"chat_id,omitempty"`
	MessageID       int64       `json:"message_id,omitempty"`
	InlineMessageID string      `json:"inline_message_id,omitempty"`
	ReplyMarkup     ReplyMarkup `json:"reply_markup,omitempty"`
}

type answerInlineQueryParams struct {
	InlineQueryID string             `json:"inline_query_id"`
	Results       InlineQueryResults `json:"results"`
	AnswerInlineQueryOptions
}
