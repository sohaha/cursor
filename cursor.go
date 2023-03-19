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
	res, err := NewRequest(http.MethodPost, convUrl, convPayload(message, language...))
	if err != nil {
		return "", err
	}

	return Parse(res.Bytes())
}

func convPayload(message string, language ...string) []byte {
	ext := ".go"
	if len(language) > 0 {
		ext = "." + language[0]
	}
	conversation := ConvReq{
		UserRequest: UserReq{
			Message:              message,
			CurrentFileName:      zstring.Rand(zstring.RandInt(3, 8), "abcdefghijklmnopqrstuvwxyz") + ext,
			CurrentRootPath:      zfile.RealPath("."),
			PrecedingCode:        []any{},
			SuffixCode:           []any{},
			CopilotCodeBlocks:    []any{},
			CustomCodeBlocks:     []any{},
			CodeBlockIdentifiers: []any{},
			MsgType:              "generate",
		},
		UserMessages: []any{},
		BotMessages:  []any{},
		ContextType:  "copilot",
		RootPath:     "",
	}
	payload, _ := json.Marshal(conversation)

	return payload
}
