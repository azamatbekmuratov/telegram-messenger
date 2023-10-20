package handler

import (
	"net/http"
	"telegram-service/domain"
	"telegram-service/service"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func SendMessage(c echo.Context) error {
	request := new(domain.MessageRequest)
	if err := c.Bind(request); err != nil {
		log.Error().Err(err).Msg("Ошибка при обработке запроса")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный запрос"})
	}

	err := service.SendMessageToTelegram(request.ChatID, request.Message)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка при отправке сообщения в Telegram")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось отправить сообщение"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Сообщение успешно отправлено"})
}
