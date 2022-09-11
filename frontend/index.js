import React, { useEffect, useState } from "react";
import { renderToRoot, getData } from "@support/render";
import { Navbar, Footer, RoundMatch, BracketMatch, Match } from "@components";
import { DateTime } from "luxon";
import { formatDate } from "@support/date";

const App = ({ matches, lastMatches, nextMatches }) => {
  const [group, setGroup] = useState({});
  const [bracket, setBracket] = useState({});

  useEffect(() => {
    let g = {};
    let b = {};
    for (const match of matches) {
      const date = formatDate(DateTime.fromISO(match.Date));
      if (match.Type == "G") {
        g[date] = g[date] ?? {
          title: date,
          datas: [],
        };

        g[date].datas.push(match);
      } else {
        b[match.Round] = b[match.Round] ?? {
          title: match.Group,
          datas: {},
        };
        b[match.Round].datas[date] = {
          title: date,
          match: match,
        };
      }
    }
    setGroup(g);
    setBracket(b);
  }, []);

  return (
    <div>
      <Navbar active="pertandingan" />
      <div className="container has-background-white">
        {!!lastMatches && lastMatches.length > 0 && (
          <div className="p-4">
            <h4 className="is-size-4 has-text-weight-semibold has-text-centered">
              Pertandingan Terakhir
            </h4>
            <p className="has-text-centered mb-4">
              {formatDate(DateTime.fromISO(lastMatches[0].Date))}
            </p>
            <div className="columns">
              {lastMatches.map((n) => (
                <div key={n.ID} className="column is-flex">
                  <Match data={n} />
                </div>
              ))}
            </div>
          </div>
        )}
        {!!nextMatches && nextMatches.length > 0 && (
          <div className="p-4">
            <h4 className="is-size-4 has-text-weight-semibold has-text-centered">
              Pertandingan Selanjutnya
            </h4>
            <p className="has-text-centered mb-4">
              {formatDate(DateTime.fromISO(nextMatches[0].Date))}
            </p>
            <div className="columns">
              {nextMatches.map((n) => (
                <div
                  key={n.ID}
                  className="column is-flex is-justify-content-center"
                >
                  <Match data={n} />
                </div>
              ))}
            </div>
          </div>
        )}
        <RoundMatch title="Penyisihan Grup" datas={group} />
        {Object.values(bracket).map(({ title, datas }) => (
          <BracketMatch key={title} title={title} datas={datas} />
        ))}
      </div>
      <Footer />
    </div>
  );
};

const data = getData();
renderToRoot(<App {...data} />);
