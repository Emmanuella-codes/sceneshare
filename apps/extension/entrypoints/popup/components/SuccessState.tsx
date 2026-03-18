import { useState } from "react";
import { CheckIcon, CopyIcon } from "./icons/icons";

interface Props {
    shortUrl: string;
    copied: boolean;
    onReset: () => void;
}

export default function SuccessState({ shortUrl, copied, onReset }: Props) {
    const [copyStatus, setCopyStatus] = useState(copied);

    const handleCopy = async () => {
        try {
            await navigator.clipboard.writeText(shortUrl);
            setCopyStatus(true);
            setTimeout(() => setCopyStatus(false), 2000);
        } catch {
            setCopyStatus(false);
        }
    };
    
    return (
        <div style={{...styles.centered, animation: "fadeUp 0.3s ease forwards" }}>
            <div style={styles.iconCircle}>
                <CheckIcon color="#16a34a" size={20} />
            </div>
        
            <p style={styles.title}>Link ready!</p>
            <p style={styles.desc}>{copyStatus ? "Copied to clipboard" : "Link ready to copy"}</p>
        
            <div style={styles.urlBox}>
                <span style={styles.urlText}>{shortUrl}</span>
                <button style={styles.copyBtn} onClick={handleCopy} title="Copy link">
                {copyStatus
                    ? <CheckIcon color="#16a34a" size={14} />
                    : <CopyIcon />
                }
                </button>
            </div>
        
            <button style={styles.ghostBtn} onClick={onReset}>
                Share another moment
            </button>
        </div>
    );
}

const styles: Record<string, React.CSSProperties> = {
    centered: {
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        textAlign: "center",
        padding: "1rem 0",
        gap: 10,
    },
    iconCircle: {
        width: 48,
        height: 48,
        borderRadius: "50%",
        background: "#dcfce7",
        border: "1px solid #bbf7d0",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        marginBottom: 2,
    },
    title: {
        fontFamily: "var(--font-display)",
        fontSize: 18,
        fontWeight: 700,
        color: "#1a1814",
        margin: 0,
    },
    desc: {
        fontSize: 13,
        color: "#706c65",
        margin: 0,
    },
    urlBox: {
        display: "flex",
        alignItems: "center",
        gap: 8,
        background: "#f4f3f0",
        border: "1px solid #e8e6e1",
        borderRadius: 8,
        padding: "8px 10px",
        width: "100%",
    },
    urlText: {
        flex: 1,
        fontFamily: "var(--font-mono)",
        fontSize: 11,
        color: "#706c65",
        overflow: "hidden",
        textOverflow: "ellipsis",
        whiteSpace: "nowrap",
    },
    copyBtn: {
        background: "none",
        border: "none",
        cursor: "pointer",
        padding: 2,
        display: "flex",
        alignItems: "center",
        color: "#706c65",
        flexShrink: 0,
    },
    ghostBtn: {
        background: "none",
        border: "1px solid #e8e6e1",
        borderRadius: 10,
        padding: "9px 16px",
        fontSize: 12,
        color: "#706c65",
        fontFamily: "var(--font-body)",
        cursor: "pointer",
        width: "100%",
        marginTop: 2,
    },
};
