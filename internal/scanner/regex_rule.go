package scanner

import "regexp"

// RegexRule implements the Rule interface using a regular expression.
type RegexRule struct {
	id          string
	description string
	pattern     *regexp.Regexp
}

func NewRegexRule(id, description, pattern string) RegexRule {
	return RegexRule{
		id:          id,
		description: description,
		pattern:     regexp.MustCompile(pattern),
	}
}

func (r RegexRule) ID() string {
	return r.id
}

func (r RegexRule) Description() string {
	return r.description
}

func (r RegexRule) Match(line string) bool {
	return r.pattern.MatchString(line)
}
