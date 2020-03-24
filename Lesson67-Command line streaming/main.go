// Credits: https://gist.github.com/hivefans/ffeaf3964924c943dd7ed83b406bbdea

package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {

	// var pods []struct {
	// 	podName        string
	// 	NoOfContainers string
	// 	status         string
	// 	restarts       string
	// 	age            string
	// }

	namespace := "fabrikam"
	serviceList := []string{"cnwebserver", "cnreverseproxy"}
	services := "'" + strings.Join(serviceList, "\\|") + "'"

	cmdLine := fmt.Sprintf("kubectl get pods --namespace=%s | grep %s", namespace, services)

	fmt.Println(cmdLine)
	cmd := exec.Command("sh", "-c", cmdLine)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	// if err := json.NewDecoder(stdout).Decode(&pods); err != nil {
	// log.Fatal(err)
	// }

	// if err := cmd.Wait(); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Pods Info=%s\n", pods)
}
