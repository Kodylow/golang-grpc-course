#!/bin/bash

protoc greet/greetpb/greet.proto --go_out=prugins=grpc:.
