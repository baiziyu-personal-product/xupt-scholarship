#!/bin/bash
source /etc/profile

if [ ! -d "./main.go" ];then
echo "main entry go file is not exit!!!"
else

echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> 开始更新仓库。"
git pull origin main

echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> 仓库更新完毕，开始编译环境。"

go mod tidy && go build

echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OK "
fi

