import type { CreateLinkPayload, Link } from '@sceneshare/types';

const API_URL = import.meta.env.VITE_API_URL ?? 'http://localhost:4006';

export async function createLink(payload: CreateLinkPayload): Promise<Link> {
    const res = await fetch(`${API_URL}/api/v1/links`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
    });
    if (!res.ok) throw new Error('API error: ${res.status');
    return res.json();
}
