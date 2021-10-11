package main

// JSON Utility - Alfred Help and Support (https://www.alfredapp.com/help/workflows/utilities/json/)
type AlfredJSON struct {
	Workflow *AlfredWorkflow `json:"alfredworkflow"`
}

type AlfredWorkflow struct {
	Variables struct {
		Link    string `json:"link"`
		Browser string `json:"browser"`
	} `json:"variables"`
}

func NewAlfredJSON(link string, browser string) *AlfredJSON {
	workflow := &AlfredWorkflow{}
	workflow.Variables.Link = link
	workflow.Variables.Browser = browser
	return &AlfredJSON{
		Workflow: workflow,
	}
}
