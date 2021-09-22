package src

import (
	"errors"
	"net"
	"strconv"
)

type struct Server {
	port int
	address string
	listener Listener
}

(s *Server) func Bind() err {

	listener, err := net.Listen("tcp", s.address + ":" + strconv(s.port))
	if(err != nil){
		return errors.New("could not bind to socket", err.Error())
	}
	s.listener = listener

}

(s *Server) func Serve(connection Conn){
	time.Sleep(5)
	buf := make([]byte, 1024)
  	// Read the incoming connection into the buffer.
  	reqLen, err := conn.Read(buf)
	connection.Write([]byte("Message received."))
	connection.Close()
}

(s *Server) func Accept() (Conn, err){
	if(s.listener == nil){
		return errors.New("no socket binding")
	}

	conn, err := s.listener.Accept()
	if(err != nil){
		return nil, errors.New("Error accepting: ", err.Error())
	}
	return conn, nil

}

(s *Server) func ValidatePort() err {
	if(s.port < 0 || s.port > 65536){
		return errors.New("invalid port")
	}
	return nil
}