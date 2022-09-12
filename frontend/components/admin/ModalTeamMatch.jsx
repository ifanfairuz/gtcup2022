import React, {
  useCallback,
  useEffect,
  useMemo,
  useRef,
  useState,
} from "react";
import { Dropdown } from "@components/form/Dropdown";
import { SetForm } from "./SetForm";

function generateOption(match) {
  return {
    value: `${match.ID}`,
    label: `${match.info.label} - ${match.info.home} vs ${match.info.away}`,
    match,
  };
}

export const ModalTeamMatch = ({ onClose, matches, teams, edit }) => {
  const submitButton = useRef();
  const matchesOptions = useMemo(
    () => matches.map((m) => generateOption(m)),
    [matches]
  );
  const teamsOptions = useMemo(
    () => [
      { value: 0, label: "No Team" },
      ...teams.map((m) => ({ value: m.ID, label: m.Name })),
    ],
    [teams]
  );
  const [match, setMatch] = useState(edit);
  const [home, setHome] = useState(edit.TeamHomeId ? edit.TeamHomeId : null);
  const [away, setAway] = useState(edit.TeamAwayId ? edit.TeamAwayId : null);

  const submit = (e) => {
    submitButton.current.click(e);
  };

  useEffect(() => {
    setHome(edit.TeamHomeId ? edit.TeamHomeId : null);
    setAway(edit.TeamAwayId ? edit.TeamAwayId : null);
  }, [match]);

  return (
    <div className="modal is-active">
      <input type="hidden" name="act" value="update_team" />
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
                <p className="title is-6 mb-1">Home</p>
                <Dropdown
                  className="w-full"
                  placeholder="Select Team"
                  options={teamsOptions}
                  value={home}
                  onChange={(id) => setHome(id)}
                />
                <input type="hidden" name="home_id" value={home} />
              </div>
              <div className="column is-full">
                <p className="title is-6 mb-1">Away</p>
                <Dropdown
                  className="w-full"
                  placeholder="Select Team"
                  options={teamsOptions}
                  value={away}
                  onChange={(id) => setAway(id)}
                />
                <input type="hidden" name="away_id" value={away} />
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
