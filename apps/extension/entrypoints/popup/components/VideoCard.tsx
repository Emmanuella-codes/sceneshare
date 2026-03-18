import { VideoInfo } from "@sceneshare/types";
import { ShareButton, Thumbnail, VideoMeta } from "./VideoCardCmp";

interface Props {
    videoInfo: VideoInfo;
    sharing: boolean;
    onShare: () => void;
}

export function VideoCard({ videoInfo, sharing, onShare }: Props) {
    return (
        <div className="">
            <Thumbnail videoInfo={videoInfo} />
            <VideoMeta videoInfo={videoInfo} />
            <ShareButton sharing={sharing} onShare={onShare} platform={videoInfo.platform} />
        </div>
    );
}
