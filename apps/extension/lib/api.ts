import type { CreateLinkPayload, Link } from '@sceneshare/types';

const API_URL = import.meta.env.VITE_API_URL ?? 'http://localhost:4006';

export async function createLink(payload: CreateLinkPayload): Promise<Link> {
  const res = await fetch(`${API_URL}/api/v1/links`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });

  if (!res.ok) {
    let message = `API error: ${res.status}`;

    try {
      const body = (await res.json()) as { message?: string };
      if (body.message) {
        message = body.message;
      }
    } catch {
      // Fall back to the HTTP status when the response body is unavailable.
    }

    throw new Error(message);
  }

  return res.json();
}
