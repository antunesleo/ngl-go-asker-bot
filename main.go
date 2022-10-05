package main

import (
	"bufio"
	"os"

	"github.com/antunesleo/ngl-go-asker-bot/askerbot"
	"github.com/antunesleo/ngl-go-asker-bot/nglclient"
	"github.com/antunesleo/ngl-go-asker-bot/termasker"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	termAsker := termasker.TermAsker{Scanner: scanner}
	nglClient := nglclient.NGLClient{URL: "https://ngl.link"}
	askerbot.Run(nglClient, termAsker)
}
