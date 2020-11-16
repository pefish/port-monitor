package command

import (
	"flag"
	"fmt"
	go_config "github.com/pefish/go-config"
	go_logger "github.com/pefish/go-logger"
	"github.com/pefish/telegram-bot-manager/pkg/telegram-sender"
	"net"
	"time"
)

type DefaultCommand struct {
}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{}
}

func (dc *DefaultCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	return nil
}

func (dc *DefaultCommand) OnExited() error {
	return nil
}

func (dc *DefaultCommand) Start() error {
	ports, err := go_config.Config.GetSlice("ports")
	if err != nil {
		return err
	}
	interval, err := go_config.Config.GetUint64("interval")
	if err != nil {
		return err
	}
	sendInterval, err := go_config.Config.GetUint64("sendInterval")
	if err != nil {
		return err
	}
	token, err := go_config.Config.GetString("/telegram/token")
	if err != nil {
		return err
	}
	chatId, err := go_config.Config.GetInt64("/telegram/chatId")
	if err != nil {
		return err
	}
	telegramSender := telegram_sender.NewTelegramSender(token)
	telegramSender.SetLogger(go_logger.Logger)
	timer := time.NewTimer(0)
	for {
		<-timer.C
		go_logger.Logger.Info("next check...")
		for _, p := range ports {
			pMap := p.(map[interface{}]interface{})
			host := pMap["host"].(string)
			port := pMap["port"].(int)
			alertMsg := pMap["alertMsg"].(string)
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
			if err != nil {
				go_logger.Logger.WarnF("check error. host: %s, port: %d, err: %s", host, port, err.Error())
				telegramSender.SendMsg(telegram_sender.MsgStruct{
					ChatId: chatId,
					Msg:    []byte(alertMsg),
				}, time.Duration(sendInterval)*time.Second)
				continue
			}
			conn.Close()
		}
		timer.Reset(time.Duration(interval) * time.Second)
	}
}
