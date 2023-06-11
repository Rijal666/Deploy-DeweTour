/** @format */

import Navbar from "../components/Navbar";
import { useState } from "react";
import Jumbotron from "../components/Jumbotron";
import Footer from "../components/Footer";
import Cards from "../components/Card";
import { useQuery } from "react-query";
import { API } from "../config/api";
function Home() {
  document.title = "DeweTour | Home";
  const [search, setSearch] = useState("");
  const [data, setData] = useState();

  let { data: trips } = useQuery("tripsCache", async () => {
    const response = await API.get("/trips");
    setData(response.data.data);
  });

  const handleSearch = (e) => {
    setSearch(e.target.value);
  };

  console.log(trips);
  return (
    <>
      <Navbar />
      <Jumbotron search={search} handleSearch={handleSearch} />
      <Cards data={data} search={search} />
      <Footer />
    </>
  );
}
export default Home;
