# manifests

This folder contains manifest files for course deployment, service, statefulset etc. Manifest files are meant to run with kubectl and course tasks run on single-node cluster with k3s and k3d.

Create cluster:
```bash
 k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
```

Apply manifest file:
```bash
kubectl apply -f manifests/<MANIFEST_FILE.yml>
```

Apply encrypted secret:
```bash
export SOPS_AGE_KEY_FILE=/path/to/key
sops --decrypt <ENCRYPTED_SECRET_MANIFEST.yml> | kubectl apply -f -
```
