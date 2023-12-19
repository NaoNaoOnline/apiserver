package systemd

const UserData = `#cloud-config

apt:
  sources:
    docker.list:
      source: deb [arch=amd64] https://download.docker.com/linux/ubuntu $RELEASE stable
      keyid: 9DC858229FC7DD38854AE2D88D81803C0EBFCD88

packages:
  - ca-certificates
  - curl
  - docker-ce
  - docker-ce-cli

write_files:
  - path: /etc/systemd/journald.conf.d/systemMaxUse.conf
    content: |
      [Journal]
      SystemMaxUse=100M
    permissions: "0644"

groups:
  - docker

system_info:
  default_user:
    groups: [docker, sudo]
    home: /home/ubuntu
    name: ubuntu

runcmd:
  - curl --location https://github.com/NaoNaoOnline/apiserver/releases/download/{{ .ApiServer.Version }}/apiserver-linux-amd64 --output /usr/local/bin/apiserver
  - chmod +x /usr/local/bin/apiserver

  - /usr/local/bin/apiserver systemd`
