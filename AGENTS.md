# Agent Context

**This repo:** `ffreis-siteops` — Go CLI that orchestrates local website builds.
It wraps `ffreis-website-compiler`, docker-compose, and AWS CLI into a single
command interface for local development. It is **not used in CI/CD pipelines**.

For the complete system map — how this repo relates to the deployer, compiler,
inventory, S3 infrastructure, and the other websites — see the private fleet
inventory repository:

> `the fleet inventory` → `AGENTS.md`

Architecture detail (CI/CD job graph, design decisions): `AGENTS.md` links to
`docs/ARCHITECTURE.md` in the same repo.

Do not look for cross-component flow documentation in this repo's README;
it covers only siteops's own CLI commands and configuration.
