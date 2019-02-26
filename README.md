# Dockerized JetBrains License Server

Educative test of my local license server

# Prerequires

1. [Git](https://git-scm.com/downloads) 2+
2. [Docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/) 18+


# How to play

Clone

```
git clone https://github.com/D-Andreevich/docker-jetbrains-license-server.git
```

Inside

```
cd docker-jetbrains-license-server

copy authtoken_ngrok.example

get authtoken in https://dashboard.ngrok.com/auth

write key ngrok in authtoken_ngrok

sh run.sh
```

Get Container IP

```
docker inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}" my-license
```

```
172.17.0.2
```

or for who uses docker-machine

```
docker-machine ip default
```

```
192.168.99.100
```


Inspecting port 4040

```
http://172.17.0.2:4040
```

![](static/ip.png)

Registering on IDE at Help > Register...

![](static/dialog.png)


# References

1. [NGrok](https://ngrok.com/)

2. [GoLang](https://golang.org/)
