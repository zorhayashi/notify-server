package cmd

import (
	"fmt"

	"github.com/zorhayashi/notify-server/config"
	util "github.com/zorhayashi/notify-server/util"
)

var (
	version = "0.0.2"
)

//Print init
func init() {
	xMsg := util.GetLineMsg()
	logo := fmt.Sprintf(` 	     _   _  __                                          
 _ __   ___ | |_(_)/ _|_   _       ___  ___ _ ____   _____ _ __ 
| '_ \ / _ \| __| | |_| | | |_____/ __|/ _ \ '__\ \ / / _ \ '__|
| | | | (_) | |_| |  _| |_| |_____\__ \  __/ |   \ V /  __/ |   
|_| |_|\___/ \__|_|_|  \__, |     |___/\___|_|    \_/ \___|_|   
             	        |___/               - version : %s -
----------------------------------------------------------------`, version)
	util.Logo(logo)
	fmt.Println(xMsg)
	fmt.Println("----------------------------------------------------------------")
	config.Configinit()
	switch config.Global.Systeam.Admin {
	case true:
		util.Info("admin mode")
		Admin()
	case false:
		util.Info("cmd mode")
		util.Info("Server Port : " + config.Global.Systeam.Port)
	}
}
