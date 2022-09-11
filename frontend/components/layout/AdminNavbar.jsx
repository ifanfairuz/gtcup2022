import React, { useRef } from "react";

export const AdminNavbar = ({ active }) => {
  const formLogout = useRef();
  const logout = () => {
    formLogout.current.submit();
  };

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
        <li className={active == "bracket" ? "is-active" : ""}>
          <a href="#" onClick={logout}>
            Logout
          </a>
        </li>
      </ul>
      <form action="/bla/logout" method="post" ref={formLogout}></form>
    </div>
  );
};
