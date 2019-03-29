package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

// UseClusterContext will use the context in kubeconfig from the clusterName variable
func (c *clusterConfig) UseClusterContext() {

	cmd := exec.Command("kubectl", "config", "use-context", c.clusterName)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("error: cannot set cluster config: %v", err)
	}

}

// SetClusterContext will set the context in kubeconfig to the clusterName variable
func (c *clusterConfig) SetClusterContext() {

	cluster := fmt.Sprintf("--cluster=%s", c.clusterName)
	user := fmt.Sprintf("--user=%s", c.userName)

	cmd := exec.Command("kubectl", "config", "set-context", c.clusterName,
		cluster,
		user)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("error: cannot set cluster config: %v", err)
	}

}

// SetClusterConfig will set the cluster config in kubeconfig
func (c *clusterConfig) SetClusterConfig() {

	server := fmt.Sprintf("--server=https://%s:%d", c.clusterName, c.port)
	insecure := fmt.Sprintf("--insecure-skip-tls-verify")

	cmd := exec.Command("kubectl", "config", "set-cluster", c.clusterName,
		server,
		insecure)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("error: cannot set cluster config: %v", err)
	}

}

// SetAuthConfig will set the cluster auth config in kubeconfig
func (c *clusterConfig) SetAuthConfig() {

	authProv := fmt.Sprintf("--auth-provider=oidc")
	authProvIdp := fmt.Sprintf("--auth-provider-arg=idp-issuer-url=%s", c.idpIssuerURL)
	authProvClientID := fmt.Sprintf("--auth-provider-arg=client-id=%s", c.clientID)
	authProvClientSecret := fmt.Sprintf("--auth-provider-arg=client-secret=%s", c.clientSecret)
	authProvIDToken := fmt.Sprintf("--auth-provider-arg=id-token=%s", c.idToken)
	authProvRefreshToken := fmt.Sprintf("--auth-provider-arg=refresh-token=%s", c.refreshToken)

	cmd := exec.Command("kubectl", "config", "set-credentials", c.userName,
		authProv,
		authProvIdp,
		authProvClientID,
		authProvClientSecret,
		authProvIDToken,
		authProvRefreshToken)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("error: cannot set kubectl OIDC credentials: %v", err)
	}
}
