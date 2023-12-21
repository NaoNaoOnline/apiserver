package systemd

const RedisServerService = `[Unit]
Description=redis server
After=docker.service
Requires=docker.service

[Service]
ExecStartPre=-/usr/bin/docker stop redis
ExecStartPre=-/usr/bin/docker rm redis
ExecStartPre=/usr/bin/docker pull redis:{{ .RedisServer.Version }}
ExecStart=/usr/bin/docker run --rm -p 127.0.0.1:6379:6379 -v {{ .RedisServer.Directory }}:/data/ --name redis redis:{{ .RedisServer.Version }} redis-server /data/redis.conf
Nice=-10
Restart=always
TimeoutStartSec=0

[Install]
WantedBy=multi-user.target
`
