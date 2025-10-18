Continuous Integration Pipeline using GitHub Actions

Deepak Tamilzhagan

# Objective

The Continuous Integration (CI) pipeline using GitHub Actions that automatically builds, tests, and containerizes a simple Go application.

The goal was to:

  Understand the principles of Continuous Integration.
  
  Implement a CI workflow with Build and Test stages.
  
  Demonstrate automated testing using unit tests.
  
  Create and publish a Docker image to a container registry.
  
  Show both success and failure CI scenarios.

# Project Overview: NumStats

NumStats is a small command-line tool written in Go that calculates basic statistics for any list of numbers.
When users enter a few numbers, it prints the Count, Sum, Mean, Median, Minimum, and Maximum values.

Example:
numstats 10 20 30 40


Output:

Count: 4
Sum: 100
Mean: 25
Median: 25
Min: 10
Max: 40


It also validates inputs — for example, typing a non-number shows:

invalid number: "x"

# Implementation Steps
### Step 1 – Application Development

Language: Go (Golang)

Implemented six key functions:
Sum, Mean, Median, Min, Max (each with test cases).

Added a CLI (main.go) that accepts user inputs and prints statistics.

### Step 2 – Unit Testing

Added 5 unit tests in stats_test.go.

Verified results with:

go test -v -cover ./internal/stats


Covered normal, edge, and invalid input scenarios.

### Step 3 – GitHub Actions (CI Workflow)

Created .github/workflows/ci.yml with two jobs:

Build & Test

Installs Go

Builds code

Runs unit tests

Docker Publish

Builds a Docker image

Pushes to GitHub Container Registry (GHCR)

### Step 4 – Docker Image

Multi-stage Dockerfile using golang:1.25-alpine → alpine:3.20.

Image published automatically after successful CI.

### Pull command:

docker pull ghcr.io/deepak-tamizhalagan/numstats-ci-demo:latest

docker run --rm ghcr.io/deepak-tamizhalagan/numstats-ci-demo:latest 5 10 15

# CI Demonstration

- ✅ Success Run	All tests passed → Docker image built and pushed	--Green

+ ❌ Fail Demo PR	One test deliberately failed to demonstrate pipeline failure	--Red

* ✅ Fix PR	Corrected test and re-ran pipeline → Passed successfully	--Green

These runs were visible under the GitHub Actions tab.

# Technologies Used

+ Go (Golang)	--Application development

+ Git & GitHub	--Version control

+ GitHub Actions	--Continuous Integration pipeline

+ Docker	--Containerization

+ GHCR (GitHub Container Registry)	--Image publishing

# Results & Learning Outcomes

+ Successfully automated build → test → docker publish stages.

+ Demonstrated CI behavior for both passing and failing test cases.


