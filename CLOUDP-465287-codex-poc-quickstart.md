# Codex Agent Skills PoC

## Checkout and install

The clone commands assume `poc/atlas-skills` has been pushed to `origin`.

```sh
git clone https://github.com/mongodb/mongodb-atlas-cli.git
cd mongodb-atlas-cli
git fetch origin poc/atlas-skills
git switch --track -c poc/atlas-skills origin/poc/atlas-skills
make install
export CODEX_HOME="${CODEX_HOME:-$HOME/.codex}"
```

If the branch is already checked out locally, just switch to `poc/atlas-skills`.

## Compare Codex before and after a standalone skill

Run the same prompt before and after installing the skill:

```sh
codex exec "Design a MongoDB schema for an online store. Explain embedding vs. references and recommend indexes."

atlas skills install mongodb-schema-design --agent codex

# Start a new Codex invocation so it loads the installed skill.
codex exec "Design a MongoDB schema for an online store. Explain embedding vs. references and recommend indexes."
```

To install every standalone skill instead, run `atlas skills install --all --agent codex`.

## Optional Atlas plugin test

```sh
atlas skills install --plugin mongodb-atlas --agent codex
codex plugin list --json
```

The Codex plugin uses the upstream app/skills manifest. It does not add a direct
entry to `codex mcp list`.

## Optional cleanup before or after testing

```sh
CODEX_HOME="${CODEX_HOME:-$HOME/.codex}"
rm -rf "$CODEX_HOME/skills/mongodb-schema-design"

# Run if the Atlas plugin was installed:
codex plugin remove mongodb-atlas@mongodb-plugins

# Also run if you want to test adding the marketplace from scratch:
codex plugin marketplace remove mongodb-plugins
```

If you used `--all`, remove the remaining standalone skills too:

```sh
CODEX_HOME="${CODEX_HOME:-$HOME/.codex}"
for skill in mongodb-atlas-stream-processing mongodb-connection mongodb-mcp-setup \
  mongodb-natural-language-querying mongodb-qe-size-estimation \
  mongodb-query-optimizer mongodb-search-and-ai; do
  rm -rf "$CODEX_HOME/skills/$skill"
done
```

The native Codex plugin does not create a direct MCP server entry, so no
`codex mcp remove` command is needed for this PoC.
