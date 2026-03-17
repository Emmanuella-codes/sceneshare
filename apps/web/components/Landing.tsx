import { steps } from "@/data/data";
import HeroLinkPreview from "./HeroLinkPreview";
import Footer from "./Footer";
import Header from "./Header";

export default function Landing() {
    return (
        <main className="overflow-x-hidden bg-white text-black">
            <Header view="landing" />
            <section className="relative flex min-h-screen flex-col items-center justify-center bg-white px-4 pt-24 pb-12 text-center sm:px-6">
                <div className="absolute inset-0 pointer-events-none" aria-hidden="true">
                    <div className="absolute top-1/2 left-1/2 h-[260px] w-[320px] -translate-x-1/2 -translate-y-1/2 rounded-full bg-amber-500/8 blur-[90px] sm:h-[400px] sm:w-[600px] sm:blur-[120px]" />
                </div>
                <div className="animate-fade-in mb-6 inline-flex max-w-full items-center gap-2 rounded-full border border-black/15 bg-white px-3 py-1.5 shadow-[0_8px_30px_rgba(0,0,0,0.04)] sm:mb-8 sm:px-4">
                    <span className="h-1.5 w-1.5 rounded-full bg-black animate-pulse-dot" />
                    <span className="text-[11px] font-mono uppercase tracking-[0.16em] text-black sm:text-xs sm:tracking-widest">
                        YouTube supported
                    </span>
                </div>
                <h1 
                    className="animate-fade-up mb-5 max-w-[10ch] font-display text-4xl leading-[0.95] tracking-tight text-black sm:mb-6 sm:max-w-none sm:text-7xl lg:text-8xl"
                    style={{ animationDelay: "0.1s", opacity: 0 }}
                >
                    Share the
                    <br />
                    <span className="italic text-amber-400">exact moment</span>
                </h1>
                <p
                    className="animate-fade-up mb-8 max-w-md text-base leading-relaxed text-black/70 sm:mb-10 sm:max-w-lg sm:text-lg"
                    style={{ animationDelay: "0.2s", opacity: 0 }}
                >
                    Stop saying &ldquo;skip to 1:23&rdquo;. SceneShare turns any timestamp
                    into a link that opens the video right where you mean.
                </p>
                <div
                    className="animate-fade-up flex w-full max-w-sm flex-col gap-3 sm:max-w-none sm:flex-row sm:justify-center"
                    style={{ animationDelay: "0.3s", opacity: 0 }}
                >
                    <a
                        href="https://chromewebstore.google.com"
                        className="inline-flex w-full items-center justify-center gap-2 rounded-full bg-amber-500 px-6 py-3 text-sm font-medium text-canvas transition-colors hover:bg-amber-400 sm:w-auto sm:px-7"
                    >
                        Add to Chrome — it&apos;s free
                    </a>
                    <a
                        href="#how-it-works"
                        className="inline-flex w-full items-center justify-center gap-2 rounded-full border border-black/15 px-6 py-3 text-sm font-medium text-black/70 transition-colors hover:border-black/30 hover:text-black sm:w-auto sm:px-7"
                    >
                        See how it works
                    </a>
                </div>
                <div
                    className="animate-fade-up mt-12 mb-4 w-full max-w-sm sm:mt-16 sm:mb-8"
                    style={{ animationDelay: "0.4s", opacity: 0 }}
                >
                    <HeroLinkPreview />
                </div>
            </section>
            <section id="how-it-works" className="border-t border-black/12 bg-[#fbfaf7] px-4 py-20 sm:px-6 sm:py-32">
                <div className="max-w-4xl mx-auto">
                    <div className="mb-6 flex justify-center">
                        <span className="max-w-full rounded-full border border-black/15 px-3 py-1 text-[11px] font-mono uppercase tracking-[0.16em] text-black/75 sm:px-4 sm:text-xs sm:tracking-[0.24em]">
                            How it works
                        </span>
                    </div>
                    <h2 className="mb-4 text-center font-display text-3xl font-bold text-black sm:text-5xl">
                        Three steps.
                    </h2>
                    <p className="mb-10 text-center text-black/65 sm:mb-16">
                        No account needed. No timestamps to type.
                    </p>
            
                    <div className="grid gap-4 sm:grid-cols-3 sm:gap-6">
                        {steps.map((step, idx) => (
                        <div
                            key={idx}
                            className="relative rounded-2xl border border-black/12 bg-white p-6 shadow-[0_12px_40px_rgba(0,0,0,0.05)]"
                        >
                            <div className="font-mono text-5xl font-bold text-black/15 mb-4 select-none">
                                0{idx + 1}
                            </div>
                            <h3 className="font-display text-xl font-semibold text-black mb-2">
                                {step.title}
                            </h3>
                            <p className="text-sm text-black/65 leading-relaxed">
                                {step.description}
                            </p>
                        </div>
                        ))}
                    </div>
                </div>
            </section>
        <Footer view="landing" />
        </main>
    );
}
