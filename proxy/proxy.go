package proxy

import (
	"log"
	"net"
)

var (
	sessions = NewSessionManager()
)

// Start запускает UDP-прокси
func Start(listenAddr, serverAddr string) {
	addr, err := net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		log.Fatalf("Ошибка resolve proxy address: %v", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Ошибка запуска прокси: %v", err)
	}
	defer conn.Close()

	buffer := make([]byte, 4096)

	log.Printf("[+] Прокси слушает на %s → %s\n", listenAddr, serverAddr)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Ошибка чтения:", err)
			continue
		}

		data := make([]byte, n)
		copy(data, buffer[:n])

		go handlePacket(conn, clientAddr, serverAddr, data)
	}
}

func handlePacket(proxyConn *net.UDPConn, clientAddr *net.UDPAddr, serverAddr string, data []byte) {
	session := sessions.Get(clientAddr, serverAddr, proxyConn)
	session.SendToServer(data)
}
