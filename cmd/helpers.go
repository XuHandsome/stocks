package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func ExecuteCommand(name string, subname string, args ...string) (string, error) {
	args = append([]string{subname}, args...)

	cmd := exec.Command(name, args...)
	bytes, err := cmd.CombinedOutput()

	return string(bytes), err
}

func Error(cmd *cobra.Command, args []string, err error) {
	_, err = fmt.Fprintf(os.Stderr, "execute %s args:%v error:%v\n", cmd.Name(), args, err)
	if err != nil {
		return
	}
	os.Exit(1)
}
