package main

import (
	"bufio"
	"os"

	"github.com/antunesleo/ngl-go-asker-bot/askerbot"
	"github.com/antunesleo/ngl-go-asker-bot/dataproviders"
	"github.com/antunesleo/ngl-go-asker-bot/nglclient"
	"github.com/antunesleo/ngl-go-asker-bot/termasker"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	termAsker := termasker.New(scanner, os.Stdout)
	dataProvider := dataproviders.CreateDataProvider("./", termAsker)
	nglClient := nglclient.New("https://ngl.link", os.Stdout)
	askerbot.Run(os.Stdout, nglClient, dataProvider)
}
