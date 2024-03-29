package analysis

import "strings"

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
func (s *State) GetLine(uri string, line int) string {
  document := s.Documents[uri]
  // return the line from the document
  documentLines := strings.Split(document, "\n")
  return documentLines[line]
}
