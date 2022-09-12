import React from "react";
import { renderToRoot, getData } from "@support/render";
import { Navbar, Footer, AdminNavbar, AdminMatch } from "@components";

const App = ({ matches, teams }) => {
  return (
    <div>
      <Navbar />
      <AdminNavbar active="match" />
      <AdminMatch datas={matches} teams={teams} />
      <Footer />
    </div>
  );
};

const data = getData();
renderToRoot(<App {...data} />);
