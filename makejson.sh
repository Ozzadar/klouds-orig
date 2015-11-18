#!/bin/bash
cat > klouds-ui.json <<EOF
{
    "id":"klouds-ui",
    "cpus": 0.01,
    "mem": 50.0,
    "instances": 1,

    "container": {
      "type": "DOCKER",
      "docker": {
        "image": "$APP_IMAGE",
        "network": "BRIDGE",
        "portMappings": [
          {
            "containerPort": 80,
            "protocol": "tcp"
          }
        ],
        "privileged": true,
        "forcePullImage": true
      }
    },
    "env": {
        
    },
    "ports":[10000],
    "labels": {
        "HAPROXY_HTTP": "true",
        "HTTP_PORT_IDX_0_NAME": "klouds-ui"
    }
}
EOF