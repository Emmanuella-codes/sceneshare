// ── Popup → Content script ────────────────────────────────────────────────────

import { VideoInfo } from "@sceneshare/types";

 
export interface GetVideoInfoMessage {
    action: "getVideoInfo";
}
   
  // ── Popup → Background ────────────────────────────────────────────────────────
   
export interface ShareMomentMessage {
    action: "shareMoment";
    videoInfo: VideoInfo;
}
   
  // ── Union type used in onMessage listeners ────────────────────────────────────
   
export type ExtensionMessage =
    | GetVideoInfoMessage
    | ShareMomentMessage;
   
  // ── Response shapes ───────────────────────────────────────────────────────────
   
export interface VideoInfoResponse {
    videoInfo: VideoInfo | null;
}

export interface ShareMomentResponse {
    success: boolean;
    shortUrl?: string;
    error?: string;
}
