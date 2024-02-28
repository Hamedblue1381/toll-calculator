package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hamedblue1381/tolling/types"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type DataReceiver struct {
	msg  chan types.OBUData
	conn *websocket.Conn
	prod DataProducer
}

func main() {
	drecv, err := NewDatReceiver()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starting the receiver ...")
	http.HandleFunc("/ws", drecv.handleWS)
	http.ListenAndServe(":30000", nil)
}
func NewDatReceiver() (*DataReceiver, error) {
	var (
		p          DataProducer
		err        error
		kafkaTopic = "obudata"
	)
	p, err = NewKafkaProducer(kafkaTopic)
	if err != nil {
		return nil, err
	}
	p = NewLogMiddleware(p)
	return &DataReceiver{
		msg:  make(chan types.OBUData, 1028),
		prod: p,
	}, nil
}
func (dr *DataReceiver) produceData(data types.OBUData) error {
	return dr.prod.ProduceData(data)
}
func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn
	go dr.wsReceieveLoop()
}
func (dr *DataReceiver) wsReceieveLoop() {
	fmt.Println("New OBU connected !")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error:", err)
			continue
		}
		err := dr.produceData(data)
		if err != nil {
			fmt.Println("Kafka Producer Error:", err)
		}
	}
}
