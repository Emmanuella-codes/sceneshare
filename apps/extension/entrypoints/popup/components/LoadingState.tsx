export function LoadingState() {
    return (
      <div style={styles.centered}>
        <div style={styles.spinner} />
      </div>
    );
}
   
const styles: Record<string, React.CSSProperties> = {
    centered: {
      display: "flex",
      flexDirection: "column",
      alignItems: "center",
      justifyContent: "center",
      padding: "2.5rem 1rem",
    },
    spinner: {
      width: 20,
      height: 20,
      border: "2px solid #e8e6e1",
      borderTopColor: "#1a1814",
      borderRadius: "50%",
      animation: "spin 0.7s linear infinite",
    },
};
