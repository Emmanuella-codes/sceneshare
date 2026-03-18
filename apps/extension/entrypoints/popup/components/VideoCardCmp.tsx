import { formatTimestamp } from "@/utils/formatTimestamp";
import { platformLabel } from "@/utils/platformLabel";
import type { Platform, VideoInfo } from "@sceneshare/types";
import { ArrowIcon } from "./icons/icons";

export function Thumbnail({ videoInfo }: { videoInfo: VideoInfo }) {
    return (
        <div style={styles.thumbWrap}>
            <img
                src={videoInfo.thumbnail}
                alt={videoInfo.title}
                style={{
                ...styles.thumb,
                opacity: styles.thumb.opacity,
                }}
            />
            <div style={styles.thumbOverlay} />
            <span style={styles.platformBadge}>
                {platformLabel(videoInfo.platform)}
            </span>
            <span style={styles.tsBadge}>
                {formatTimestamp(videoInfo.timestamp_s)}
            </span>

        </div>
    );
}

export function VideoMeta({ videoInfo }: { videoInfo: VideoInfo }) {
    return (
        <div style={styles.metaBlock}>
            <p style={styles.videoTitle} title={videoInfo.title}>
                {videoInfo.title}
            </p>
            <p style={styles.videoSubtitle}>
                Sharing moment at{" "}
                <span style={styles.timestamp}>
                {formatTimestamp(videoInfo.timestamp_s)}
                </span>
            </p>
        </div>
    );
}

export function ShareButton({
    sharing,
    onShare,
    platform,
}: {
    sharing: boolean;
    onShare: () => void;
    platform: Platform;
}) {
    return (
        <button 
            style={{...styles.btn, opacity: sharing ? 0.7 : 1, cursor: sharing ? "not-allowed" : "pointer" }}
            onClick={onShare}
            disabled={sharing}
        >
            {sharing ? (
                <>
                    <span style={styles.spinner} />
                    Creating link…
                </>
            ) : (
                <>
                    Share on {platformLabel(platform)}
                    <ArrowIcon />
                </>
            )}
        </button>
    )
}

const styles: Record<string, React.CSSProperties> = {
    thumbWrap: {
      position: "relative",
      borderRadius: 12,
      overflow: "hidden",
      marginBottom: 12,
      aspectRatio: "16 / 9",
      background: "#f4f3f0",
    },
    thumb: {
      width: "100%",
      height: "100%",
      objectFit: "cover",
      display: "block",
    },
    thumbOverlay: {
      position: "absolute",
      inset: 0,
      background: "linear-gradient(to top, rgba(26,24,20,0.45) 0%, transparent 55%)",
    },
    platformBadge: {
      position: "absolute",
      top: 8,
      left: 8,
      background: "#FF0000",
      color: "#fff",
      fontFamily: "var(--font-body)",
      fontSize: 10,
      fontWeight: 500,
      padding: "2px 7px",
      borderRadius: 5,
    },
    tsBadge: {
      position: "absolute",
      bottom: 8,
      right: 8,
      background: "rgba(250,250,248,0.92)",
      border: "1px solid rgba(26,24,20,0.1)",
      borderRadius: 6,
      padding: "3px 8px",
      fontFamily: "var(--font-mono)",
      fontSize: 12,
      color: "#d97706",
    },
    metaBlock: {
      marginBottom: 12,
    },
    videoTitle: {
      fontFamily: "var(--font-display)",
      fontSize: 15,
      fontWeight: 700,
      color: "#1a1814",
      marginBottom: 3,
      lineHeight: 1.3,
      display: "-webkit-box",
      WebkitLineClamp: 2,
      WebkitBoxOrient: "vertical",
      overflow: "hidden",
      margin: "0 0 3px",
    },
    videoSubtitle: {
      fontSize: 12,
      color: "#706c65",
      margin: 0,
    },
    timestamp: {
      fontFamily: "var(--font-mono)",
      color: "#d97706",
    },
    btn: {
      width: "100%",
      display: "flex",
      alignItems: "center",
      justifyContent: "center",
      gap: 7,
      background: "#1a1814",
      color: "#fafaf8",
      border: "none",
      borderRadius: 10,
      padding: "11px 16px",
      fontSize: 13,
      fontWeight: 500,
      fontFamily: "var(--font-body)",
      transition: "background 0.15s",
    },
    spinner: {
      display: "inline-block",
      width: 13,
      height: 13,
      border: "2px solid rgba(250,250,248,0.3)",
      borderTopColor: "#fafaf8",
      borderRadius: "50%",
      animation: "spin 0.7s linear infinite",
    },
};
