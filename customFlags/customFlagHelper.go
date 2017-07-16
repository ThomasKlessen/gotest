package customFlags

import (
	"fmt"
	"flag"
)

func showHelp() {
	fmt.Print(`

Usage: CLI Template [OPTIONS]

Options:
	-h, --help		Prints help
	-d, --debug		Set debugMode (default: true)
`)
}

func getHeader() {
	// Generated with http://patorjk.com/software/taag/#p=display&f=Isometric1&t=Insitu
	// for the Lolz
	fmt.Println(`
                  ___           ___                       ___           ___
      ___        /\__\         /\  \          ___        /\  \         /\__\
     /\  \      /::|  |       /::\  \        /\  \       \:\  \       /:/  /
     \:\  \    /:|:|  |      /:/\ \  \       \:\  \       \:\  \     /:/  /
     /::\__\  /:/|:|  |__   _\:\~\ \  \      /::\__\      /::\  \   /:/  /  ___
  __/:/\/__/ /:/ |:| /\__\ /\ \:\ \ \__\  __/:/\/__/     /:/\:\__\ /:/__/  /\__\
 /\/:/  /    \/__|:|/:/  / \:\ \:\ \/__/ /\/:/  /       /:/  \/__/ \:\  \ /:/  /
 \::/__/         |:/:/  /   \:\ \:\__\   \::/__/       /:/  /       \:\  /:/  /
  \:\__\         |::/  /     \:\/:/  /    \:\__\       \/__/         \:\/:/  /
   \/__/         /:/  /       \::/  /      \/__/                      \::/  /
                 \/__/         \/__/                                   \/__/    `)
}

func setFlag(flag *flag.FlagSet) {
	flag.Usage = func() {
		getHeader()
		showHelp()
	}
}