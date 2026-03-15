# ADR-001: Stack Decision

**Date:** 2026-03-15  
**Status:** Accepted

## Decision
React + Vite (PWA) for frontend, Go + Gin for backend,
PostgreSQL via Supabase for database.

## Reasons
- React: already known, PWA avoids app store complexity
- Go: learning on the job at Entraste, double benefit
- PostgreSQL: better analytics support, Supabase gives
  auth + storage for free
- Supabase: eliminates need for separate auth service

## Alternatives discarded
- React Native: too much overhead for MVP validation
- MySQL: no integrated ecosystem, weaker analytics
- Express/Node: Go is better for concurrent API workloads

## Update — 2026-03-15
Backend hosting changed from Railway to DigitalOcean App Platform.
Reason: $200 credit via GitHub Student Pack covers 16 months at no cost.