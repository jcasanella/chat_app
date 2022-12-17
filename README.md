# Chat App

[![codecov](https://codecov.io/gh/jcasanella/chat_app/branch/main/graph/badge.svg?token=VIU7H2NELQ)](https://codecov.io/gh/jcasanella/chat_app)

The purpose of this project is to write a Chat Application using Golang and React. 

## Architecture 

TBD

## Technologies to Learn

* Golang
* Database to store the msgs and users
* Redis to store the latest messages
* Flex to deploy into K8S (EKS)
* Github Actions
* Test Coverage with CodeCov
* React

## Roadmap

[03/12/2022]
. Basic github actions => Testing, Lint and Vet 
. Enable test coverage using CodeCov as part of the Github Actions

[09/12/2022]
. First Golang router => Created only for testing purposes

[11/12/2022]
. Improve Dockerfile with multistages and non root user

[18/12/2022]
. Github action to push docker image into ECS

