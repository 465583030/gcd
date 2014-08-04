package gcd

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io"
	"log"
	"net"
	"sync/atomic"
)

type TargetInfo struct {
	Description          string `json:"description"`
	DevtoolsFrontendUrl  string `json:"devtoolsFrontendUrl"`
	FaviconUrl           string `json:"faviconUrl"`
	Id                   string `json:"id"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	Url                  string `json:"url"`
	WebSocketDebuggerUrl string `json:"webSocketDebuggerUrl"`
}

type gcdMessage struct {
	ReplyCh chan string // json response channel
	Id      int64       // id to map response channels to send chans
	Data    []byte      // the data to send out over the websocket channel
	Type    string      // event name type.
}

type ChromeTarget struct {
	sync.RWMutex
	replyDispatcher map[int64]chan *gcdMessage
	conn            *websocket.Conn
	console         *ChromeConsole
	input           *ChromeInput
	network         *ChromeNetwork
	Target          *TargetInfo
	sendCh          chan *gcdMessage
	doneCh          chan bool
	sendId          int64
	/*
		debugger    *ChromeDebugger
		dom         *ChromeDom
		domDebugger *ChromeDomDebugger
		page        *ChromePage
		runtime     *ChromeRuntime
		timeline    *ChromeTimeline
	*/
}

func newChromeTarget(port string, target *TargetInfo) *ChromeTarget {
	conn := wsConnection(fmt.Sprintf("localhost:%s", port), target.WebSocketDebuggerUrl)
	replier := make(map[int64]chan *gcdMessage)
	sendCh := make(chan *gcdMessage)
	doneCh := make(chan bool)
	chromeTarget := &ChromeTarget{conn: conn, console: nil, Target: target, sendCh: sendCh, replyDispatcher: replier, doneCh: doneCh, sendId: 0}
	chromeTarget.listen()
	return chromeTarget
}

func (c *ChromeTarget) getId() int64 {
	fmt.Printf("sendId: %d\n", c.sendId)
	return atomic.AddInt64(&c.sendId, 1)
}

func (c *ChromeTarget) listen() {
	go c.listenRead()
	go c.listenWrite()
}

func (c *ChromeTarget) listenWrite() {
	log.Println("Listening write for client")
	for {
		select {
		// send message to the client
		case msg := <-c.sendCh:
			log.Println("Send:", string(msg))
			websocket.Message.Send(c.conn, string(msg))
		// receive done from listenRead
		case <-c.doneCh:
			fmt.Println("listenWrite doneCh true, returning...")
			c.doneCh <- true // for listenRead method
			return
		}
	}
}

// Listen read request via chanel
func (c *ChromeTarget) listenRead() {
	log.Println("Listening read from client")
	for {
		select {
		// receive done from listenWrite
		case <-c.doneCh:
			fmt.Println("listenRead doneCh true, returning")
			c.doneCh <- true
			return
		// read data from websocket connection
		default:
			var msg string
			err := websocket.Message.Receive(c.conn, &msg)
			if err == io.EOF {
				fmt.Println("EOF, sending doneCh")
				c.doneCh <- true
				return
			} else if err != nil {
				log.Printf("error in websocket read: %v\n", err)
				c.doneCh <- true
			} else {
				c.parseResponse(msg)
			}
		}
	}
}

func (c *ChromeTarget) parseResponse(msg string) {
	log.Printf("parseResponse got msg: %s\n", msg)
}

func wsConnection(addr, url string) *websocket.Conn {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	config, errConfig := websocket.NewConfig(url, "http://localhost")
	if errConfig != nil {
		log.Fatalf("error building config: %v\n", err)
	}
	client, errWS := websocket.NewClient(config, conn)
	if errWS != nil {
		log.Fatalf("WebSocket handshake error: %v\n", errWS)
	}
	return client
}