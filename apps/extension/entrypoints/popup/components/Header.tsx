export function Header() {
    return (
      <div style={styles.header}>
        <span style={styles.logo}>
          Scene<span style={{ color: "#d97706" }}>Share</span>
        </span>
      </div>
    );
}
   
const styles: Record<string, React.CSSProperties> = {
    header: {
      display: "flex",
      alignItems: "center",
      justifyContent: "space-between",
      padding: "14px 16px 12px",
      borderBottom: "1px solid #f0ede8",
    },
    logo: {
      fontFamily: "var(--font-display)",
      fontSize: 18,
      color: "#1a1814",
      letterSpacing: "-0.02em",
    },
};
