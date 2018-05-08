package iso8583SDK

import (
	"fmt"
	"net"
	"time"
)

func send(reqMsg []byte, config *Config) ([]byte, error) {
	conn, err := net.Dial("tcp", config.Host)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Duration(config.TimeOut) * time.Second))
	_, err = conn.Write(reqMsg)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, 64)
	respMsg := make([]byte, 0)
	totalLen := 0
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return nil, err
		}
		if len(respMsg) == 0 {
			if n < 2 {
				return nil, fmt.Errorf("the response length is less than 2 !")
			}
			totalLen = (int(buf[0]) << 8) + int(buf[1]) + 2
			fmt.Println(totalLen)
		}
		respMsg = append(respMsg, buf[:n]...)
		if len(respMsg) >= totalLen {
			break
		}
	}
	return respMsg, nil
}
