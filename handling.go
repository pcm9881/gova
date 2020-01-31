package gova

import (
	"fmt"
	"os/exec"
)

func PrintJavaVersion() string {
	out, err := exec.Command("sh", "-c", "echo $(java -version 2>&1)").Output()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("%s\n", out)
	return fmt.Sprintf("%s", out)
}

func ExecJar() string {
	out, err := exec.Command("java", "-cp", "./example/lib/GoTest.jar", "com.pcm9881.gotest.GoTest").Output()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("%s\n", out)
	return fmt.Sprintf("%s", out)
}
