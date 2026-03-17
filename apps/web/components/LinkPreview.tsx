"use client";

import type { Link } from "@sceneshare/types";
import Image from "next/image";
import { useEffect, useState } from "react";
import Header from "./Header";
import { ArrowIcon, Spinner, VideoIcon } from "./icons/Icons";
import { platformColor, platformLabel } from "@/utils/utils";
import Footer from "./Footer";

const REDIRECT_DELAY = 100;
const API_URL = process.env.NEXT_PUBLIC_API_URL;

interface Props {
  link: Link;
}

export default function LinkPreview({ link }: Props) {
  const [countdown, setCountdown] = useState(REDIRECT_DELAY);
  const [redirecting, setRedirecting] = useState(false);

  useEffect(() => {
    if (countdown <= 0) {
      window.location.href = `${API_URL}/r/${link.short_code}`;
      return;
    }

    const timer = setTimeout(() => setCountdown((c) => c - 1), 1000);
    return () => clearTimeout(timer);
  }, [countdown, link.short_code]);

  const handleRedirect = () => {
    setRedirecting(true);
    window.location.href = `${API_URL}/r/${link.short_code}`;
  };

  return (
    <div className="min-h-screen bg-white">
      <main className="relative flex min-h-[calc(100vh-1px)] flex-col items-center justify-center px-4 py-24 sm:px-6 sm:py-12">
        <div className="fixed inset-0 pointer-events-none" aria-hidden="true">
          <div className="absolute top-1/2 left-1/2 h-[260px] w-[320px] -translate-x-1/2 -translate-y-1/2 rounded-full bg-amber-500/8 blur-[90px] sm:h-[300px] sm:w-[500px] sm:blur-[100px]" />
        </div>
        <div className="relative w-full max-w-lg animate-fade-up">
          <Header view="redirect" link={link} />
        </div>
        <div className="relative w-full max-w-lg overflow-hidden rounded-2xl border border-black/12 bg-[#17171c] shadow-2xl glow-amber">
          <div className="relative aspect-video bg-elevated overflow-hidden">
            {link.thumbnail ? (
              <>
                <Image
                  src={link.thumbnail}
                  alt={link.title ?? "Video thumbnail"}
                  fill
                  className="object-cover"
                  priority
                />
                <div className="absolute inset-0 bg-linear-to-t from-surface/60 via-transparent to-transparent" />
              </>
            ) : (
              <div className="absolute inset-0 bg-linear-to-br from-elevated to-canvas flex items-center justify-center">
                <VideoIcon />
              </div>
            )}
            <div
              className="absolute top-4 left-4 rounded-md px-2.5 py-1 text-xs font-semibold text-white"
              style={{ backgroundColor: platformColor(link.platform) }}
            >
              {platformLabel(link.platform)}
            </div>
            <div className="absolute bottom-4 right-4 rounded-lg border border-white/10 bg-canvas/80 px-3 py-1.5 backdrop-blur-sm">
              <span className="font-mono text-sm font-medium text-amber-400">{link.timestamp_fmt}</span>
            </div>
          </div>
          <div className="p-5">
            {link.title && (
              <h1 className="mb-1 line-clamp-2 font-display text-xl font-semibold text-text-primary">{link.title}</h1>
            )}
            <p className="text-sm text-text-secondary">
              Shared moment · starts at <span className="font-mono text-amber-400">{link.timestamp_fmt}</span>
            </p>
          </div>
          <div className="flex flex-col gap-3 px-5 pb-5">
            <button
              onClick={handleRedirect}
              disabled={redirecting || countdown <= 0}
              className="flex w-full items-center justify-center gap-2 rounded-xl bg-amber-500 py-3.5 text-sm font-semibold text-canvas transition-colors hover:bg-amber-400 disabled:cursor-not-allowed disabled:opacity-60"
            >
              {redirecting || countdown <= 0 ? (
                <>
                  <Spinner />
                  Opening {platformLabel(link.platform)}...
                </>
              ) : (
                <>
                  Watch on {platformLabel(link.platform)}
                  <ArrowIcon />
                </>
              )}
            </button>
            <div className="h-0.5 w-full overflow-hidden rounded-full bg-elevated">
              <div
                className="h-full rounded-full bg-amber-500/50 transition-all duration-1000 ease-linear"
                style={{
                  width: `${((REDIRECT_DELAY - countdown) / REDIRECT_DELAY) * 100}%`,
                }}
              />
            </div>
          </div>
        </div>
      </main>
      <Footer view="redirect" link={link} />
    </div>
  );
}
