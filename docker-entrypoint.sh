#!/bin/sh

easy-gate /etc/easy-gate/easy-gate.json &
nginx -g "daemon off;"