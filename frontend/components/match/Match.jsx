import React, { useMemo } from "react";
import { Ball } from "@components/icons/Ball";
import { Crown } from "@components/icons/Crown";
import { Set } from "./Set";

function generateMatch(data) {
  let info = {
    label: data.Label,
    sublabel: `Grup ${data.Group}`,
    home: data.TeamHome.Name,
    away: data.TeamAway.Name,
    agregat: "",
    winner: "",
    sets: data.Sets,
  };

  // bracket
  if (data.Type == "B") {
    const temp = data.Label.split("|");
    info.label = temp[0];
    info.sublabel = "";
    if (!data.TeamAwayId || !data.TeamHomeId) {
      info.home = temp[1];
      info.away = temp[2];
    }
  }

  // done
  if (data.Done) {
    const poinHome = data.Sets.filter(
      (s) => s.Winner === data.TeamHomeId
    ).length;
    const poinAway = data.Sets.filter(
      (s) => s.Winner === data.TeamAwayId
    ).length;
    info.agregat = `${poinHome}-${poinAway}`;
    info.winner =
      poinHome > poinAway ? "home" : poinHome < poinAway ? "away" : "";

    info.sets = info.sets.map((s) => ({
      ...s,
      winner:
        s.Winner == data.TeamHomeId
          ? "home"
          : s.Winner == data.TeamAwayId
          ? "away"
          : "",
    }));
  }

  return info;
}

export const Match = ({ data, className }) => {
  const item = useMemo(() => generateMatch(data), [data]);

  return (
    <div className={`p-4 flex-1 is-rounded ${className || ""}`}>
      <div className="columns is-gapless">
        <div className="column is-flex is-flex-direction-column">
          <p style={{ opacity: 0.7 }} className="is-size-6">
            {item.label}
          </p>
          <p style={{ opacity: 0.7 }} className="is-size-6 mb-2">
            {item.sublabel}
          </p>
          <div className="flex-1 is-flex is-flex-direction-column is-justify-content-space-around">
            <div className="flex-1 has-text-centered is-flex is-flex-direction-row is-align-items-center mb-4">
              <Ball className="is-size-4 has-text-primary mr-2" />
              <h3 className="has-text-weight-medium">{item.home}</h3>
              {item.winner == "home" && (
                <Crown className="has-text-warning ml-2" />
              )}
            </div>
            <div className="flex-1 has-text-centered is-flex is-flex-direction-row is-align-items-center mb-4">
              <Ball className="is-size-4 has-text-link mr-2" />
              <h3 className="has-text-weight-medium">{item.away}</h3>
              {item.winner == "away" && (
                <Crown className="has-text-warning ml-2" />
              )}
            </div>
          </div>
        </div>
        <div className="column is-one-third is-flex is-flex-direction-column">
          <p style={{ opacity: 0.7 }} className="is-size-6 mb-2">
            {data.Done ? `Agregat: ${item.agregat}` : ""}
          </p>
          {item.sets.map((s) => (
            <Set key={s.ID} data={s} />
          ))}
        </div>
      </div>
    </div>
  );
};
