import React, { useCallback, useMemo, useState } from "react";
import { DownloadImageIcon } from "@components/icons/DownloadImageIcon";
import { ShareIcon } from "@components/icons/ShareIcon";
import { filename } from "@support/string";

const shareImage = (url) => {
  return fetch(url, { mode: 'no-cors', referrerPolicy: 'no-referrer' })
    .then((res) => res.blob())
    .then((blob) => {
      return new File([blob], filename(url), {
        type: blob.type,
        lastModified: new Date().getTime(),
      });
    })
    .then((file) => {
      const data = { files: [file] };
      if (window.navigator.share && window.navigator.canShare(data)) {
        window.navigator.share(data);
      }
    })
    .catch((err) => {
      console.error(err);
    });
};

export const ShareButtons = ({ size, image, inverted }) => {
  const image_url = useMemo(() => `/assets/match/${image}`, [image]);
  const btnsize = useMemo(() => size || "normal", [size]);
  const [loading, setLoadingState] = useState({
    share: false,
    download: false,
  });

  const setLoading = (key, load) =>
    setLoadingState((s) => ({
      ...s,
      [key]: load,
    }));
  const onShare = useCallback(() => {
    setLoading("share", true);
    shareImage(image_url).finally(() => {
      setLoading("share", false);
    });
  }, [image_url]);

  if (!image || image == "") return null;
  return (
    <div className="field has-addons">
      <p className="control">
        <a
          className={`button is-link ${
            inverted ? "is-inverted" : ""
          } is-${btnsize} ${loading.download && !inverted ? "is-loading" : ""}`}
          disabled={loading.share || loading.download}
          href={image_url}
          download="gtcup2022_match.jpg"
        >
          <DownloadImageIcon className={`icon is-${btnsize} mr-1`} />
          <span>{loading.download ? "loading..." : "Unduh"}</span>
        </a>
      </p>
      {!!window.navigator.share && (
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
      )}
    </div>
  );
};