# file-encrypt

## run

```sh
$ cd $project_root
$ docker container run \
    -i -t --rm \
    --tmpfs /.cache \
    --network host \
    -v `pwd`:/work \
    golang-work:latest \
    bash
```
