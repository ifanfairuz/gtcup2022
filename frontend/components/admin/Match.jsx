import React, { useMemo, useState } from "react";
import { ModalInputMatch } from "./ModalInputMatch";
import { ModalTeamMatch } from "./ModalTeamMatch";
import { DateTime } from "luxon";
import { formatDate } from "@support/date";

function generateInfo(match) {
  const d = DateTime.fromISO(match.Date);
  const date = formatDate(d);
  if (match.Type == "B") {
    const temp = match.Label.split("|");
    if (match.TeamAwayId && match.TeamHomeId) {
      return {
        ...match,
        info: {
          label: temp[0],
          home: match.TeamHome.Name,
          away: match.TeamAway.Name,
          date,
          d,
        },
      };
    } else {
      return {
        ...match,
        info: {
          label: temp[0],
          home: temp[1],
          away: temp[2],
          date,
          d,
        },
      };
    }
  }

  return {
    ...match,
    info: {
      label: match.Label,
      home: match.TeamHome.Name,
      away: match.TeamAway.Name,
      date,
      d,
    },
  };
}

export const Match = ({ datas, teams }) => {
  const [modalInput, setModalInput] = useState(null);
  const [modalTeam, setModalTeam] = useState(null);
  const matches = useMemo(() => datas.map((d) => generateInfo(d)), [datas]);

  return (
    <form
      action="/bla/match/update"
      method="post"
      className="has-background-white"
    >
      <div className="p-2">
        <div className="table-container">
          <table className="table is-striped is-fullwidth is-narrow">
            <thead>
              <tr>
                <th width="100px">#</th>
                <th width="50px">Id</th>
                <th>Info</th>
                <th width="200px">Tanding</th>
                <th width="180px">Home</th>
                <th width="180px">Away</th>
                <th width="180px">Pemenang</th>
                <th style={{ minWidth: "300px" }}>Set</th>
                <th width="180px">Image</th>
              </tr>
            </thead>
            <tbody>
              {matches.map((m) => {
                return (
                  <tr key={m.ID}>
                    <td>
                      <div className="field has-addons">
                        <p className="control">
                          <button
                            type="button"
                            className="button is-small is-warning"
                            onClick={() => setModalInput(m)}
                          >
                            edit
                          </button>
                        </p>
                        <p className="control">
                          <a
                            type="button"
                            className="button is-small is-success"
                            href={`/bla/genimage?date=${m.info.d.toFormat('yyyy-MM-dd')}`}
                          >
                            gen
                          </a>
                        </p>
                        {m.Type == "B" && (
                          <p className="control">
                            <button
                              type="button"
                              className="button is-small is-info"
                              onClick={() => setModalTeam(m)}
                            >
                              tim
                            </button>
                          </p>
                        )}
                      </div>
                    </td>
                    <td>{m.ID}</td>
                    <td>{m.info.label}</td>
                    <td>
                      <div className="is-flex is-align-items-center">
                        <div
                          className={`dot p-1 mr-1 is-rounded-all has-background-${
                            m.Done ? "success" : "danger"
                          }`}
                        />
                        <span>{m.info.date}</span>
                      </div>
                    </td>
                    <td>{m.info.home}</td>
                    <td>{m.info.away}</td>
                    <td>
                      {m.Done
                        ? m.Winner == m.TeamHomeId
                          ? m.info.home
                          : m.Winner == m.TeamAwayId
                          ? m.info.away
                          : "-"
                        : "-"}
                    </td>
                    <td>
                      <strong className="mr-1">
                        {m.Sets.filter((s) => s.Winner == m.TeamHomeId).length}-
                        {m.Sets.filter((s) => s.Winner == m.TeamAwayId).length}
                      </strong>
                      {m.Sets.map((s) => (
                        <span key={s.ID} className="mr-1">
                          <strong>#{s.Key}</strong>[{s.Home} - {s.Away}]
                        </span>
                      ))}
                    </td>
                    <td>{m.Image || "NO"}</td>
                  </tr>
                );
              })}
            </tbody>
          </table>
        </div>
      </div>
      {!!modalInput && (
        <ModalInputMatch
          onClose={() => setModalInput(null)}
          edit={
            modalInput
              ? typeof modalInput !== "boolean"
                ? modalInput
                : null
              : null
          }
          matches={matches}
        />
      )}
      {!!modalTeam && (
        <ModalTeamMatch
          onClose={() => setModalTeam(null)}
          edit={
            modalTeam
              ? typeof modalTeam !== "boolean"
                ? modalTeam
                : null
              : null
          }
          matches={matches}
          teams={teams}
        />
      )}
    </form>
  );
};
