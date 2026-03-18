import { WarningIcon } from "./icons/icons";

interface Props {
    message: string;
    onRetry: () => void;
}
   
export function ErrorState({ message, onRetry }: Props) {
    return (
      <div style={{ ...styles.centered, animation: "fadeUp 0.3s ease forwards" }}>
        <div style={styles.iconCircle}>
          <WarningIcon />
        </div>
        <p style={styles.title}>Couldn't create link</p>
        <p style={styles.errorMsg}>{message}</p>
        <button style={styles.btn} onClick={onRetry}>
          Try again
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
      padding: "1.5rem 0",
      gap: 8,
    },
    iconCircle: {
      width: 44,
      height: 44,
      borderRadius: "50%",
      background: "#fee2e2",
      border: "1px solid #fecaca",
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
    errorMsg: {
      fontSize: 12,
      color: "#dc2626",
      margin: 0,
      maxWidth: 220,
      lineHeight: 1.5,
    },
    btn: {
      display: "flex",
      alignItems: "center",
      justifyContent: "center",
      background: "#1a1814",
      color: "#fafaf8",
      border: "none",
      borderRadius: 10,
      padding: "10px 24px",
      fontSize: 13,
      fontWeight: 500,
      fontFamily: "var(--font-body)",
      cursor: "pointer",
      marginTop: 4,
    },
};
