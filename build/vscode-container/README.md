# docker build

```sh
$ SUDO_USER=$(/usr/bin/logname)
$ docker build \
    --build-arg UID=$(id -u $SUDO_USER) \
    -t golang-work:latest .
```
