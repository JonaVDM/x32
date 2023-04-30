package cmd

import (
	// "strconv"
	// "time"

	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// wavesCmd represents the waves command
var wavesCmd = &cobra.Command{
	Use:   "waves",
	Short: "Makes some slides do a bit of a wave",
	Args:  cobra.RangeArgs(1, 2),
	RunE:  runWavesCommand,
}

func init() {
	rootCmd.AddCommand(wavesCmd)

	wavesCmd.Flags().IntP("mix", "m", -1, "The mixbus to execute this on")
	wavesCmd.Flags().IntP("amount", "a", 100, "The amount of times it needs to change value")
}

func runWavesCommand(cmd *cobra.Command, args []string) error {
	loops, err := cmd.Flags().GetInt("amount")
	if err != nil {
		return err
	}

	mixbus, err := cmd.Flags().GetInt("mix")
	if err != nil {
		return err
	}

	start, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	end := start
	if len(args) > 1 {
		end, err = strconv.Atoi(args[1])
		if err != nil {
			return err
		}
	}

	values := waves()

	for i := 0; i < loops; i++ {
		for channel := start; channel <= end; channel++ {
			index := (len(values) + i + (channel-start)*4) % len(values)

			if mixbus > 0 {
				client.SetMixbusFader(channel, mixbus, values[index])
			} else {
				client.SetFader(channel, values[index])
			}
		}
		time.Sleep(time.Millisecond * 75)
	}

	return nil
}

// waves gives back the values to make a slider go from down to up back down
func waves() []float32 {
	return []float32{0.000000,
		0.035191,
		0.070381,
		0.105572,
		0.140762,
		0.175953,
		0.211144,
		0.246334,
		0.281525,
		0.316716,
		0.351906,
		0.387097,
		0.422287,
		0.457478,
		0.492669,
		0.527859,
		0.563050,
		0.598240,
		0.633431,
		0.668622,
		0.703812,
		0.739003,
		0.774194,
		0.809384,
		0.844575,
		0.879765,
		0.914956,
		0.950147,
		0.985337,
		1.000000,
		0.964809,
		0.929619,
		0.894428,
		0.859238,
		0.824047,
		0.788856,
		0.753666,
		0.718475,
		0.683284,
		0.648094,
		0.612903,
		0.577713,
		0.542522,
		0.507331,
		0.472141,
		0.436950,
		0.401760,
		0.366569,
		0.331378,
		0.296188,
		0.260997,
		0.225806,
		0.190616,
		0.155425,
		0.120235,
		0.085044,
		0.049853,
		0.014663,
		0.000000,
	}
}
