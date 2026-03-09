package main
import (
  "github.com/urfave/cli"
  "os"
  "fmt"
)

func main(){
	app := cli.NewApp()
	app.Name = "greeting app"
	app.Usage = "say a greeting"


	author := cli.Author{
        Name:  "Gilles",
        Email: "gilles@gmail.com",
    }

	app.Action = func(c *cli.Context) error {
		fmt.Println("the author is: ", author)
		return nil
	}

	
	//app.HideHelp = true
	app.Run(os.Args)
}