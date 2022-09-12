import React, {
  useCallback,
  useEffect,
  useMemo,
  useRef,
  useState,
} from "react";
import { Dropdown } from "@components/form/Dropdown";
import { SetForm } from "./SetForm";
import { DateTime } from "luxon";

function generateOption(match) {
  return {
    value: `${match.ID}`,
    label: `${match.info.label} - ${match.info.home} vs ${match.info.away}`,
    match,
  };
}

export const ModalInputMatch = ({ onClose, matches, edit }) => {
  const submitButton = useRef();
  const matchesOptions = useMemo(
    () => matches.map((m) => generateOption(m)),
    [matches]
  );
  const [match, setMatch] = useState(edit);
  const [sets, setSets] = useState([]);
  const [date, setDate] = useState(
    edit ? DateTime.fromISO(edit.Date).toFormat("yyyy-MM-dd'T'HH:mm") : ""
  );

  const addSet = useCallback(() => {
    if (!match) return;
    setSets((s) => [
      ...s,
      {
        ID: null,
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

  const submit = (e) => {
    submitButton.current.click(e);
  };

  useEffect(() => {
    setSets(match ? match.Sets || [] : []);
    setDate(
      match ? DateTime.fromISO(match.Date).toFormat("yyyy-MM-dd'T'HH:mm") : ""
    );
  }, [match]);

  return (
    <div className="modal is-active">
      <input type="hidden" name="act" value="update_set" />
      <div className="modal-background"></div>
      <div className="modal-content">
        <div className="card">
          <div className="card-content" style={{ minHeight: "300px" }}>
            <div className="columns is-multiline">
              <div className="column is-full">
                <p className="title is-6 mb-1">Match</p>
                <Dropdown
                  className="w-full"
                  placeholder="Select Match"
                  options={matchesOptions}
                  value={match ? `${match.ID}` : null}
                  onChange={(id, { match }) => setMatch(match)}
                />
                <input
                  type="hidden"
                  name="match_id"
                  value={match ? `${match.ID}` : ""}
                />
              </div>
              <div className="column is-full">
                <p className="title is-6 mb-1">Tanggal</p>
                <input
                  type="datetime-local"
                  className="input"
                  value={date}
                  onChange={(e) => setDate(e.target.value)}
                />
                <input
                  type="hidden"
                  name="match_date"
                  value={DateTime.fromFormat(date, "yyyy-MM-dd'T'HH:mm").toRFC2822()}
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
                <input
                  type="hidden"
                  name="sets_json"
                  value={JSON.stringify(sets)}
                />
              </div>
              <div className="column is-full">
                <p className="title is-6">Pertandingan</p>
                <label className="checkbox">
                  <input
                    type="checkbox"
                    className="mr-1"
                    defaultChecked={match.Done}
                    name="match_done"
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
            <a href="#" onClick={(e) => submit(e)} className="card-footer-item">
              Save
            </a>
          </footer>
          <button type="submit" style={{ display: "none" }} ref={submitButton}>
            submit
          </button>
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
