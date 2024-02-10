package models

// TemplateData holds data from handlers to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // Data is a map of string that points tp 'any' other data
	CSRFToken string                 //Cross Site Request Forgery Token
	Flash     string
	Warning   string
	Error     string
}
