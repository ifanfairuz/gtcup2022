import React, { useCallback, useMemo, useState } from "react";
import { DownloadImageIcon } from "@components/icons/DownloadImageIcon";
import { ShareIcon } from "@components/icons/ShareIcon";

export const ShareButtons = ({ size, date, inverted }) => {
  const btnsize = useMemo(() => size || "normal", [size]);
  const [loading, setLoadingState] = useState({
    share: false,
    download: false,
  });

  const setLoading = (key, load) =>
    setLoadingState((s) => ({ ...s, [key]: load }));
  const onDownload = useCallback(() => {
    setLoading("download", true);
    setTimeout(() => {
      setLoading("download", false);
    }, 1000);
  }, [date]);
  const onShare = useCallback(() => {
    setLoading("share", true);
    setTimeout(() => {
      setLoading("share", false);
    }, 1000);
  }, [date]);

  return (
    <div className="field has-addons">
      <p className="control">
        <button
          className={`button is-link ${
            inverted ? "is-inverted" : ""
          } is-${btnsize} ${loading.download && !inverted ? "is-loading" : ""}`}
          disabled={loading.share || loading.download}
          type="button"
          onClick={onDownload}
        >
          <DownloadImageIcon className={`icon is-${btnsize} mr-1`} />
          <span>{loading.download ? "loading..." : "Unduh"}</span>
        </button>
      </p>
      <p className="control">
        <button
          className={`button is-info ${
            inverted ? "is-inverted" : ""
          } is-${btnsize} ${loading.share && !inverted ? "is-loading" : ""}`}
          disabled={loading.share || loading.download}
          type="button"
          onClick={onShare}
        >
          <ShareIcon className={`icon is-${btnsize} mr-1`} />
          <span>{loading.share ? "loading..." : "Bagikan"}</span>
        </button>
      </p>
    </div>
  );
};
