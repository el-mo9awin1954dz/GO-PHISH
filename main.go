package main

import (
	"os"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-webkit/webkit"
)
// YOU CAN SETU A PHISHING HTTP(S) SERVER YOU CAN ALSO ADD ICON AS A FAKE BROWSER 

const HTML_STRING = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>LOGIN</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <div class="container">
        <img src="rocket.png" alt="rocket" class="rocket">
        <div class="text">
            <h1>LOGIN</h1>
            <p>LANDING PAGE</p>
        </div>
        <form  class="form">
            <div class="animated-input">
                <input type="text" placeholder="UUSERNAME">
                <input type="password" placeholder="PASSWORD">
            </div>
            <div class="check">
                <div>
                    <input type="checkbox" id="check">
                    <label for="check" class="disc"></label>
                    <label for="check" class="remember">REMEMBER</label>
                </div>
                <p class="forget"><a href="#">Forget Password ?</a></p>
            </div>
        </form>
        <button class="btn" type="submit">LOGIN</button>
        <p class="account"><a href="#">CREATE ACCOUNT ?</a></p>
    </div>
</body>
</html>

`

const IFRAME_EMBED = `
<style> *{ margin : 0; padding : 0; } </style>
<iframe width="500" height="500" src="https://www.youtube.com/embed/BHt8APB3E0I" title="DAK - Tacchini (Officiel Music vidéo) (Explicite) Prod By @KersBeats&amp; @houssemmerzouga" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>
`

func main() {
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("PHISHING REDTEAM")
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.NewVBox(false, 1)

	entry := gtk.NewEntry()
	url := os.Args[1]
	entry.SetText(url)
	vbox.PackStart(entry, false, false, 0)

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)

	webview := webkit.NewWebView()
	webview.Connect("load-committed", func() {
		entry.SetText(webview.GetUri())
	})
	swin.Add(webview)

	vbox.Add(swin)

	entry.Connect("activate", func() {
		webview.LoadUri(entry.GetText())
	})
	button := gtk.NewButtonWithLabel("WHO")
	button.Clicked(func() {
		webview.LoadString("PHISHING REDTEAM", "text/plain", "utf-8", ".")
	})
	vbox.PackStart(button, false, false, 0)

	button = gtk.NewButtonWithLabel("LEAD LOGIN")
	button.Clicked(func() {
		webview.LoadHtmlString(HTML_STRING, ".")
	})
	vbox.PackStart(button, false, false, 0)

	button = gtk.NewButtonWithLabel("MUSIC")
	button.Clicked(func() {
		webview.LoadHtmlString(MAP_EMBED, ".")
	})
	vbox.PackStart(button, false, false, 0)

	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()

	proxy := os.Getenv("HTTP_PROXY")
	if len(proxy) > 0 {
		soup_uri := webkit.SoupUri(proxy)
		webkit.GetDefaultSession().Set("proxy-uri", soup_uri)
		soup_uri.Free()
	}
	entry.Emit("activate")
	gtk.Main()
}
