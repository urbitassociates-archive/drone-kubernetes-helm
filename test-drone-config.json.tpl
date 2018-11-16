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
      "kubeconfig": "{{KUBE_CONFIG}}",
      "credentials": {
        "certificate-authority": "{{KUBE_CA}}",
        "client-certificate": "{{KUBE_CLIENT_CERT}}",
        "client-key": "{{KUBE_CLIENT_KEY}}"
      }
    },
    "commands": [
      {
        "repo": {
          "command": "add",
          "args": [
            "urbit",
            "http://charts.urb-it.io"
          ]
        }
      },
      {
        "install": {
          "chart": "urbit/developer-urbit-com",
          "release": "testing-drone-kubernetes-helm-plugin",
          "flags": [
            {
              "dry-run": "true",
              "namespace": "test"
            }
          ]
        }
      }
    ]
  }
}
