/** @format */
import { Button, Form, InputGroup } from "react-bootstrap";

const Jumbotron = ({ search, handleSearch }) => {
  return (
    <>
      <div>
        <img
          src="/images/bg.png"
          alt="#"
          width="100%"
          style={{ position: "absolute", top: "0", zIndex: "-1" }}
        />
        <div style={{ padding: "0 70px" }}>
          <div data-aos="fade-up" data-aos-duration="1000">
            <h1
              style={{
                color: "white",
                fontWeight: "bold",
                fontSize: "70px",
                marginBottom: "0",
              }}
            >
              Explore
            </h1>
            <p style={{ color: "white", fontSize: "50px" }}>
              your amazing city together
            </p>
            <p style={{ color: "white" }}>find great places to holiday</p>
          </div>
          <div data-aos="fade-up" data-aos-duration="1000">
            <InputGroup className="mb-5">
              <Form.Control
                aria-label="Recipient's username"
                onChange={(value) => handleSearch(value)}
                value={search}
              />
              <Button
                style={{
                  backgroundColor: "#FFAF00",
                  padding: "0 30px",
                  border: "none",
                }}
              >
                Search
              </Button>
            </InputGroup>
          </div>
          <div style={{ display: "flex", justifyContent: "space-between" }}>
            <img
              src="/images/Card1.png"
              alt="#"
              data-aos="flip-left"
              data-aos-duration="1000"
            />
            <img
              src="/images/card2.png"
              alt="#"
              data-aos="flip-left"
              data-aos-duration="1000"
            />
            <img
              src="/images/card3.png"
              alt="#"
              data-aos="flip-left"
              data-aos-duration="1000"
            />
            <img
              src="/images/card4.png"
              alt="#"
              data-aos="flip-left"
              data-aos-duration="1000"
            />
          </div>
          <h1 className="fw-bold text-center my-5">Group Tour</h1>
        </div>
      </div>
    </>
  );
};

export default Jumbotron;
