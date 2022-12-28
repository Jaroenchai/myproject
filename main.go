package main

import (
	"fmt"

	"github.com/Jaroenchai/cinema/movie"
	"github.com/Jaroenchai/cinema/ticket"
)

func main() {
	movieName := movie.FindMovieName("tt1")
	fmt.Println(movieName)
	movie.ReviewMovie(movieName, 8.5)
	ticket.BuyTicket(movieName)
}
