import { Platform } from "@sceneshare/types";

export function formatTimestamp(s: number): string {
    const h = Math.floor(s / 3600);
    const m = Math.floor((s % 3600) / 60);
    const sec = s % 60;
    if (h > 0) {
        return `${h}:${m.toString().padStart(2, '0')}:${sec.toString().padStart(2, '0')}`;
    }
    return `${m}:${sec.toString().padStart(2, '0')}`;
}

export function platformLabel(platform: Platform): string {
    const map: Record<string, string> = {
        youtube: "YouTube",
        // maybe more platforms in the future
    };
    return map[platform] ?? platform;
}

export function platformColor(platform: Platform): string {
    const map: Record<string, string> = {
        youtube: "#FF0000",
    }
    return map[platform] ?? "#888";
}