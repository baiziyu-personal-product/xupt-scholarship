#!/bin/bash
source /etc/profile

echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> start update ECS server code and build for your service "

cd /project/xupt-scholarship/ && git pull  && go build && supervisorctl restart xupt

echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> finish all works"

