import React from "react";
import { renderToRoot, getData } from "@support/render";
import { Navbar, Footer, AdminNavbar, AdminMatch } from "@components";

const App = ({ matches }) => {
  return (
    <div>
      <Navbar />
      <AdminNavbar active="match" />
      <AdminMatch datas={matches} />
      <Footer />
    </div>
  );
};

const data = getData();
renderToRoot(<App {...data} />);
