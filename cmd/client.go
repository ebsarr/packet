package cmd

import (
	"github.com/packethost/packngo"
)

// NewPacketClient returns a *packngo.Client ready for API calls
func NewPacketClient() (*packngo.Client, error) {
	k, err := GetAPIKey()
	if err != nil {
		return nil, err
	}

	packetClient := packngo.NewClient("", k, nil)
	return packetClient, nil
}
