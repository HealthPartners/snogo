# Snogo

A Prometheus webhook receiver that generates ServiceNow incidents.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

- golang

```
yum/apt-get/brew install go
```

### Installing

Setup your local environment and produce the snogo binary.

```
$ export GOPATH='/home/me/some/path'
$ go get -u github.com/healthpartnersoss/snogo
```

_This will produce a working binary._

Use environment variables to change the behavior of Snogo.

```
export SERVICE_NOW_INSTANCE_NAME='service_now_org'
export SERVICE_NOW_USERNAME='me'
export SERVICE_NOW_PASSWORD='secretme'
```

Execute snogo.

```
./snogo
```

When executed, Snogo will begin listening on port 8080 and is ready to receive
Alertmanager events.

The last step (external to running this binary) is to point an
[alertmanager webhook receiver](https://prometheus.io/docs/alerting/configuration/#webhook_config) 
at Snogo.

```
- name: service-now
  webhook_configs:
  - send_resolved: false
    url: http://localhost:8080/

```

## Running the tests

We use the built-in `go test` mechanism for coverage.

```
go test
```

## Deployment

- Execute the binary as a system user
- Open port 8080 (by default)
- 

## Built With

* [Golang]()

## Contributing

Please read [CONTRIBUTING.md](:https//example.com) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/healthpartnersoss/snogo/tags). 

## Authors

* **Jesse Olson** - *Prometheus work* - [HealthPartners](https://github.com/healthpartnersoss)
* **James McShane** - *ServiceNow work* - [HealthPartners](https://github.com/healthpartnersoss)

See also the list of [contributors](https://github.com/healthpartnersoss/snogo/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
