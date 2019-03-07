# HTTP Assertion Tool

While tools like Cypress have evolved to provide concise, dependable end-to-end tests, these tools rely on driving a UI and asserting its state. What is lacking is a similarly concise, similarly dependable, similarly _enjoyable_ tool for driving an HTTP API and asserting its responses. This tool aims to meet that need.

Like most TD tools, this tool aims to codify the opinions of its authors for not only a technology, but its influence on the _process_ and _practice_ of making software. This tool is opinionated.

## Features

- Makes requests and asserts responses.
- Tests can be preceded by setup, and postceded by teardown.
  - Setup and teardown can be either request/response pairs (e.g. POST /resource) or running executables (e.g. DB migrations).
- Aspects of a test can be generated by user-provided code.

## Goals

- Language agnostic: This should work just as well for an ASP.NET project as a Rails one.
- Comprehensive: All aspects of a response can be asserted (status code, headers, entity, trailers, etc.).
- Strict by default: When something is not specified, the defaults should be as restrictive as possible, e.g. keys not specified in JSON entities are _forbidden_, rather than ignored.
- Concise: If such a default can be provided, it will be.
- Works for REST and GraphQL, etc.: Rather than tying the requests and responses to Resources or Queries, tests should be HTTP-specific, but no moreso.
- Speed: This tool should be fast enough to, if the underlying service is as performant, be used during feature development.

## Design

- Byte-for-byte comparisons work well:
  - Render request and response pair as template strings.
  - Make request.
  - Compare response, byte by byte, with the rendered version. For every template _matcher_ encountered, run the matcher.
- Needs to be smart about parsing headers, etc.
  - Example: Date header. It should be ignored from response assertions, but that requires _parsing_ the headers to know that.
  - Example: Host header. It should be dynamic, based on the host used for testing.
  - Example: HTTP version. It should be optional in response expectations, asserted if and only if present.
- Large bodies: Needs to be able to pull in other files via fixture, etc.

## Alternatives

- Superagent
- Frisby
- Dredd
- Roll-your-own

## TODO

- Request bodies
- Rendering templates before parsing
- Matchers while diffing
- Default body diffs to application/octet-stream; diff byte by byte
- Linting
- BATS
- Rename to Diplomat
- Make HTTP/V optional? Assert it and make it a first party instead?
- Debug flag
- UX
- Streaming
- Trailers
