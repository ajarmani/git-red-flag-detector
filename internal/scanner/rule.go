package scanner

// Rule defines a red flag matching rule.
type Rule interface {
	ID() string
	Description() string
	Match(line string) bool
}
