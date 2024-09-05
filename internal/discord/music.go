package discord

import (
	"alti-radio/common/logger"
	"os"
	"time"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

type VoiceController struct {
	VoiceConnection *discordgo.VoiceConnection

	stop chan bool
}

func (d *Discord) JoinChannel() *VoiceController {
	vc, err := d.Session.ChannelVoiceJoin(d.guildId, d.channelID, false, true)
	if err != nil {
		logger.PrintError(4, "Join Voice Channel Error", err)
	}
	d.VoiceConnection = vc
	return &VoiceController{
		VoiceConnection: vc,

		stop: make(chan bool),
	}
}

func (voiceController *VoiceController) Play(filename string) {
	_, err := os.Stat(filename)
	if err != nil {
		logger.PrintError(4, "file is not extist: "+filename, err)
		return
	}
	dgvoice.PlayAudioFile(voiceController.VoiceConnection, filename, voiceController.stop)

	time.Sleep(time.Second * 30)
	voiceController.Stop()
}

func (voiceController *VoiceController) Stop() {
	<-voiceController.stop
}
