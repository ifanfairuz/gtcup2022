import React, { useCallback } from "react";

export const SetForm = ({ set, match, onChange, onDelete }) => {
  const change = useCallback(
    (key, value) => {
      !!onChange && onChange({ ...set, [key]: value });
    },
    [set, onChange]
  );

  return (
    <div className="column is-full b-all">
      <div className="field is-horizontal">
        <div className="field-label is-small">
          <label className="label">Set Ke</label>
        </div>
        <div className="field-body">
          <div className="field">
            <p className="control">
              <input
                className="input is-small"
                type="number"
                placeholder="Set"
                value={set.Key}
                onChange={(e) => change("Key", parseInt(e.target.value))}
              />
            </p>
          </div>
        </div>
      </div>
      <div className="field is-horizontal">
        <div className="field-label is-small">
          <label className="label">{match ? match.info.home : ""}</label>
        </div>
        <div className="field-body">
          <div className="field">
            <p className="control">
              <input
                className="input is-small"
                type="number"
                placeholder="Poin"
                value={set.Home}
                onChange={(e) => change("Home", parseInt(e.target.value))}
              />
            </p>
          </div>
        </div>
      </div>
      <div className="field is-horizontal">
        <div className="field-label is-small">
          <label className="label">{match ? match.info.away : ""}</label>
        </div>
        <div className="field-body">
          <div className="field">
            <p className="control">
              <input
                className="input is-small"
                type="number"
                placeholder="Poin"
                value={set.Away}
                onChange={(e) => change("Away", parseInt(e.target.value))}
              />
            </p>
          </div>
        </div>
      </div>
      <div className="field is-horizontal">
        <div className="field-label is-small">
          <label className="label">Deskripsi</label>
        </div>
        <div className="field-body">
          <div className="field">
            <p className="control">
              <textarea
                className="textarea is-small"
                rows={2}
                value={set.Desc}
                onChange={(e) => change("Desc", e.target.value)}
              ></textarea>
            </p>
          </div>
        </div>
      </div>
      <button
        type="button"
        className="button is-small is-danger"
        onClick={onDelete}
      >
        delete
      </button>
    </div>
  );
};
