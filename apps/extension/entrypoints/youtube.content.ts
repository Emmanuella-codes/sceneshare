import type { GetVideoInfoMessage, VideoInfoResponse } from '@/types/messages';

export default defineContentScript({
  matches: [
    'https://www.youtube.com/watch*',
    'https://m.youtube.com/watch*',
    'https://www.youtube.com/shorts/*',
    'https://m.youtube.com/shorts/*',
  ],
  main() {
    browser.runtime.onMessage.addListener((msg, _sender, sendResponse) => {
      const message = msg as GetVideoInfoMessage;
      if (message.action === 'getVideoInfo') {
        const video = document.querySelector<HTMLVideoElement>('video');
        const videoID = getYouTubeVideoId();

        if (!video || !videoID) {
          sendResponse({ videoInfo: null } satisfies VideoInfoResponse);
          return;
        }

        sendResponse({
          videoInfo: {
            platform: 'youtube',
            content_id: videoID,
            timestamp_s: Math.floor(video.currentTime),
            title: document.title.replace(' - YouTube', '').trim(),
            thumbnail: `https://i.ytimg.com/vi/${videoID}/hqdefault.jpg`,
          },
        } satisfies VideoInfoResponse);
      }
    });
  },
});

function getYouTubeVideoId(): string | null {
  const { pathname, search } = window.location;

  if (pathname === '/watch') {
    return new URLSearchParams(search).get('v');
  }

  if (pathname.startsWith('/shorts/')) {
    return pathname.split('/')[2] ?? null;
  }

  return null;
}
