package cmd

import (
	"context"
	"fmt"
	"github/yudgxe/leadgen.market/pkg/config"
	"github/yudgxe/leadgen.market/pkg/utils"
	"github/yudgxe/leadgen.market/service/cache"
	"github/yudgxe/leadgen.market/service/server"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Long:  "Start server",
	Run: func(cmd *cobra.Command, args []string) {
		// todo: wrap in app package 
		
		config.MustReadConfig(utils.MustGet(cmd.Flags().GetString("config"))) // todo: default value

		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.Level(viper.GetInt("logger.level")))

		if err := cache.SetupRedis(
			context.TODO(),
			redis.Options{
				Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
				Password: viper.GetString("redis.password"),
			},
			viper.GetInt("redis.ttl_min"),
		); err != nil {
			log.Errorf("error on init redis - %v", err) // todo: mock if redis not setup
			return
		}

		srv, err := server.New(
			viper.GetString("server.host"),
			viper.GetString("server.port"),
		)
		if err != nil {
			log.Errorf("error on init server - %v", err)
			return
		}

		srv.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().String("config", "config.toml", "path to config file")
}
