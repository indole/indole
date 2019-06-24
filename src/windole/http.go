package main

import (
	"net/http"
	"os/exec"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(html))
	})
}

func init() {
	http.HandleFunc("/act/open_system_proxy", func(w http.ResponseWriter, r *http.Request) {
		err := exec.Command("reg", "add", "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings", "/v", "ProxyEnable", "/t", "REG_DWORD", "/d", "1", "/f").Run()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(200)
		w.Write([]byte(ok))
	})
}

func init() {
	http.HandleFunc("/act/close_system_proxy", func(w http.ResponseWriter, r *http.Request) {
		err := exec.Command("reg", "add", "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings", "/v", "ProxyEnable", "/t", "REG_DWORD", "/d", "0", "/f").Run()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(200)
		w.Write([]byte(ok))
	})
}

const ok = "OK"

const html = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Indole Console</title>
</head>

<body>
    <h1>Commands</h1>
    <form action="/act/open_system_proxy" target="main">
        <button type="submit">Open System Proxy</button>
    </form>
    <form action="/act/close_system_proxy" target="main">
        <button type="submit">Close System Proxy</button>
    </form>
    <h1>Results</h1>
    <iframe name="main"></frame>
</body>

</html>
`
