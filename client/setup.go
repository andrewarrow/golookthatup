package client

import "github.com/dontpanic92/wxGo/wx"

var ui_query wx.TextCtrl

func setupStart(f *TheFrame) {
	row := wx.NewBoxSizer(wx.HORIZONTAL)
	ui_query = wx.NewTextCtrl(f.frame, wx.ID_ANY, "", wx.DefaultPosition, wx.NewSize(180, 25), 0)
	row.Add(ui_query, 0, wx.ALL|wx.EXPAND, 5)
	f.sizer.Add(row, 0, wx.ALL|wx.EXPAND, 5)
	//wx.Bind(f.frame, wx.EVT_TEXT, f.evtTextChange, ui_query.GetId())
	go f.checkQueryEvery()
	f.frame.Layout()
}
