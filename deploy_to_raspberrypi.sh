#!/bin/bash

MCP23017_source=`cat deploy.json | jq '. | .raspberry.source'`
MCP23017_build_name=`cat deploy.json | jq '. | .raspberry.build.name'`
MCP23017_target_user=`cat deploy.json | jq '. | .raspberry.target.user'`
MCP23017_target_host=`cat deploy.json | jq '. | .raspberry.target.host'`
MCP23017_target_path=`cat deploy.json | jq '. | .raspberry.target.path'`

MCP23017_source="${MCP23017_source//\"}"
MCP23017_build_name="${MCP23017_build_name//\"}"
MCP23017_target_user="${MCP23017_target_user//\"}"
MCP23017_target_host="${MCP23017_target_host//\"}"
MCP23017_target_path="${MCP23017_target_path//\"}"

echo "Compiling source $MCP23017_source ...";
GOOS=linux GOARCH=arm GOARM=5 go build -o $MCP23017_build_name $MCP23017_source

echo "Deploying to $MCP23017_target_host:$MCP23017_target_path ...";
scp $MCP23017_build_name $MCP23017_target_user@$MCP23017_target_host:$MCP23017_target_path

echo "Deployed!"
