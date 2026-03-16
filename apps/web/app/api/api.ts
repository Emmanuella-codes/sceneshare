import type { CreateLinkPayload, Link, Stats } from '@sceneshare/types';

const API_URL = process.env.NEXT_PUBLIC_API_URL

export class ApiError extends Error {
    constructor(
        public status: number,
        public code: string,
        message: string,
    ) {
        super(message);
        this.name = 'APIError';
    }
}

export class NotFoundError extends ApiError {}
export class ExpiredError extends ApiError {}

async function handleResponse<T>(res: Response): Promise<T> {
    if (res.ok) return res.json() as Promise<T>;

    let code = "UNKNOWN_ERROR";
    let message = `HTTP ${res.status}`;

    try {
        const body = await res.json();
        code = body.code ?? code;
        message = body.message ?? message;
    } catch (error) {

    }

    if (res.status === 404) throw new NotFoundError(res.status, code, message);
    if (res.status === 410) throw new ExpiredError(res.status, code, message);
    throw new ApiError(res.status, code, message);
}

export async function createLink(payload: CreateLinkPayload): Promise<Link> {
    const res = await fetch(`${API_URL}/api/v1/links`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
    });
    return handleResponse<Link>(res);
}

export async function getLink(code: string): Promise<Link> {
    const res = await fetch(`${API_URL}/api/v1/links/${code}`, {
        next: { revalidate: 60 },
    });
    return handleResponse<Link>(res);
}

export async function deleteLink(code: string): Promise<void> {
    const res = await fetch(`${API_URL}/api/v1/links/${code}`, {
        method: 'DELETE',
    });
    if (!res.ok && res.status !== 404) {
        await handleResponse(res);
    }
}

export async function getStats(code: string): Promise<Stats> {
    const res = await fetch(`${API_URL}/api/v1/links/${code}/stats`, {
        next: { revalidate: 60 },
    });
    return handleResponse<Stats>(res);
}

export async function redirect(code: string): Promise<void> {
    const res = await fetch(`${API_URL}/api/v1/links/${code}`, {
        method: 'GET',
    });
    if (!res.ok) throw new Error('API error: ${res.status');
    return res.json();
}
