# interval-runner

指定されたコマンドを定期的に実行するDockerコンテナ
Docker container that periodically executes specified commands

## Usage

```
docker run hiko/interval-runner:1.0.0 echo hello
```

### with docker client

Create Dockerfile like this...

```
FROM hiko/interval-runner:1.0.0

# Install docker client
# https://docs.docker.com/engine/install/debian/

RUN apt-get update -y && \
    apt-get install -y ca-certificates curl gnupg lsb-release
RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
RUN echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" \
  | tee /etc/apt/sources.list.d/docker.list > /dev/null
RUN apt-get update && apt-get install -y docker-ce-cli
VOLUME ["/var/run/docker.sock"]
```

build...

```
docker build -t docker-runner .
```

and run.

```
docker run --rm --name docker-runner
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v $PWD/letsencrypt/log:/var/log/letsencrypt \
  -v $PWD/letsencrypt/etc:/etc/letsencrypt \
  -v $PWD/public:/var/www/html \
  docker-runner \
    --interval 168h \
    docker run --rm --name certbot \
      -v /var/log/letsencrypt:/var/log/letsencrypt \
      -v /etc/letsencrypt:/etc/letsencrypt \
      -v /var/www/html:/var/www/html \
      certbot/certbot:v1.26.0 certonry --webroot -w /var/www/html -d example.com
```
