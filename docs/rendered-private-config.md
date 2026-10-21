# Application configuration from private environment values

Traefik's static configuration is maintained as a public `.in` YAML template. `!env HOMELAB_...` marks a whole YAML scalar;
`../homelab/scripts/render_private_yaml.py` resolves it from the environment or a private
env file. The generated original filename is ignored by Git.

Use `../homelab/scripts/private-yaml.env.example` for key names. The real prepared values
remain on BenBeast. The renderer accepts plain, single-quoted or JSON-quoted env
values and does not execute a shell, evaluate expressions or recursively expand
values. Environment values override the supplied env file.

From the `homelab-config` checkout, with a sibling `homelab` checkout and
the output absent:

```sh
python3 ../homelab/scripts/render_private_yaml.py \
  --template traefik/config/traefik.yaml.in \
  --env-file /private/path/application-values.env \
  --output traefik/config/traefik.yaml
```

The env file must be owner-only. Output is a new mode-0600 file. Existing files
and symlink output paths are refused: prepare a separate output for review,
then use the normal deployment procedure to replace an existing configuration.
No service is started or reloaded by the renderer.

Rendering is a deployment-input step, not a publication file list. Normal Git
tracking and `.gitignore` determine source inclusion. Live files are preserved
and production has not been switched to these templates.

Reference: Traefik Go templates apply to dynamic files, not static configuration:
https://doc.traefik.io/traefik/reference/dynamic-configuration/file/
