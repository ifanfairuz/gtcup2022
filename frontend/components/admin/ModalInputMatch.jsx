import React, { useCallback, useEffect, useMemo, useState } from "react";
import { Dropdown } from "@components/form/Dropdown";
import { SetForm } from "./SetForm";

function generateOption(match) {
  return {
    value: `${match.ID}`,
    label: `${match.info.label} - ${match.info.home} vs ${match.info.away}`,
    match,
  };
}

export const ModalInputMatch = ({ onClose, matches, edit }) => {
  const matchesOptions = useMemo(
    () => matches.map((m) => generateOption(m)),
    [matches]
  );
  const [match, setMatch] = useState(edit);
  const [sets, setSets] = useState([]);

  const addSet = useCallback(() => {
    if (!match) return;
    setSets((s) => [
      ...s,
      {
        ID: null,
        MatchId: match.ID,
        Key: s.length + 1,
        Home: 0,
        Away: 0,
        Desc: "",
      },
    ]);
  }, [match]);
  const deleteSet = (key) => {
    setSets((s) => {
      let ss = [...s];
      ss.splice(key, 1);
      return ss;
    });
  };

  useEffect(() => {
    setSets(match ? match.Sets || [] : []);
  }, [match]);

  return (
    <div className="modal is-active">
      <div className="modal-background"></div>
      <div className="modal-content">
        <div className="card">
          <div className="card-content" style={{ minHeight: "300px" }}>
            <div className="columns is-multiline">
              <div className="column is-full">
                <p className="title is-6 mb-1">Teams</p>
                <Dropdown
                  className="w-full"
                  placeholder="Select Match"
                  options={matchesOptions}
                  value={match ? `${match.ID}` : null}
                  onChange={(id, { match }) => setMatch(match)}
                />
              </div>
              <div className="column is-full">
                <p className="title is-6">Sets</p>
                <div className="columns is-multiline">
                  {sets.map((set, i) => (
                    <SetForm
                      key={set.ID || `new-${i}`}
                      match={match}
                      set={set}
                      onChange={(se) =>
                        setSets((s) => {
                          let ss = [...s];
                          ss[i] = se;
                          return ss;
                        })
                      }
                      onDelete={() => deleteSet(i)}
                    />
                  ))}
                  <div className="column is-full">
                    <button
                      type="button"
                      className="button is-fullwidth is-dark"
                      onClick={addSet}
                    >
                      Add Set
                    </button>
                  </div>
                </div>
              </div>
              <div className="column is-full">
                <p className="title is-6">Pertandingan</p>
                <label class="checkbox">
                  <input
                    type="checkbox"
                    className="mr-1"
                    defaultChecked={match.Done}
                  />
                  Selesai
                </label>
              </div>
            </div>
          </div>
          <footer className="card-footer">
            <a
              href="#"
              className="card-footer-item"
              onClick={(e) => !!onClose && onClose(e)}
            >
              Cancel
            </a>
            <a href="#" className="card-footer-item">
              Save
            </a>
          </footer>
        </div>
      </div>
      <button
        type="button"
        onClick={(e) => !!onClose && onClose(e)}
        className="modal-close is-large"
        aria-label="close"
      ></button>
    </div>
  );
};
