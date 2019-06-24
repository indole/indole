package main

import (
	"net/http"
	"os"
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

func init() {
	http.HandleFunc("/act/start_command", func(w http.ResponseWriter, r *http.Request) {
		cmd := r.URL.Query().Get("cmd")
		go exec.Command("cmd", "/c", cmd).Run()
		w.WriteHeader(200)
		w.Write([]byte(ok))
	})
}

func init() {
	http.HandleFunc("/act/exec_command", func(w http.ResponseWriter, r *http.Request) {
		cmd := r.URL.Query().Get("cmd")
		err := exec.Command("cmd", "/c", cmd).Run()
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
	http.HandleFunc("/act/exit", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(0)
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
    <form action="/act/start_command" target="main">
        <input type="text" name="cmd" value="indole.exe &lt; config.xml"/>
        <button type="submit">Launch Indole (Start CMD)</button>
    </form>
    <form action="/act/exec_command" target="main">
        <input type="text" name="cmd" value="taskkill /im indole.exe /F"/>
        <button type="submit">Kill Indole (Exec CMD)</button>
    </form>
    <form action="/act/exit" target="main">
        <button type="submit">Exit</button>
    </form>
    <h1>Results</h1>
    <iframe name="main"></frame>
</body>

</html>
`
