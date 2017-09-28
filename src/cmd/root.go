package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

//
// RootCmd represents the base command
//
var RootCmd = &cobra.Command{
	Use:   "dbs",
	Short: "Simple restful app with different dbs repositories",
	Run:   root,
}

func root(cmd *cobra.Command, args []string) {
}

//
// Execute is application main entry point
//
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Printf("dbs: failed to execute root cmd: %s", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./dist/etc/config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./etc/")
		viper.SetConfigName("config")
		viper.SetConfigType("json")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("dbs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("dbs: failed to read config file:", viper.ConfigFileUsed())
	}

	log.Println("dbs: kstarting with config file:", viper.ConfigFileUsed())

}
