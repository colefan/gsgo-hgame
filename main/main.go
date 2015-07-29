package main

import "github.com/colefan/gsgo-hgame/gs"

func main() {

	mnt := gs.NewGameManager()
	if mnt.Init("") {
		mnt.Run()
	}
	mnt.ShutDown()
}
