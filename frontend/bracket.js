import React, { useEffect, useState } from "react";
import { renderToRoot, getData } from "@support/render";
import { formatDate } from "@support/date";
import { Navbar, Bracket } from "@components";
import { DateTime } from "luxon";

function generateMatch(match, aditional = {}) {
  const temp = match.Label.split("|");
  if (match.TeamAwayId && match.TeamHomeId) {
    return {
      label: temp[0],
      date: match.date,
      home: match.TeamHome.Name,
      away: match.TeamAway.Name,
      ...aditional,
    };
  } else {
    return {
      label: temp[0],
      date: match.date,
      home: temp[1],
      away: temp[2],
      ...aditional,
    };
  }
}

const App = ({ matches }) => {
  const [data, setData] = useState([]);

  useEffect(() => {
    let b = {};
    let third;
    for (const match of matches) {
      const date = formatDate(DateTime.fromISO(match.Date));
      if (match.Round == 3) {
        third = {
          id: match.ID,
          match: generateMatch(match, { date }),
        };
        continue;
      }
      b[match.Round] = b[match.Round] ?? {
        title: "",
        seeds: [],
      };
      b[match.Round].seeds.push({
        id: match.ID,
        match: generateMatch(match, { date }),
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
