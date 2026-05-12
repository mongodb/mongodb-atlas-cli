/**
 * Atlas pledge extension for pi.
 *
 * Restricts atlas-cli to the configured pledge profile at session start.
 *
 * Install with:
 *   atlas hook install pi [--profile <profile>]
 *
 * Then restart pi or run /reload.
 */

import type { ExtensionAPI } from "@mariozechner/pi-coding-agent";
import { createBashTool, createLocalBashOperations } from "@mariozechner/pi-coding-agent";

const ATLAS_PLEDGE_PROFILE = "{{.Profile}}";
const ATLAS_TIMEOUT_MS = 10_000;

export default async function atlasPledgeExtension(pi: ExtensionAPI) {
	const cwd = process.cwd();

	// Set the pledge at session start. Errors are swallowed so a missing or
	// misconfigured atlas binary never prevents pi from starting.
	try {
		await pi.exec("atlas", ["pledge", "set", ATLAS_PLEDGE_PROFILE, "--yes"], {
			cwd,
			timeout: ATLAS_TIMEOUT_MS,
		});
	} catch {
		// Intentionally ignored.
	}

	pi.registerTool(
		createBashTool(cwd, {
			operations: createLocalBashOperations(),
		}),
	);
}
