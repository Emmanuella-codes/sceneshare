import { platforms, steps } from "@/data/data";
import HeroLinkPreview from "./HeroLinkPreview";
import Footer from "./Footer";

export default function Landing() {
    return (
        <main>
            <nav 
                className="fixed top-0 inset-x-0 z-50 flex items-center justify-between px-6 py-5 border-b border-border-subtle bg-canvas/80 backdrop-blur-md"
            >
                <span className="font-display text-xl text-text-primary tracking-tight">
                    Scene<span className="text-amber-500">Share</span>
                </span>
                <a
                    href="https://chromewebstore.google.com"
                    className="text-sm text-text-secondary hover:text-text-primary transition-colors"
                >
                    Get the extension →
                </a>
            </nav>

            <section className="relative flex flex-col items-center justify-center min-h-screen px-6 text-center pt-20">
                <div className="absolute inset-0 pointer-events-none" aria-hidden="true">
                    <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[600px] h-[400px] rounded-full bg-amber-500/5 blur-[120px]" />
                </div>
                <div className="animate-fade-in mb-8 inline-flex items-center gap-2 rounded-full border border-amber-500/20 bg-amber-500/5 px-4 py-1.5">
                    <span className="h-1.5 w-1.5 rounded-full bg-amber-500 animate-pulse-dot" />
                    <span className="text-xs font-mono text-amber-400 tracking-widest uppercase">
                        YouTube supported
                    </span>
                </div>
                <h1 
                    className="animate-fade-up font-display text-5xl sm:text-7xl lg:text-8xl font-bold leading-[0.95] tracking-tight text-text-primary mb-6"
                    style={{ animationDelay: "0.1s", opacity: 0 }}
                >
                    Share the
                    <br />
                    <span className="italic text-amber-400">exact moment</span>
                </h1>
                <p
                    className="animate-fade-up max-w-lg text-lg text-text-secondary leading-relaxed mb-10"
                    style={{ animationDelay: "0.2s", opacity: 0 }}
                >
                    Stop saying &ldquo;skip to 1:23&rdquo;. SceneShare turns any timestamp
                    into a link that opens the video right where you mean.
                </p>
                <div
                    className="animate-fade-up flex flex-col sm:flex-row gap-3"
                    style={{ animationDelay: "0.3s", opacity: 0 }}
                >
                    <a
                        href="https://chromewebstore.google.com"
                        className="inline-flex items-center gap-2 rounded-full bg-amber-500 px-7 py-3 text-sm font-medium text-canvas hover:bg-amber-400 transition-colors"
                    >
                        Add to Chrome — it&apos;s free
                    </a>
                    <a
                        href="#how-it-works"
                        className="inline-flex items-center gap-2 rounded-full border border-border px-7 py-3 text-sm font-medium text-text-secondary hover:text-text-primary hover:border-border transition-colors"
                    >
                        See how it works
                    </a>
                </div>
                <div
                    className="animate-fade-up mt-16 w-full max-w-sm"
                    style={{ animationDelay: "0.4s", opacity: 0 }}
                >
                    <HeroLinkPreview />
                </div>
            </section>
            <section id="how-it-works" className="py-32 px-6 border-t border-border-subtle">
                <div className="max-w-4xl mx-auto">
                    <h2 className="font-display text-4xl sm:text-5xl font-bold text-text-primary mb-4 text-center">
                        Three steps.
                    </h2>
                    <p className="text-text-secondary text-center mb-16">
                        No account needed. No timestamps to type.
                    </p>
            
                    <div className="grid sm:grid-cols-3 gap-6">
                        {steps.map((step, idx) => (
                        <div
                            key={idx}
                            className="relative rounded-2xl border border-border bg-surface p-6"
                        >
                            <div className="font-mono text-5xl font-bold text-border mb-4 select-none">
                                0{idx + 1}
                            </div>
                            <h3 className="font-display text-xl font-semibold text-text-primary mb-2">
                                {step.title}
                            </h3>
                            <p className="text-sm text-text-secondary leading-relaxed">
                                {step.description}
                            </p>
                        </div>
                        ))}
                    </div>
                </div>
            </section>
            <section className="py-20 px-6 border-t border-border-subtle">
            <div className="max-w-4xl mx-auto text-center">
                <p className="text-xs font-mono text-text-muted uppercase tracking-widest mb-8">
                    Platform support
                </p>
                <div className="flex flex-wrap justify-center gap-4">
                    {platforms.map((p) => (
                    <div
                        key={p.name}
                        className={`flex items-center gap-2.5 rounded-full border px-4 py-2 text-sm ${
                            p.live
                                ? "border-amber-500/30 bg-amber-500/5 text-amber-400"
                                : "border-border bg-surface text-text-muted"
                        }`}
                    >
                        <span
                            className={`h-1.5 w-1.5 rounded-full ${p.live ? "bg-amber-500" : "bg-text-secondary"}`}
                        />
                        {p.name}
                        {!p.live && (
                            <span className="text-xs text-text-secondary">soon</span>
                        )}
                    </div>
                    ))}
                </div>
            </div>
        </section>
        <Footer />
        </main>
    );
}