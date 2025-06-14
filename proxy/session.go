package proxy

import (
	"log"
	"net"
	"sync"
	"time"
)

type Session struct {
	ClientAddr *net.UDPAddr
	ServerConn *net.UDPConn
	ServerAddr *net.UDPAddr
	ProxyConn  *net.UDPConn
	LastSeen   time.Time
	closed     chan struct{}
	once       sync.Once
}

func (s *Session) SendToServer(data []byte) {
	s.LastSeen = time.Now()
	_, err := s.ServerConn.Write(data)
	if err != nil {
		log.Println("Ошибка при отправке на сервер:", err)
	}
}

type SessionManager struct {
	sync.Mutex
	sessions map[string]*Session
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*Session),
	}
}

func (m *SessionManager) Get(client *net.UDPAddr, serverAddr string, proxyConn *net.UDPConn) *Session {
	m.Lock()
	defer m.Unlock()

	key := client.String()
	if session, ok := m.sessions[key]; ok {
		return session
	}

	serverUDPAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		log.Println("Ошибка при разрешении адреса сервера:", err)
		return nil
	}

	serverConn, err := net.DialUDP("udp", nil, serverUDPAddr)
	if err != nil {
		log.Println("Ошибка при подключении к серверу:", err)
		return nil
	}

	session := &Session{
		ClientAddr: client,
		ServerConn: serverConn,
		ServerAddr: serverUDPAddr,
		ProxyConn:  proxyConn,
		LastSeen:   time.Now(),
		closed:     make(chan struct{}),
	}

	// Чтение от сервера
	go func() {
		buf := make([]byte, 4096)
		for {
			select {
			case <-session.closed:
				return
			default:
				serverConn.SetReadDeadline(time.Now().Add(5 * time.Second))
				n, err := serverConn.Read(buf)
				if err != nil {
					if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
						continue
					}
					log.Println("Ошибка чтения от сервера:", err)
					session.close()
					return
				}
				_, err = proxyConn.WriteToUDP(buf[:n], client)
				if err != nil {
					log.Println("Ошибка отправки клиенту:", err)
					session.close()
					return
				}
			}
		}
	}()

	// Таймаут — 30 секунд без активности
	go func() {
		for {
			select {
			case <-time.After(10 * time.Second):
				if time.Since(session.LastSeen) > 30*time.Second {
					session.close()
					return
				}
			case <-session.closed:
				return
			}
		}
	}()

	m.sessions[key] = session
	return session
}

func (s *Session) close() {
	s.once.Do(func() {
		close(s.closed)
		s.ServerConn.Close()
	})
}
