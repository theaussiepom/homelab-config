# Compose deployment values

The canonical Compose files use required `${HOMELAB_...:?message}` references
for literal deployment addresses. Each project's `.env.example` lists its keys.
Put real values in the existing private project environment file used by
`scripts/compose.py`, or pass the equivalent private `--env-file` explicitly.
An absent value must fail interpolation; do not substitute guessed defaults.

The private prepared additions remain on BenBeast. Merge those keys without
overwriting the existing project env files before deploying this source.
The candidate's substituted YAML was compared with the original YAML to verify
that addresses, ports, service settings and stable identities resolve unchanged.

These are branch preparation changes. No Compose apply or service restart was
performed. WiNet's service block is maintained by its separate work.
