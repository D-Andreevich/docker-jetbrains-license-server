FROM golang:alpine

RUN apk update && apk add unzip curl jq dos2unix

# Install ngrok (latest official stable from https://ngrok.com/download).
ADD https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip /ngrok.zip
RUN set -x \
 && unzip -o /ngrok.zip -d /bin \
 && rm -f /ngrok.zip

# Add config script.
COPY ngrok.yml /home/ngrok/.ngrok2/
COPY entrypoint.sh /

COPY main.go /home/ngrok
COPY static /home/ngrok/static


# Create non-root user.
RUN set -x \
 && echo 'ngrok:x:6737:6737:Ngrok user:/home/ngrok:/bin/false' >> /etc/passwd \
 && echo 'ngrok:x:6737:' >> /etc/group \
 && chown -R ngrok:ngrok /home/ngrok \
 && chmod -R go=u,go-w /home/ngrok \
 && chmod go= /home/ngrok \
 && dos2unix /entrypoint.sh \
 && chmod +x /entrypoint.sh

USER ngrok
WORKDIR /home/ngrok

EXPOSE 4040 1234

CMD ["/entrypoint.sh"]
