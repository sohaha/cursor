package cursor

import (
	"encoding/json"
	"net/http"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zstring"
)

const (
	convUrl = "https://aicursor.com/conversation"
)

func Conv(message string, language ...string) (response string, err error) {
	return ConvRequest(convPayload(message, "generate", language...))
}

func ConvRequest(data *ConvReq) (response string, err error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	res, err := NewRequest(http.MethodPost, convUrl, payload)
	if err != nil {
		return "", err
	}
	return Parse(res.Bytes())
}

func convPayload(message, msgType string, language ...string) *ConvReq {
	ext := ".go"
	if len(language) > 0 {
		ext = "." + language[0]
	}

	return &ConvReq{
		UserRequest: UserReq{
			Message:              message,
			CurrentFileName:      zstring.Rand(zstring.RandInt(3, 8), "abcdefghijklmnopqrstuvwxyz") + ext,
			CurrentRootPath:      zfile.RealPath("."),
			PrecedingCode:        []any{},
			SuffixCode:           []any{},
			CopilotCodeBlocks:    []any{},
			CustomCodeBlocks:     []any{},
			CodeBlockIdentifiers: []any{},
			MsgType:              msgType,
		},
		UserMessages: []any{},
		BotMessages:  []any{},
		ContextType:  "copilot",
		RootPath:     zfile.RealPath("."),
	}
}
