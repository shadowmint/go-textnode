package textnode

import (
	"golang.org/x/text/language"
)

const (
	Equal = iota
	NotEqual = iota
	LessThan = iota
	LessThanEq = iota
	GreaterThan = iota
	GreaterThanEq = iota
)

// Env is a set of Status objects
type Env struct {
	Status     *Status
	Language   *language.Tag
	Matcher    language.Matcher
	Stylesheet *StyleSheet
}

// NewEnv returns a new Env with a matcher for the given set of supported languages.
func NewEnv(status *Status, tags... language.Tag) *Env {
	return &Env{
		Status: status,
		Language: nil,
		Matcher: language.NewMatcher(tags),
		Stylesheet: NewStyleSheet()}
}

// SelectLanguage selects the best language tag based on the given locale tag,
// and return a new clone of the Env that uses that language.
// eg. "en" -> language.AmericanEnglish
func (env *Env) SelectLanguage(languageTags... string) *Env {
	userRequestedTags := make([]language.Tag, len(languageTags))
	for i := 0; i < len(languageTags); i++ {
		userRequestedTags[i] = language.Make(languageTags[i])
	}
	tag, _, _ := env.Matcher.Match(userRequestedTags...)
	return &Env{
		Status: env.Status,
		Language: &tag,
		Stylesheet: env.Stylesheet,
		Matcher: env.Matcher}
}
