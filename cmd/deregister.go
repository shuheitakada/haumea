package cmd

import (
	"github.com/shuheitakada/haumea/elbv2"
	"github.com/spf13/cobra"
)

// deregisterCmd represents the deregister command
var deregisterCmd = &cobra.Command{
	Use:   "deregister",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		client := elbv2.NewClient(role)
		switch len(args) {
		case 1:
			for _, server := range config[args[0]].(map[string]interface{}) {
				targetGroupArn := server.(map[string]interface{})["target_group_arn"]
				targets := server.(map[string]interface{})["targets"]
				client.DeregisterTargets(targetGroupArn.(string), targets.([]interface{}))
			}
		case 2:
			server := config[args[0]].(map[string]interface{})[args[1]]
			targetGroupArn := server.(map[string]interface{})["target_group_arn"]
			targets := server.(map[string]interface{})["targets"]
			client.DeregisterTargets(targetGroupArn.(string), targets.([]interface{}))
		}
	},
}

func init() {
	rootCmd.AddCommand(deregisterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deregisterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deregisterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
