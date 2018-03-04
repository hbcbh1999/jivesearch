package instant

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/jivesearch/jivesearch/instant/contributors"
	"golang.org/x/text/language"
)

// Minify is an instant answer
type Minify struct {
	Answer
}

func (m *Minify) setQuery(req *http.Request, q string) answerer {
	m.Answer.setQuery(req, q)
	return m
}

func (m *Minify) setUserAgent(req *http.Request) answerer {
	return m
}

func (m *Minify) setLanguage(lang language.Tag) answerer {
	m.language = lang
	return m
}

func (m *Minify) setType() answerer {
	m.Type = "minify"
	return m
}

func (m *Minify) setContributors() answerer {
	m.Contributors = contributors.Load(
		[]string{
			"brentadamson",
		},
	)
	return m
}

func (m *Minify) setRegex() answerer {
	triggers := []string{
		"minify", "minifier", "pretty", "prettifier",
	}

	t := strings.Join(triggers, "|")
	m.regex = append(m.regex, regexp.MustCompile(fmt.Sprintf(`^(?P<trigger>%s)$`, t)))
	m.regex = append(m.regex, regexp.MustCompile(fmt.Sprintf(`^(?P<trigger>%s) .*$`, t)))
	m.regex = append(m.regex, regexp.MustCompile(fmt.Sprintf(`^.* (?P<trigger>%s)$`, t)))

	return m
}

func (m *Minify) solve() answerer {
	// The caller is expected to provide the solution when triggered, preferably in JavaScript
	return m
}

func (m *Minify) setCache() answerer {
	m.Cache = true
	return m
}

func (m *Minify) tests() []test {
	typ := "minify"

	contrib := contributors.Load([]string{"brentadamson"})

	d := Data{
		Type:         typ,
		Triggered:    true,
		Contributors: contrib,
		Cache:        true,
	}

	tests := []test{
		{
			query:    "minify javascript",
			expected: []Data{d},
		},
		{
			query:    "pretty",
			expected: []Data{d},
		},
		{
			query:    `css prettifier`,
			expected: []Data{d},
		},
	}

	return tests
}
