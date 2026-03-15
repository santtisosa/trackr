# ADR-003: Backend Setup

**Date:** 2026-03-15
**Status:** Accepted

## Decision
Go + Gin as the HTTP framework for the REST API.

## Reasons
- Go: learning on the job at Entraste, double benefit
- Gin: most popular Go HTTP framework, good performance,
  simple routing and middleware
- golang-migrate: standard migration tool for Go
- sqlc: type-safe SQL queries, avoids magic ORMs

## Alternatives discarded
- Express/Node: Go is better for concurrent API workloads
- GORM: too much magic, hides SQL complexity
- Echo: similar to Gin but smaller community