package systemd

const CaddyProxyService = `[Unit]
Description=caddy proxy
After=docker.service
Requires=docker.service

[Service]
ExecStartPre=-/usr/bin/docker stop caddy
ExecStartPre=-/usr/bin/docker rm caddy
ExecStartPre=/usr/bin/docker pull caddy:{{ .CaddyProxy.Version }}
ExecStart=/usr/bin/docker run --rm --cap-add=NET_ADMIN --network=host -p 0.0.0.0:80:80 -p 0.0.0.0:443:443 -p 0.0.0.0:443:443/udp -v {{ .CaddyProxy.Directory }}:/data/ --name caddy caddy:{{ .CaddyProxy.Version }} caddy reverse-proxy --from api.naonao.io:443 --to 127.0.0.1:7777
Restart=always
TimeoutStartSec=0

[Install]
WantedBy=multi-user.target`
