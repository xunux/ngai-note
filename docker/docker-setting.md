# docker registry mirror address setting

vim /etc/docker/daemon.json

```
{
  "registry-mirrors": [
    "https://hub-mirror.c.163.com",
    "https://ccr.ccs.tencentyun.com"
  ]
}

```

then restart docker service use `systemctl restart docker.service`

at last, verify docker mirror use `docker info` command


