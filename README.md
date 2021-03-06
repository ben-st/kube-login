# kube-login

kube-login will fetch an id-token and a resfresh-token from an oidc provider(tested against keycloak) and write it to the kubeconfig file with the idp config. It will also create a context for the given clustername.
After kube-login, even when no kubeconfig file exists, you will be able to access the k8s cluster api with kubectl.

## prerequisites

kubectl binary in $PATH, because kubeconfig patching relies on locally available kubectl.
Here is room for improvement, but you will need local kubectl binary anyway.

## quick start

move the kube-login.yml or toml to your homedir:

`mv .kube-login.yml ~/.kube-login.yml`

and update the key value pairs according to your setup
e.g. username, password , clustername ...

``` bash
kube-login
2019/03/29 20:34:59 Using config file: /Users/ben-st/.kube-login.yml
2019/03/29 20:34:59 username: admin@keycloak.devlocal
2019/03/29 20:34:59 cluster: minikube.dev
```

if you just want to show your id- and refresh token you can use the --show-token flag

`kube-login --show-token`

## Configuration

This supports the following options.

```

login to keycloak and generate/update kubeconfig with id and refresh token

Usage:
  kube-login [flags]

Flags:
      --clientid string         clientid for idp
      --clientsecret string     client secret for idp
  -c, --clustername string      clustername fqdn e.g api.kubernetes.example
      --config string           config file (default is $HOME/.kube-login.yml)
  -h, --help                    help for kube-login
      --idp-issuer-url string   idp/oidc fqdn
      --insecure-cluster        if set insecure tls to cluster in kubeconfig will be set, use with caution (default true)
      --insecure-oidc           if set insecure tls to oidc provider will be used, use with caution
  -p, --password string         password for keycloak
      --port int                port for apiserver (default 6443)
      --show-token              show keycloak token and exit
  -t, --toggle                  Help message for toggle
  -u, --username string         username for keycloak

```

All commandline parameters can be set in the kube-login config.
See .kube-login.yml-example for reference.

```yaml
username: "admin@keycloak.devlocal"
password: "admin"
clustername: "minikube.dev"
port: 8443
idp-issuer-url: "https://keycloak.devlocal/auth/realms/k8s"
clientid: "oidckube"
clientsecret: "797a53e4-5824-49f9-a71a-bf1e2a50404f"
insecure-oidc: false
insecure-cluster: true
```

## ENV Variables

You can set variables like clustername oder password as env variables.
You have to export them with the `KUBE_` prefix and in **UPPERCASE**

`export KUBE_USERNAME=foo`

`export KUBE_PASSWORD=bar`

### configfile path

You can set the commandline parameter `--config` to point the config file.
Defaults to `~/.kube-login.yml`.

you can use yaml or toml

### CA Certificates

you can skip ca certificate checks with `kube-login --insecure-oidc --insecure-cluster`
--insecure-oidc is for self signed certificates for oidc server and
--insecure-cluster is for self signed k8s api certificate

## development

if you want to build it yourself just download with `go get`

`go get github.com/ben-st/kube-login`

you can build it with just `go build`

i tested against keycloak in minikube with the help of:
<https://github.com/mrbobbytables/oidckube>

## Contributing

This is an open source software licensed under MIT License.
All contributions are welcome

1. fork
2. create a feature branch
3. update the code
4. create a pull request