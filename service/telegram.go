package service

import (
	"fmt"
	"net/http"
	"net/url"
	"telegram-service/config"

	"github.com/rs/zerolog/log"
)

const telegramAPIURL = "https://api.telegram.org/bot%s/sendMessage"

func SendMessageToTelegram(chatID, message string) error {
	apiURL := fmt.Sprintf(telegramAPIURL, config.BotToken)
	resp, err := http.PostForm(apiURL, url.Values{
		"chat_id": {chatID},
		"text":    {message},
	})
	if err != nil {
		log.Error().Err(err).Msg("Ошибка при отправке сообщения в Telegram")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errMsg := fmt.Sprintf("Не удалось отправить сообщение в Telegram, код статуса: %d", resp.StatusCode)
		log.Error().Msg(errMsg)
		return fmt.Errorf(errMsg)
	}
	return nil
}
