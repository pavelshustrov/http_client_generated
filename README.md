# API Client Generator with OpenAPI and oapi-codegen

This repository demonstrates an **API-first** approach to client generation using OpenAPI and the `oapi-codegen` library by DeepMap (NVIDIA). The goal is to avoid writing manual HTTP call code and instead rely on automatically generated client code based on an OpenAPI specification (`openapi.yml`).

## Overview

With this approach, API endpoints are documented using the OpenAPI format, and clients are generated automatically using the `oapi-codegen` tool. The generated code allows developers to plug in custom HTTP clients. This means you can use the standard Go HTTP client, or integrate more advanced 3rd-party solutions like [Gojek Heimdall](https://github.com/gojek/heimdall) for enhanced reliability and functionality.

### Why Use a 3rd-Party HTTP Client?

Using a 3rd-party HTTP client like Heimdall provides several benefits over the standard Go HTTP client:
- **Retries**: Automatic retries on transient failures.
- **Timeouts**: Fine-grained control over request timeouts.
- **Circuit Breakers**: Helps handle service degradation or failures gracefully.
- **Metrics**: Built-in support for monitoring and logging.

These features improve the resilience, observability, and performance of your applications, so you don't have to reinvent the wheel when handling common HTTP-related concerns.

## How It Works

1. **Document your API**: Define your API endpoints in the OpenAPI format (`openapi.yml`).
2. **Generate Client Code**: Use `oapi-codegen` to generate Go client code from the OpenAPI specification.
3. **Customize HTTP Client**: The generated client code lets you integrate your custom HTTP client. You can use the default `net/http` client or a more advanced client like Heimdall.
