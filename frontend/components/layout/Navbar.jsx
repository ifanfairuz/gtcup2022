import React from "react";

export const Navbar = ({ active }) => {
  return (
    <div>
      <nav
        className="navbar is-dark"
        role="navigation"
        aria-label="main navigation"
      >
        <div className="navbar-brand">
          <span className="navbar-item">
            <span className="is-flex is-flex-direction-row is-align-items-center">
              <img
                src="/assets/img/logo.svg"
                width="50"
                height="150"
                alt="logo"
              />
              <h3 className="is-size-5 has-text-weight-medium">GTCup 2022</h3>
            </span>
          </span>
        </div>
      </nav>
      <div className="tabs has-background-white">
        <ul>
          <li className={active == "pertandingan" ? "is-active" : ""}>
            <a href="/">Pertandingan</a>
          </li>
          <li className={active == "bracket" ? "is-active" : ""}>
            <a href="/bracket">Bracket</a>
          </li>
          <li className={active == "klasemen" ? "is-active" : ""}>
            <a href="/klasemen">Klasemen</a>
          </li>
        </ul>
      </div>
    </div>
  );
};
