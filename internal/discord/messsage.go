package discord

import (
	"alti-radio/common/logger"
	"alti-radio/database"
	"context"
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type MessageCommand struct {
	Command string `json:"command"`
	Message string `json:"Message"`
	Target  string `json:"target"`

	Args []string `json:"args"`
}

func (d *Discord) AddMessage() {
	d.Session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		messageCreate(s, m, d.guildId, d.channelID)
	})
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate, guildId string, channelId string) {
	var err error
	// ignore bot chat
	if m.Author.ID == s.State.User.ID {
		return
	}

	nickName := m.Member.Nick
	if nickName == "" {
		nickName, err = GetDisplayName(s, guildId, m.Author.ID)
		if err != nil {
			logger.PrintError(4, "can't get user DisplayName: "+m.Author.ID, err)
		}
	}

	// set channel for user it matched 1:1
	userChannel, _ := s.UserChannelCreate(m.Author.ID)

	commandMap := GetMessageCommand()
	command, exist := commandMap[m.Content]
	if !exist {
		_, err := s.ChannelMessageSend(channelId, "존재하지 않는 명령어에요.")
		logger.PrintLog(2, "Invalid Command: ", fmt.Sprintf("id: %s name: %s, message:%s response:%s", m.Author.ID, nickName, m.Content, "존재하지 않는 명령어에요."))
		if err != nil {
			logger.PrintError(4, "channel error: "+channelId, err)
		}
		return
	}
	var arguments []interface{}
	for _, arg := range command.Args {
		arguments = append(arguments, GetArgument(arg, nickName))
	}

	message := fmt.Sprintf(command.Message, arguments...)
	switch command.Target {
	case "all":
		_, err = s.ChannelMessageSend(channelId, message)
	case "user":
		_, err = s.ChannelMessageSend(userChannel.ID, message)
	}

	if err != nil {
		logger.PrintError(4, "Command Error: "+command.Command, err)
	}
	logger.PrintLog(2, "ChatBot Command: "+command.Command, fmt.Sprintf("id: %s name: %s, message:%s response:%s", m.Author.ID, nickName, m.Content, message))
}

func GetMessageCommand() map[string]*MessageCommand {
	messageCommands := make(map[string]*MessageCommand)

	q := database.RadioDB.GetQuery()
	commands, err := q.GetCommand(context.Background())
	if err != nil {
		logger.PrintError(4, "Database Error ", err)
	}

	for _, command := range commands {
		mc := &MessageCommand{
			Command: command.Command,
			Message: command.Message,
			Target:  command.Target,

			Args: command.Args,
		}
		messageCommands[command.Command] = mc
	}
	return messageCommands
}

func GetArgument(args string, nickName string) interface{} {
	switch args {
	case "random":
		return rand.Intn(100)
	case "displayName":
		return nickName
	case "schedule":
		schedule, err := database.RadioDB.GetQuery().GetScheduleToday(context.Background())
		if err != nil {
			return "현재 예정된 일정 이 없습니다"
		}
		return fmt.Sprintf("%s - %s \n 게스트: %s \n\n %s", schedule.RunningTime.Time.Format("2006 01 02"), schedule.Title, strings.Join(schedule.Guest, ","), schedule.Description)
	default:
		return ""
	}
}

func GetDisplayName(s *discordgo.Session, guildID string, clientId string) (string, error) {
	member, err := s.GuildMember(guildID, clientId)
	if err != nil {
		return "", fmt.Errorf("멤버 조회 실패: %v", err)
	}

	return member.DisplayName(), nil
}
