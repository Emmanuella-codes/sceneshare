import { VideoOffIcon } from "./icons/icons";

interface Props {
    title?: string;
    description?: string;
}

export default function UnsupportedState({
    title = "Unsupported page",
    description = "Open a YouTube video and click the extension icon to share a moment.",
}: Props) {
    return (
        <div style={styles.centered}>
          <div style={styles.iconCircle}>
            <VideoOffIcon />
          </div>
          <p style={styles.title}>{title}</p>
          <p style={styles.desc}>{description}</p>
        </div>
      );
}

const styles: Record<string, React.CSSProperties> = {
    centered: {
      display: "flex",
      flexDirection: "column",
      alignItems: "center",
      textAlign: "center",
      padding: "1.5rem 1rem",
      gap: 8,
    },
    iconCircle: {
      width: 44,
      height: 44,
      borderRadius: "50%",
      background: "#f4f3f0",
      border: "1px solid #e8e6e1",
      display: "flex",
      alignItems: "center",
      justifyContent: "center",
      marginBottom: 4,
    },
    title: {
      fontFamily: "var(--font-display)",
      fontSize: 17,
      fontWeight: 700,
      color: "#1a1814",
      margin: 0,
    },
    desc: {
      fontSize: 13,
      color: "#706c65",
      lineHeight: 1.55,
      maxWidth: 220,
      margin: 0,
    },
};
