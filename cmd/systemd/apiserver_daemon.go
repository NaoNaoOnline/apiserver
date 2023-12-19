package systemd

const ApiserverDaemonService = `[Unit]
Description=apiserver daemon
After=redis.server.service
Requires=redis.server.service

[Service]
EnvironmentFile=/etc/environment
WorkingDirectory={{ .ApiServer.Directory }}
ExecStartPre=/bin/bash -c "until docker run --rm --network host redis:{{ .RedisServer.Version }} redis-cli ping; do sleep 5; done"
ExecStart=/usr/local/bin/apiserver daemon
ExecReload=/bin/kill -HUP $MAINPID
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
`
