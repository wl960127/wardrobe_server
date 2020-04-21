package wsservice


import (
	"log"
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

// Connection .
type Connection struct {
	wsConn *websocket.Conn
	inChan chan []byte
	outChan chan []byte
	closeChan chan byte

	mutex sync.Mutex
	isClosed bool
}

// InitConnection .
func InitConnection(wsCoon *websocket.Conn)(conn *Connection,err error){
	conn = &Connection{
		wsConn:wsCoon,
		inChan:make(chan []byte,1000),
		outChan:make(chan []byte,1000),
		closeChan:make(chan byte,1),
	}

	//启动一个读协程
	go conn.readLoop()

	//启动写协程
	go conn.writeLoop()

	return
}


//ReadMessage  api使用 .
func (conn *Connection) ReadMessage() (data []byte,err error){
	select {
	case data = <- conn.inChan:
		log.Printf("读取到客户端的信息:%s", string(data));
		case<- conn.closeChan:
			err = errors.New("connection is closed")
	}

	return
}

// WriteMessage .
func (conn *Connection) WriteMessage(data []byte) (err error){

	select {
	case conn.outChan <- data:
		// log.Printf("发送到客户端的信息:%s", string(data));
		case <- conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}


// Close .
func (conn *Connection) Close(){
	//wsConn.Close()是线程安全的,可重入的close【特例】
	conn.wsConn.Close()

	//一个channel只能关闭一次，所以保证close(conn.closeChan)只执行一次【通过状态位】
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

//内部实现
func (conn *Connection) readLoop(){
	var (
		data []byte
		err error
	)

	for{
		if _,data,err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}

		//阻塞在这里，等待inchan有空闲的位置
		select {
			case  conn.inChan <- data:

				case <- conn.closeChan:
					//当closechan被关闭的时候
					goto ERR
		}

	}
	ERR:
		conn.Close()
}


func (conn *Connection) writeLoop(){
	var (
		data []byte
		err error
	)
	for{
		select {
		case data = <- conn.outChan:
			case <- conn.closeChan:
				goto ERR
		}

		if err =conn.wsConn.WriteMessage(websocket.TextMessage,data); err != nil {
			goto ERR
		}
	}
	ERR:
		conn.Close()
}