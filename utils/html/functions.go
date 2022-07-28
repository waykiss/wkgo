package utilshtml

import (
	"bytes"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/svg"
	utilstrings "goapp/utils/strings"
	"regexp"
	"strings"
)

func MinifyHtml(v string) string {
	m := minify.New()
	m.AddFunc("image/svg+html", svg.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]html$"), html.Minify)
	r := strings.NewReader(v)
	buf := &bytes.Buffer{}
	err := m.Minify("text/html", buf, r)
	if err != nil {
		return "Erro ao minificar HTML"
	}
	return utilstrings.RemoveExtraSpaces(buf.String())
}
