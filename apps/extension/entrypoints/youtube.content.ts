export default defineContentScript({
  matches: ['https://www.youtube.com/watch*'],
  main() {
    browser.runtime.onMessage.addListener((msg, _sender, sendResponse) => {
      if (msg.action === 'getVideoInfo') {
        const video = document.querySelector<HTMLVideoElement>('video');
        const params = new URLSearchParams(window.location.search);
        const videoID = params.get('v');

        if (!video || !videoID) {
          sendResponse(null);
          return;
        }

        sendResponse({
          platform: 'youtube',
          contentId: videoID,
          timestamp_s: Math.floor(video.currentTime),
          title: document.title.replace(' - YouTube', '').trim(),
          thumbnail: `https://i.ytimg.com/vi/${videoID}/hqdefault.jpg`,
        })
      }
    })
  },
});
