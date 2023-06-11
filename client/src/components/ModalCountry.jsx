/** @format */
import { Modal, Button, Form, NavLink } from "react-bootstrap";
import { useState } from "react";
import { useMutation } from "react-query";
import { API } from "../config/api";
import Swal from "sweetalert2";

export const Country = (props) => {
  const [formCountry, setFormCountry] = useState({
    name: "",
  });

  const { name } = formCountry;

  const OnChangeHandler = (e) => {
    setFormCountry({
      ...formCountry,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = useMutation(async (e) => {
    try {
      e.preventDefault();

      const response = await API.post("/country", formCountry);
      console.log("add country success : ", response);
      Swal.fire({
        position: "center",
        icon: "success",
        title: "Add country Success",
        showConfirmButton: false,
        timer: 1500,
      });
      setFormCountry({
        name: "",
      });
    } catch (error) {
      Swal.fire({
        position: "center",
        icon: "error",
        title: "Add country Failed",
        showConfirmButton: false,
        timer: 1500,
      });
      console.log("Add country failed : ", error);
    }
    props.onHide();
  });

  return (
    <>
      <Modal
        show={props.show}
        onHide={props.onHide}
        aria-labelledby="contained-modal-title-vcenter"
        centered
      >
        <img
          src="/images/palm.png"
          alt="#"
          width="30%"
          style={{ position: "absolute" }}
        />
        <img
          src="/images/hibiscus.png"
          alt="#"
          style={{ position: "absolute", left: "433px", borderRadius: "6px" }}
        />
        <div className="px-5 pb-3">
          <p className="fs-3 fw-bold text-center " style={{ paddingTop: 50 }}>
            Country
          </p>
          <Form className="mt-4" onSubmit={(e) => handleSubmit.mutate(e)}>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Control
                className="p-2 mb-3"
                onChange={OnChangeHandler}
                name="name"
                value={name}
                type="text"
              />
            </Form.Group>
            <NavLink
              href="/IncomeTrip"
              className="d-flex justify-content-center"
            >
              <Button
                type="submit"
                style={{
                  backgroundColor: "#FFAF00",
                  fontWeight: "bold",
                  border: "none",
                  padding: "10px 40px",
                  marginBottom: "30px",
                }}
              >
                ADD COUNTRY
              </Button>
            </NavLink>
          </Form>
        </div>
      </Modal>
    </>
  );
};
export default Country;
