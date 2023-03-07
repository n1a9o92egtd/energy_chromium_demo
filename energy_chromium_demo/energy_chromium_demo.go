// https://github.com/energye/energy
// https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/634d3bb5a749ba0318943eb4
// https://energy.yanghy.cn/#/example/6342d986401bfe4d0cdf6067/634d3bd5a749ba0318943eb6
package main

import (
	"github.com/energye/energy/cef"
)

func main() {
	//Global initialization must be called by every application
	cef.GlobalInit(nil, nil)
	//Create application
	cefApp := cef.NewApplication(nil)
	//Set URL
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	cef.BrowserWindow.Config.Icon = "resources/icon.ico"
	cef.BrowserWindow.Config.Title = "energy tutorial"
	config := cef.BrowserWindow.Config.ChromiumConfig()
	config.SetEnableDevTools(false)
	config.SetEnableMenu(false)
	config.SetEnableOpenUrlTab(false)
	config.SetEnableViewSource(false)
	config.SetEnableWindowPopup(false)
	//Run App
	cef.Run(cefApp)
}
