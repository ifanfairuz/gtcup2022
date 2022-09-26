package images

import (
	"io"
	"strconv"
	"strings"
	"time"

	svg "github.com/ajstarks/svgo"
	"github.com/ifanfairuz/gtcup2022/repositories/match"
	"github.com/ifanfairuz/gtcup2022/support"
)

var QuarterDimension = Dimesion{W: 800, H: 1132}

func genDateQuarter(s *svg.SVG, date time.Time) {
	tgl := strings.ToUpper(WEEKDAY_ID[date.Weekday()]+date.Format(" | 02 ")+MONTH_ID[date.Month()]+date.Format(" 2006 | 15:04 WIB"))
	s.Text(QuarterDimension.MidW(), 760, tgl, mergeStyles("font-size=\"30pt\"", "ff-montserrat", "fw-extrabold", "ta-middle", "fill-white")...)
}

func genMatchQuarter(s *svg.SVG, m match.Match) {
	home := strings.ToUpper(m.TeamHome.Name)
	away := strings.ToUpper(m.TeamAway.Name)
	midW := QuarterDimension.MidW()
	p := 30
	rC := 75
	xH := midW-rC
	xH = p+((xH-p)/2)
	xA := midW+rC
	xA = QuarterDimension.W-p-((QuarterDimension.W-p-xA)/2)
	y := 847
	s.Text(xH, y, home, mergeStyles("font-size=\"35pt\"", "ff-montserrat", "fw-extrabold", "ta-middle", "fill-white")...)
	s.Text(xA, y, away, mergeStyles("font-size=\"35pt\"", "ff-montserrat", "fw-extrabold", "ta-middle", "fill-white")...)
}

func genMatchesQuarter(s *svg.SVG, mm []match.Match) {
	y := 910
	x := 30
	anchor := "ta-start"
	for i, m := range mm {
		if i == 2 {
			anchor = "ta-end"
			x = QuarterDimension.W - 30
			y = 910
		}
		labels := strings.Split(m.Label, "|")
		home := strings.ToUpper(m.TeamHome.Name)
		away := strings.ToUpper(m.TeamAway.Name)
		s.Text(x, y, strings.ToUpper(labels[0]), mergeStyles("font-size=\"18pt\"", "ff-montserrat", "fw-extrabold", anchor, "fill-green")...)
		y += 25
		if m.TeamHomeId <= 0 {
			home = labels[1]
		}
		if m.TeamAwayId <= 0 {
			away = labels[2]
		}
		var wH, wA int
		if m.Done {
			for _, s := range m.Sets {
				if s.Winner == m.TeamHomeId {
					wH++
				} else if s.Winner == m.TeamAwayId {
					wA++
				}
			}
		}
		xx := x
		if anchor == "ta-start" {
			xx += 180
		} else {
			xx -= 325
		}
		s.Text(x, y, strings.ToUpper(home), mergeStyles("font-size=\"15pt\"", "ff-montserrat", "fw-regular", anchor, "fill-white")...)
		if m.Done {
			xxx := xx
			s.Text(xxx, y, "("+strconv.Itoa(wH)+")", mergeStyles("font-size=\"15pt\"", "ff-montserrat", "fw-regular", "ta-middle", "fill-white")...)
			for i := 0; i < 5; i++ {
				sc := "0"
				if len(m.Sets) > i {
					sc = strconv.Itoa(m.Sets[i].Home)
				}
				xxx += 30
				s.Text(xxx, y, sc, mergeStyles("font-size=\"15pt\"", "ff-montserrat", "fw-regular", "ta-middle", "fill-white")...)
			}
		}
		y += 20
		s.Text(x, y, strings.ToUpper(away), mergeStyles("font-size=\"15pt\"", "ff-montserrat", "fw-regular", anchor, "fill-white")...)
		if m.Done {
			xxx := xx
			s.Text(xxx, y, "("+strconv.Itoa(wA)+")", mergeStyles("font-size=\"15pt\"", "ff-montserrat", "fw-regular", "ta-middle", "fill-white")...)
			for i := 0; i < 5; i++ {
				sc := "0"
				if len(m.Sets) > i {
					sc = strconv.Itoa(m.Sets[i].Away)
				}
				xxx += 30
				s.Text(xxx, y, sc, mergeStyles("font-size=\"15pt\"", "ff-montserrat", "fw-regular", "ta-middle", "fill-white")...)
			}
		}
		y += 40
	}
}

func genFooterQuarter(s *svg.SVG) {
}

func GenSVGQuarter(writer io.Writer, m match.Match, mm []match.Match) {
	wh := QuarterDimension.WString()+" "+QuarterDimension.HString()
	s := svg.New(writer)
	s.Start(QuarterDimension.W, QuarterDimension.H, "viewBox=\"0 0 "+ wh +"\" enable-background=\"new 0 0 "+ wh +"\" xml:space=\"preserve\"")
	s.Style("text/css", Css("montserrat"))
	s.Image(0, 0, QuarterDimension.W, QuarterDimension.H, B_BG)
	s.Group("stroke:none;")
	genDateQuarter(s, m.Date.In(support.JAKARTA_TZ))
	genMatchQuarter(s, m)
	genMatchesQuarter(s, mm)
	genFooterQuarter(s)
	s.Gend()
	s.End()
}