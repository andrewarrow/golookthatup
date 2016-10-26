package client

import "github.com/dontpanic92/wxGo/wx"

var ui_query wx.TextCtrl

func setupStart(f *TheFrame) {
	row := wx.NewBoxSizer(wx.HORIZONTAL)
	ui_query = wx.NewTextCtrl(f.frame, wx.ID_ANY, "", wx.DefaultPosition, wx.NewSize(80, 25), 0)
	row.Add(ui_query, 0, wx.ALL|wx.EXPAND, 5)
	f.sizer.Add(row, 0, wx.ALL|wx.EXPAND, 5)
	f.frame.Layout()
}
