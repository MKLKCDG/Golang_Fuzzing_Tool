package option

import (
	"fmt"
	"net/http"
	"os"
)

type Options struct {
	ShowHelp     bool
	TargetUrl    string
	WordlistPath string
	Post         bool
	Data         string
	FilterCode   int
	Slow         int
	Medium       string
	Host         string
}

func (o Options) GetRequest(v string) {

	resp, err := http.Get(v)
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != o.FilterCode {
		fmt.Println("URL: ", v)
		fmt.Println("Status code: ", resp.StatusCode)

	}

	//fmt.Println("Status code: ", resp.StatusCode)
	//fmt.Println("URL: ", v)

	//fmt.Println(resp)
	//fmt.Print(o.TargetUrl)
	//fmt.Println(main.satir)
	//if resp.StatusCode != 404 {
	//	fmt.Println(resp)
	//}

}

func (o Options) PostRequest(v string) {
	resp, _ := http.Post(v, "application/x-www-form-urlencoded", nil)

	if resp.StatusCode != o.FilterCode {
		fmt.Println("URL: ", v)
		fmt.Println("Status code: ", resp.StatusCode)

	}

}

func (o Options) DisplayHelp() {

	if !o.ShowHelp {
		return
	}
	usage := `  
	                                                                         
 HHHHHHHHH     HHHHHHHHHEEEEEEEEEEEEEEEEEEEEEELLLLLLLLLLL             PPPPPPPPPPPPPPPPP   
 H:::::::H     H:::::::HE::::::::::::::::::::EL:::::::::L             P::::::::::::::::P  
 H:::::::H     H:::::::HE::::::::::::::::::::EL:::::::::L             P::::::PPPPPP:::::P 
 HH::::::H     H::::::HHEE::::::EEEEEEEEE::::ELL:::::::LL             PP:::::P     P:::::P
   H:::::H     H:::::H    E:::::E       EEEEEE  L:::::L                 P::::P     P:::::P
   H:::::H     H:::::H    E:::::E               L:::::L                 P::::P     P:::::P
   H::::::HHHHH::::::H    E::::::EEEEEEEEEE     L:::::L                 P::::PPPPPP:::::P 
   H:::::::::::::::::H    E:::::::::::::::E     L:::::L                 P:::::::::::::PP  
   H:::::::::::::::::H    E:::::::::::::::E     L:::::L                 P::::PPPPPPPPP    
   H::::::HHHHH::::::H    E::::::EEEEEEEEEE     L:::::L                 P::::P            
   H:::::H     H:::::H    E:::::E               L:::::L                 P::::P            
   H:::::H     H:::::H    E:::::E       EEEEEE  L:::::L         LLLLLL  P::::P            
 HH::::::H     H::::::HHEE::::::EEEEEEEE:::::ELL:::::::LLLLLLLLL:::::LPP::::::PP          
 H:::::::H     H:::::::HE::::::::::::::::::::EL::::::::::::::::::::::LP::::::::P          
 H:::::::H     H:::::::HE::::::::::::::::::::EL::::::::::::::::::::::LP::::::::P          
 HHHHHHHHH     HHHHHHHHHEEEEEEEEEEEEEEEEEEEEEELLLLLLLLLLLLLLLLLLLLLLLLPPPPPPPPPP   
  
 -data string
        To send a POST request
 -fc int
        Qutput filter for  response statu (default 400)
 -med 
        To activate medium speed mode
 -post
        To send a Post request
 -slow
        sTo activate slow speed mode
 -u string
        URL to send request (dont't forget to add FUZZ)
 -w string
        Path of wordlist to use		


 `
	fmt.Printf("%s", usage)
	os.Exit(0)
}
