# clira

**WARNING**: This is work in progress, don't trust to use this yet.

## Usage

### Setup

Use the following environment variables to configure clira, with the values for your corresponding setup:
```bash
export JIRA_HOST=http://localhost:8080
export JIRA_USER=admin
export JIRA_PASSWORD=admin
```

### List issues

```bash
clira ls
```

### Get issue details

```bash
clira get KEY-1
```

## Development setup

You'll need:
- Either [minikube](https://kubernetes.io/docs/setup/minikube/) or [docker-compose](https://docs.docker.com/compose/), and
- a [MyAtlassian account](https://id.atlassian.com/login), with either a valid license or the ability to get a new evaluation license for 90 days.

Clone the repo:
```bash
git clone https://github.com/ifosch/clira
cd clira
```

### With minikube

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
kubectl apply -f manifests/local-jira.yaml
```

Wait until the pod is finally running, and then you'll get your JIRA to be setup:
```bash
minikube service jira
```

### With docker-compose

This setup uses the port 8080 in your computer. Edit the docker-compose.yml file to change it.

If you prefer to use docker-compose:
```bash
cd manifests
docker-compose up
```

Finally, once the container is running fine, open a browser and point it to https://localhost:8008.

### Project setup

Go ahead with standard JIRA setup which will require you to provide a license from Atlassian, which can be for evaluation.
Then you'll have to create a user account, and a project.
Ensure to have a Scrum software project, with sprints

## Development commands

For testing, use `go test ./...` in the root directory
For formatting, use your editor of choice, or run `gofmt -w .`, in the root directory.
For running you can use `go run main.go`.
