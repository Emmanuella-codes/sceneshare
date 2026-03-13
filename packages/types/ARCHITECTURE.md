**Scope**
Shared TypeScript types for API contracts used by web and extension clients.

**Current Architecture**
1. Type-only package exported from `src/index.ts`
2. Used by extension; web app not yet wired to it

**Data Flow**
1. Clients import types for request/response payloads.
2. API returns data matching these types.

