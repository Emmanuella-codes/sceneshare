export default function HeroLinkPreview() {
    return (
        <div className="rounded-2xl border border-border bg-surface overflow-hidden text-left shadow-2xl">
            {/* Fake thumbnail */}
            <div className="relative aspect-video bg-elevated overflow-hidden">
                <div className="absolute inset-0 bg-linear-to-br from-amber-500/10 to-transparent" />
                <div className="absolute inset-0 flex items-center justify-center">
                <div className="w-12 h-12 rounded-full bg-white/10 flex items-center justify-center backdrop-blur-sm">
                    <svg
                        className="w-5 h-5 text-white ml-0.5"
                        fill="currentColor"
                        viewBox="0 0 24 24"
                    >
                    <path d="M8 5v14l11-7z" />
                    </svg>
                </div>
                </div>
                {/* Timestamp badge */}
                <div className="absolute bottom-3 right-3 rounded-md bg-canvas/90 px-2 py-0.5 font-mono text-xs text-amber-400 backdrop-blur-sm border border-border">
                    1:23
                </div>
                {/* Platform badge */}
                <div className="absolute top-3 left-3 rounded-md bg-[#FF0000]/90 px-2 py-0.5 text-xs font-medium text-white">
                    YouTube
                </div>
            </div>
            {/* Info */}
            <div className="p-4">
                <p className="text-sm font-medium text-text-primary truncate">
                    Someone shared a moment with you
                </p>
                <p className="text-xs text-text-secondary mt-1">
                    Opens at 1:23 · sceneshare.app/r/ab3x7f2
                </p>
            </div>
        </div>
    );
}
