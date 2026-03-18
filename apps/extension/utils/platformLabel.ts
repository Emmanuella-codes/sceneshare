import { Platform } from "@sceneshare/types";

export function platformLabel(platform: Platform): string {
    const labels: Record<string, string> = {
        youtube: "YouTube",
    };
    return labels[platform] ?? platform;
}
