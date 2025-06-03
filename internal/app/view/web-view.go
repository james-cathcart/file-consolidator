package view

import "net/http"

type WebView struct {
	handler http.Handler
}

func NewWebView(h http.Handler) View {
	return &WebView{
		handler: h,
	}
}

func (v *WebView) ShowScanResults() {

}

func (v *WebView) ShowOptions() {

}

func (v *WebView) ShowMigrationResults() {

}
