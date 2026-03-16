export type Platform = 'youtube' | 'netflix' | 'prime';

export interface Link {
    id: string;
    short_code: string;
    platform: Platform;
    content_id: string;
    timestamp_s: number;
    timestamp_fmt: string;
    title: string | null;
    click_count: number;
    created_at: string;
    expires_at: string | null;
}

export interface CreateLinkPayload {
    platform: Platform;
    content_id: string;
    timestamp_s: number;
    title?: string;
    thumbnail?: string;
    expires_in?: string;
}

export interface VideoInfo {
    platform: Platform;
    content_id: string;
    title: string;
    thumbnail: string;
    timestamp_s: number;
}

export interface Stats {
    short_code: string;
    click_count: number;
    created_at: string;
}
