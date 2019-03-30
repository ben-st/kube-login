package cmd

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type clusterConfig struct {
	idpIssuerURL    string
	clientID        string
	clientSecret    string
	refreshToken    string
	idToken         string
	clusterName     string
	userName        string
	insecureCluster bool
	port            int
}

type authConfig struct {
	ClientID     string
	ClientSecret string
	User         string
	Password     string
	URL          string
	insecureOIDC bool
}

var cfgFile string
var username string
var password string
var clusterName string
var clientID string
var clientSecret string
var idpIssuerURL string
var insecureOIDC bool
var insecureCluster bool
var port int
var showToken bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kube-login",
	Short: "login to keycloak and generate/update kubeconfig with id and refresh token",
	Run: func(cmd *cobra.Command, args []string) {

		// construct URL with oidc string
		URL := fmt.Sprintf("%s/protocol/openid-connect/token", viper.GetString("idp-issuer-url"))

		// initialize config struct for oidc request
		c := authConfig{
			URL:          URL,
			ClientID:     viper.GetString("clientID"),
			ClientSecret: viper.GetString("clientSecret"),
			User:         viper.GetString("username"),
			Password:     viper.GetString("password"),
			insecureOIDC: viper.GetBool("insecure-oidc"),
		}

		// initialize clusterconfig struct for later kubeconfig manipulation
		cc := clusterConfig{
			idpIssuerURL:    viper.GetString("idp-issuer-url"),
			clientID:        viper.GetString("clientid"),
			clientSecret:    viper.GetString("clientsecret"),
			userName:        viper.GetString("username"),
			clusterName:     viper.GetString("clustername"),
			insecureCluster: viper.GetBool("insecure-cluster"),
			port:            viper.GetInt("port"),
		}

		// give user information about his current user and cluster
		// maybe move to loglevel info and additional logger
		log.Printf("username: %v \n", cc.userName)
		log.Printf("cluster: %v \n", cc.clusterName)

		// get id and refresh token from keycloak as tokenSet struct
		tokenSet := c.GetTokenSet()

		cc.idToken = tokenSet.IDToken
		cc.refreshToken = tokenSet.RefreshToken

		// if show-token is provided as a flag, only output tokens and exit
		if viper.GetBool("show-token") {
			fmt.Println("id-token is: ", cc.idToken)
			fmt.Println("refresh-token is: ", cc.refreshToken)
			os.Exit(0)
		}

		// set oidc config
		// and patch kubeconfig
		cc.SetAuthConfig()
		cc.SetClusterConfig()
		cc.SetClusterContext()
		cc.UseClusterContext()

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kube-login.yml)")
	rootCmd.PersistentFlags().BoolVar(&showToken, "show-token", false, "show keycloak token and exit")

	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "username for keycloak")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password for keycloak")
	rootCmd.PersistentFlags().StringVarP(&clusterName, "clustername", "c", "", "clustername fqdn e.g api.kubernetes.example")
	rootCmd.PersistentFlags().IntVar(&port, "port", 6443, "port for apiserver")

	rootCmd.PersistentFlags().StringVar(&clientID, "clientid", "", "clientid for idp")
	rootCmd.PersistentFlags().StringVar(&clientSecret, "clientsecret", "", "client secret for idp")
	rootCmd.PersistentFlags().StringVar(&idpIssuerURL, "idp-issuer-url", "", "idp/oidc fqdn")
	rootCmd.PersistentFlags().BoolVar(&insecureOIDC, "insecure-oidc", false, "if set insecure tls to oidc provider will be used, use with caution")
	rootCmd.PersistentFlags().BoolVar(&insecureCluster, "insecure-cluster", true, "if set insecure tls to cluster in kubeconfig will be set, use with caution")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Viper Bind Flags
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("clustername", rootCmd.PersistentFlags().Lookup("clustername"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("clientid", rootCmd.PersistentFlags().Lookup("clientid"))
	viper.BindPFlag("clientsecret", rootCmd.PersistentFlags().Lookup("clientsecret"))
	viper.BindPFlag("idp-issuer-url", rootCmd.PersistentFlags().Lookup("idp-issuer-url"))
	viper.BindPFlag("insecureoidc", rootCmd.PersistentFlags().Lookup("insecure-oidc"))
	viper.BindPFlag("insecurecluster", rootCmd.PersistentFlags().Lookup("insecure-cluster"))
	viper.BindPFlag("show-token", rootCmd.PersistentFlags().Lookup("show-token"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kube-login" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kube-login")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}
