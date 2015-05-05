package main

import ( 
  "fmt"
  "net"
  "log"
  "time"
)

type ZigbeeDriver interface {
  ZclSetOnOff(newValue int, zNode ZigbeeNode) error
  ZclGetOnOff() bool
  ToggleOnOff()
  BroadcastState() int
  GotoLightLevel(newValue int, zNode ZigbeeNode) error
}

type ZigbeePacket struct {
  msgType byte
  seq byte
  cmd [2]byte
  msgLen byte
  payload []byte
  netAddr [2]byte
  longAddr [8]byte
}

type ZigbeeNode struct {
  transitionTime int
  hardwareAddr int
}

type LightBulbGELink struct {
  isOn bool
  lightLevel int
  ZigbeeNode
  ZigbeePacket
}

func (l LightBulbGELink) ZclInitDefaultValues() {
  l.isOn = false
  l.lightLevel = 0xff
  l.ZigbeeNode.hardwareAddr = 0x0000
  l.ZigbeeNode.transitionTime = 1
}

func (l LightBulbGELink) ZclSetHardwareAddress(addr int) {
  l.ZigbeeNode.hardwareAddr = addr
}

func (l LightBulbGELink) ZclSetOnOff(newValue int, zNode ZigbeeNode) int {
  if newValue == 1 {
    l.isOn = true
  } else {
    l.isOn = false
  }
  // TODO: Broadcast
  return l.BroadcastState()
}

func (l LightBulbGELink) ZclGetOnOff() bool {
  // Broadcast to get value from the GE Link bulb directly
  return l.isOn
}

func (l LightBulbGELink) ToggleOnOff() int {
  l.isOn = !l.isOn
  return l.BroadcastState()
}

func (l LightBulbGELink) GotoLightLevel(newValue int, zNode ZigbeeNode) {
  l.lightLevel = newValue
  l.transitionTime = zNode.transitionTime
  l.hardwareAddr = zNode.hardwareAddr
  l.BroadcastState()
}

func (l LightBulbGELink) BroadcastState() int {
  // Hardcoding for testing
  // buf := new(bytes.Buffer)
  // binary.Write(buf, binary.LittleEndian, d)
  msgType := "\x43"
  seq := "\x01"
  cmd := "\x04\x00"
  msgLen := "\x0f"
  payload := "\x00"
  netAddr := "\xde\xdf"
  longAddr := "\x7c\xe5\x24\x00\x00\x0b\x6a\xf0"
  fucking := "\x01\x01\x00\x00"
  hardcore := msgType + seq + cmd + msgLen + payload + netAddr + longAddr + fucking
  fmt.Printf("% x"+ " \n", hardcore)

  // Listen on TCP port 1234
  urlport := "192.168.2.177:1234"
  fmt.Println("Writing to %s", urlport)
  addr, _ := net.ResolveTCPAddr("tcp", urlport)
  conn, err := net.DialTCP("tcp", nil, addr)
  if err != nil {
      log.Fatal(err)
  }
  defer conn.Close()
  err = conn.SetNoDelay(true)
  if err != nil {
      log.Fatal(err)
  }
  // Wait for a connection.
  conn.Write([]byte(hardcore))
  fmt.Println("Wrote to pipeline")
  time.Sleep(30 * time.Second)
  return 1
}

func main() {
  thisBulb := LightBulbGELink{}
  thisBulb.BroadcastState()
}
