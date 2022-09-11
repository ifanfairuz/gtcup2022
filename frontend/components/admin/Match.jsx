import React, { useMemo, useState } from "react";
import { ModalInputMatch } from "./ModalInputMatch";
import { DateTime } from "luxon";
import { formatDate } from "@support/date";

function generateInfo(match) {
  const date = formatDate(DateTime.fromISO(match.Date));
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
    },
  };
}

export const Match = ({ datas }) => {
  const [modalInput, setModalInput] = useState(null);
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
                <th width="50px">#</th>
                <th width="50px">Id</th>
                <th>Info</th>
                <th width="200px">Tanding</th>
                <th width="180px">Home</th>
                <th width="180px">Away</th>
                <th width="180px">Pemenang</th>
                <th style={{ minWidth: "300px" }}>Set</th>
              </tr>
            </thead>
            <tbody>
              {matches.map((m) => {
                const { ID, info, Done, Winner, TeamHomeId, TeamAwayId } = m;
                return (
                  <tr key={ID}>
                    <td>
                      <button
                        type="button"
                        className="button is-small is-warning"
                        onClick={() => setModalInput(m)}
                      >
                        edit
                      </button>
                    </td>
                    <td>{ID}</td>
                    <td>{info.label}</td>
                    <td>
                      <div className="is-flex is-align-items-center">
                        <div
                          className={`dot p-1 mr-1 is-rounded has-background-${
                            Done ? "success" : "danger"
                          }`}
                        />
                        <span>{info.date}</span>
                      </div>
                    </td>
                    <td>{info.home}</td>
                    <td>{info.away}</td>
                    <td>
                      {Done
                        ? Winner == TeamHomeId
                          ? info.home
                          : Winner == TeamAwayId
                          ? info.away
                          : "-"
                        : "-"}
                    </td>
                    <td>
                      <strong>#{} </strong>
                      <span>
                        [H({}), A({})]
                      </span>
                    </td>
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
    </form>
  );
};
