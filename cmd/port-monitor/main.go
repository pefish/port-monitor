package main

import (
	"github.com/pefish/go-commander"
	"github.com/pefish/port-monitor/cmd/port-monitor/command"
	"github.com/pefish/port-monitor/version"
	"log"
)

func main() {
	commanderInstance := commander.NewCommander(version.AppName, version.Version, version.AppName+" 是一个模板，祝你玩得开心。作者：pefish")
	//commanderInstance.RegisterSubcommand("client", client.NewClient())
	//commanderInstance.RegisterSubcommand("server", server.NewServer())
	commanderInstance.RegisterDefaultSubcommand(command.NewDefaultCommand())
	err := commanderInstance.Run()
	if err != nil {
		log.Fatal(err)
	}
}
