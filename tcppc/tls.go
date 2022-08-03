package tcppc

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type NewConn struct {
	net.Conn
	r io.Reader
}

func (c NewConn) Read(p []byte) (int, error) {
	return c.r.Read(p)
}

func HandleTLSSession(conn *tls.Conn, writer *RotWriter, timeout int) {
	defer conn.Close()
	defer counter.dec()
	counter.inc()

	var src, dst *net.TCPAddr
	src = conn.RemoteAddr().(*net.TCPAddr)
	dst = conn.LocalAddr().(*net.TCPAddr)

	flow := NewTLSFlow(src, dst)
	session := NewSession(flow)

	log.Printf("TLS: Established: %s (#Sessions: %d)\n", session, counter.count())

	var length uint
	var err error

	buf := make([]byte, 4096)

	for {
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))

		length, err := conn.Read(buf)
		if err != nil {
			break
		}

		data := make([]byte, length)
		copy(data, buf[:length])
		if length > 10 {
			if string(data[0:5]) == "POST " || string(data[0:4]) == "GET " || string(data[0:5]) == "HEAD " || string(data[0:8]) == "OPTIONS " || string(data[0:7]) == "DELETE " || string(data[0:4]) == "PUT " || string(data[0:6]) == "TRACE " || string(data[0:8]) == "CONNECT " {
				conn.Write([]byte(data1))
				a := session.String()
				conn.Write([]byte("\n" + a))

				conn.Write([]byte("\ncontent-length: " + strconv.Itoa(len(data2))))
				conn.Write([]byte("\n\n"))
				conn.Write([]byte(data2))
				log.Printf("TLS: Send: %s\n", session)
			} else {
				//log.Printf("TLS: noSend: %s : %q\n", session, data[0:5])
			}

		} else {
			//log.Printf("TLS: short: %s : %q\n", session, data)
		}

		session.AddPayload(data)

		log.Printf("TLS: Received: %s: %q (%d bytes)\n", session, buf[:length], length)
	}

	if writer != nil {
		outputJson, err := json.Marshal(session)
		if err == nil {
			log.Printf("Wrote data: %s\n", session)
			writer.Write(outputJson)
		} else {
			log.Printf("Failed to encode data as json: %s\n", err)
		}
	}

	if length == 0 {
		log.Printf("Closed: %s (#Sessions: %d)\n", session, counter.count())
	} else {
		log.Printf("Aborted: %s %s (#Sessions: %d)\n", session, err, counter.count())
	}
}
