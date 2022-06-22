# Finally, the first step into the world of Go programming.
# Lets take a look at the basic structure of a go file. To know more about installation, execution etc., please
# check the Getting Started page


# Why are we using this main package here? :/
# So the below import tells the Go compiler that this should be compiled as an executable program rather than
# a shared library. By convention, Executable programs (the ones with the main package) are called Commands and others
# are called simply Packages
package main

# Now importig the fmt package. Why?
# This is required as it is used for formatted input and output with functions
import "fmt"

func main(){
    fmt.Println("Hello World :) ")
}

# To run, "go run hello_world.go" and voila!!