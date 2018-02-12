#!/bin/bash
## Created by Anthony Juan               ##
## anthonyjuan.github.io                 ##
########################################### 


set -e

Package=( 
		"github.com/astaxie/beego"
		"github.com/smartystreets/goconvey/convey"
		"github.com/mattn/go-sqlite3"
    "github.com/beego/bee"
		)

arrayPackage=${#Package[@]}

for ((k=0; k<${arrayPackage}; k++));
do
	echo "GO GET ${Package[$k]}"
	go get ${Package[$k]} 2>&1
done