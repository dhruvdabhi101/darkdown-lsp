package analysis

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(document, content string) {
	s.Documents[document] = content
}
func (s *State) UpdateDocument(uri, content string) {
	s.Documents[uri] = content
}
