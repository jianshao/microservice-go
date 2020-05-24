package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os/exec"
)

func main()  {
	http.HandleFunc("/build/object", handle)
	http.ListenAndServe(":9090", nil)
}

func handle(w http.ResponseWriter, req *http.Request)  {
	if err := req.ParseForm(); err != nil {
		logrus.Error("parse failed %s", err)
		return
	}

	name := req.FormValue("app")
	cmd := exec.Command(fmt.Sprintf("/home/eric/%s/%s.sh", name, name))
	if output, err := cmd.Output(); err == nil {
		w.Write([]byte(fmt.Sprintf("build %s success\n", name)))
	} else {
		w.Write([]byte(fmt.Sprintf("build %s failed: %s, %s\n", name, err, output)))
	}

}
