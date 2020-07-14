#!/usr/bin/env bash

declare -ri N="${N:-3}"
declare -r  TAG="${TAG:-latest}"

NETWORK=lachesis

LIMIT_CPU=$(echo "scale=2; 1/$N" | bc)
LIMIT_IO=$(echo "500/$N" | bc)

PORT_BASE=3000
RPCP_BASE=4000
WSP_BASE=4500

TEST_ACCS_START=1000
TEST_ACCS_COUNT=10000

PROMETHEUS_URL=http://prometheus:9090/
PROMETHEUS_JOB=node0

TARGET_HOST=node0:19090
ALERTMANAGER_HOST=alertmanager:9093

SMTP_HOST=localhost:25
SMTP_FROM=alert@lachesis.net
SMTP_AUTH_USER=
SMTP_AUTH_PASS=
EMAIL_ADMIN=admin@sample.com
