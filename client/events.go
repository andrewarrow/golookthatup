package client

import "github.com/dontpanic92/wxGo/wx"
import "os"

func (f *TheFrame) evtThread(e wx.Event) {
	te := wx.ToThreadEvent(e)
	switch {
	case te.GetInt() == 1:
	case te.GetInt() == 2:
	}
}

func (f *TheFrame) evtQuit(wx.Event) {
	os.Exit(0)
}

func (f *TheFrame) evtLogin(wx.Event) {
	go func() {
		f.SendEvent(1)
	}()
}
