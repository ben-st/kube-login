# kube-login

fetch token from keycloak and update kubeconfig

## quick start

rename .kube-login.yml.example to ~/.kube-login.yml
and update the key value pairs according to your setup
e.g. username, password ...

## prerequisites

kubectl binary in $PATH, because kubeconfig patching relies on locally available kubectl

```
% kube-login
2019/03/29 20:34:59 Using config file: /Users/ben-st/.kube-login.yml
2019/03/29 20:34:59 username: admin@keycloak.devlocal
2019/03/29 20:34:59 cluster: minikube.dev
```

## Configuration

This supports the following options.

```

  kube-login [OPTIONS]

Application Options:
      --username or -u      = username for idp provider
      --password or -p      = password for idp provider
      --clustername or -c   = k8s api fqdn
      --port                = Port for k8s api (default: 6443)
      --idp-issuer-url      = Issuer URL of the provider
      --client-id           = Client ID of the provider
      --client-secret       = Client Secret of the provider
      --insecure-oidc  If set, the idp`s certificate will not be checked for validity, use with caution.
      --insecure-cluster If set, the k8s api certificate will not be checked for validity, use with caution.

Help Options:
  -h, --help        Show help message

```

All commandline parameters can be set in the kube-login config.

Key | Value
----|-----------|------
`username`          |   username for idp provider
`password`          |   password for idp provider
`clustername`       |   k8s api fqdn
`port`              |   k8s api port (default 6443)
`idp-issuer-url`    |   Issuer URL of the provider.
`client-id`         |   Client ID of the provider.
`client-secret`     |   Client Secret of the provider.
`insecure-oidc`     |   Insecure Connection to idp provider
`insecure-cluster`  |   Insecure Connection to k8s api

### configfile path

You can set the commandline parameter `config` to point the config file.
Defaults to `~/.kube-login.yml`.

### CA Certificates

you can skip ca certificate checks with `kube-login --insecure-oidc --insecure-cluster`
--insecure-oidc is for self signed certificates for oidc server and
--insecure-cluster is for self signed k8s api certificate

## Contributing

This is an open source software licensed under MIT License.
All coontributions are welcome

1. fork
2. create a feature branch
3. update the code
4. create a pull request