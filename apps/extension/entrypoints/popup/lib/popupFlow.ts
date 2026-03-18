import type { VideoInfo } from '@sceneshare/types';
import type { ShareMomentResponse, VideoInfoResponse } from '@/types/messages';

export type TabInspectionResult =
  | { status: 'ready'; videoInfo: VideoInfo }
  | { status: 'unsupported' | 'no-video' };

// resolves the latest playable moment from the active tab.
export async function inspectCurrentTab(): Promise<TabInspectionResult> {
  let response:  VideoInfoResponse;
  const [tab] = await browser.tabs.query({
    active: true,
    currentWindow: true,
  });

  if (!tab?.id) {
    return { status: 'no-video' };
  }

  const url = tab.url ?? '';
  if (!isSupportedYouTubeUrl(url)) {
    return { status: 'unsupported' };
  }

  try {
    response = (await browser.tabs.sendMessage(tab.id, {
      action: 'getVideoInfo',
    })) as VideoInfoResponse;
  } catch (error) {
    return { status: 'no-video' };
  }

  if (!response?.videoInfo) {
    return { status: 'no-video' };
  }

  return {
    status: 'ready',
    videoInfo: response.videoInfo,
  };
}

// creates a short link for the latest captured video info.
export async function shareMoment(videoInfo: VideoInfo): Promise<string> {
  const response = (await browser.runtime.sendMessage({
    action: 'shareMoment',
    videoInfo,
  })) as ShareMomentResponse;

  if (!response.success || !response.shortUrl) {
    throw new Error(response.error ?? 'Something went wrong');
  }

  return response.shortUrl;
}

function isSupportedYouTubeUrl(url: string): boolean {
  try {
    const parsed = new URL(url);
    const host = parsed.hostname.replace(/^www\./, '');

    if (host !== 'youtube.com' && host !== 'm.youtube.com') {
      return false;
    }

    return parsed.pathname === '/watch' || parsed.pathname.startsWith('/shorts/');
  } catch {
    return false;
  }
}
