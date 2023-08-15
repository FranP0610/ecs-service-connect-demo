# service-connect-demo
Two apis written in go to test AWS ECS service connect

## Getting Started

### Dependencies

- Go version 1.20+

[//]: # (- Only works on ECS task)

### Installing

- Clone the repo: `git clone https://github.com/FranP0610/service-connect-demo.git`
- Navigate into the cloned project: `cd <repo-name>`


### Executing program

1. **create the docker image**: I have uploaded a template, all you have to do is change the values.
    ```bash
    docker built -t <name> -f <path to api> . 
    docker build -t service-connect-b -f cmd/api-b/Dockerfile .
    ```