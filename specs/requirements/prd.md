# greeter — PRD

## Problem Statement

Teams building new components on this platform need a minimal, known-good example of a Go HTTP service to copy conventions from — project layout, request handling, and response shape. Without a small reference service, each new Go component re-derives these conventions from scratch, leading to inconsistency across the organization. E2E marker s0p4-1004a.

## Solution

Greeter is a small Go HTTP service that exposes a single endpoint, `GET /hello?name=X`, returning a JSON greeting for the given name. It follows the conventions established in `app-factory-kaj/e2e-reference`, so it can serve as a lightweight, working reference for how a Go service on this platform is structured and built.

## Actors

- **Caller** — another internal service or component in the organization that invokes the greeter endpoint to obtain a JSON greeting. *(assumed)*

## User Stories

1. As a Caller, I want to send a name to the greeter service, so that I receive back a JSON greeting addressed to that name.
2. As a Caller, I want a sensible response when I omit the name, so that the service behaves predictably rather than failing unclearly.

## Product Decisions

- Callers: the service is intended to be called by other internal services/components in the organization, not by external developers or end users directly. *assumed*
- Authentication: the `/hello` endpoint is open — it does not require caller authentication. *assumed*
- Conventions: the service's project structure and build conventions follow `app-factory-kaj/e2e-reference`.
- No sign-in, user interface, or persistence is part of this service — it is a stateless, single-endpoint API.

## Out of Scope

- Any user interface or web application — greeter is an API-only service.
- Persisting greetings, names, or request history.
- Supporting languages/locales other than a single greeting format.
- Rate limiting, authentication, or authorization on the endpoint.
- Any endpoint other than `GET /hello`.

## Open Questions

1. Should `GET /hello` without a `name` parameter return an error (e.g. 400) or a generic greeting (e.g. "Hello, World")? *(deferred — default/generic-greeting behavior will be decided at design time unless the user specifies otherwise)*

## Further Notes

This project exists primarily as a small, conventions-following reference service; scope is intentionally minimal.