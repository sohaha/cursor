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
		UserRequest  UserReq        `json:"userRequest"`
		UserMessages []*UserMessage `json:"userMessages"`
		BotMessages  []*BotMessage  `json:"botMessages"`
		ContextType  string         `json:"contextType"`
		RootPath     string         `json:"rootPath"`
	}
	BotMessage struct {
		Sender         string `json:"sender"`
		SentAt         int64  `json:"sentAt"`
		ConversationId string `json:"conversationId"`
		Type           string `json:"type"`
		Message        string `json:"message"`
		LastToken      string `json:"lastToken"`
		Finished       bool   `json:"finished"`
		CurrentFile    string `json:"currentFile"`
		Interrupted    bool   `json:"interrupted"`
	}
	UserMessage struct {
		Sender           string        `json:"sender"`
		SentAt           int64         `json:"sentAt"`
		Message          string        `json:"message"`
		ConversationId   string        `json:"conversationId"`
		OtherCodeBlocks  []interface{} `json:"otherCodeBlocks"`
		CodeSymbols      []interface{} `json:"codeSymbols"`
		Selection        Selection     `json:"selection"`
		CurrentFile      string        `json:"currentFile"`
		PrecedingCode    string        `json:"precedingCode"`
		ProcedingCode    string        `json:"procedingCode"`
		CurrentSelection string        `json:"currentSelection"`
		MsgType          string        `json:"msgType"`
	}
	Selection struct {
		From int `json:"from"`
		To   int `json:"to"`
	}
)
