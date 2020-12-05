package cmd

import (
	"github.com/shuheitakada/haumea/elbv2"
	"github.com/spf13/cobra"
)

// healthcheckCmd represents the healthcheck command
var healthcheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Check health of the specified targets or all of your targets",
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		client := elbv2.NewClient(role)
		switch len(args) {
		case 1:
			var targetGroupArns []string
			for _, server := range config[args[0]].(map[string]interface{}) {
				targetGroupArns = append(targetGroupArns, server.(map[string]interface{})["target_group_arn"].(string))
			}
			client.DescribeAllTargetHealth(targetGroupArns)
		case 2:
			server := config[args[0]].(map[string]interface{})[args[1]]
			targetGroupArn := server.(map[string]interface{})["target_group_arn"]
			client.DescribeTargetHealth(targetGroupArn.(string))
		}
	},
}

func init() {
	rootCmd.AddCommand(healthcheckCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// healthcheckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// healthcheckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
