import React, { useCallback, useMemo, useState } from "react";
import { DownloadImageIcon } from "@components/icons/DownloadImageIcon";
import { ShareIcon } from "@components/icons/ShareIcon";
import { Canvg } from "canvg";

const getImage = (date) => {
  const d = date.toFormat("yyyy-MM-dd");
  const c = document.getElementById("cvs");
  if (!c) return Promise.reject();

  return fetch("/share/image?date="+d)
  .then(res => res.text())
  .then(res => {
    const ctx = c.getContext("2d");
    const v = Canvg.fromString(ctx, res);
    v.start();
    return v.ready()
    .then(() => {
      return {uri: c.toDataURL("image/jpg"), name: `match_${d}.jpg`};
    });
  });
};
const share = (date) => {
  return getImage(date)
    .then(({ uri, name }) => {
      return fetch(uri)
        .then((res) => res.blob())
        .then((blob) => {
          return new File([blob], name, {
            type: blob.type,
            lastModified: new Date().getTime(),
          });
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
    });;
};
const download = (date) => {
  return getImage(date)
    .then(({uri, name}) => {
      const a = document.createElement("a");
      a.href = uri;
      a.download = name
      document.getElementById("opt").appendChild(a);
      a.click();
      setTimeout(() => a.remove(), 1000);
    })
    .catch(err => {
      console.error(err);
    });
};

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
    download(date)
    .finally(() => {
      setLoading("download", false);
    })
  }, [date]);
  const onShare = useCallback(() => {
    setLoading("share", true);
    share(date)
    .finally(() => {
      setLoading("share", false);
    })
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
