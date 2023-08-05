package Epoll

import (
	"golang.org/x/sys/unix"
	"log"
	"net"
	"reflect"
	"sync"
	"syscall"
	"context"
)


type Epoll struct {
	fd          int
	connections map[int]net.Conn
	contexts 	  map[int]context.Context
	lock        *sync.RWMutex
}

var epoller *Epoll

// Singleton
func GetEpollInstance() *Epoll{
	if epoller != nil {
		return epoller
	}

	// Start epoll
	var err error
	epoller, err = MkEpoll()
	if err != nil {
		panic(err)
	}
	return epoller
}

func MkEpoll() (*Epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &Epoll{
		fd:          fd,
		lock:        &sync.RWMutex{},
		contexts: make(map[int]context.Context),
		connections: make(map[int]net.Conn),
	}, nil
}

func (e *Epoll) Add(conn net.Conn, ctx context.Context) error {
	// Extract file descriptor associated with the connection
	fd := websocketFD(conn)
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, &unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, Fd: int32(fd)})
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	e.connections[fd] = conn
	e.contexts[fd] = ctx
	//if len(e.connections)%100 == 0 {
		log.Printf("Total number of connections: %v", len(e.connections))
	//}
	return nil
}

func (e *Epoll) Remove(conn net.Conn) error {
	fd := websocketFD(conn)
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	delete(e.connections, fd)
	//if len(e.connections)%100 == 0 {
		log.Printf("Total number of connections: %v", len(e.connections))
	//}
	return nil
}

func (e *Epoll) Wait() ([]net.Conn, error) {
	events := make([]unix.EpollEvent, 100)
	n, err := unix.EpollWait(e.fd, events, 100)
	if err != nil {
		return nil, nil
	}
	e.lock.RLock()
	defer e.lock.RUnlock()
	var connections []net.Conn
	for i := 0; i < n; i++ {
		conn := e.connections[int(events[i].Fd)]
		connections = append(connections, conn)
	}
	return connections, nil
}

func (e *Epoll) GetContext(conn net.Conn) context.Context {
	fd := websocketFD(conn)
	e.lock.RLock()
	defer e.lock.RUnlock()
	return e.contexts[fd]
}

func websocketFD(conn net.Conn) int {
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")

	return int(pfdVal.FieldByName("Sysfd").Int())
}

func GetIdFromConn(conn net.Conn) int {
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")

	return int(pfdVal.FieldByName("Sysfd").Int())
}

func GetConnById(connId int) net.Conn {
	return epoller.connections[connId]
}