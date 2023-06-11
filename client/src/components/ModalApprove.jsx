/** @format */

import { Button, Modal, Form } from "react-bootstrap";

export const ModalApprove = (props) => {
  return (
    <Modal show={props.show} onHide={props.onHide} size="xl" centered>
      <div
        style={{
          boxShadow: "2px 2px 20px grey",
          padding: "50px",
          borderRadius: "10px",
        }}
      >
        <div className="d-flex justify-content-between mb-4">
          <img src="/images/icon.png" alt="" />
          <div className="text-end">
            <h1 className="fw-bold ">Booking</h1>
            <h4>Saturday, 22 july 2020</h4>
          </div>
        </div>
        <div className="d-flex justify-content-between">
          <div className="d-flex gap-5">
            <div>
              <h3 className="fw-bold">6D/4N Fun Tassie Vacation</h3>
              <p
                style={{
                  color: "#959595",
                  fontWeight: "bold",
                  marginBottom: "30px",
                }}
              >
                Bekasi
              </p>
              <span
                style={{
                  color: "#3CF71E",
                  backgroundColor: "#E9FDEB",
                  fontWeight: "bold",
                  padding: "5px",
                  borderRadius: "10px",
                }}
              >
                Approve
              </span>
            </div>
            <div>
              <div style={{ display: "flex", gap: "70px" }}>
                <div>
                  <h5 className="fw-bold">Date trip</h5>
                  <p style={{ color: "#959595", fontWeight: "bold" }}>
                    26 agustus 2020
                  </p>
                </div>
                <div>
                  <h5 className="fw-bold">Duration</h5>
                  <p style={{ color: "#959595", fontWeight: "bold" }}>
                    6 Day 4 Night
                  </p>
                </div>
              </div>
              <div style={{ display: "flex", gap: "55px" }}>
                <div>
                  <h5 className="fw-bold">Accomodation</h5>
                  <p style={{ color: "#959595", fontWeight: "bold" }}>
                    26 agustus 2020
                  </p>
                </div>
                <div>
                  <h5 className="fw-bold">Transportation</h5>
                  <p style={{ color: "#959595", fontWeight: "bold" }}>
                    6 Day 4 Night
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div className="text-center me-5">
            <img src="/images/barcode.svg" alt="#" />
            <h4>TCK0101</h4>
          </div>
        </div>
      </div>
    </Modal>
  );
};
