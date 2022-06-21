package engine

import (
	"bufio"
	"io"
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
type IListenerOnStruct[T any] func(client *Client, dto T)

type Server struct {
	Clients        Clients
	Port           uint16
	eventsOnString []IListenerOnString
	eventsOnError  []IListenerOnError
	eventsOnStruct []IListenerOnStruct
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
			s.CauseError(err)
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
		if err == io.EOF { // exit
			return
		} else if err != nil {
			s.CauseError(err)
			return
		}
		data := strings.ReplaceAll(input, string(DIVIDER), "")
		for _, listener := range s.eventsOnString {
			listener(client, data)
		}

		// for _, listener := range s.eventsOnStruct {

		// }
	}
}

func (s *Server) CauseError(err error) {
	for _, listener := range s.eventsOnError {
		listener(err)
	}
}

func (s *Server) OnString(listener IListenerOnString) {
	s.eventsOnString = append(s.eventsOnString, listener)
}

func (s *Server) OnError(listener IListenerOnError) {
	s.eventsOnError = append(s.eventsOnError, listener)
}

type S struct{}

// Identity is a simple identity method that works for any type.
func (S) Identity[T any](v T) T { return v }