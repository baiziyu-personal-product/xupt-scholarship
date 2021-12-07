#!/bin/bash
source /etc/profile


echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> start update loca repo"
git pull origin main

echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> local repo is lastet, start build go application"

go mod tidy && go build

echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OK "

