package cmd

import (
	"github.com/xuyunfeng12388/gin_vue/config"
	"github.com/xuyunfeng12388/gin_vue/db"
	"github.com/xuyunfeng12388/gin_vue/model"
	"github.com/xuyunfeng12388/gin_vue/router"
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	logger  = &logrus.Logger{}
	rootCmd = &cobra.Command{}
)

func initConfig() {
	config.MustInit(os.Stdout, cfgFile) // 配置初始化
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config/dev.yaml", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().Bool("debug", true, "开启debug")
	viper.SetDefault("gin.mode", rootCmd.PersistentFlags().Lookup("debug"))
}

func Execute() error {
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		_, err := db.Mysql(
			viper.GetString("db.hostname"),
			viper.GetInt("db.port"),
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.dbname"),
		)
		if err != nil {
			return err
		}

		db.DB.AutoMigrate(&model.User{})

		defer db.DB.Close()

		r := router.SetupRouter()
		r.Run()

		port := viper.GetString("server.port")
		println(port)
		log.Println("port = *** =", port)
		return http.ListenAndServe(port, nil) // listen and serve
	}

	return rootCmd.Execute()

}

