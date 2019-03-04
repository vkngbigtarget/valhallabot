package models

// Server from BM
type Server struct {
	Type       string            `json:"type"`
	ID         string            `json:"id"`
	Attributes *ServerAttributes `json:"attributes"`
}

// ServerAttributes for server
type ServerAttributes struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	IP         string `json:"ip"`
	Players    int    `json:"players"`
	MaxPlayers int    `json:"maxPlayers"`
	Rank       int    `json:"rank"`
	Status     string `json:"status"`
}
