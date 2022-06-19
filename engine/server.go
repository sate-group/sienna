package engine

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

const (
	DIVIDER = '\n'
)

type IListenerOnString func(client *Client, data string)
type IListenerOnError func(err error)

type Server struct {
	Clients        Clients
	Port           uint16
	eventsOnString []IListenerOnString
	eventsOnError  []IListenerOnError
}

func NewServer(options *ServerOptions) *Server {
	server := &Server{
		Clients: Clients{
			repo: make(map[uuid.UUID]*Client),
		},
		Port:           options.getPort(),
		eventsOnString: []IListenerOnString{},
		eventsOnError:  []IListenerOnError{},
	}

	return server
}

func (s *Server) Listen() error {
	address := ":" + strconv.Itoa(int(s.Port))
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error in accepting waits for")
		}
		id := uuid.New()
		client := &Client{
			Id:   id,
			Conn: conn,
		}
		s.Clients.Add(client)
		go s.HandleClient(client)
	}
}

func (s *Server) HandleClient(client *Client) {
	id := client.Id
	conn := client.Conn
	defer func() {
		conn.Close()
		s.Clients.Remove(id)
	}()
	for {
		input, err := bufio.NewReader(conn).ReadString(DIVIDER)
		if err != nil {
			log.Print(err)
			return
		}
		data := strings.ReplaceAll(input, string(DIVIDER), "")
		for _, listener := range s.eventsOnString {
			listener(client, data)
		}
	}
}

func (s *Server) OnString(listener IListenerOnString) {
	s.eventsOnString = append(s.eventsOnString, listener)
}

func (s *Server) OnError(listener IListenerOnError) {
	s.eventsOnError = append(s.eventsOnError, listener)
}
