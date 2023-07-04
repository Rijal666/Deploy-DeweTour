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
        <div
          style={{ padding: "0 70px" }}
          data-aos="fade-up"
          data-aos-duration="2000">
          <div>
            <h1
              className="lg:text-[90px] text-[40px] lg:my-40"
              style={{
                color: "white",
                fontWeight: "bold",
                marginBottom: "0",
              }}>
              Explore
            </h1>
            <p
              className="lg:text-[70px] text-[20px]"
              style={{ color: "white" }}>
              your amazing city together
            </p>
            <p style={{ color: "white" }} className="lg:text-[20px]">
              find great places to holiday
            </p>
          </div>
          <div>
            <InputGroup className="mb-5">
              <Form.Control
                className="lg:p-3"
                aria-label="Recipient's username"
                onChange={(value) => handleSearch(value)}
                value={search}
              />
              <Button
                style={{
                  backgroundColor: "#FFAF00",
                  padding: "0 30px",
                  border: "none",
                }}>
                Search
              </Button>
            </InputGroup>
          </div>
          <div className="hidden lg:flex lg:justify-around">
            <img src="/images/Card1.png" width="20%" alt="#" />
            <img width="20%" src="/images/card2.png" alt="#" />
            <img width="20%" src="/images/card3.png" alt="#" />
            <img width="20%" src="/images/card4.png" alt="#" />
          </div>
          <h1 className="fw-bold text-center my-5 lg:text-[60px]">
            Group Tour
          </h1>
        </div>
      </div>
    </>
  );
};

export default Jumbotron;
