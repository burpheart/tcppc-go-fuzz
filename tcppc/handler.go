package tcppc

import (
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net"
	"syscall"
	"time"
)

func StartFuzzServer(host string, port int, config *tls.Config, writer *RotWriter, timeout int) {
	log.Printf("Server Mode: TLS\n")
	log.Printf("Listen: %s:%d\n", host, port)

	addr := &net.TCPAddr{
		IP:   net.ParseIP(host),
		Port: port,
	}

	tcpLn, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen TCP socket: %s\n", err)
	}
	defer tcpLn.Close()

	file, err := tcpLn.File()
	if err != nil {
		log.Fatalf("Failed to get a file descriptor of the listener: %s\n", err)
	}
	defer file.Close()

	fd := int(file.Fd())
	if err := syscall.SetsockoptInt(fd, syscall.SOL_IP, syscall.IP_TRANSPARENT, 1); err != nil {
		log.Fatalf("Failed to set socket option (IP_TRANSPARENT): %s\n", err)
	}
	if err := syscall.SetsockoptInt(fd, syscall.SOL_IP, syscall.IP_RECVORIGDSTADDR, 1); err != nil {
		log.Fatalf("Failed to set socket option (IP_RECVORIGDSTADDR): %s\n", err)
	}
	//ln := tls.NewListener(tcpLn, config)
	log.Printf("Start TLS server.\n")

	for {
		//conn, err := ln.Accept()
		conn, err := tcpLn.AcceptTCP()
		if err != nil {
			log.Fatalf("Failed to accept a new connection: %s\n", err)
			continue
		}
		go handler(conn, writer, timeout, config)
		//go HandleTLSSession(conn.(*tls.Conn), writer, timeout)
	}
}

func handler(conn net.Conn, writer *RotWriter, timeout int, config *tls.Config) {
	log.Printf("New Conn: %s -> %s)\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
	conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	firstByte := []byte{0, 0}
	_, err := conn.Read(firstByte)
	if err != nil {
		conn.Close()
		return
	}
	conn2 := NewConn{
		conn,
		io.MultiReader(bytes.NewReader(firstByte), conn),
	}
	if bytes.Equal(firstByte[0:2], []uint8{22, 3}) {
		conn3 := tls.Server(conn2, config)
		go HandleTLSSession(conn3, writer, timeout)
	} else {
		go HandleTCPSession(conn2, writer, timeout)
	}
}
