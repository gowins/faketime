package main

import (
	"fmt"

	"github.com/gowins/faketime/time"
	"github.com/spf13/cobra"
)

var (
	pid int
	sec int
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "faketime",
		Short: "faketime is simulated time",
		Run: func(cmd *cobra.Command, args []string) {

			if pid <= 0 {
				fmt.Println("Error: flag needs an argument: --pid")
				_ = cmd.Help()
				return
			}

			if sec == 0 {
				fmt.Println("Error: flag needs an argument: --sec")
				_ = cmd.Help()
				return
			}

			if err := time.ModifyTime(pid, int64(sec), 0, 1); err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("OK")
		},
	}

	rootCmd.Flags().IntVar(&pid, "pid", 0, "process identification number")
	rootCmd.Flags().IntVar(&sec, "sec", 0, "TV_SEC_DELTA")

	_ = rootCmd.Execute()
}
