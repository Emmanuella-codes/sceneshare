import { useEffect, useState } from 'react';
import './App.css';
import { Status } from '@/lib/types';
import type { VideoInfo } from '@sceneshare/types';
import { Header } from './components/Header';
import { LoadingState } from './components/LoadingState';
import UnsupportedState from './components/UnsupportedState';
import { VideoCard } from './components/VideoCard';
import SuccessState from './components/SuccessState';
import { ErrorState } from './components/ErrorState';
import { inspectCurrentTab, shareMoment } from './lib/popupFlow';

function App() {
  const [status, setStatus] = useState<Status>("loading");
  const [videoInfo, setVideoInfo] = useState<VideoInfo | null>(null);
  const [shortUrl, setShortUrl] = useState<string | null>(null);
  const [copied, setCopied] = useState(false);
  const [errorMsg, setErrorMsg] = useState<string | null>(null);

  const syncCurrentVideoInfo = async () => {
    try {
      const result = await inspectCurrentTab();

      if (result.status === "unsupported") {
        setVideoInfo(null);
        setStatus("unsupported");
        return null;
      }

      if (result.status === "no-video") {
        setVideoInfo(null);
        setStatus("no-video");
        return null;
      }

      if (result.status === "ready") {
        setVideoInfo(result.videoInfo);
        setStatus("ready");
        return result.videoInfo;
      }

      return null;
    } catch (error) {
      setVideoInfo(null);
      setErrorMsg(error instanceof Error ? error.message : "Could not inspect the current tab");
      setStatus("error");
      return null;
    }
  };

  useEffect(() => {
    void syncCurrentVideoInfo();
  }, []);

  const handleShare = async () => {
    setErrorMsg(null); 

    const latestVideoInfo = await syncCurrentVideoInfo();
    if (!latestVideoInfo) {
      return;
    }

    setStatus("sharing");

    try {
      const nextShortUrl = await shareMoment(latestVideoInfo);

      let didCopy = false;
      try {
        await navigator.clipboard.writeText(nextShortUrl);
        didCopy = true;
      } catch {
        didCopy = false;
      }
      setCopied(didCopy);
      setShortUrl(nextShortUrl);
      setStatus("done");
    } catch (error) {
      setErrorMsg(error instanceof Error ? error.message : "Something went wrong");
      setStatus("error");
    }
  };
 
  const handleRetry = () => {
    setErrorMsg(null);
    setStatus("loading");
    void syncCurrentVideoInfo();
  };
 
  const handleReset = () => {
    setShortUrl(null);
    setCopied(false);
    setErrorMsg(null);
    setStatus("loading");
    void syncCurrentVideoInfo();
  };

  return (
    <div style={styles.root}>
      <Header />
      <div style={styles.body}>
        {status === "loading" && <LoadingState />}
        {status === "unsupported" && (
          <UnsupportedState
            title="Unsupported page"
            description="Open a YouTube watch page or Short and click the extension icon to share a moment."
          />
        )}
        {status === "no-video" && (
          <UnsupportedState
            title="No video detected"
            description="Refresh the YouTube page, start the video, then open the extension again."
          />
        )}
        {(status === "ready" || status === "sharing") && videoInfo && (
          <VideoCard
            videoInfo={videoInfo}
            sharing={status === "sharing"}
            onShare={handleShare}
          />
        )}
        {status === "done" && shortUrl && (
          <SuccessState shortUrl={shortUrl} copied={copied} onReset={handleReset} />
        )}
        {status === "error" && (
          <ErrorState message={errorMsg ?? "Unknown error"} onRetry={handleRetry} />
        )}
      </div>
    </div>
  );
}

export default App;

const styles: Record<string, React.CSSProperties> = {
  root: {
    width: 320,
    background: "#fafaf8",
    display: "flex",
    flexDirection: "column",
  },
  body: {
    padding: "14px 16px 16px",
  },
};
