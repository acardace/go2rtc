package debug

import (
	"github.com/AlexxIT/go2rtc/cmd/api"
	"github.com/AlexxIT/go2rtc/cmd/streams"
	"github.com/AlexxIT/go2rtc/pkg/core"
)

func Init() {
	api.HandleFunc("api/stack", stackHandler)

	streams.HandleFunc("null", nullHandler)
}

func nullHandler(string) (core.Producer, error) {
	return nil, nil
}
