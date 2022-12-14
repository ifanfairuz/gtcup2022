import React from "react";
import { renderToRoot, getData } from "@support/render";
import { Navbar, Footer, AdminTeam, AdminNavbar } from "@components";

const App = ({ teams }) => {
  return (
    <div>
      <Navbar />
      <AdminNavbar active="team" />
      <AdminTeam datas={teams} />
      <Footer />
    </div>
  );
};

const data = getData();
renderToRoot(<App {...data} />);
