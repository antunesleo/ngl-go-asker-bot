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
	termAsker := termasker.TermAsker{Scanner: scanner, Writer: os.Stdout}
	nglClient := nglclient.NGLClient{URL: "https://ngl.link", Writer: os.Stdout}
	askerbot.Run(os.Stdout, nglClient, termAsker)
}
