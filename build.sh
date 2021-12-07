#!/bin/bash
source /etc/profile

echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> start "

cd /project/xupt-scholarship/ && git pull && go mod tidy && go build

echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> finish "

