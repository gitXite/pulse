# Pulse

<div style="flex">
  <img src="web/public/icon-reversed-1024.png" alt="Icon" width="200" />
  <p align="left">
    <img src="https://img.shields.io/badge/TypeScript-%E2%9C%94-blue?logo=go" alt="Go">
    <img src="https://img.shields.io/github/last-commit/gitXite/pulse" alt="Last Commit">
    <img src="https://img.shields.io/github/languages/top/gitXite/pulse" alt="Top Language">
  </p>
</div>

### A lightweight HTTP load-testing tool written in Go

Pulse is a load-testing tool I'm building as a long-term project alongside my bachelor's degree in computer science.

The goal is to start with a small, focused CLI and gradually evolve Pulse into a capable load-testing system. 
The project is intentionally being developed iteratively, with each stage introducing new engineering challenges rather than trying to build everything at once.

## Project direction

The initial version of Pulse will focus on the fundamentals of HTTP load testing:

* Generating configurable HTTP traffic
* Controlling concurrency and request rates
* Collecting and aggregating performance metrics
* Measuring latency, throughput, errors, and other useful statistics
* Providing clear results through a CLI

From there, the project will grow toward a configurable dashboard and eventually a distributed architecture with multiple workers.

The rough progression is:

**CLI → Configurable Dashboard → Distributed Load Testing**

## Why I'm building it

Pulse is as much an engineering project as it is a load-testing tool.

I want to use the project to develop a deeper understanding of areas that become increasingly important as systems grow in complexity:

* System design and architecture
* Concurrency and resource management
* Performance engineering
* Statistics and data aggregation
* Networking and HTTP
* Distributed systems
* Observability
* Designing software that can handle increasing load

The project will also serve as a practical way to apply concepts from my studies and evaluate how design decisions hold up as the system becomes more complex.

## Technology

Pulse is written in **Go**, chosen primarily for its concurrency model, performance, and suitability for building networked systems.

The project will remain deliberately lightweight and dependency-conscious where practical. I want to understand the systems Pulse relies on rather than hiding the interesting parts behind large frameworks.

## Project status

Pulse is currently in the early development and planning stage.

The first milestone is a working CLI capable of generating HTTP load against a test API and producing useful measurements. Features will be added incrementally as the underlying design matures.

This repository will document that progression, including architectural decisions, experiments, and lessons learned along the way.

---

Pulse is a project about building a load tester, but more importantly, about learning how to design and build systems that can handle load.
