package lsp

type DocumentColorPresentationRequest struct {
	Request
	Params DocumentColorPresentationParams `json:"params"`
}

type DocumentColorPresentationParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Color        Color                  `json:"color"`
	Range        Range                  `json:"range"`
}

type DocumentColorPresentationResponse struct {
	Response
	Result []ColorPresentation `json:"result"`
}

type ColorPresentation struct {
	Label               string      `json:"label"`
	TextEdit            *TextEdit   `json:"textEdit,omitempty"`
	AdditionalTextEdits []*TextEdit `json:"additionalTextEdits,omitempty"`
}

type TextEdit struct {
	Range   Range  `json:"range"`
	NewText string `json:"newText"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type Color struct {
	Red   float64 `json:"red"`
	Green float64 `json:"green"`
	Blue  float64 `json:"blue"`
	Alpha float64 `json:"alpha"`
}
