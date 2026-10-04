# Domain Model

Greeter has a single, transient concept: the greeting it produces for a caller-supplied name. Nothing is persisted.

```mermaid
erDiagram
    GREETING {
        string name
        string message
    }
```

- **Greeting** — not stored; computed per request from the `name` query parameter (or a generic default when omitted) and returned as the response body.

