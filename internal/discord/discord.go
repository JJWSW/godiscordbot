package discord

import (
	"alti-radio/common/logger"

	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	Session         *discordgo.Session
	Message         *discordgo.Message
	VoiceConnection *discordgo.VoiceConnection

	token     string
	channelID string
	guildId   string
}


func New(token string) *Discord {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.PrintError(3, "Invalid Token", err)
	}
	s.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers | discordgo.IntentsGuildMessages

	err = s.Open()
	if err != nil {
		logger.PrintError(3, "Discord Open Error", err)
	}

	logger.PrintLog(2, "Session Opened by Bot Token: ", token)
	return &Discord{
		Session: s,

		token:   token,
		guildId: "",
	}
}

func (d *Discord) SetChannelId(channelId string) *Discord {
	d.channelID = channelId
	return d
}

func (d *Discord) addHandler(handler interface{}) {
	d.Session.AddHandler(handler)
}

func (d *Discord) SetSession(session *discordgo.Session) *Discord {
	d.Session = session
	return d
}

func (d *Discord) SetMessage(message *discordgo.Message) *Discord {
	d.Message = message
	return d
}

func (d *Discord) SetVoiceConnection(voiceConnection *discordgo.VoiceConnection) *Discord {
	d.VoiceConnection = voiceConnection
	return d
}
