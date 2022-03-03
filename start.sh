#!/bin/bash

usage() {
    echo "Setup and start wager-be"
    echo
    echo "Usage:"
    echo "  start.sh [command]"
    echo 
    echo "Available commands:"
    echo "  bootstrap   Setup all dependencies"
    echo "  start       Start wager-be server"
}

bootstrap() {
    echo "----- SETUP DOCKER -----"
    make docker-down
    make docker-build-local
    make docker-up
}

start() {
    echo "----- START SERVER -----"
    make docker-up-db
    make build
    . ./local.env
    ./bin/wager-be server
}

if [[ $# == 0 ]]; then
    usage
    exit
fi

while [[ $# -gt 0 ]]; do 
    key="$1"
    case $key in
        build)              build
                            shift
                            ;;
        bootstrap)          bootstrap
                            shift
                            ;;
        start)              start
                            shift
                            ;;
        -h | --help)        usage
                            exit
                            ;;
        *)                  usage
                            exit 1
    esac
    shift
done