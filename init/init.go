package init

import "github.com/colefan/gsgo/logs"

var Log *logs.Logger4g

func init() {
	Log = logs.NewLogger("gslogs", 10)
	Log.SetAppender("console", `{"level":0,"prefix":"[gs]"}`)
}
