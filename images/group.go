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

var GrupDimension = Dimesion{W: 1000, H: 1415}

func genTitleGroup(s *svg.SVG, x int, y int) (int, int) {
	s.Text(x, y, "QUALIFICATION", mergeStyles("font-size=\"84px\"", "ff-montserrat", "fw-black")...)
	y += 85
	s.Text(x, y, "MATCH", mergeStyles("font-size=\"110px\"", "ff-montserrat", "fw-black")...)
	y += 70
	return x, y
}

func genDateGroup(s *svg.SVG, x int, y int, date time.Time) (int, int) {
	px := 8
	py := 10
	tgl := strings.ToUpper(date.Format("02 ")+MONTH_ID[date.Month()-1]+date.Format(" 2006"))
	s.Rect(x, y-30-py, WEEKDAY_W[date.Weekday()]+(px*2), 30+(py*2), genStyles("fill-blue")...)
	s.Text(x+px, y, strings.ToUpper(WEEKDAY_ID[date.Weekday()]), mergeStyles("font-size=\"42px\"", "ff-montserrat", "fw-semibold", "fill-white")...)
	s.Text(x, y+35, tgl, mergeStyles("font-size=\"24px\"", "ff-montserrat", "fw-bold")...)
	s.Text(x+400, y-py, date.Format("15:04"), mergeStyles("font-size=\"42px\"", "ff-montserrat", "fw-extrabold", "fill-blue", "ta-end")...)
	s.Text(x+400, y-py+35, "WIB", mergeStyles("font-size=\"42px\"", "ff-montserrat", "fw-extrabold", "fill-blue", "ta-end")...)
	y += 100
	return x, y
}

func genMatchGroup(s *svg.SVG, x int, y int, mm []match.Match) (int, int) {
	px := 7
	py := 5
	ovr := 0
	for i, m := range mm {
		istoday := m.Date.In(support.JAKARTA_TZ).Format("2006-01-02") == time.Now().In(support.JAKARTA_TZ).Format("2006-01-02")
		title := "SEASON "+strconv.Itoa(i+1)
		if !istoday {
			title = "MATCH #"+strconv.Itoa(i+1)
			ovr = 20
		}
		s.Rect(x, y-18-py, 125+ovr+(px*2), 18+(py*2), genStyles("fill-blue")...)
		s.Text(x+px, y, title, mergeStyles("font-size=\"24px\"", "ff-montserrat", "fw-semibold", "fill-white")...)
		y += 50

		home := strings.ToUpper(m.TeamHome.Name)
		away := strings.ToUpper(m.TeamAway.Name)
		if m.Done {
			var wH, wA int
			for _, s := range m.Sets {
				if s.Winner == m.TeamHomeId {
					wH++
				} else if s.Winner == m.TeamAwayId {
					wA++
				}
			}
			home += " ("+ strconv.Itoa(wH) +")"
			away += " ("+ strconv.Itoa(wA) +")"
		}
		s.Text(x, y, home, mergeStyles("font-size=\"42px\"", "ff-montserrat", "fw-black")...)
		y += 50
		s.Text(x, y, away, mergeStyles("font-size=\"42px\"", "ff-montserrat", "fw-black")...)
		y += 80
	}
	return x, y
}

func genKlasemenGroup(s *svg.SVG, x int, y int, g string, kl []match.GrupKlasemen) (int, int) {
	c := []struct{name string;x int;Anchor string;get func (k match.GrupKlasemen) string}{
		{"NO", 0, "start", func (k match.GrupKlasemen) string {
			return strconv.Itoa(k.Pos)
		}},
		{"TEAM", 40, "start", func (k match.GrupKlasemen) string {
			return strings.ToUpper(k.Team.Name)
		}},
		{"P", 180, "middle", func (k match.GrupKlasemen) string {
			return strconv.Itoa(k.P)
		}},
		{"M", 30, "middle", func (k match.GrupKlasemen) string {
			return strconv.Itoa(k.M)
		}},
		{"K", 30, "middle", func (k match.GrupKlasemen) string {
			return strconv.Itoa(k.K)
		}},
		{"AS", 40, "middle", func (k match.GrupKlasemen) string {
			return strconv.Itoa(k.AS)
		}},
		{"A-SC", 50, "middle", func (k match.GrupKlasemen) string {
			return strconv.Itoa(k.ASC)
		}},
		{"POIN", 60, "middle", func (k match.GrupKlasemen) string {
			return strconv.Itoa(k.Poin)
		}},
	}
	s.Text(x, y, "KLASEMEN GRUP "+g, mergeStyles("font-size=\"32px\"", "ff-montserrat", "fw-bold")...)
	y += 30
	xc := x
	for _, v := range c {
		xc += v.x
		s.Text(xc, y, v.name, mergeStyles("font-size=\"20px\"", "ff-montserrat", "fw-bold", "ta-"+v.Anchor)...)
	}
	y += 30
	for _, k := range kl {
		xc := x
		for _, v := range c {
			xc += v.x
			s.Text(xc, y, v.get(k), mergeStyles("font-size=\"20px\"", "ff-montserrat", "fw-semibold", "ta-"+v.Anchor)...)
		}
		y += 25
	}
	return x, y
}

func genFooterGroup(s *svg.SVG, x int, y int) (int, int) {
	yy := GrupDimension.H-350
	s.Text(x, yy, "LAPANGAN KITERAN", mergeStyles("font-size=\"32px\"", "ff-montserrat", "fw-bold")...)
	yy += 20
	s.Text(x, yy, "RT.05, RW.04 LINGKUNGAN GENENG TIMUR", mergeStyles("font-size=\"18px\"", "ff-montserrat", "fw-medium")...)
	yy += 20
	s.Text(x, yy, "KELURAHAN LEDUG, PRIGEN - PASURUAN", mergeStyles("font-size=\"18px\"", "ff-montserrat", "fw-medium")...)
	yy += 20
	genQr(s, x, yy)
	yy += 230
	s.Text(x, yy, "OFFICIAL WEBSITE", mergeStyles("font-size=\"24px\"", "ff-montserrat", "fw-bold")...)
	yy += 20
	s.Text(x, yy, "https://gtcup2022.herokuapp.com", mergeStyles("font-size=\"18px\"", "ff-montserrat", "fw-medium")...)

	return x, yy
}

func GenSVGGroup(writer io.Writer, m []match.Match, k []match.GrupKlasemen) {
	var (
		x = 50
		y = 250
	)
	
	wh := GrupDimension.WString()+" "+GrupDimension.HString()
	s := svg.New(writer)
	s.Start(GrupDimension.W, GrupDimension.H, "viewBox=\"0 0 "+ wh +"\" enable-background=\"new 0 0 "+ wh +"\" xml:space=\"preserve\"")
	s.Style("text/css", Css("montserrat"))
	s.Image(0, 0, GrupDimension.W, GrupDimension.H, G_BG)
	s.Group("stroke:none;")
	x, y = genTitleGroup(s, x, y)
	x, y = genDateGroup(s, x, y, m[0].Date.In(support.JAKARTA_TZ))
	x, y = genMatchGroup(s, x, y, m)
	genKlasemenGroup(s, x, y, m[0].Group, k)
	genFooterGroup(s, x, y)
	s.Gend()
	s.End()
}