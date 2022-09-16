import React, { useState } from "react";
import {
  Bracket as BracketComp,
  Seed,
  SeedItem,
  SeedTeam,
} from "react-brackets";
import { Crown } from "@components/icons/Crown";

const CustomSeedItem = ({ match, third }) => {
  return (
    <div
      className={!!third ? "is-absolute-desktop w-full-mobile mb-4" : "w-full"}
      style={{ bottom: "30px" }}
    >
      <p
        className="has-text-centered p-1 mb-3"
        style={{
          color: "rgb(143, 143, 143)",
          fontWeight: "400",
          fontSize: "1em",
        }}
      >
        {match.label}
      </p>
      <SeedItem>
        <a href={`/#match-${match.id}`} className="has-text-white">
          <SeedTeam style={{ minWidth: "300px" }}>
            <span>
              {match.home || "NO TEAM "}
              {match.winner == "home" && (
                <Crown className="has-text-warning ml-2" />
              )}
            </span>
            {!!match.done && <span>{match.poin.home}</span>}
          </SeedTeam>
          <SeedTeam style={{ minWidth: "300px" }}>
            <span>
              {match.away || "NO TEAM "}
              {match.winner == "away" && (
                <Crown className="has-text-warning ml-2" />
              )}
            </span>
            {!!match.done && <span>{match.poin.away}</span>}
          </SeedTeam>
        </a>
      </SeedItem>
      <p className="has-text-centered p-1">{match.date}</p>
    </div>
  );
};

const CustomSeed = ({ seed, breakpoint }) => {
  return (
    <Seed
      mobileBreakpoint={breakpoint}
      className={`is-relative py-0 ${
        !seed.third
          ? ""
          : "is-flex-direction-column-reverse is-flex-direction-column-desktop"
      }`}
    >
      <CustomSeedItem match={seed.match} />
      {!!seed.third && <CustomSeedItem match={seed.third.match} third={true} />}
    </Seed>
  );
};

export const Bracket = ({ rounds }) => {
  const [active, setActive] = useState(0);

  return (
    <div
      className="py-4 flex-1 is-flex-desktop is-align-content-center-desktop"
      style={{ minHeight: "80vh" }}
    >
      <BracketComp
        bracketClassName="is-justify-content-center h-full"
        swipeableProps={{
          className: "flex-1",
          onSwitching: (active) => setActive(active),
        }}
        rounds={rounds}
        renderSeedComponent={CustomSeed}
      />
      <div className="is-hidden-desktop is-relative">
        <div
          className="is-flex is-flex-direction-row is-justify-content-center is-align-content-center is-absolute w-full"
          style={{ bottom: "20px" }}
        >
          {rounds.map((r, i) => (
            <div
              key={`indicator-swipe-${i}`}
              className={`p-1 is-rounded-all mx-1 ${
                active == i
                  ? "has-background-dark"
                  : "has-background-grey-light"
              }`}
            ></div>
          ))}
        </div>
      </div>
    </div>
  );
};
