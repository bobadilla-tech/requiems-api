// Server-only module — never import from client components.
// The REQUIEMS_API env var has no NEXT_PUBLIC_ prefix so Next.js
// will never include it in the browser bundle.

import { env } from "@/env";

export interface EmailValidationData {
	email: string;
	valid: boolean;
	syntax_valid: boolean;
	mx_valid: boolean;
	disposable: boolean;
	normalized: string;
	domain: string;
	/** Domain suggestion only (e.g. "gmail.com"). Reconstruct full email with localPart + "@" + suggestion. */
	suggestion: string | null;
}

interface EmailValidationResponse {
	data: EmailValidationData;
	metadata: { timestamp: string; trace_id?: string };
}

export type ValidateEmailResult =
	| { ok: true; data: EmailValidationData }
	| { ok: false; error: string; status: number };

/**
 * Calls POST /v1/email/validate on the Requiems API.
 * Safe to call from Route Handlers, Server Actions, and Better Auth hooks only.
 */
export async function validateEmail(email: string): Promise<ValidateEmailResult> {
	let response: Response;

	try {
		response = await fetch("https://api.requiems.xyz/v1/email/validate", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				"requiems-api-key": env.REQUIEMS_API,
			},
			body: JSON.stringify({ email }),
			cache: "no-store", // never cache validation results
		});
	} catch {
		return {
			ok: false,
			error: "Could not reach the Requiems API. Check your network.",
			status: 503,
		};
	}

	if (!response.ok) {
		let errorCode = `API error ${response.status}`;
	
    try {
			const body = await response.json();
			if (body?.error) errorCode = body.error;
		} catch {
			// ignore parse error
		}
		return { ok: false, error: errorCode, status: response.status };
	}

	const json: EmailValidationResponse = await response.json();
	return { ok: true, data: json.data };
}
