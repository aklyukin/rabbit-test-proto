#!/bin/bash
protoc --go_out=plugins=grpc:. master/*.proto

protoc --go_out=plugins=grpc:. node/*.proto