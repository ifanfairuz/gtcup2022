import React, { useEffect, useMemo, useState } from "react";
import { renderToRoot, getData } from "@support/render";
import {
  Navbar,
  Footer,
  RoundMatch,
  BracketMatch,
  Match,
  ShareButtons,
} from "@components";
import { DateTime } from "luxon";
import { formatDate } from "@support/date";

const App = ({ matches, lastMatches, nextMatches }) => {
  const [group, setGroup] = useState({});
  const [bracket, setBracket] = useState({});

  const nextMatchesData = useMemo(
    () => {
      const has = nextMatches && nextMatches.length > 0;
      if (!has) return null;
      const date = DateTime.fromISO(nextMatches[0].Date);
      const label = date.hasSame(DateTime.now(), "day")
        ? "Pertandingan Hari Ini"
        : "Akan Datang";
      return { has, date, date_format: formatDate(date), label };
    },
    [nextMatches]
  );
  const lastMatchesData = useMemo(
    () => {
      const has = lastMatches && lastMatches.length > 0;
      if (!has) return null;
      const date = DateTime.fromISO(lastMatches[0].Date);
      const label = `Pertandingan ${lastMatches[0].Done ? "Terakhir" : "Terdekat"}`;
      return { has, date, date_format: formatDate(date), label };
    },
    [lastMatches]
  );

  useEffect(() => {
    let g = {};
    let b = {};
    for (const match of matches) {
      const date = DateTime.fromISO(match.Date);
      const title = formatDate(date);
      if (match.Type == "G") {
        g[title] = g[title] ?? {
          title,
          date,
          datas: [],
        };

        g[title].datas.push(match);
      } else {
        b[match.Round] = b[match.Round] ?? {
          title: match.Group,
          date,
          datas: {},
        };
        b[match.Round].datas[title] = {
          title,
          date,
          match: match,
        };
      }
    }
    setGroup(g);
    setBracket(b);
  }, []);

  const onRender = () => {
    if (window.location.hash) {
      const el = document.querySelector(window.location.hash);
      if (el) {
        el.scrollIntoView({ behavior: "smooth" });
        const uri = window.location.toString();
        history.replaceState(
          {},
          document.title,
          uri.substring(0, uri.indexOf("#"))
        );
      }
    }
  };

  return (
    <div ref={onRender}>
      <Navbar active="pertandingan" />
      <div className="container has-background-white">
        {!!nextMatchesData && (
          <>
            <hr className="m-0" />
            <div className="p-4">
              <h4 className="is-size-4 has-text-weight-semibold has-text-centered">
                {nextMatchesData.label}
              </h4>
              <div className="is-flex is-flex-direction-column is-align-items-center mb-2">
                <p className="has-text-centered">{nextMatchesData.date_format}</p>
                <ShareButtons inverted date={nextMatchesData.date} />
              </div>
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
          </>
        )}
        {!!lastMatchesData && (
          <div className="p-4">
            <h4 className="is-size-4 has-text-weight-semibold has-text-centered">
              {lastMatchesData.label}
            </h4>
            <div className="is-flex is-flex-direction-column is-align-items-center mb-2">
              <p className="has-text-centered">{lastMatchesData.date_format}</p>
              <ShareButtons inverted date={lastMatchesData.date} />
            </div>
            <div className="columns">
              {lastMatches.map((n) => (
                <div key={n.ID} className="column is-flex">
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
