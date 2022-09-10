import React from "react";
import { renderToRoot, getData } from "@support/render";
import { Navbar, Footer, Klasemen } from "@components";

const App = ({ klasemen }) => {
  return (
    <div>
      <Navbar active="klasemen" />
      <div className="container has-background-white">
        <div className="p-4 content">
          <div className="columns is-gapless">
            <div className="column">
              <ul>
                <li>
                  <span className="tag is-dark">P</span> Jumlah Pertandingan
                </li>
                <li>
                  <span className="tag is-success">M</span> Jumlah Kemenangan
                </li>
                <li>
                  <span className="tag is-danger">K</span> Jumlah Kekalahan
                </li>
              </ul>
            </div>
            <div className="column">
              <ul>
                <li>
                  <span className="tag is-info">SM</span> Jumlah Set Kemenangan
                </li>
                <li>
                  <span className="tag is-warning">SK</span> Jumlah Set
                  Kekalahan
                </li>
                <li>
                  <span className="tag is-primary">AS</span> Jumlah Agregat Set
                  (<span className="tag is-info">SM</span> -
                  <span className="tag is-warning">SK</span>)
                </li>
              </ul>
            </div>
            <div className="column">
              <ul>
                <li>
                  <span className="tag is-light is-info">SC-M</span> Jumlah Skor
                  Kemenangan
                </li>
                <li>
                  <span className="tag is-light is-warning">SC-K</span> Jumlah
                  Skor Kekalahan
                </li>
                <li>
                  <span className="tag is-light is-primary">A-SC</span> Jumlah
                  Agregat Skor (
                  <span className="tag is-light is-info">SC-M</span> -
                  <span className="tag is-light is-warning">SC-K</span>)
                </li>
              </ul>
            </div>
          </div>
        </div>
        <Klasemen title="Grup A" datas={klasemen.A} />
        <Klasemen title="Grup B" datas={klasemen.B} />
        <Klasemen title="Grup C" datas={klasemen.C} />
        <Klasemen title="Grup D" datas={klasemen.D} />
      </div>
      <Footer />
    </div>
  );
};

const data = getData();
renderToRoot(<App {...data} />);
