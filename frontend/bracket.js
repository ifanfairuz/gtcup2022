import React, { useEffect, useState } from "react";
import { renderToRoot, getData } from "@support/render";
import { formatDate } from "@support/date";
import { Navbar, Bracket } from "@components";
import { DateTime } from "luxon";

function generateMatch(match, aditional = {}) {
  const date = formatDate(DateTime.fromISO(match.Date));
  const temp = match.Label.split("|");
  let info = {
    date,
    id: match.ID,
    label: temp[0],
    home: match.TeamHome.Name,
    away: match.TeamAway.Name,
    agregat: "",
    winner: "",
    poin: { home: 0, away: 0 },
    done: match.Done,
    ...aditional,
  };
  if (!match.TeamAwayId || !match.TeamHomeId) {
    info.home = temp[1];
    info.away = temp[2];
  }

  // done
  if (match.Done) {
    const poinHome = match.Sets.filter(
      (s) => s.Winner === match.TeamHomeId
    ).length;
    const poinAway = match.Sets.filter(
      (s) => s.Winner === match.TeamAwayId
    ).length;
    info.agregat = `${poinHome}-${poinAway}`;
    info.winner =
      poinHome > poinAway ? "home" : poinHome < poinAway ? "away" : "";
    info.poin = {
      home: poinHome,
      away: poinAway,
    };
  }

  return info;
}

const App = ({ matches }) => {
  const [data, setData] = useState([]);

  useEffect(() => {
    let b = {};
    let third;
    for (const match of matches) {
      if (match.Round == 3) {
        third = {
          id: match.ID,
          match: generateMatch(match),
        };
        continue;
      }
      b[match.Round] = b[match.Round] ?? {
        title: "",
        seeds: [],
      };
      b[match.Round].seeds.push({
        id: match.ID,
        match: generateMatch(match),
        third: match.Round == 4 ? third : undefined,
      });
    }
    setData(Object.values(b));
  }, []);

  return (
    <div
      className="is-flex is-flex-direction-column"
      style={{ background: "#f8f9fa", height: "100vh" }}
    >
      <Navbar active="bracket" />
      <Bracket rounds={data} />
    </div>
  );
};

const data = getData();
renderToRoot(<App {...data} />);
