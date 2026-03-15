# ADR-002: Frontend Setup

**Date:** 2026-03-15
**Status:** Accepted

## Decision
React + Vite with TypeScript, Tailwind CSS v4, and shadcn/ui (Base + Nova preset).

## Reasons
- TypeScript: industry standard, catches errors at compile time
- Tailwind v4: latest version, utility-first CSS
- shadcn/ui: accessible components, no external dependency lock-in,
  components live in your own codebase

## Alternatives discarded
- JavaScript: less safe, worse tooling support
- CSS Modules: more verbose, slower to develop
- Material UI / Ant Design: heavier, harder to customize
