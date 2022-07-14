
// to download all package
// fmt : default package for text, standard from (format package)
//https://github.com/rnadjqaim
//go programs are organized into package
//golang standard library, provides different core packages for us to use

// a package : collection of source files.

import "fmt"

//main function is the entrypoint of golang
// we can use it by firstly importing it.
//go run : complies and runs the codes
//minmum projects strucuture and how to run go program
func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50 //constant : the vzlue cant be changed
	const remainingTickets = 50
	fmt.Printf("wlecome to %v booking application", conferenceName)
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("we have total of %v tic kets and %v are avaliable\n", conferenceTickets, remainingTickets)

}
