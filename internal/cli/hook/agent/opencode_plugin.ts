/**
 * Atlas pledge plugin for opencode.
 *
 * Restricts atlas-cli to the configured pledge profile at session start.
 *
 * Install with:
 *   atlas hook install opencode [--profile <profile>]
 *
 * Opencode loads plugins from ~/.config/opencode/plugins/ automatically.
 */

const ATLAS_PLEDGE_PROFILE = "{{.Profile}}";

export const AtlasPledgePlugin = async ({ $ }: { $: typeof import("bun").$ }) => {
	return {
		event: async ({ event }: { event: { type: string } }) => {
			if (event.type === "session.created") {
				try {
					await $`atlas pledge set ${ATLAS_PLEDGE_PROFILE} --yes`.quiet();
				} catch {
					// Intentionally ignored: a missing or misconfigured atlas binary
					// must never block session creation.
				}
			}
		},
	};
};
