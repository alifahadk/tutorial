# Plugin Development Tutorial

This repository provides a tutorial for developing new plugins for Blueprint. The repository covers three major use-cases for plugins:

+ Instrumenting services in an application's workflow
+ Adding new methods/APIs to services in an application's workflow
+ Modifying the function signatures of methods for services in an application's workflow

The repository provides a plugin implementation for each of the aforementioned types along with a simple two-service application to demonstrate how to use the plugins in an application. The repository is structured as follows:

+ [plugins/tutorial](plugins/tutorial): Package that contains the implementation of the `tutorial` plugins.
+ [examples/helloworld](examples/helloworld): Module that provides the two-service application and demonstrates the use of the `tutorial` plugins.


# Run Loadgen

While in loadgen directory

```sh
go run loadgen.go --rps=10000 --duration=1s --total=1000

```

# Example Metafor DSL JSON

```json
{
    "servers": [
        {
            "name": "52",
            "qsize": 100,
            "threadpool": 1,
            "apis": {
                "insert": {
                    "processing_rate": 10,
                    "downstream_services": []
                },
                "get": {
                    "processing_rate": 10,
                    "downstream_services": []
                },
                "put": {
                    "processing_rate": 20,
                    "downstream_services": []
                },
                "list": {
                    "processing_rate": 2,
                    "downstream_services": []
                }
            }
        },
        {
            "name": "server2",
            "qsize": 20,
            "threadpool": 1,
            "apis": {
                "insert": {
                    "processing_rate": 10,
                    "downstream_services": [
                        {
                            "source": "server2",
                            "target": "server2",
                            "api": "rd",
                            "blocking": true,
                            "timeout": 10,
                            "retry": 3
                        }
                    ]
                },
                "get": {
                    "processing_rate": 10,
                    "downstream_services": []
                },
                "put": {
                    "processing_rate": 20,
                    "downstream_services": []
                },
                "list": {
                    "processing_rate": 2,
                    "downstream_services": []
                }
            }
        }
    ],
    "sources": [
        {
            "name": "client",
            "api": "insert",
            "arrival_rate": 9.5,
            "timeout": 3,
            "retries": 4
        }
    ]
}
```
