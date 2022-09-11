import React from "react";

export const AdminNavbar = ({ active }) => {
  return (
    <div className="tabs has-background-white mb-0">
      <ul>
        <li className={active == "team" ? "is-active" : ""}>
          <a href="/bla">Team</a>
        </li>
        <li className={active == "match" ? "is-active" : ""}>
          <a href="/bla/match">Match</a>
        </li>
        <li className={active == "bracket" ? "is-active" : ""}>
          <a href="/bla/bracket">Bracket</a>
        </li>
      </ul>
    </div>
  );
};
