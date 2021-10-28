#!/bin/bash

protoc greet/greetpb/greet.proto --go_out=prugins=grpc:.
protoc calculator/calculatorpb/calculator.proto --go_out=prugins=grpc:.

