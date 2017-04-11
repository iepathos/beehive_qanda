// Questions and Answers (audio responses)

// qanda models
type Question struct {
	Prompt string `json:"prompt"`
}

type AudioArgument struct {
	Username string `json:"username"`
	Audio    byte   `json:"audio"`
	Question string `json:"question"`
}

// qanda views

// get list of questions

// post audio argument
