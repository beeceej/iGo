package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/nsf/termbox-go"
)

type repl struct {
	state replState
}

type replState struct {
	x       int
	y       int
	currMsg string
	prevMsg string
}

func (r *repl) PrintHeader() {
	r.state.currMsg = `
/##################################\
# Welcome to iGo                   #
# A go interpreter written in go   #
\##################################/
`
	r.state = r.tbprintstate(termbox.ColorCyan, termbox.ColorDefault)
	termbox.Flush()
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	initialX := x
	for _, c := range msg {
		if c == '\n' {
			y++
			x = initialX
			termbox.SetCell(x, y, c, fg, bg)
			continue
		}
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func (r repl) tbprintstate(fg, bg termbox.Attribute) replState {
	xBegin := r.state.x
	// yBegin := r.state.y // Not needed just yet
	xCurr := r.state.x
	yCurr := r.state.y

	if xCurr == 0 {
		termbox.SetCell(xCurr, yCurr, '>', fg, bg)
		xCurr += 2
	}

	for _, c := range r.state.currMsg {
		if c == '\n' {
			yCurr++
			xCurr = xBegin
			termbox.SetCell(xCurr, yCurr, c, fg, bg)
			continue
		}

		termbox.SetCell(xCurr, yCurr, c, fg, bg)
		xCurr++

	}

	return replState{
		x:       xCurr,
		y:       yCurr,
		prevMsg: r.state.currMsg,
		currMsg: "",
	}
}

var current string
var currentEvent termbox.Event

func mouseButtonStr(k termbox.Key) string {
	switch k {
	case termbox.MouseLeft:
		return "MouseLeft"
	case termbox.MouseMiddle:
		return "MouseMiddle"
	case termbox.MouseRight:
		return "MouseRight"
	case termbox.MouseRelease:
		return "MouseRelease"
	case termbox.MouseWheelUp:
		return "MouseWheelUp"
	case termbox.MouseWheelDown:
		return "MouseWheelDown"
	}
	return "Key"
}

func modStr(m termbox.Modifier) string {
	var out []string
	if m&termbox.ModAlt != 0 {
		out = append(out, "ModAlt")
	}
	if m&termbox.ModMotion != 0 {
		out = append(out, "ModMotion")
	}
	return strings.Join(out, " | ")
}

func (r *repl) redrawAll() {
	const coldef = termbox.ColorDefault

	switch currentEvent.Type {
	case termbox.EventKey:
		tbprint(0, 0, coldef, coldef,
			fmt.Sprintf("EventKey: k: %d, c: %c, mod: %s", currentEvent.Key, currentEvent.Ch, modStr(currentEvent.Mod)))
		if currentEvent.Key == termbox.KeyEnter {
			r.state = replState{
				x:       0,
				y:       r.state.y + 1,
				prevMsg: r.state.currMsg,
				currMsg: "",
			}
			r.state = r.tbprintstate(coldef, coldef)
			t := map[string]string{
				"raw": current,
			}
			b, err := json.Marshal(t)
			if err != nil {
				tbprint(0, 0, coldef, coldef, err.Error())
				break
			}
			tbprint(0, 0, coldef, coldef, string(b))

			req, _ := http.NewRequest("POST", "http://localhost:9999/interpret", bytes.NewReader(b))
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				tbprint(0, 0, coldef, coldef, err.Error())
				current = ""
				break
			}
			defer resp.Body.Close()
			b, _ = ioutil.ReadAll(resp.Body)
			var res map[string]string
			json.Unmarshal(b, &res)
			r.state = replState{
				x:       0,
				y:       r.state.y,
				prevMsg: r.state.currMsg,
				currMsg: fmt.Sprintf("%s", res["raw"]),
			}
			r.state = r.tbprintstate(coldef, coldef)
			current = ""
			break
		}
		if currentEvent.Key == termbox.KeyBackspace || currentEvent.Key == termbox.KeyBackspace2 {
			tbprint(0, 0, coldef, coldef,
				fmt.Sprintf("HIT BACKSPACE: EventKey: k: %d, c: %c, mod: %s", currentEvent.Key, currentEvent.Ch, modStr(currentEvent.Mod)))
			r.state = replState{
				x:       r.state.x - 1,
				y:       r.state.y,
				prevMsg: r.state.currMsg,
				currMsg: current[:len(current)-2],
			}
			r.state = r.tbprintstate(coldef, coldef)
			break
		}

		r.state = replState{
			x:       0,
			y:       r.state.y,
			prevMsg: r.state.currMsg,
			currMsg: fmt.Sprintf("%s", current),
		}
		r.state = r.tbprintstate(coldef, coldef)

	case termbox.EventNone:
		tbprint(0, 6, coldef, coldef, "EventNone")
	}
	termbox.Flush()
}

func main() {
	err := termbox.Init()
	// fmt.Println(err.Error())
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
	r := &repl{
		state: replState{
			x: 0, y: 1, currMsg: "", prevMsg: "",
		},
	}

	r.PrintHeader()
	r.redrawAll()
	data := make([]byte, 0, 64)
mainloop:
	for {
		if cap(data)-len(data) < 32 {
			newdata := make([]byte, len(data), len(data)+32)
			copy(newdata, data)
			data = newdata
		}
		beg := len(data)
		d := data[beg : beg+32]
		switch ev := termbox.PollRawEvent(d); ev.Type {
		case termbox.EventRaw:
			if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
				data = make([]byte, 0, 64)
				current = string(data)
				for {
					ev := termbox.ParseEvent(data)
					if ev.N == 0 {
						break
					}
					currentEvent = ev
					copy(data, data[currentEvent.N:])
					data = data[:len(data)-currentEvent.N]
				}
			} else {
				data = data[:beg+ev.N]
				current += string(data)
				if current == "quit()" {
					break mainloop
				}
				for {
					ev := termbox.ParseEvent(data)
					if ev.N == 0 {
						break
					}
					currentEvent = ev
					copy(data, data[currentEvent.N:])
					data = data[:len(data)-currentEvent.N]
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		r.redrawAll()
	}
}
