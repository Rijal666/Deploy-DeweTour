/** @format */

import Card from "react-bootstrap/Card";
import { NavLink } from "react-bootstrap";
import { useQuery } from "react-query";
import { API } from "../config/api";

const Cards = ({ data, search }) => {
  console.log(data, "ini data kontol doni ");
  // console.log(search, "ini search ");

  const rupiah = (number) => {
    return new Intl.NumberFormat("id-ID", {
      style: "currency",
      currency: "IDR",
    }).format(number);
  };
  return (
    <>
      {data?.length !== 0 ? (
        <div
          data-aos="fade-up"
          data-aos-duration="3000"
          style={{
            marginBottom: "50px",
            display: "flex",
            justifyContent: "center",
            gap: "30px",
            flexWrap: "wrap",
          }}>
          {data
            ?.filter((itemSearch) => {
              if (search === "") {
                return itemSearch;
              } else if (
                itemSearch.country.name
                  .toLowerCase()
                  .includes(search.toLocaleLowerCase())
              ) {
                return itemSearch;
              } else if (
                itemSearch.title
                  .toLowerCase()
                  .includes(search.toLocaleLowerCase())
              ) {
                return itemSearch;
              }
            })
            .map((trip, i) => {
              return (
                <Card
                  key={i}
                  className="lg:w-[480px] w-[300px]"
                  style={{
                    padding: "10px",
                    boxShadow: "2px 2px 20px grey",
                  }}>
                  <NavLink href={`/Detail/${trip?.id}`}>
                    <span
                      className="text-decoration-none absolute ml-[230px] mt-[30px] lg:ml-[410px]"
                      style={{
                        background: "white",
                        borderRadius: "5px 0 0 5px",
                        textAlign: "center",
                        width: "50px",
                        height: "30px",
                      }}>
                      {trip.quota}
                    </span>
                    <Card.Img
                      variant="top"
                      src={trip?.image}
                      width="100px"
                      height="200px"
                    />
                    <Card.Body>
                      <Card.Title>{trip?.title}</Card.Title>
                      <div className="d-flex" style={{ marginBottom: "-20px" }}>
                        <Card.Text style={{ color: "#FFAF00" }}>
                          {rupiah(trip?.price)}
                        </Card.Text>
                        <Card.Text
                          className="ms-auto"
                          style={{ color: "#878787" }}>
                          {trip?.country.name}
                        </Card.Text>
                      </div>
                    </Card.Body>
                  </NavLink>
                </Card>
              );
            })}
        </div>
      ) : (
        <h1>Trip not found</h1>
      )}
    </>
  );
};

export default Cards;
