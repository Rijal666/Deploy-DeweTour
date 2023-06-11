/** @format */

import Card from "react-bootstrap/Card";
import { NavLink } from "react-bootstrap";
import { useQuery } from "react-query";
import { API } from "../config/api";

function Cards() {
  let { data: trips } = useQuery("tripsCache", async () => {
    const response = await API.get("/trips");
    return response.data.data;
  });

  console.log(trips);

  const rupiah = (number) => {
    return new Intl.NumberFormat("id-ID", {
      style: "currency",
      currency: "IDR",
    }).format(number);
  };
  return (
    <>
      <div
        data-aos="fade-up"
        data-aos-duration="1000"
        style={{
          marginBottom: "50px",
          display: "flex",
          justifyContent: "center",
          gap: "30px",
          flexWrap: "wrap",
        }}
      >
        {trips?.map((trip, i) => {
          return (
            <Card
              key={i}
              style={{
                width: "300px",
                padding: "10px",
                boxShadow: "2px 2px 20px grey",
              }}
            >
              <NavLink href={`/Detail/${trip?.id}`}>
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
                    <Card.Text className="ms-auto" style={{ color: "#878787" }}>
                      {trip?.country.name}
                    </Card.Text>
                  </div>
                </Card.Body>
              </NavLink>
            </Card>
          );
        })}
      </div>
    </>
  );
}

export default Cards;
