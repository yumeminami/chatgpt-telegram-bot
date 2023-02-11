package chatgpt

var (
	completion_url   = "https://api.openai.com/v1/completions"
	edit_url         = "https://api.openai.com/v1/edits"
	completion_model = "text-davinci-003"
	edit_model       = "text-davinci-edit-001"
)

type CompletionRequest struct {
	Engine           string   `json:"engine"`
	Prompt           string   `json:"prompt"`
	MaxTokens        int      `json:"max_tokens"`
	Temperature      float64  `json:"temperature"`
	TopP             float64  `json:"top_p"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	PresencePenalty  float64  `json:"presence_penalty"`
	Stop             []string `json:"stop"`
}

type CompletionResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	// Choices is a list of choices, ordered from most to least likely.
	Choices []struct {
		Text  string `json:"text"`
		Index int    `json:"index"`
	} `json:"choices"`
}

type EditRequest struct {
	Engine      string  `json:"engine"`
	Input       string  `json:"input"`
	Instruction string  `json:"instruction"`
	Temperature float64 `json:"temperature"`
	TopP        float64 `json:"top_p"`
}

type EditResponse struct {
	Object  string `json:"object"`
	Created int    `json:"created"`
	Choices []struct {
		Text  string `json:"text"`
		Index int    `json:"index"`
	} `json:"choices"`
}

func GetCompletionsRequest(prompt string) CompletionRequest {
	return CompletionRequest{
		Engine:           completion_model,
		Prompt:           prompt,
		MaxTokens:        64,
		Temperature:      0.9,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
}

func GetEditRequest(input string, instruction string) EditRequest {
	return EditRequest{
		Engine:      edit_model,
		Input:       input,
		Instruction: instruction,
		Temperature: 0.9,
		TopP:        1,
	}
}
