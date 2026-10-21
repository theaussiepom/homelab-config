> Publication copy: deployment addresses and domains below are examples; restore private inputs before use.

# Unbound source and retained state

Use the canonical `stacks/benbeast/network/compose.yaml` entrypoint, project `network`, service `unbound`. Check `docs/migration/progress.json` for deployed status. Source preparation alone does not migrate a running container.

This directory holds editable `unbound.conf` and public `root.hints`. Its whole-directory bind preserves normal file replacement by editors. The configuration currently enables authenticated remote control; its root-hints and trust-anchor directives remain unchanged. The presence of `root.hints` or `var/root.key` does not itself establish active DNSSEC validation.

Private keys and certificates stay at `/mnt/nas/docker/network/unbound`: `unbound_server.key`, `unbound_server.pem`, `unbound_control.key` and `unbound_control.pem`. Each is mounted read-only over a zero-byte target placeholder here, at its existing path inside the Unbound container. Do not put key contents in this repository. Git exclusions are not access control. The Codex view sees the empty host placeholders, not the private files mounted in Unbound's namespace; existing Codex workspace guards remain intact.

The `dev` and `var` directories also remain in the original private/state tree, mounted read-write into Unbound. Empty directories here are only mount targets. Startup maintains chroot device nodes and the public root-anchor state in those retained directories. The ignored `unbound.pid` here is transient runtime state. Do not replace the base directory while its container mount is active.

All seven mounts refuse missing sources. Verify the NAS export `192.0.2.1:/volume1/docker-beast` before startup. Same-looking Pironman paths use the separate `docker-pi-study` export and are not shared state.

Private file binds retain their file inodes: after an atomic replacement of a key or certificate, recreate only Unbound to adopt it. A reload alone may retain the previous bound file. Preserve all four related credentials consistently; never regenerate keys or disable control merely to complete a source migration.

For a source rollback, stop the exact current writer, preserve its latest public configuration and current retained state, and recreate only Unbound from the original definition with the same image. Copy any newer `unbound.conf` and `root.hints` back to the old directory before that start; retained keys and `var`/`dev` already contain latest state. Never restore stale state as a shortcut. Use the verified private recovery record and check DNS plus authenticated control.
