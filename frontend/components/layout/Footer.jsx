import React from "react";

export const Footer = ({ active }) => {
  return (
    <footer className="footer mt-10 has-background-dark has-text-light">
      <div className="columns">
        <div className="column">
          <img
            src="/assets/img/logo.svg"
            alt="logo"
            className="image is-64x64 mb-2 mx-auto"
          />
          <h4 className="has-text-weight-semibold is-size-4 has-text-centered">
            GTCup 2022
          </h4>
          <p className="has-text-centered">
            © 2022 FPPGT™. All Rights Reserved.
          </p>
        </div>
      </div>
    </footer>
  );
};
