#!/bin/bash

echo 'copy authtoken_ngrok'
echo "web_addr: 0.0.0.0:4040 
authtoken: "$(grep -E '(\b^\w{43}$\b)' authtoken_ngrok) > data/ngrok.yml

cd data

echo -n "Консоль показываем?"

read -t 10 -p "Press ENTER or wait ten seconds (y/n)? " item
case "$item" in
    y|Y) echo "Ввели «y», продолжаем..."
        docker-compose up --force-recreate jetbrains-license-server 
        ;;
    n|N) echo "Ввели «n», завершаем..."
        docker-compose up -d --force-recreate jetbrains-license-server 
        ;;
    *) echo "Ничего не ввели. Выполняем действие по умолчанию..."
        docker-compose up -d --force-recreate jetbrains-license-server 
        ;;
esac
