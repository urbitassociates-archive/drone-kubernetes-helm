# drone-kubernetes-helm

Drone plugin to deploy or update a project on Kubernetes using Helm. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Binary

Build the binary using `make`:

```
make deps build
```

### Example

```sh
./drone-kubernetes-helm <<EOF
{
      "repo": {
              "clone_url": "git://github.com/drone/drone",
              "owner": "drone",
              "name": "drone",
              "full_name": "drone/drone"
      },
      "system": {
              "link_url": "https://beta.drone.io"
      },
      "build": {
              "number": 22,
              "status": "success",
              "started_at": 1421029603,
              "finished_at": 1421029813,
              "message": "Update the Readme",
              "author": "johnsmith",
              "author_email": "john.smith@gmail.com",
              "event": "push",
              "branch": "master",
              "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
              "ref": "refs/heads/master"
      },
      "workspace": {
              "root": "/drone/src",
              "path": "/drone/src/github.com/drone/drone"
      },
      "vargs": {
        "config": {
          "kubeconfig": "",
          "credentials": {
            "certificate-authority": "",
            "client-certificate": "",
            "client-key": ""
          }
        },
        "commands": [
          {
            "install": {
              "chart": "ims-api",
              "release": "test-ims-api",
              "flags": [
                {
                  "dry-run": true
                }
              ]
            }
          }
        ]
      }
}
EOF
```

## Docker

Build the container using `make`:

```
make deps docker
```

### Example

```sh
docker run -i mandrean/drone-kubernetes-helm <<EOF
{
      "repo": {
              "clone_url": "git://github.com/drone/drone",
              "owner": "drone",
              "name": "drone",
              "full_name": "drone/drone"
      },
      "system": {
              "link_url": "https://beta.drone.io"
      },
      "build": {
              "number": 22,
              "status": "success",
              "started_at": 1421029603,
              "finished_at": 1421029813,
              "message": "Update the Readme",
              "author": "johnsmith",
              "author_email": "john.smith@gmail.com",
              "event": "push",
              "branch": "master",
              "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
              "ref": "refs/heads/master"
      },
      "workspace": {
              "root": "/drone/src",
              "path": "/drone/src/github.com/drone/drone"
      },
      "vargs": {
        "config": {
          "kubeconfig": "",
          "credentials": {
            "certificate-authority": "",
            "client-certificate": "",
            "client-key": ""
          }
        },
        "commands": [
          {
            "install": {
              "chart": "ims-api",
              "release": "test-ims-api",
              "flags": [
                {
                  "dry-run": true
                }
              ]
            }
          }
        ]
      }
}
EOF
```
