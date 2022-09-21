package images

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	svg "github.com/ajstarks/svgo"
	"github.com/ifanfairuz/gtcup2022/repositories/match"
	"github.com/ifanfairuz/gtcup2022/support"
)

const FONT_FAMILY = "font-family:'Montserrat', sans-serif"
var MONTH_ID = []string{"","Januari","Februari","Maret","April","Mei","Juni","Juli","September","Oktober","November","Desember"}
var WEEKDAY_ID = []string{"Minggu","Senin","Selasa","Rabu","Kamis","Jum'at","Sabtu"}
var WEEKDAY_W = []int{185,133,170,127,140,160,145}
var (
	w = 1000
	h = 1415
)

func genTitleGroup(s *svg.SVG, x int, y int, date time.Time) (int, int) {
	titles := []string{"TODAY'S", "MATCHES"}
	istoday := date.In(support.JAKARTA_TZ).Format("2006-01-02") == time.Now().In(support.JAKARTA_TZ).Format("2006-01-02")
	if !istoday {
		titles[0] = "PENYISIHAN"
		titles[1] = "GRUP"
	}
	s.Text(x, y, titles[0], "font-size:100px;font-weight:900;"+FONT_FAMILY)
	y += 85
	s.Text(x, y, titles[1], "font-size:110px;font-weight:900;"+FONT_FAMILY)
	y += 70
	return x, y
}
func genTitleBracket(s *svg.SVG, x int, y int, title string) (int, int) {
	t := strings.SplitN(strings.ToUpper(title), " ", 2)
	s.Text(x, y, t[0], "font-size:100px;font-weight:900;"+FONT_FAMILY)
	y += 80
	s.Text(x, y, t[1], "font-size:100px;font-weight:900;"+FONT_FAMILY)
	y += 60
	return x, y
}

func genDate(s *svg.SVG, x int, y int, date time.Time) (int, int) {
	px := 8
	py := 10
	tgl := strings.ToUpper(date.Format("02 ")+MONTH_ID[date.Month()]+date.Format(" 2006"))
	s.Rect(x, y-30-py, WEEKDAY_W[date.Weekday()]+(px*2), 30+(py*2), "fill:#025577")
	s.Text(x+px, y, strings.ToUpper(WEEKDAY_ID[date.Weekday()]), "font-size:42px;font-weight:600;fill:#fff;"+FONT_FAMILY)
	s.Text(x, y+35, tgl, "font-size:24px;font-weight:700;"+FONT_FAMILY)
	s.Text(x+400, y-py, date.Format("15:04"), "font-size:42px;font-weight:800;fill:#025577;text-anchor:end;"+FONT_FAMILY)
	s.Text(x+400, y-py+35, "WIB", "font-size:42px;font-weight:800;fill:#025577;text-anchor:end;"+FONT_FAMILY)
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
		s.Rect(x, y-18-py, 125+ovr+(px*2), 18+(py*2), "fill:#025577")
		s.Text(x+px, y, title, "font-size:24px;font-weight:600;fill:#fff;"+FONT_FAMILY)
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
		s.Text(x, y, home, "font-size:42px;font-weight:900;"+FONT_FAMILY)
		y += 40
		s.Text(x, y, away, "font-size:42px;font-weight:900;"+FONT_FAMILY)
		y += 80
	}
	return x, y
}

func genMatchBracket(s *svg.SVG, x int, y int, m match.Match) (int, int) {
	y += 80
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
	s.Text(x+280, y, home, "font-size:62px;font-weight:900;text-anchor:middle;"+FONT_FAMILY)
	y += 87
	s.Text(x+280, y, "VS", "font-size:72px;font-weight:900;text-anchor:middle;"+FONT_FAMILY)
	y += 80
	s.Text(x+280, y, away, "font-size:62px;font-weight:900;text-anchor:middle;"+FONT_FAMILY)
	y += 80
	return x, y
}

func genKlasemen(s *svg.SVG, x int, y int, g string, kl []match.GrupKlasemen) (int, int) {
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
	s.Text(x, y, "KLASEMEN GRUP "+g, "font-size:32px;font-weight:700;"+FONT_FAMILY)
	y += 30
	xc := x
	for _, v := range c {
		xc += v.x
		s.Text(xc, y, v.name, "font-size:20px;font-weight:700;text-anchor:"+v.Anchor+";"+FONT_FAMILY)
	}
	y += 30
	for _, k := range kl {
		xc := x
		for _, v := range c {
			xc += v.x
			s.Text(xc, y, v.get(k), "font-size:20px;font-weight:600;text-anchor:"+v.Anchor+";"+FONT_FAMILY)
		}
		y += 25
	}
	return x, y
}

func GenSVG(writer io.Writer, m []match.Match, k []match.GrupKlasemen) {
	var (
		x = 50
		y = 250
	)
	
	s := svg.New(writer)
	s.Start(w, h, "viewBox=\"0 0 1000 1415\" enable-background=\"new 0 0 1000 1415\" xml:space=\"preserve\"")
	s.Style("text/css", "@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@400;600;700;800;900&display=swap');")
	s.Image(0, 0, w, h, BG)
	s.Group("stroke:none;")
	if m[0].Type == "B" {
		x, y = genTitleBracket(s, x, y, m[0].Group)
		y += 20
	} else {
		x, y = genTitleGroup(s, x, y, m[0].Date)
	}
	x, y = genDate(s, x, y, m[0].Date)
	if m[0].Type == "B" {
		genMatchBracket(s, x, y, m[0])
	} else {
		x, y = genMatchGroup(s, x, y, m)
		genKlasemen(s, x, y, m[0].Group, k)
	}
	y = h-350
	s.Text(x, y, "LAPANGAN KITERAN", "font-size:32px;font-weight:700;"+FONT_FAMILY)
	y += 20
	s.Text(x, y, "RT.05, RW.04 LINGKUNGAN GENENG TIMUR", "font-size:18px;font-weight:600;"+FONT_FAMILY)
	y += 20
	s.Text(x, y, "KELURAHAN LEDUG, PRIGEN - PASURUAN", "font-size:18px;font-weight:600;"+FONT_FAMILY)
	y += 20
	genQr(s, x, y)
	y += 230
	s.Text(x, y, "OFFICIAL WEBSITE", "font-size:24px;font-weight:700;"+FONT_FAMILY)
	y += 20
	s.Text(x, y, "https://gtcup2022.herokuapp.com", "font-size:18px;font-weight:600;"+FONT_FAMILY)
	s.Gend()
	s.End()
}

func convertToImage(svgfile string, resfile string) error {
	in, err := filepath.Abs(svgfile)
	if err != nil {
		log.Println(err)
		return err
	}
	out, err := filepath.Abs(resfile)
	if err != nil {
		log.Println(err)
		return err
	}

	cmd := exec.Command("convert", in, out)
	e := cmd.Run()
	if e != nil {
		log.Println(e)
		return e
	}

	return nil
}

func RemoveOldImage(m []match.Match)  {
	for _, m2 := range m {
		p := path.Join("public", "assets", "match", m2.Image)
		if _, e := os.Stat(p); e == nil {
			os.Remove(p)
		}
	}
}

func GenImage(m []match.Match, k []match.GrupKlasemen) (string, error) {
	svgname := m[0].Date.Format("2006-01-02")+".svg"
	resname := strconv.FormatInt(time.Now().UnixMilli(), 10)+".jpg"
	svgfile := path.Join("images", "result", svgname)
	resfile := path.Join("public", "assets", "match", resname)

	if _, e := os.Stat(svgfile); e == nil {
		os.Remove(svgfile)
	}

	f, err := os.Create(svgfile)
	if err != nil {
		log.Println(err)
		return resname, err
	}
	defer f.Close()
	GenSVG(f, m, k)
	
	e := convertToImage(svgfile, resfile)
	if e != nil {
		log.Println(e)
		return resname, e
	}

	return resname, nil
}