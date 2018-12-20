# Gray release

Gray release based on Istio and Kubernetes.

- [Introduction](#introduction)
- [Repositories](#repositories)
- [Issue management](#issue-management)

In addition, here are some other documents you may wish to read:

- [Istio ](https://github.com/istio/istio) - an open platform to connect, manage, and secure microservices
- [Kubernetes ](https://github.com/kubernetes/kubernetes) - an open source system for managing containerized
applications across multiple hosts; providing basic mechanisms for deployment, maintenance, and scaling of applications.
## Introduction

Istio is an open platform for providing a uniform way to integrate
microservices, manage traffic flow across microservices, enforce policies
and aggregate telemetry data. Istio's control plane provides an abstraction
layer over the underlying cluster management platform, such as Kubernetes,
Mesos, etc.

Visit [istio.io](https://istio.io) for in-depth information about using Istio.

Istio is composed of these components:

- **Envoy** - Sidecar proxies per microservice to handle ingress/egress traffic
   between services in the cluster and from a service to external
   services. The proxies form a _secure microservice mesh_ providing a rich
   set of functions like discovery, rich layer-7 routing, circuit breakers,
   policy enforcement and telemetry recording/reporting
   functions.

  > Note: The service mesh is not an overlay network. It
  > simplifies and enhances how microservices in an application talk to each
  > other over the network provided by the underlying platform.

- **Mixer** - Central component that is leveraged by the proxies and microservices
   to enforce policies such as authorization, rate limits, quotas, authentication, request
   tracing and telemetry collection.

- **Pilot** - A component responsible for configuring the proxies at runtime.

- **Citadel** - A centralized component responsible for certificate issuance and rotation.

- **Node Agent** - A per-node component responsible for certificate issuance and rotation.

- **Galley**- Central component for validating, ingesting, aggregating, transforming and distributing config within Istio.

Istio currently supports Kubernetes and Consul-based environments. We plan support for additional platforms such as
Cloud Foundry, and Mesos in the near future.



