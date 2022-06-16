package main

import (
	"fmt"
	"image/color"
	"log"
	"net/http"
	"os"

	g "github.com/AllenDang/giu"
	"github.com/gorilla/mux"

	"github.com/spf13/cobra"
)

var (
	source2 string
	data2   string
	target2 string
	rootCmd = &cobra.Command{
		Use:   "21e8-go-node",
		Short: "21e8 Miner/Node coded in Golang.",
		Long:  `21e8 Miner/Node coded in Golang.`,
		Args:  cobra.MinimumNArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Heyo what up")
		},
	}
	imguicmd = &cobra.Command{
		Use:   "imgui",
		Short: "Launch imgui node interface",
		Long:  `Launch imgui node interface`,
		// Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			wnd := g.NewMasterWindow("Go Node", 1200, 800, 0)
			wnd.Run(func() {
				wnd.SetBgColor(color.Black)
				w, h := wnd.GetSize()
				loop(float32(w), float32(h))
			})
		},
	}
	runner = &cobra.Command{
		Use:   "run",
		Short: "Run node on given port",
		Long:  `Launch imgui node interface`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r := mux.NewRouter()
			r.HandleFunc("/api/v2/mine", MineReq).Methods("POST")
			r.HandleFunc("/api/v2/hashwall", Hashwall).Methods("PUT")
			r.HandleFunc("/api/v2/data/{dataHash}", GetData).Methods("GET")
			r.HandleFunc("/api/v2/index/{sourceHash}", GetIndex).Methods("GET")
			r.HandleFunc("/api/v2/trie/{target}", TriePrefixLookup).Methods("GET")
			r.HandleFunc("/api/v2/raw/{rotation}", GetRaw).Methods("GET")
			r.HandleFunc("/api/v2/json/{rotation}", GetRaw).Methods("GET")
			r.HandleFunc("/binary", PostBinary).Methods("POST")
			port := args[0]
			fmt.Printf("Running node on port %s\n", args[0])
			// go func() {
			// 	http.ListenAndServe(":2180", nil)
			// 	fmt.Println("WS Server Started on port 2180")
			// }()

			log.Fatal(http.ListenAndServe(":"+port, r))
		},
	}
	mining = &cobra.Command{
		Use:   "mine",
		Short: "Mine data into a keyword",
		Long:  `21e8 Miner/Node coded in Golang.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			Mine(args[0], args[1], args[2], 0, true)
		},
	}
	cpu = &cobra.Command{
		Use:   "cpu",
		Short: "CPU Mine data into a keyword",
		Long:  `21e8 Miner/Node coded in Golang.`,
		// Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			Mine(args[0], args[1], args[2], 0, true)
		},
	}
	gpu = &cobra.Command{
		Use:   "gpu",
		Short: "GPU Mine data into a keyword",
		Long:  `21e8 Miner/Node coded in Golang.`,
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			HelloCpp(args[0], args[1], args[2])
		},
	}
)

func Execute() {
	// rootCmd.PersistentFlags().BoolVarP(&pFlag, "pFlag", "p", false, "")
	rootCmd.AddCommand(imguicmd)
	rootCmd.AddCommand(runner)
	mining.AddCommand(cpu)
	mining.AddCommand(gpu)
	rootCmd.AddCommand(mining)
	cpu.Flags().StringVarP(&source2, "source", "s", "source", "Source to be mined into")
	cpu.Flags().StringVarP(&data2, "data", "d", "data", "Data to be mine into keyword")
	cpu.Flags().StringVarP(&target2, "target", "t", "21e8", "Prefix to be matched")
	gpu.PersistentFlags().StringVarP(&source2, "source", "s", "source", "Source to be mined into")
	gpu.PersistentFlags().StringVarP(&data2, "data", "d", "data", "Data to be mine into keyword")
	gpu.PersistentFlags().StringVarP(&target2, "target", "t", "21e8", "Prefix to be matched")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
