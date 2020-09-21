package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/hhyouke/server/api"
	"github.com/hhyouke/server/conf"
	"github.com/hhyouke/server/logger"
	"github.com/hhyouke/server/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var profile string
var (
	// apiServerConfig conf.APIServerConfig
	// appLoggerConfig conf.AppLoggerConfig

	gconf conf.GlobalConfiguration
)
var defaultProfile string = "local"

var (
	rootCmd = &cobra.Command{
		Use:   "useage",
		Short: "useage short",
		Long:  `usage long`,
		Run: func(cmd *cobra.Command, args []string) {
			apiServe()
		},
	}
)

// NewRootCmd create new root cmd
func NewRootCmd() *cobra.Command {
	return rootCmd
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&profile, "config", "", "config file name")
}

func initConfig() {
	if profile == "" {
		log.Println("no profile input, using the default profile(local)")
		viper.SetConfigName(defaultProfile)
	} else {
		viper.SetConfigName(profile)
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // root of the working directory
	// readin config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("load config file error: %v\n", err.Error())
	}

	if err := viper.Unmarshal(&gconf); err != nil {
		log.Fatalf("unmarshall configuration error, %v\n", err.Error())
	}
}

func apiServe() {
	hostAndPort := fmt.Sprintf("%v:%v", gconf.API.Host, gconf.API.Port)
	appLogger := logger.NewLogger(gconf.Logging.Level)
	dblogger := appLogger.With(zap.String("component", "db"))
	db, err := models.Connect(&gconf.DB, dblogger)
	if err != nil {
		log.Fatalf("Error opening database: %+v", err)
	}
	defer db.Close() // close database connections when server shutdown
	api := api.NewAPI(context.Background(), &gconf.API, appLogger, db)
	api.ListenAndServe(hostAndPort)
}
