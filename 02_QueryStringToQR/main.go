package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")
var templ = template.Must(template.New("qr").Parse(tplString))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	log.Printf("Listening: localhost:8080 ")
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("q"))
}

const tplString = `
<html>
	<head>
	<title>QR Link Generator</title>
	</head>
	<body>
		{{if .}}
		<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
		<br>
		{{.}}
		<br>
		<br>
		{{end}}
		<form action="/" name="f" method="GET">
			<input maxLength="1024" size="50" name="q" value="" title="Text to QR Encode">
			<input type="submit" value="Generate QR" name="qr">
		</form>
	</body>
</html>
`
