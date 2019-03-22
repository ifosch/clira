# clira

**WARNING**: This is work in progress, don't trust to use this yet.

# Commands

For testing, use `go test ./...` in the root directory
For formatting, use your editor of choice, or run `gofmt -w .`, in the root directory.
For running you can use `go run main/clira.go`, but ensure you get a working config file.

## Development setup

You'll need [minikube](https://kubernetes.io/docs/setup/minikube/), and a [MyAtlassian account](https://id.atlassian.com/login), with either a valid license or the ability to get a new evaluation license for 90 days.

Clone the repo:
```bash
git clone https://github.com/ifosch/clira
cd clira
```

Ensure your minikube has, at least 3 cpus and >2048 Mb of memory:
```bash
minikube stop
minikube config set memory 4096
minikube config set cpus 3
minikube delete
minikube start
```

Once minikube has enough resources:
```bash
kubectl apply -f manifests/loca-jira.yaml
```

Wait until the pod is finally running, and then you'll get your JIRA to be setup:
```bash
minikube service jira
```

Go ahead with standard JIRA setup (will require you to provide a license from Atlassian).
