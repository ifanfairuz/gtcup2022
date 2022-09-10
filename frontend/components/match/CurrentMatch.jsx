import React from "react";
import { Ball } from "@components/icons/Ball";
import { Crown } from "@components/icons/Crown";

const Set = () => {
  return (
    <div className="columns is-mobile is-vcentered is-centered is-gapless mb-0">
      <p className="column has-text-left">
        25
        {/* <Crown className="has-text-warning" /> */}
      </p>
      <p className="column has-text-centered">Set 1</p>
      <p className="column has-text-right">19</p>
    </div>
  );
};

export const CurrentMatch = (props) => {
  return (
    <div {...props}>
      <div className="columns is-centered is-mobile mb-0">
        <div className="column has-text-centered">
          <Ball className="is-size-1 has-text-primary" />
          <h3 className="is-size-4-desktop is-size-5 has-text-weight-medium">
            Joko Tingkir
          </h3>
          <p className="is-size-4-desktop is-size-5">
            2{/* <Crown className="has-text-warning" /> */}
          </p>
        </div>
        <div className="column is-one-fifth">
          <h1 className="is-size-2-desktop is-size-4 has-text-weight-semibold has-text-centered">
            VS
          </h1>
          {/* <div className="is-hidden-mobile">
            <Set />
            <Set />
            <Set />
          </div> */}
        </div>
        <div className="column has-text-centered">
          <Ball className="is-size-1 has-text-link" />
          <h3 className="is-size-4-desktop is-size-5 has-text-weight-medium">
            Lorem ipsusm
          </h3>
          <p className="is-size-4-desktop is-size-5">1</p>
        </div>
      </div>
      {/* <div className="is-hidden-desktop px-4">
        <Set />
        <Set />
        <Set />
      </div> */}
    </div>
  );
};
