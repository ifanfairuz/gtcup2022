import React, { useMemo } from "react";
import { Ball } from "@components/icons/Ball";
import { Crown } from "@components/icons/Crown";
import { Set } from "./Set";

export const Match = ({ data }) => {
  const item = useMemo(() => {
    if (data.Type == "B") {
      const temp = data.Label.split("|");
      if (data.TeamAwayId && data.TeamHomeId) {
        return {
          label: temp[0],
          sublabel: "",
          home: data.TeamHome.Name,
          away: data.TeamAway.Name,
        };
      } else {
        return {
          label: temp[0],
          sublabel: "",
          home: temp[1],
          away: temp[2],
        };
      }
    }
    return {
      label: data.Label,
      sublabel: `Grup ${data.Group}`,
      home: data.TeamHome.Name,
      away: data.TeamAway.Name,
    };
  }, [data]);

  return (
    <div className="p-4">
      <div className="columns">
        <div className="column is-flex is-flex-direction-column">
          <p className="has-text-grey is-size-6">{item.label}</p>
          <p className="has-text-grey is-size-6">{item.sublabel}</p>
          <p className="has-text-grey is-size-6 mb-4">{/* Agregat: 2-1 */}</p>
          <div className="flex-1 is-flex is-flex-direction-column is-justify-content-space-around">
            <div className="flex-1 has-text-centered is-flex is-flex-direction-row is-align-items-center mb-4">
              <Ball className="is-size-4 has-text-primary mr-2" />
              <h3 className="has-text-weight-medium">{item.home}</h3>
            </div>
            <div className="flex-1 has-text-centered is-flex is-flex-direction-row is-align-items-center mb-4">
              <Ball className="is-size-4 has-text-link mr-2" />
              <h3 className="has-text-weight-medium">{item.away}</h3>
              {/* <Crown className="has-text-warning" /> */}
            </div>
          </div>
        </div>
        {/* <div className="column is-one-third is-flex is-flex-direction-column">
          <Set />
          <Set />
          <Set />
        </div> */}
      </div>
    </div>
  );
};
