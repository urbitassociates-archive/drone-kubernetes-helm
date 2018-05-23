Probably not used, we use the official version of this one. //Simon Forsman, 2018-05-23

# drone-kubernetes-helm

Drone plugin to deploy or update a project on Kubernetes using Helm. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Vargs

API server IP or hostname:
```
    config:
      kubeconfig:
        api_server:
```
Authentication (only X509/client certificate for now):
```
      authentication:
        x509:
          certificate_authority: <CA.crt base64 string> // cat CA.crt | base64
          client_certificate: <client.crt base64 string> // cat client.crt | base64
          client_key: <client.key base64 string> // cat client.key | base64
```
Commands. Tries to follow the Helm CLI command names and flags closely:
```
    commands:
    - install:
        chart: nginx
        release: dev-nginx
        flags:
        - dry-run: true
    - get:
        subcommand: values
        release: dev-nginx
        flags:
        - all: true
        - output: ./values-dev.yaml
    - upgrade:
        chart: nginx
        release: dev-nginx
        flags:
        - namespace: default
        - values: ./values-dev.yaml
        - set: imageTag=1.10-alpine
```

## Example .drone.yaml
```
deploy:
  kubernetes-helm:
    image: mandrean/drone-kubernetes-helm
    when:
      event: push
      branch: master
    config:
      kubeconfig:
        api_server: https://192.168.99.100:8443
      authentication:
        x509:
          certificate_authority: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM2RENDQWRDZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFsTVJFd0R3WURWUVFLRXdocmRXSmwKTFdGM2N6RVFNQTRHQTFVRUF4TUhhM1ZpWlMxallUQWVGdzB4TmpFeE1EY3hNVEV4TVRCYUZ3MHlOakV4TURVeApNVEV4TVRCYU1DVXhFVEFQQmdOVkJBb1RDR3QxWW1VdFlYZHpNUkF3RGdZRFZRUURFd2RyZFdKbExXTmhNSUlCCklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUFzQWQ3WGxiQ3c3YU9tR1Z1anhDSzIwMzIKb1FiMGg0MC9jZWN1MUxObWJ0OG5ON1h1NHNuSXdpWEplbEg5K0d1N0duQy9ha3R1ei9JYUFWYlZzNmJRdElVVgpiRVQxQ25GcEg5UGR2MFZiNjRGS2U1TFc3ZFNtYlJsSzVDazhVQTB2dVlTZmhvakJSTE5uakFGNmMyRjNJd2FtCnBGTGxoc3plb3JOZThCeWVnQUN2Q2NuZS8zTGdkY3ptL0FMSGlZdGc5YU5pNTIxbjg5QW5rYWdXNlplTU1rRXYKbXUwZDRReFJMOUkvcW9tQ3ZlK2hUWXZuRy92dm5pQmNHaWRFdTYydnU2eGRGNmY1NkovK0Y3U2c0ZEZaYjhlbApKcnpFUzVMREZMOWJ2RE8zVGEvdjcvbWV5L1k5bkhTQ3FyUnVGdFV4VnFvZkdHM05sSnFNSFl5cnRpeTdhd0lECkFRQUJveU13SVRBT0JnTlZIUThCQWY4RUJBTUNBcVF3RHdZRFZSMFRBUUgvQkFVd0F3RUIvekFOQmdrcWhraUcKOXcwQkFRc0ZBQU9DQVFFQU5ieFdRK1c0REgrOU51bGhLeExONmlXMDBxU3JEZkdRREJuZ1VudjBxR1FDRUpVTAo5VTNrZ0Z1WjhFZXJGU085SXBvOXVDdEI3MnE1NkZyOVJjd1VuV0M2Z2kwTDVPbTlKT3owbGl0Q2RtR1cwbXRsCmM0dEFTM0RGc1FaM3ZTdmZZV21mcDFrcUFqdGFSYWJ4NEpkSVcrOTZ6TUxobUttQ0VqZU5GRXhtdEpVeWRJZ2MKSTJJK3hOZnBlK2l1TU5ib1NPTm90aU9RcjdZVTVJNTlZUlYzYms5b2FxcG9JSVU1cmduZW1KMWRWWkQzZ2szNAorY0dkUnVJZFhUbjBJdHZNeWRUUnFTcVB2enQvUk5NQlFaTW5RK25wZCtneWNBbE12eGE2cDJhWlljUzF6akhUCkU3dlROSTQ2Mk0vdHRMTjQvTUdiSWJMaHQ4eWRkc0tIV01aWlh3PT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
          client_certificate: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNzekNDQVpzQ0NRRHQ2UkxWSTE2T1JEQU5CZ2txaGtpRzl3MEJBUVVGQURBbE1SRXdEd1lEVlFRS0V3aHIKZFdKbExXRjNjekVRTUE0R0ExVUVBeE1IYTNWaVpTMWpZVEFlRncweE5qRXlNRGd4T1RNNE16bGFGdzB4TnpFeQpNRGd4T1RNNE16bGFNQkl4RURBT0JnTlZCQU1UQjNWeVltbDBZMmt3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBCkE0SUJEd0F3Z2dFS0FvSUJBUUM1Q1NOeW9hUk1mOHVMcDE3eW1oaHY2ZnlZb1lEZGg4K3UwVmdxbTVYL1lJd3IKUVJydnBRYTNDRFdFS0RQZVE1eVE1cExBTlBQdlJvZ3d5M2t5cW93em5OSmJqR01XaVBvWVFTcEloS1FNSlVTdQpRV1AyNlNEQ2ZRbW5jSnVSVG82MGpxeWxaai9lV1NnMFd4K0dkTmlUR1MwVEVrbVA0SUtLcWFIeGFTbkVDeW91CkZHZ1pTN0QrcER2T1FCM0FKalMwdVFoemxZWWFkMWRzSEZNbzFZdXJ1OWtBQ0VVekYrNUppTTNQSlBvR1JvcFIKL05WYzM0RDc1VjB4aWpxWlJBbUxwR2dvTHN0Sis3dUxDNGxEV0IzemN1RXRqNmtiNUtZRWoyK3N0QWc4M3BLSgp3WG9lSlNOZXB2YmVjc1ZoL0NjWVNiamNNaEFOQ2s1WkJISHpDV09MQWdNQkFBRXdEUVlKS29aSWh2Y05BUUVGCkJRQURnZ0VCQUFRWkZjOU83NmZXRGFVOUFONk1wYVBBOXF4UUEyL3RPVk9XRWJzUWZWbks1T0s0SmlkNVk5LzAKbGFVWlNLTC80TlRQVVVUZ1NPajJCQVBoM05JdnVBdFQyQ1VYRG5FaDZKUExxclZsZVEyQ2tPTFBMc2JvdEZhYwplNGtIOFlyL1cydjZKM21HOUhCUEJJaWVsbVc0NVc0Nng1djF5UWlMcElHdjVuMW5uKy9VR2swbDlISmd0bHgrCm9NekYvVHpWNHMzdFNzUzFrREg5aytuakJYNlh5d3ZDTkoydGsxUmNPZVB2SW5xdjdyYTN2NExNM1lTWEV4czYKSFRVTHBHNVFndWt1RkZpTUtBS21SMitIeGFNMmVQQXkxekZxYWdlQUZVbHN0UXhQNVQ5ZEUyeVo5czlzU0lzNgpwVW9JQys3V0dQTmp0S3BvdzJyKzBQOGFBdFZkWEhNPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
          client_key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBdVFramNxR2tUSC9MaTZkZThwb1liK244bUtHQTNZZlBydEZZS3B1Vi8yQ01LMEVhCjc2VUd0d2cxaENnejNrT2NrT2FTd0RUejcwYUlNTXQ1TXFxTU01elNXNHhqRm9qNkdFRXFTSVNrRENWRXJrRmoKOXVrZ3duMEpwM0Nia1U2T3RJNnNwV1kvM2xrb05Gc2ZoblRZa3hrdEV4SkpqK0NDaXFtaDhXa3B4QXNxTGhSbwpHVXV3L3FRN3prQWR3Q1kwdExrSWM1V0dHbmRYYkJ4VEtOV0xxN3ZaQUFoRk14ZnVTWWpOenlUNkJrYUtVZnpWClhOK0ErK1ZkTVlvNm1VUUppNlJvS0M3TFNmdTdpd3VKUTFnZDgzTGhMWStwRytTbUJJOXZyTFFJUE42U2ljRjYKSGlValhxYjIzbkxGWWZ3bkdFbTQzRElRRFFwT1dRUng4d2xqaXdJREFRQUJBb0lCQVFDY0dzT1hJQnUybGxJbwp2Y2x5cnVKUytIcXNZZ1NQNE5ZcnpGMnZoSmRsWGhTaklVZ1NTWTJDdVNBOUlKV3h2Q1RJY2wzNFhqUTE5N0ZLClNUODBxWUdpd1hrTzF6OTVjWkpkQ0EwZUpSa3BUZi9GYTFGa3E0V0J6MjlubmE4QlJkOUxJTnN5cHpMVzZTenMKRHJ4bitRZ0dBY1Z2UTR4Z0g2N0NRUjVveHFuL3cyUkRpenVWQ1lUMUpqYzE3dlgxM2hDbzVsNkgrUmdBL2VhRApBaGNrdTBLRzZwa0R3WUE5L2k1Tlg0VGc2ZElEN3ZXU01hSDluUi9WRExxcmV2YnZnSG5PT0xseUV4R3RJL3NGCnFucVFxTGI0eW4yMTY1V2xkdlJwNUk2akNJekJINVZqKzgrOEVFRXU4SVB5cEhDd3cyT2xzV0FOMnZURGxnQm4KajZ6MVl0V1JBb0dCQU4zSjVMTWpXeEIwUys4OGZDamNHQzFjSGVwYm5pckhXYXcrR3ltbytBOVpKMlV2UU9nNApzMndIQmxhcXdNaU9UU0dmbUtxd0t1ZnJEM3J2WC9BNmlVWTYyRjFQdWZRK2dNTUFrbU4yemk3dlYxSFdCbnN3CjJGcnpiTmNpRGhTU3BCREFTVk12YVhXNWdQK2hBVFpjeXlnWjVDTXUrTUlqakllQ3JjWlk4aHA5QW9HQkFOV1QKN09PTVAxbEFkZ3FaUE5tUFdkNHBSLzJrNGFRNjl5OXdtZ0NIUEFid1M2WUY1d3JsditEa2UxTVhCb3lVM3Z1RApVelhXTkczZHBzSjNqV1JWYXB2NzZqM2JUQU0rUE1tVmNJLzBFS0E3Wk5mazZtTG81L2M5TmFIc2xFMkFFd0p6CmUzTTF2YW5hSld5eEhKdXNhYkhjS0NtYWlvMFFoMjlwd3ZkNXdVeW5Bb0dBQzhsK0VSTXc5TWZwZlRadXR0RXoKcTcxNGpZcis5ZkVRVC9vaEFXN01lQ3haenFQYlJEdzNOT2VPcTY1NWZtOHBwRDdTSTBnbmo1bkxnZElVL0RSdwpOVDVOWDNBc1J0SEhrQldJc2lhUFFLbFJyN1M3TlhMY0hNRlJLSUhUMDc0VFlCeUlDUmE4K1JlNXhsd3RMMUZ4CkxwbHBxWUVHa1hMSU5pOTR3dERaVlJVQ2dZQlV3UUgxZnBjNC9OcWE5QnB3bjNGak52Q0ptQit1dzNPS0VONGMKTFk1RmxwLytmME1qVU83bStPUnpvYVNJcng4Wm9oQ29RWnZHcVhuZW5BQ3crekIyTys3Rm96dXo0Y1BQbnd6dgpJMFJod1pBUUdKaG1yZFEzaWNPNXdSOU03ZkVkUE9TVllKTW1UeG9nMnR2bWJ2SDJrYzRpVEdDRkFEVXVva0tyCllGYXo2d0tCZ0Q3QVRtWVZrOWgxTkdLdlgrelh4dDhseTFnM2JmTkpkbTh1N050Q3JyY1doQ0U3ZHRkWTdwOHUKVEd6QUpudVNRZU1BODRwdWVNNzlZQWVIZHdmVEQxZFd6SU5BNjhVVUd6ZTNST1JySEpzaTQvZVIvSGZpNnJlZgpyMlZ2elJieDBTSHd0YWpwRHR3eHY5NmZsSHNtZG1Ld2t2Z2E5NXZEQ3M1YVBOOHpqVE51Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    commands:
    - install:
        chart: nginx
        release: dev-nginx
        flags:
        - dry-run: true
    - get:
        subcommand: values
        release: dev-nginx
        flags:
        - all: true
        - output: ./values-dev.yaml
    - upgrade:
        chart: nginx
        release: dev-nginx
        flags:
        - namespace: default
        - values: ./values-dev.yaml
        - set: imageTag=1.10-alpine

```

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
          "kubeconfig": {
            "api_server": "https://192.168.99.100:8443"
          },
          "authentication": {
            "x509": {
              "certificate_authority": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM2RENDQWRDZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFsTVJFd0R3WURWUVFLRXdocmRXSmwKTFdGM2N6RVFNQTRHQTFVRUF4TUhhM1ZpWlMxallUQWVGdzB4TmpFeE1EY3hNVEV4TVRCYUZ3MHlOakV4TURVeApNVEV4TVRCYU1DVXhFVEFQQmdOVkJBb1RDR3QxWW1VdFlYZHpNUkF3RGdZRFZRUURFd2RyZFdKbExXTmhNSUlCCklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUFzQWQ3WGxiQ3c3YU9tR1Z1anhDSzIwMzIKb1FiMGg0MC9jZWN1MUxObWJ0OG5ON1h1NHNuSXdpWEplbEg5K0d1N0duQy9ha3R1ei9JYUFWYlZzNmJRdElVVgpiRVQxQ25GcEg5UGR2MFZiNjRGS2U1TFc3ZFNtYlJsSzVDazhVQTB2dVlTZmhvakJSTE5uakFGNmMyRjNJd2FtCnBGTGxoc3plb3JOZThCeWVnQUN2Q2NuZS8zTGdkY3ptL0FMSGlZdGc5YU5pNTIxbjg5QW5rYWdXNlplTU1rRXYKbXUwZDRReFJMOUkvcW9tQ3ZlK2hUWXZuRy92dm5pQmNHaWRFdTYydnU2eGRGNmY1NkovK0Y3U2c0ZEZaYjhlbApKcnpFUzVMREZMOWJ2RE8zVGEvdjcvbWV5L1k5bkhTQ3FyUnVGdFV4VnFvZkdHM05sSnFNSFl5cnRpeTdhd0lECkFRQUJveU13SVRBT0JnTlZIUThCQWY4RUJBTUNBcVF3RHdZRFZSMFRBUUgvQkFVd0F3RUIvekFOQmdrcWhraUcKOXcwQkFRc0ZBQU9DQVFFQU5ieFdRK1c0REgrOU51bGhLeExONmlXMDBxU3JEZkdRREJuZ1VudjBxR1FDRUpVTAo5VTNrZ0Z1WjhFZXJGU085SXBvOXVDdEI3MnE1NkZyOVJjd1VuV0M2Z2kwTDVPbTlKT3owbGl0Q2RtR1cwbXRsCmM0dEFTM0RGc1FaM3ZTdmZZV21mcDFrcUFqdGFSYWJ4NEpkSVcrOTZ6TUxobUttQ0VqZU5GRXhtdEpVeWRJZ2MKSTJJK3hOZnBlK2l1TU5ib1NPTm90aU9RcjdZVTVJNTlZUlYzYms5b2FxcG9JSVU1cmduZW1KMWRWWkQzZ2szNAorY0dkUnVJZFhUbjBJdHZNeWRUUnFTcVB2enQvUk5NQlFaTW5RK25wZCtneWNBbE12eGE2cDJhWlljUzF6akhUCkU3dlROSTQ2Mk0vdHRMTjQvTUdiSWJMaHQ4eWRkc0tIV01aWlh3PT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=",
              "client_certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNzekNDQVpzQ0NRRHQ2UkxWSTE2T1JEQU5CZ2txaGtpRzl3MEJBUVVGQURBbE1SRXdEd1lEVlFRS0V3aHIKZFdKbExXRjNjekVRTUE0R0ExVUVBeE1IYTNWaVpTMWpZVEFlRncweE5qRXlNRGd4T1RNNE16bGFGdzB4TnpFeQpNRGd4T1RNNE16bGFNQkl4RURBT0JnTlZCQU1UQjNWeVltbDBZMmt3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBCkE0SUJEd0F3Z2dFS0FvSUJBUUM1Q1NOeW9hUk1mOHVMcDE3eW1oaHY2ZnlZb1lEZGg4K3UwVmdxbTVYL1lJd3IKUVJydnBRYTNDRFdFS0RQZVE1eVE1cExBTlBQdlJvZ3d5M2t5cW93em5OSmJqR01XaVBvWVFTcEloS1FNSlVTdQpRV1AyNlNEQ2ZRbW5jSnVSVG82MGpxeWxaai9lV1NnMFd4K0dkTmlUR1MwVEVrbVA0SUtLcWFIeGFTbkVDeW91CkZHZ1pTN0QrcER2T1FCM0FKalMwdVFoemxZWWFkMWRzSEZNbzFZdXJ1OWtBQ0VVekYrNUppTTNQSlBvR1JvcFIKL05WYzM0RDc1VjB4aWpxWlJBbUxwR2dvTHN0Sis3dUxDNGxEV0IzemN1RXRqNmtiNUtZRWoyK3N0QWc4M3BLSgp3WG9lSlNOZXB2YmVjc1ZoL0NjWVNiamNNaEFOQ2s1WkJISHpDV09MQWdNQkFBRXdEUVlKS29aSWh2Y05BUUVGCkJRQURnZ0VCQUFRWkZjOU83NmZXRGFVOUFONk1wYVBBOXF4UUEyL3RPVk9XRWJzUWZWbks1T0s0SmlkNVk5LzAKbGFVWlNLTC80TlRQVVVUZ1NPajJCQVBoM05JdnVBdFQyQ1VYRG5FaDZKUExxclZsZVEyQ2tPTFBMc2JvdEZhYwplNGtIOFlyL1cydjZKM21HOUhCUEJJaWVsbVc0NVc0Nng1djF5UWlMcElHdjVuMW5uKy9VR2swbDlISmd0bHgrCm9NekYvVHpWNHMzdFNzUzFrREg5aytuakJYNlh5d3ZDTkoydGsxUmNPZVB2SW5xdjdyYTN2NExNM1lTWEV4czYKSFRVTHBHNVFndWt1RkZpTUtBS21SMitIeGFNMmVQQXkxekZxYWdlQUZVbHN0UXhQNVQ5ZEUyeVo5czlzU0lzNgpwVW9JQys3V0dQTmp0S3BvdzJyKzBQOGFBdFZkWEhNPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==",
              "client_key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBdVFramNxR2tUSC9MaTZkZThwb1liK244bUtHQTNZZlBydEZZS3B1Vi8yQ01LMEVhCjc2VUd0d2cxaENnejNrT2NrT2FTd0RUejcwYUlNTXQ1TXFxTU01elNXNHhqRm9qNkdFRXFTSVNrRENWRXJrRmoKOXVrZ3duMEpwM0Nia1U2T3RJNnNwV1kvM2xrb05Gc2ZoblRZa3hrdEV4SkpqK0NDaXFtaDhXa3B4QXNxTGhSbwpHVXV3L3FRN3prQWR3Q1kwdExrSWM1V0dHbmRYYkJ4VEtOV0xxN3ZaQUFoRk14ZnVTWWpOenlUNkJrYUtVZnpWClhOK0ErK1ZkTVlvNm1VUUppNlJvS0M3TFNmdTdpd3VKUTFnZDgzTGhMWStwRytTbUJJOXZyTFFJUE42U2ljRjYKSGlValhxYjIzbkxGWWZ3bkdFbTQzRElRRFFwT1dRUng4d2xqaXdJREFRQUJBb0lCQVFDY0dzT1hJQnUybGxJbwp2Y2x5cnVKUytIcXNZZ1NQNE5ZcnpGMnZoSmRsWGhTaklVZ1NTWTJDdVNBOUlKV3h2Q1RJY2wzNFhqUTE5N0ZLClNUODBxWUdpd1hrTzF6OTVjWkpkQ0EwZUpSa3BUZi9GYTFGa3E0V0J6MjlubmE4QlJkOUxJTnN5cHpMVzZTenMKRHJ4bitRZ0dBY1Z2UTR4Z0g2N0NRUjVveHFuL3cyUkRpenVWQ1lUMUpqYzE3dlgxM2hDbzVsNkgrUmdBL2VhRApBaGNrdTBLRzZwa0R3WUE5L2k1Tlg0VGc2ZElEN3ZXU01hSDluUi9WRExxcmV2YnZnSG5PT0xseUV4R3RJL3NGCnFucVFxTGI0eW4yMTY1V2xkdlJwNUk2akNJekJINVZqKzgrOEVFRXU4SVB5cEhDd3cyT2xzV0FOMnZURGxnQm4KajZ6MVl0V1JBb0dCQU4zSjVMTWpXeEIwUys4OGZDamNHQzFjSGVwYm5pckhXYXcrR3ltbytBOVpKMlV2UU9nNApzMndIQmxhcXdNaU9UU0dmbUtxd0t1ZnJEM3J2WC9BNmlVWTYyRjFQdWZRK2dNTUFrbU4yemk3dlYxSFdCbnN3CjJGcnpiTmNpRGhTU3BCREFTVk12YVhXNWdQK2hBVFpjeXlnWjVDTXUrTUlqakllQ3JjWlk4aHA5QW9HQkFOV1QKN09PTVAxbEFkZ3FaUE5tUFdkNHBSLzJrNGFRNjl5OXdtZ0NIUEFid1M2WUY1d3JsditEa2UxTVhCb3lVM3Z1RApVelhXTkczZHBzSjNqV1JWYXB2NzZqM2JUQU0rUE1tVmNJLzBFS0E3Wk5mazZtTG81L2M5TmFIc2xFMkFFd0p6CmUzTTF2YW5hSld5eEhKdXNhYkhjS0NtYWlvMFFoMjlwd3ZkNXdVeW5Bb0dBQzhsK0VSTXc5TWZwZlRadXR0RXoKcTcxNGpZcis5ZkVRVC9vaEFXN01lQ3haenFQYlJEdzNOT2VPcTY1NWZtOHBwRDdTSTBnbmo1bkxnZElVL0RSdwpOVDVOWDNBc1J0SEhrQldJc2lhUFFLbFJyN1M3TlhMY0hNRlJLSUhUMDc0VFlCeUlDUmE4K1JlNXhsd3RMMUZ4CkxwbHBxWUVHa1hMSU5pOTR3dERaVlJVQ2dZQlV3UUgxZnBjNC9OcWE5QnB3bjNGak52Q0ptQit1dzNPS0VONGMKTFk1RmxwLytmME1qVU83bStPUnpvYVNJcng4Wm9oQ29RWnZHcVhuZW5BQ3crekIyTys3Rm96dXo0Y1BQbnd6dgpJMFJod1pBUUdKaG1yZFEzaWNPNXdSOU03ZkVkUE9TVllKTW1UeG9nMnR2bWJ2SDJrYzRpVEdDRkFEVXVva0tyCllGYXo2d0tCZ0Q3QVRtWVZrOWgxTkdLdlgrelh4dDhseTFnM2JmTkpkbTh1N050Q3JyY1doQ0U3ZHRkWTdwOHUKVEd6QUpudVNRZU1BODRwdWVNNzlZQWVIZHdmVEQxZFd6SU5BNjhVVUd6ZTNST1JySEpzaTQvZVIvSGZpNnJlZgpyMlZ2elJieDBTSHd0YWpwRHR3eHY5NmZsSHNtZG1Ld2t2Z2E5NXZEQ3M1YVBOOHpqVE51Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="
            }
          }
        },
        "commands": [
          {
            "install": {
              "chart": "nginx",
              "release": "dev-nginx",
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
          "kubeconfig": {
            "api_server": "https://192.168.99.100:8443"
          },
          "authentication": {
            "x509": {
              "certificate_authority": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM2RENDQWRDZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFsTVJFd0R3WURWUVFLRXdocmRXSmwKTFdGM2N6RVFNQTRHQTFVRUF4TUhhM1ZpWlMxallUQWVGdzB4TmpFeE1EY3hNVEV4TVRCYUZ3MHlOakV4TURVeApNVEV4TVRCYU1DVXhFVEFQQmdOVkJBb1RDR3QxWW1VdFlYZHpNUkF3RGdZRFZRUURFd2RyZFdKbExXTmhNSUlCCklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUFzQWQ3WGxiQ3c3YU9tR1Z1anhDSzIwMzIKb1FiMGg0MC9jZWN1MUxObWJ0OG5ON1h1NHNuSXdpWEplbEg5K0d1N0duQy9ha3R1ei9JYUFWYlZzNmJRdElVVgpiRVQxQ25GcEg5UGR2MFZiNjRGS2U1TFc3ZFNtYlJsSzVDazhVQTB2dVlTZmhvakJSTE5uakFGNmMyRjNJd2FtCnBGTGxoc3plb3JOZThCeWVnQUN2Q2NuZS8zTGdkY3ptL0FMSGlZdGc5YU5pNTIxbjg5QW5rYWdXNlplTU1rRXYKbXUwZDRReFJMOUkvcW9tQ3ZlK2hUWXZuRy92dm5pQmNHaWRFdTYydnU2eGRGNmY1NkovK0Y3U2c0ZEZaYjhlbApKcnpFUzVMREZMOWJ2RE8zVGEvdjcvbWV5L1k5bkhTQ3FyUnVGdFV4VnFvZkdHM05sSnFNSFl5cnRpeTdhd0lECkFRQUJveU13SVRBT0JnTlZIUThCQWY4RUJBTUNBcVF3RHdZRFZSMFRBUUgvQkFVd0F3RUIvekFOQmdrcWhraUcKOXcwQkFRc0ZBQU9DQVFFQU5ieFdRK1c0REgrOU51bGhLeExONmlXMDBxU3JEZkdRREJuZ1VudjBxR1FDRUpVTAo5VTNrZ0Z1WjhFZXJGU085SXBvOXVDdEI3MnE1NkZyOVJjd1VuV0M2Z2kwTDVPbTlKT3owbGl0Q2RtR1cwbXRsCmM0dEFTM0RGc1FaM3ZTdmZZV21mcDFrcUFqdGFSYWJ4NEpkSVcrOTZ6TUxobUttQ0VqZU5GRXhtdEpVeWRJZ2MKSTJJK3hOZnBlK2l1TU5ib1NPTm90aU9RcjdZVTVJNTlZUlYzYms5b2FxcG9JSVU1cmduZW1KMWRWWkQzZ2szNAorY0dkUnVJZFhUbjBJdHZNeWRUUnFTcVB2enQvUk5NQlFaTW5RK25wZCtneWNBbE12eGE2cDJhWlljUzF6akhUCkU3dlROSTQ2Mk0vdHRMTjQvTUdiSWJMaHQ4eWRkc0tIV01aWlh3PT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=",
              "client_certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNzekNDQVpzQ0NRRHQ2UkxWSTE2T1JEQU5CZ2txaGtpRzl3MEJBUVVGQURBbE1SRXdEd1lEVlFRS0V3aHIKZFdKbExXRjNjekVRTUE0R0ExVUVBeE1IYTNWaVpTMWpZVEFlRncweE5qRXlNRGd4T1RNNE16bGFGdzB4TnpFeQpNRGd4T1RNNE16bGFNQkl4RURBT0JnTlZCQU1UQjNWeVltbDBZMmt3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBCkE0SUJEd0F3Z2dFS0FvSUJBUUM1Q1NOeW9hUk1mOHVMcDE3eW1oaHY2ZnlZb1lEZGg4K3UwVmdxbTVYL1lJd3IKUVJydnBRYTNDRFdFS0RQZVE1eVE1cExBTlBQdlJvZ3d5M2t5cW93em5OSmJqR01XaVBvWVFTcEloS1FNSlVTdQpRV1AyNlNEQ2ZRbW5jSnVSVG82MGpxeWxaai9lV1NnMFd4K0dkTmlUR1MwVEVrbVA0SUtLcWFIeGFTbkVDeW91CkZHZ1pTN0QrcER2T1FCM0FKalMwdVFoemxZWWFkMWRzSEZNbzFZdXJ1OWtBQ0VVekYrNUppTTNQSlBvR1JvcFIKL05WYzM0RDc1VjB4aWpxWlJBbUxwR2dvTHN0Sis3dUxDNGxEV0IzemN1RXRqNmtiNUtZRWoyK3N0QWc4M3BLSgp3WG9lSlNOZXB2YmVjc1ZoL0NjWVNiamNNaEFOQ2s1WkJISHpDV09MQWdNQkFBRXdEUVlKS29aSWh2Y05BUUVGCkJRQURnZ0VCQUFRWkZjOU83NmZXRGFVOUFONk1wYVBBOXF4UUEyL3RPVk9XRWJzUWZWbks1T0s0SmlkNVk5LzAKbGFVWlNLTC80TlRQVVVUZ1NPajJCQVBoM05JdnVBdFQyQ1VYRG5FaDZKUExxclZsZVEyQ2tPTFBMc2JvdEZhYwplNGtIOFlyL1cydjZKM21HOUhCUEJJaWVsbVc0NVc0Nng1djF5UWlMcElHdjVuMW5uKy9VR2swbDlISmd0bHgrCm9NekYvVHpWNHMzdFNzUzFrREg5aytuakJYNlh5d3ZDTkoydGsxUmNPZVB2SW5xdjdyYTN2NExNM1lTWEV4czYKSFRVTHBHNVFndWt1RkZpTUtBS21SMitIeGFNMmVQQXkxekZxYWdlQUZVbHN0UXhQNVQ5ZEUyeVo5czlzU0lzNgpwVW9JQys3V0dQTmp0S3BvdzJyKzBQOGFBdFZkWEhNPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==",
              "client_key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBdVFramNxR2tUSC9MaTZkZThwb1liK244bUtHQTNZZlBydEZZS3B1Vi8yQ01LMEVhCjc2VUd0d2cxaENnejNrT2NrT2FTd0RUejcwYUlNTXQ1TXFxTU01elNXNHhqRm9qNkdFRXFTSVNrRENWRXJrRmoKOXVrZ3duMEpwM0Nia1U2T3RJNnNwV1kvM2xrb05Gc2ZoblRZa3hrdEV4SkpqK0NDaXFtaDhXa3B4QXNxTGhSbwpHVXV3L3FRN3prQWR3Q1kwdExrSWM1V0dHbmRYYkJ4VEtOV0xxN3ZaQUFoRk14ZnVTWWpOenlUNkJrYUtVZnpWClhOK0ErK1ZkTVlvNm1VUUppNlJvS0M3TFNmdTdpd3VKUTFnZDgzTGhMWStwRytTbUJJOXZyTFFJUE42U2ljRjYKSGlValhxYjIzbkxGWWZ3bkdFbTQzRElRRFFwT1dRUng4d2xqaXdJREFRQUJBb0lCQVFDY0dzT1hJQnUybGxJbwp2Y2x5cnVKUytIcXNZZ1NQNE5ZcnpGMnZoSmRsWGhTaklVZ1NTWTJDdVNBOUlKV3h2Q1RJY2wzNFhqUTE5N0ZLClNUODBxWUdpd1hrTzF6OTVjWkpkQ0EwZUpSa3BUZi9GYTFGa3E0V0J6MjlubmE4QlJkOUxJTnN5cHpMVzZTenMKRHJ4bitRZ0dBY1Z2UTR4Z0g2N0NRUjVveHFuL3cyUkRpenVWQ1lUMUpqYzE3dlgxM2hDbzVsNkgrUmdBL2VhRApBaGNrdTBLRzZwa0R3WUE5L2k1Tlg0VGc2ZElEN3ZXU01hSDluUi9WRExxcmV2YnZnSG5PT0xseUV4R3RJL3NGCnFucVFxTGI0eW4yMTY1V2xkdlJwNUk2akNJekJINVZqKzgrOEVFRXU4SVB5cEhDd3cyT2xzV0FOMnZURGxnQm4KajZ6MVl0V1JBb0dCQU4zSjVMTWpXeEIwUys4OGZDamNHQzFjSGVwYm5pckhXYXcrR3ltbytBOVpKMlV2UU9nNApzMndIQmxhcXdNaU9UU0dmbUtxd0t1ZnJEM3J2WC9BNmlVWTYyRjFQdWZRK2dNTUFrbU4yemk3dlYxSFdCbnN3CjJGcnpiTmNpRGhTU3BCREFTVk12YVhXNWdQK2hBVFpjeXlnWjVDTXUrTUlqakllQ3JjWlk4aHA5QW9HQkFOV1QKN09PTVAxbEFkZ3FaUE5tUFdkNHBSLzJrNGFRNjl5OXdtZ0NIUEFid1M2WUY1d3JsditEa2UxTVhCb3lVM3Z1RApVelhXTkczZHBzSjNqV1JWYXB2NzZqM2JUQU0rUE1tVmNJLzBFS0E3Wk5mazZtTG81L2M5TmFIc2xFMkFFd0p6CmUzTTF2YW5hSld5eEhKdXNhYkhjS0NtYWlvMFFoMjlwd3ZkNXdVeW5Bb0dBQzhsK0VSTXc5TWZwZlRadXR0RXoKcTcxNGpZcis5ZkVRVC9vaEFXN01lQ3haenFQYlJEdzNOT2VPcTY1NWZtOHBwRDdTSTBnbmo1bkxnZElVL0RSdwpOVDVOWDNBc1J0SEhrQldJc2lhUFFLbFJyN1M3TlhMY0hNRlJLSUhUMDc0VFlCeUlDUmE4K1JlNXhsd3RMMUZ4CkxwbHBxWUVHa1hMSU5pOTR3dERaVlJVQ2dZQlV3UUgxZnBjNC9OcWE5QnB3bjNGak52Q0ptQit1dzNPS0VONGMKTFk1RmxwLytmME1qVU83bStPUnpvYVNJcng4Wm9oQ29RWnZHcVhuZW5BQ3crekIyTys3Rm96dXo0Y1BQbnd6dgpJMFJod1pBUUdKaG1yZFEzaWNPNXdSOU03ZkVkUE9TVllKTW1UeG9nMnR2bWJ2SDJrYzRpVEdDRkFEVXVva0tyCllGYXo2d0tCZ0Q3QVRtWVZrOWgxTkdLdlgrelh4dDhseTFnM2JmTkpkbTh1N050Q3JyY1doQ0U3ZHRkWTdwOHUKVEd6QUpudVNRZU1BODRwdWVNNzlZQWVIZHdmVEQxZFd6SU5BNjhVVUd6ZTNST1JySEpzaTQvZVIvSGZpNnJlZgpyMlZ2elJieDBTSHd0YWpwRHR3eHY5NmZsSHNtZG1Ld2t2Z2E5NXZEQ3M1YVBOOHpqVE51Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="
            }
          }
        },
        "commands": [
          {
            "install": {
              "chart": "nginx",
              "release": "dev-nginx",
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