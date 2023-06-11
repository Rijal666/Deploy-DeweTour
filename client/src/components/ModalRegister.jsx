/** @format */

import { Button, Modal, Form } from "react-bootstrap";
import React, { useState } from "react";
import { useMutation } from "react-query";
import { API } from "../config/api";
import Swal from "sweetalert2";

export const Register = (props) => {
  const title = "Home";
  document.title = "DeweTour | " + title;

  const [formRegister, setFormRegister] = useState({
    fullname: "",
    email: "",
    password: "",
    phone: "",
    address: "",
  });

  const { fullname, email, password, phone, address } = formRegister;

  const OnChangeHandler = (e) => {
    setFormRegister({
      ...formRegister,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = useMutation(async (e) => {
    try {
      e.preventDefault();

      const response = await API.post("/register", formRegister);

      console.log("register success : ", response);

      Swal.fire({
        position: "center",
        icon: "success",
        title: "Register Success",
        showConfirmButton: false,
        timer: 1500,
      });
      setFormRegister({
        fullname: "",
        email: "",
        password: "",
        phone: "",
        address: "",
      });
    } catch (error) {
      Swal.fire({
        position: "center",
        icon: "error",
        title: "Register Failed",
        showConfirmButton: false,
        timer: 1500,
      });
      console.log("register failed : ", error);
    }
    props.onHide();
  });

  // const handleRegister = () => {
  //   let users = JSON.parse(localStorage.getItem("users")) || [];
  //   users.push(user);
  //   localStorage.setItem("users", JSON.stringify(users));
  //   props.onHide(true);
  // };

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
            Register
          </p>
          <Form className="mt-4" onSubmit={(e) => handleSubmit.mutate(e)}>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Label className="fw-bold">Full Name</Form.Label>
              <Form.Control
                className="p-2 mb-3"
                onChange={OnChangeHandler}
                name="fullname"
                value={fullname}
                type="text"
              />
              <Form.Label className="fw-bold">Email</Form.Label>
              <Form.Control
                className="p-2 mb-3"
                onChange={OnChangeHandler}
                name="email"
                value={email}
                type="email"
              />
              <Form.Label className="fw-bold">Password</Form.Label>

              <Form.Control
                type="password"
                onChange={OnChangeHandler}
                name="password"
                value={password}
              />
              <Form.Label className="fw-bold">Phone</Form.Label>
              <Form.Control
                className="p-2 mb-3"
                onChange={OnChangeHandler}
                name="phone"
                value={phone}
                type="number"
              />
              <Form.Label className="fw-bold">Address</Form.Label>
              <Form.Control
                className=" mb-3"
                as="textarea"
                onChange={OnChangeHandler}
                name="address"
                value={address}
                style={{
                  height: "100px",
                  resize: "none",
                }}
              />
            </Form.Group>
            <Button
              type="submit"
              className="fw-bold border-0 w-100 py-2 mt-3"
              style={{ backgroundColor: "#FFAF00" }}
            >
              Register
            </Button>
          </Form>

          <p className="text-center mt-3">
            Already have an account ? Klik{" "}
            <span
              onClick={props.onClick}
              className="fw-bold"
              style={{ cursor: "pointer" }}
            >
              Here
            </span>
          </p>
        </div>
      </Modal>
    </>
  );
};

export default Register;
