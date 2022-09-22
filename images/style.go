package images

import "strings"

var css = map[string]string{
	"montserrat": "@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500;600;700;800;900');",
}

var styles = map[string]string{
	"ff-montserrat": "font-family=\"Montserrat, sans-serif\" shape-rendering=\"geometricPrecision\"",
	"fw-regular": "font-weight=\"400\"",
	"fw-medium": "font-weight=\"500\"",
	"fw-semibold": "font-weight=\"600\"",
	"fw-bold": "font-weight=\"700\"",
	"fw-extrabold": "font-weight=\"800\"",
	"fw-black": "font-weight=\"900\"",
	"fill-white": "fill=\"#ffffff\"",
	"fill-blue": "fill=\"#025577\"",
	"fill-green": "fill=\"#3ab54b\"",
	"stroke-white": "stroke=\"#ffffff\"",
	"stroke-blue": "stroke=\"#025577\"",
	"ta-start": "text-anchor=\"start\"",
	"ta-middle": "text-anchor=\"middle\"",
	"ta-end": "text-anchor=\"end\"",
}

func Css(keys ...string) string {
	s := []string{}
	for _, v := range keys {
		s = append(s, css[v])
	}

	return strings.Join(s, " ")
}

func genStyles(keys ...string) []string {
	s := []string{}
	for _, v := range keys {
		s = append(s, styles[v])
	}

	return s
}

func mergeStyles(fix string, keys ...string) []string {
	s := []string{fix}
	for _, v := range keys {
		s = append(s, styles[v])
	}

	return s
}