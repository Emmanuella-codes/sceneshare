import { createLink } from '@/lib/api';
import type { ShareMomentMessage, ShareMomentResponse } from '@/types/messages';

export default defineBackground(() => {
  browser.runtime.onMessage.addListener((message, _sender, sendResponse) => {
    const payload = message as ShareMomentMessage;

    if (payload.action !== 'shareMoment') {
      return;
    }

    void (async () => {
      try {
        const link = await createLink({
          platform: payload.videoInfo.platform,
          content_id: payload.videoInfo.content_id,
          timestamp_s: payload.videoInfo.timestamp_s,
          title: payload.videoInfo.title,
          thumbnail: payload.videoInfo.thumbnail,
        });

        sendResponse({
          success: true,
          shortUrl: link.short_url,
        } satisfies ShareMomentResponse);
      } catch (error) {
        sendResponse({
          success: false,
          error: error instanceof Error ? error.message : 'Failed to share moment',
        } satisfies ShareMomentResponse);
      }
    })();

    return true;
  });
});
