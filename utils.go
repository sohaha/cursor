package cursor

type (
	UserReq struct {
		Message              string `json:"message"`
		CurrentRootPath      string `json:"currentRootPath"`
		CurrentFileName      string `json:"currentFileName"`
		CurrentFileContents  string `json:"currentFileContents"`
		PrecedingCode        []any  `json:"precedingCode"`
		CurrentSelection     any    `json:"currentSelection"`
		SuffixCode           []any  `json:"suffixCode"`
		CopilotCodeBlocks    []any  `json:"copilotCodeBlocks"`
		CustomCodeBlocks     []any  `json:"customCodeBlocks"`
		CodeBlockIdentifiers []any  `json:"codeBlockIdentifiers"`
		MsgType              string `json:"msgType"`
		MaxOrigLine          any    `json:"maxOrigLine"`
	}
	ConvReq struct {
		UserRequest  UserReq `json:"userRequest"`
		UserMessages []any   `json:"userMessages"`
		BotMessages  []any   `json:"botMessages"`
		ContextType  string  `json:"contextType"`
		RootPath     string  `json:"rootPath"`
	}
)
