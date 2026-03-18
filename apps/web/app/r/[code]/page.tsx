import LinkPreview from "@/components/LinkPreview";
import { ExpiredError, getLink, NotFoundError } from "@/app/api/api";
import Footer from "@/components/Footer";
import Header from "@/components/Header";
import type { Link } from "@sceneshare/types";

type PageProps = {
  params: Promise<{
    code: string;
  }>;
  searchParams?: Promise<{
    preview?: string;
  }>;
};

function buildPreviewLink(code: string): Link {
  return {
    id: "preview-link",
    short_code: code,
    short_url: `http://localhost:3001/r/${code}`,
    platform: "youtube",
    content_id: "dQw4w9WgXcQ",
    timestamp_s: 83,
    timestamp_fmt: "1:23",
    title: "SceneShare preview moment",
    thumbnail: "https://i.ytimg.com/vi/dQw4w9WgXcQ/hqdefault.jpg",
    click_count: 42,
    owner_token: null,
    created_at: new Date().toISOString(),
    expires_at: null,
  };
}

function LinkState({
  title,
  message,
}: {
  title: string;
  message: string;
}) {
  return (
    <div className="min-h-screen bg-white">
      <main className="relative flex min-h-[calc(100vh-1px)] flex-col items-center justify-center px-4 py-24 sm:px-6 sm:py-12">
        <div className="fixed inset-0 pointer-events-none" aria-hidden="true">
          <div className="absolute top-1/2 left-1/2 h-[260px] w-[320px] -translate-x-1/2 -translate-y-1/2 rounded-full bg-amber-500/8 blur-[90px] sm:h-[300px] sm:w-[500px] sm:blur-[100px]" />
        </div>
        <Header view="landing" />
        <div className="relative w-full max-w-lg rounded-2xl border border-black/12 bg-[#ffffff] p-8 text-center shadow-2xl glow-amber">
          <h1 className="font-display text-3xl text-gray-600">{title}</h1>
          <p className="mt-3 text-sm text-black">{message}</p>
        </div>
      </main>
      <Footer view="landing" />
    </div>
  );
}

export default async function RedirectPage({ params, searchParams }: PageProps) {
  const { code } = await params;
  const query = searchParams ? await searchParams : undefined;

  if (query?.preview === "1") {
    return <LinkPreview link={buildPreviewLink(code)} />;
  }

  let link = null;
  let state: "ready" | "not-found" | "expired" | "error" = "ready";

  try {
    link = await getLink(code);
  } catch (error) {
    if (error instanceof NotFoundError) {
      state = "not-found";
    } else if (error instanceof ExpiredError) {
      state = "expired";
    } else {
      state = "error";
    }
  }

  if (state === "ready" && link) {
    return <LinkPreview link={link} />;
  }

  if (state === "not-found") {
    return (
      <LinkState
        title="Link not found"
        message="This SceneShare link does not exist or may have been removed."
      />
    );
  }

  if (state === "expired") {
    return (
      <LinkState
        title="Link expired"
        message="This SceneShare link has expired and can no longer be opened."
      />
    );
  }

  return (
    <LinkState
      title="Unable to load link"
      message="There was a problem loading this SceneShare link. Please try again in a moment."
    />
  );
}
