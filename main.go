package main

import (
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/icccm"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/DanieleDaccurso/goxdo"
)

var winList []xproto.Window

var cycleCount = 0
var lastClass = ""

func attach(X *xgbutil.XUtil, m Mapping) {
	err := keybind.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			if m.Type == TYPE_SEQUENCE {
				sendStroke(X, m.Mapping, m.Value)
			}
			if m.Type == TYPE_RUN_OR_RAISE {
				runOrRaise(X, m.Class, m.Command)
			}
			if m.Type == TYPE_COMMAND {
				runCommand(m.Command)
			}
			if m.Type == TYPE_TYPE {
				sendText(X, m.Mapping, m.Text)
			}
		}).Connect(X, X.RootWin(), m.Mapping, true)
	handle(err)
}

func main() {
	X, err := xgbutil.NewConn()
	handle(err)
	keybind.Initialize(X)
	config := loadConfig(os.Args[1])
	for _, com := range config.Keys {
		attach(X, com)
	}
	for _, m := range config.Mappings {
		attach(X, m)
	}
	xevent.Main(X)
}

func runOrRaise(X *xgbutil.XUtil, matchName string, cmd string) {
	clientids, err := ewmh.ClientListStackingGet(X)
	handle(err)
	fillList := false
	if !strings.EqualFold(lastClass, matchName) {
		winList = nil
		fillList = true
		cycleCount = 0
	}
	if fillList {
		for _, clientid := range clientids {
			class, err := icccm.WmClassGet(X, clientid)
			if err != nil {
				log.Println(err)
				continue
			}
			if strings.Contains(strings.ToLower(class.Instance), strings.ToLower(matchName)) && fillList {
				winList = append(winList, clientid)
			}
		}
		winList = reverseStack(winList)
	}
	if len(winList) == 0 {
		runCommand(cmd)
	} else {
		raise(X)
	}
}

func raise(X *xgbutil.XUtil) {
	ewmh.ActiveWindowReq(X, winList[cycleCount])
	c, err := icccm.WmClassGet(X, winList[cycleCount])
	if err != nil {
		log.Println(err)
		return
	}
	lastClass = c.Instance
	cycleCount++
	if cycleCount == len(winList) {
		cycleCount = 0
	}
}

func sendStroke(X *xgbutil.XUtil, from string, to string) {
	keys := strings.Split(from, "-")
	active, err := ewmh.ActiveWindowGet(X)
	handle(err)
	xdo := goxdo.NewXdo()
	w := goxdo.Window(active)
	for _, k := range keys {
		if k == "Mod3" {
			//TODO make Mod key configurable
			xdo.SendKeysequenceWindowUp(w, "Hyper_L", 0)
		} else {
			xdo.SendKeysequenceWindowUp(w, k, 0)
		}
	}
	xdo.SendKeysequenceWindow(w, strings.Join(strings.Split(to, "-"), "+"), 0)
	xdo.SendKeysequenceWindowDown(w, "Hyper_L", 0)
}

func sendText(X *xgbutil.XUtil, from string, text string) {
	keys := strings.Split(from, "-")
	command := "xdotool "
	for _, k := range keys {
		if k == "Mod3" {
			command += "keyup Hyper_L "
		} else {
			command += "keyup " + k + " "
		}
	}
	for _, c := range strings.Split(text, "") {
		command += "key " + lookup(c) + " "
	}
	runCommand(command)
}
