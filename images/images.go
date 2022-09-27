package images

import (
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ifanfairuz/gtcup2022/repositories/match"
)

type Dimesion struct {
	W int
	H int
}

func (d *Dimesion) WString() string {
	return strconv.Itoa(d.W)
}
func (d *Dimesion) HString() string {
	return strconv.Itoa(d.H)
}
func (d *Dimesion) MidW() int {
	return d.W/2
}
func (d *Dimesion) MidH() int {
	return d.H/2
}

var WEEKDAY_ID = [...]string{"Minggu","Senin","Selasa","Rabu","Kamis","Jum'at","Sabtu"}
var WEEKDAY_W = [...]int{185,133,170,127,140,160,145}
var MONTH_ID = [...]string{"","Januari","Februari","Maret","April","Mei","Juni","Juli","September","Oktober","November","Desember"}

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
		if m2.Image != "" {
			p := path.Join("public", "assets", "match", m2.Image)
			if _, e := os.Stat(p); e == nil {
				os.Remove(p)
			}
		}
	}
}

func GenImage(m []match.Match, k []match.GrupKlasemen, mm []match.Match) (string, error) {
	svgname := m[0].Date.Format("2006-01-02")+".svg"
	resname := strconv.FormatInt(time.Now().UnixNano(), 10)+".jpg"
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
	sm := m[0]
	if sm.Type == "G" {
		GenSVGGroup(f, m, k)
	} else if sm.Type == "B" {
		if sm.Round == 1 {
			GenSVGQuarter(f, m[0], mm)
		}
	}
	
	e := convertToImage(svgfile, resfile)
	if e != nil {
		log.Println(e)
		return resname, e
	}

	return resname, nil
}