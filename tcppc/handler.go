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

const data1 = `HTTP/1.1 200 OK
Server: nginx/1.14.2
Content-Type: text/html
Connection: keep-alive
Vary: Accept-Encoding
ETag: W/"6203b88e-264"
X-Content-Type-Options: nosniff`

const data2 = `
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
<!-
root
uid=0(root) gid=0(root) groups=0(root)

root:x:0:0:root:/root:/bin/bash
bin:x:1:1:bin:/bin:/sbin/nologin
daemon:x:2:2:daemon:/sbin:/sbin/nologin
adm:x:3:4:adm:/var/adm:/sbin/nologin
lp:x:4:7:lp:/var/spool/lpd:/sbin/nologin
sync:x:5:0:sync:/sbin:/bin/sync
shutdown:x:6:0:shutdown:/sbin:/sbin/shutdown
halt:x:7:0:halt:/sbin:/sbin/halt
mail:x:8:12:mail:/var/spool/mail:/sbin/nologin
operator:x:11:0:operator:/root:/sbin/nologin
games:x:12:100:games:/usr/games:/sbin/nologin
ftp:x:14:50:FTP User:/var/ftp:/sbin/nologin
nobody:x:99:99:Nobody:/:/sbin/nologin
success
SUCCCESS
MikroTik
RouterOS
->`

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
		io.MultiReader(bytes.NewReader(firstByte), conn), //由于之前已经读取了一部分数据 需要覆盖 reader 还原buffer
	}
	if bytes.Equal(firstByte[0:2], []uint8{22, 3}) { //TLS 首包特征  TODO 其他协议模拟
		conn3 := tls.Server(conn2, config)
		go HandleTLSSession(conn3, writer, timeout)
	} else {
		go HandleTCPSession(conn2, writer, timeout)
	}
}
