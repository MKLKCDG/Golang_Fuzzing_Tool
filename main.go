package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"src/option"
	"strconv"
	"time"
)

func main() {
	var options option.Options

	flag.BoolVar(&options.ShowHelp, "help", false, "shows usage details")                          //tamam
	flag.StringVar(&options.TargetUrl, "u", "", "URL to send request (dont't forget to add FUZZ)") //tamam
	flag.StringVar(&options.WordlistPath, "w", "", "Path of wordlist to use")                      //tamam
	flag.BoolVar(&options.Post, "post", false, "To send a Post request")                           //tamam
	flag.StringVar(&options.Data, "data", "", "To send a POST request")                            //bu post ile aynı açıklama anlamadım
	flag.IntVar(&options.FilterCode, "fc", 400, "Qutput filter for  response statu")               //tamam
	flag.IntVar(&options.Slow, "slow", 0, "sTo activate slow speed mode")                          //tamam
	flag.StringVar(&options.Medium, "med", "", "To activate medium speed mode")

	flag.Parse()

	options.DisplayHelp()

	//şurda bi post atmayı deniyelim
	if options.Post {
		//StatusFilter := strconv.Itoa(options.FilterCode)
		usage := `

		              _________
		            /         /.
    .-------------.        /_________/ |
   /             / |       |         | |
  /+============+\ |       | |====|  | |
  ||C:\>Post    || |       |         | |
  ||            || |       | |====|  | |
  ||            || |       |   ___   | |
  ||            || |       |  |166|  | |
  ||            ||/@@@     |   ---   | |
  \+============+/     @   |_________|./.
                       @          ..  ....'
    ..................@      __.'.'  ''
   /oooooooooooooooo//     ///
  /................//    /_/
  ------------------

       Method: POST
       URL: ` + options.TargetUrl + `
       Status Filter: ` + strconv.Itoa(options.FilterCode) + `
`

		go fmt.Printf("%s", usage)

		file, _ := os.Open(options.WordlistPath)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			satir := scanner.Text()
			options.PostRequest(options.TargetUrl + "/" + satir)
			time.Sleep(time.Duration(options.Slow) * time.Second)

		}

		os.Exit(0)

	}

	//şöyle get atalım
	if options.TargetUrl != "" {

		usage := `
		_______________                         |*\_/*|________
		|  ___________  |     .-.     .-.      ||_/-\_|______  |
		| |           | |    .****. .****.     | |           | |
		| |   0   0   | |    .*****.*****.     | |   0   0   | |
		| |     -     | |     .*********.      | |     -     | |
		| |   \___/   | |      .*******.       | |   \___/   | |
		| |___     ___| |       .*****.        | |___________| |
		|_____|\_/|_____|        .***.         |_______________|
		  _|__|/ \|_|_.............*.............._|________|_
		 / ********** \                          / ********** \
		/ ************ \                        / ***********  \
		----------------                        ----------------

	                          Method: Get
				  Url: ` + options.TargetUrl + `
				  Status Filter: ` + strconv.Itoa(options.FilterCode) + `
`

		fmt.Printf("%s", usage)

		file, _ := os.Open(options.WordlistPath)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			satir := scanner.Text()
			options.GetRequest(options.TargetUrl + "/" + satir)
			time.Sleep(time.Duration(options.Slow) * time.Second)
		}

	}

}
