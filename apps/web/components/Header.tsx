"use client";

import type { Link } from "@sceneshare/types";
import { useEffect, useState } from "react";

export default function Header({ view, link }: { view: "landing" | "redirect"; link?: Link }) {
    const [scrolled, setScrolled] = useState(false);

    useEffect(() => {
        if (view !== "landing") return;

        const onScroll = () => setScrolled(window.scrollY > 12);
        onScroll();
        window.addEventListener("scroll", onScroll, { passive: true });
        return () => window.removeEventListener("scroll", onScroll);
    }, [view]);

    return (
        <nav 
            className={`fixed top-0 inset-x-0 z-50 flex items-center justify-between px-6 py-5 transition-all duration-300 ${
                view === "landing"
                    ? scrolled
                        ? "border-b border-black/10 bg-white/80 backdrop-blur-md"
                        : "border-b border-transparent bg-transparent"
                    : "border-b border-border-subtle bg-canvas/80 backdrop-blur-md"
            }`}
        >
            <span className={`font-display text-xl tracking-tight ${view === "landing" ? "text-black" : "text-text-primary"}`}>
                Scene<span className="text-amber-500">Share</span>
            </span>
            {view === "landing" && (
               <a
                    href="https://chromewebstore.google.com"
                    className="text-sm text-black/65 hover:text-black transition-colors"
                >
                    Get the extension →
                </a> 
            )}
            {view === "redirect" && (
                <div className="flex items-center gap-2 text-xs text-text-muted font-mono">
                    <span
                        className="h-1.5 w-1.5 rounded-full bg-amber-500 animate-pulse-dot"
                    />
                    {link?.click_count} {link?.click_count === 1 ? "view" : "views"}
                </div>
            )}
            
        </nav>
    );
}
