/*
@Time : 2020/3/26 13:56
@Author : ZhouHui2
@File : main
@Software: GoLand
*/
package main

import (
	"bufio"
	"flag"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var (
	addr    = flag.String("addr", "10.0.108.164:8080", "http service address")
	cmdPath string
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 8192

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Time to wait before force close on connection.
	closeGracePeriod = 10 * time.Second
)

func pumpStdin(ws *websocket.Conn, w io.Writer) {
	defer ws.Close()
	ws.SetReadLimit(maxMessageSize)
	_ = ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		message = append(message, '\n')
		if _, err := w.Write(message); err != nil {
			break
		}
	}
}

func pumpStdout(ws *websocket.Conn, r io.Reader, done chan struct{}) {
	defer func() {
	}()
	s := bufio.NewScanner(r)
	for s.Scan() {
		_ = ws.SetWriteDeadline(time.Now().Add(writeWait))
		if err := ws.WriteMessage(websocket.TextMessage, s.Bytes()); err != nil {
			_ = ws.Close()
			break
		}
	}
	if s.Err() != nil {
		log.Println("scan:", s.Err())
	}
	close(done)

	_ = ws.SetWriteDeadline(time.Now().Add(writeWait))
	_ = ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	time.Sleep(closeGracePeriod)
	_ = ws.Close()
}

func internalError(ws *websocket.Conn, msg string, err error) {
	log.Println(msg, err)
	_ = ws.WriteMessage(websocket.TextMessage, []byte("Internal server error."))
}

var upgrader = websocket.Upgrader{}

func ping(ws *websocket.Conn, done chan struct{}) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
				log.Println("ping:", err)
			}
		case <-done:
			return
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer func() {
		_ = ws.Close()
	}()

	outr, outw, err := os.Pipe()
	if err != nil {
		internalError(ws, "stdout:", err)
		return
	}
	defer func() {
		_ = outr.Close()
	}()
	defer func() {
		_ = outw.Close()
	}()

	inr, inw, err := os.Pipe()
	if err != nil {
		internalError(ws, "stdin:", err)
		return
	}
	defer func() {
		_ = inr.Close()
	}()
	defer func() {
		_ = inw.Close()
	}()

	proc, err := os.StartProcess(cmdPath, flag.Args(), &os.ProcAttr{
		Files: []*os.File{inr, outw, outw},
	})
	if err != nil {
		internalError(ws, "start:", err)
		return
	}

	_ = inr.Close()
	_ = outw.Close()

	stdoutDone := make(chan struct{})
	go pumpStdout(ws, outr, stdoutDone)
	go ping(ws, stdoutDone)

	pumpStdin(ws, inw)

	// Some commands will exit when stdin is closed.
	_ = inw.Close()

	// Other commands need a bonk on the head.
	if err := proc.Signal(os.Interrupt); err != nil {
		log.Println("inter:", err)
	}

	select {
	case <-stdoutDone:
	case <-time.After(time.Second):
		// A bigger bonk on the head.
		if err := proc.Signal(os.Kill); err != nil {
			log.Println("term:", err)
		}
		<-stdoutDone
	}

	if _, err := proc.Wait(); err != nil {
		log.Println("wait:", err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("must specify at least one argument")
	}
	var err error
	cmdPath, err = exec.LookPath(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
