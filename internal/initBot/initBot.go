package initbot

import (
	tg "github.com/OvyFlash/telegram-bot-api"

	"log"
	"strconv"
)

func New(token string, adminsChatIDstr string) (int64, tg.UpdatesChannel, *tg.BotAPI) {
  adminsChatIDi, err := strconv.Atoi(adminsChatIDstr)
  if err != nil {
    log.Fatal("can't get admins chat id")
  }
  adminsChatID := int64(adminsChatIDi)

  bot, err := tg.NewBotAPI(token)
  if err != nil {
    log.Fatal("cant create bot")
  }

  log.Printf("авторизован под %s", bot.Self.UserName)

  updateConfig := tg.NewUpdate(0)
  updateConfig.Timeout = 60

  updates := bot.GetUpdatesChan(updateConfig)
  
  return adminsChatID, updates, bot
}

func CreateVars(u tg.Update) (chatID, usrID, msgID int64, msgText, usrName string, msgPhoto []tg.PhotoSize, msgVideo *tg.Video, msgVoice *tg.Voice, msgVideoNote *tg.VideoNote) {
  chatID = u.Message.Chat.ID
  msgText = u.Message.Text
  usrName = u.SentFrom().UserName
  usrID = u.SentFrom().ID
  msgPhoto = u.Message.Photo
  msgID = int64(u.Message.MessageID)
  msgVideo = u.Message.Video
  msgVoice = u.Message.Voice
  msgVideoNote = u.Message.VideoNote

  return chatID, usrID, msgID, msgText, usrName, msgPhoto, msgVideo, msgVoice, msgVideoNote
}
