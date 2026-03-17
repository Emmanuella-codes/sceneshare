import type { Metadata } from "next";
import { DM_Mono, DM_Sans, Playfair_Display, } from "next/font/google";
import "./globals.css";

const displayFont = Playfair_Display({
  subsets: ["latin"],
  variable: "--font-display",
  display: "swap",
});
 
const bodyFont = DM_Sans({
  subsets: ["latin"],
  variable: "--font-body",
  display: "swap",
});
 
const monoFont = DM_Mono({
  subsets: ["latin"],
  weight: ["400", "500"],
  variable: "--font-mono",
  display: "swap",
});

export const metadata: Metadata = {
  title: {
    default: "Sceneshare - Share the moment",
    template: "%s | Sceneshare",
  },
  description: "Share any moment with a single link.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${displayFont.variable} ${bodyFont.variable} ${monoFont.variable}`}
      >
        {children}
      </body>
    </html>
  );
}
