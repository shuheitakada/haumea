package cmd

import (
	"github.com/shuheitakada/haumea/elbv2"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Registers the specified targets or all of your targets",
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		client := elbv2.NewClient(role)
		switch len(args) {
		case 1:
			for _, server := range config[args[0]].(map[string]interface{}) {
				targetGroupArn := server.(map[string]interface{})["target_group_arn"]
				targets := server.(map[string]interface{})["targets"]
				client.RegisterTargets(targetGroupArn.(string), targets.([]interface{}))
			}
		case 2:
			server := config[args[0]].(map[string]interface{})[args[1]]
			targetGroupArn := server.(map[string]interface{})["target_group_arn"]
			targets := server.(map[string]interface{})["targets"]
			client.RegisterTargets(targetGroupArn.(string), targets.([]interface{}))
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
