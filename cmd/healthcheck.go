package cmd

import (
	"github.com/shuheitakada/haumea/elbv2"
	"github.com/spf13/cobra"
)

// healthcheckCmd represents the healthcheck command
var healthcheckCmd = &cobra.Command{
	Use:   "healthcheck",
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
