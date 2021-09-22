package live_rtmp

import (
	"log"

	. "github.com/logrusorgru/aurora"
	"github.com/qnsoft/live_sdk"
	. "github.com/qnsoft/live_utils"
)

var config = struct {
	ListenAddr string
	ChunkSize  int
}{":1935", 512}

func init() {
	live_sdk.InstallPlugin(&live_sdk.PluginConfig{
		Name:   "LiveRtmp",
		Config: &config,
		Run:    run,
	})
}
func run() {
	Print(Green("server LiveRtmp start at"), BrightBlue(config.ListenAddr))
	log.Fatal(ListenRtmp(config.ListenAddr))
}
