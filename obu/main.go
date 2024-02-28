package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hamedblue1381/tolling/types"
)

const wsEndPoint = "ws://127.0.0.1:30000/ws"

var sendInterval = time.Second

func genLatLong() (lat float64, long float64) {
	return genCoord(), genCoord()
}
func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func main() {
	obuIDs := generateOBUIDs(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndPoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		for _, obuID := range obuIDs {
			lat, long := genLatLong()
			data := types.OBUData{
				OBUID: obuID,
				Lat:   lat,
				Long:  long,
			}
			fmt.Printf("%+v\n", data)
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(sendInterval)
	}
}

func generateOBUIDs(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Int()
	}
	return ids
}
func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
