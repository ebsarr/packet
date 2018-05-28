package cmd

import (
	"github.com/ebsarr/packngo"

	"github.com/ebsarr/packet/extpackngo"
)

// NewPacketClient returns a *packngo.Client ready for API calls
func NewPacketClient() (*packngo.Client, error) {
	k, err := GetAPIKey()
	if err != nil {
		return nil, err
	}

	packetClient := packngo.NewClientWithAuth("", k, nil)
	return packetClient, nil
}

// NewExtPacketClient returns a *extpackngo.Client ready for API calls
func NewExtPacketClient() (*extpackngo.Client, error) {
	k, err := GetAPIKey()
	if err != nil {
		return nil, err
	}

	packetClient := extpackngo.NewClient("", k, nil)
	return packetClient, nil
}
