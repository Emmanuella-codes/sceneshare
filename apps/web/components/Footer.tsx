import { platforms } from "@/data/data";
import { platformLabel } from "@/utils/utils";
import { Link } from "@sceneshare/types";

export default function Footer({ view, link }: { view: "landing" | "redirect"; link?: Link } ) {
    return (
        <footer className="border-t border-black/12 bg-white py-8 px-6 text-center">
            <div className="px-4 py-16 sm:px-6 sm:py-20">
                <div className="max-w-4xl mx-auto text-center">
                    <div className="mb-8 flex justify-center">
                        <p className="max-w-full rounded-full border border-black/15 px-3 py-1 text-[11px] font-mono uppercase tracking-[0.16em] text-black/75 sm:px-4 sm:text-xs sm:tracking-[0.24em]">
                        Platform support
                        </p>
                    </div>
                    <div className="flex flex-wrap justify-center gap-3 sm:gap-4">
                        {platforms.map((p) => (
                        <div
                            key={p.name}
                            className={`flex items-center gap-2.5 rounded-full border px-4 py-2 text-sm ${
                                p.live
                                    ? "border-black/15 bg-white text-black"
                                    : "border-black/12 bg-[#fbfaf7] text-black/55"
                            }`}
                        >
                            <span
                                className={`h-1.5 w-1.5 rounded-full ${p.live ? "bg-black" : "bg-black/35"}`}
                            />
                            {p.name}
                            {!p.live && (
                                <span className="text-xs text-black/45">soon</span>
                            )}
                        </div>
                        ))}
                    </div>
                </div>
            </div>
            {view === "landing" && (
                <p className="text-xs text-black/55 font-mono">SceneShare · Not affiliated with any streaming platform</p>
            )}
            {view === "redirect" && (
                <p className="mt-6 text-center text-xs text-black/55">
                You&apos;ll need to be logged in to{" "}
                {platformLabel(link?.platform ?? "youtube")} to watch.
              </p>
            )}
        </footer>
    );
}
