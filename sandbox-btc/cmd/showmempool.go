package cmd

import (
	"fmt"

	"github.com/asciiu/appa/sandbox-btc/rpc"
	"github.com/spf13/cobra"
)

func init() {
	showMempoolCmd.Flags().IntVar(&jsonrpcPort, "jsonrpc-port", 9334, "JSON-RPC port to connect to.")
}

var showMempoolCmd = &cobra.Command{
	Use: "showmempool",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := rpc.NewClient(jsonrpcPort)
		if err != nil {
			return err
		}
		defer c.Close()

		var reply string
		if err := c.Call("RPC.GetMempool", nil, &reply); err != nil {
			return err
		}

		fmt.Println(reply)

		return nil
	},
}
