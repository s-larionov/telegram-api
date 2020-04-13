package helpers

import (
	"github.com/s-larionov/telegram-api/models"
)

func DistributeKeyboardRows(buttons []models.KeyboardButton, columns int) [][]models.KeyboardButton {
	var keyboard [][]models.KeyboardButton

	for i := 0; i < len(buttons); i += columns {
		if (i + 2) >= len(buttons) {
			keyboard = append(keyboard, buttons[i:])
		} else {
			keyboard = append(keyboard, buttons[i:i+columns])
		}
	}

	return keyboard
}

func DistributeInlineKeyboardRows(buttons []models.InlineKeyboardButton, columns int) [][]models.InlineKeyboardButton {
	var keyboard [][]models.InlineKeyboardButton

	for i := 0; i < len(buttons); i += columns {
		if (i + 2) >= len(buttons) {
			keyboard = append(keyboard, buttons[i:])
		} else {
			keyboard = append(keyboard, buttons[i:i+columns])
		}
	}

	return keyboard
}
