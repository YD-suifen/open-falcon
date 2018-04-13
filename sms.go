package main


import (
	"net/http"


	"github.com/toolkits/web/param"

	"fmt"
	"io/ioutil"

)

func configSmsrRoutes() {


	http.HandleFunc("/sender/sms", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()


		tos := param.MustString(r, "tos")
		content := param.MustString(r, "content")

		resp, err := http.Get("http://test.live.replays.net/api.2.0.php?/member/captchaforop/&send_to="+tos+"&msg="+content)
		defer resp.Body.Close()
		if err !=nil {
			fmt.Println(err)
		}
		body, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println(w, body)


	})

}

func main()  {

	configSmsrRoutes()

	http.ListenAndServe("127.0.0.1:4040", nil)



}
