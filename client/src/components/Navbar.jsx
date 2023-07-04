/** @format */
import React, { useState, useEffect, useContext } from "react";
import { useNavigate } from "react-router-dom";
import { Button, Nav, Navbar, NavDropdown } from "react-bootstrap";
import { UserContext } from "../context/userContext";
import { API, setAuthToken } from "../config/api";
import ModalLogin from "./ModalLogin";
import ModalRegister from "./ModalRegister";
import Swal from "sweetalert2";
import ImgProfile from "../assets/image/myProfile.png";

function Navbars() {
  let navigate = useNavigate();
  const [state, dispatch] = useContext(UserContext);
  const [isLoading, setIsLoading] = useState(true);
  const [showLogin, setShowLogin] = useState(false);
  const [showRegister, setShowRegister] = useState(false);

  const handleClose = () => {
    setShowLogin(false);
    setShowRegister(false);
  };

  const handleShowLogin = () => {
    handleClose(true);
    setShowLogin(true);
  };

  const handleShowRegister = () => {
    handleClose(true);
    setShowRegister(true);
  };

  useEffect(() => {
    // Redirect auth but just when isLoading is false
    if (!isLoading) {
      if (state.isLogin === false) {
        navigate("/");
      }
    }
  }, [isLoading]);

  useEffect(() => {
    if (localStorage.token) {
      setAuthToken(localStorage.token);
      checkUser();
    } else {
      setIsLoading(false);
    }
  }, []);

  const checkUser = async () => {
    try {
      const response = await API.get("/check-auth");
      console.log("check user success : ", response);
      // get user data
      let payload = response.data.data;
      // get token from localstorage
      payload.token = localStorage.token;
      // send data to useContext
      dispatch({
        type: "USER_SUCCESS",
        payload,
      });
      setIsLoading(false);
    } catch (error) {
      console.log("check user failed : ", error);
      dispatch({
        type: "AUTH_ERROR",
      });
      setIsLoading(false);
    }
  };

  const logout = () => {
    console.log(state);
    dispatch({
      type: "LOGOUT",
    });
    Swal.fire({
      position: "center",
      icon: "success",
      title: "Logout Success",
      showConfirmButton: false,
      timer: 1500,
    });
    navigate("/");
  };

  return (
    <>
      <Navbar
        data-aos="fade-up"
        data-aos-duration="1000"
        collapseOnSelect
        expand="md"
        className="lg:px-40 px-10 flex flex-row"
        style={{
          backgroundImage: `url("/images/bg.png")`,
          backgroundSize: "cover",
          width: "100%",
          height: "100%",
        }}>
        <img
          src="/images/icon.png"
          className="lg:w-[19%] w-[28%]"
          alt="#"></img>
        <Nav className="lg:ms-auto ml-5 gap-x-2 lg:gap-x-4 flex-row">
          {state.isLogin === true ? (
            state.user.is_admin === true ? (
              <>
                <NavDropdown
                  title={
                    <img
                      src="/images/blank-profile.png"
                      alt=""
                      style={{
                        width: "50px",
                        height: "50px",
                        border: "solid orange",
                      }}
                      data-aos="flip-left"
                      data-aos-easing="ease-out-cubic"
                      data-aos-duration="2000"
                      className="rounded-circle"
                    />
                  }>
                  <NavDropdown.Item href="/IncomeTrip">
                    <img
                      src="/images/journey.svg"
                      alt=""
                      style={{ width: "20px" }}
                    />
                    <span className="ms-3 fw-bold">Trip</span>
                  </NavDropdown.Item>
                  <NavDropdown.Divider />
                  <NavDropdown.Item onClick={logout} href="/">
                    {" "}
                    <img
                      src="/images/logout.svg"
                      alt=""
                      style={{ width: "20px" }}
                    />
                    <span className="ms-3 fw-bold">Logout</span>
                  </NavDropdown.Item>
                </NavDropdown>
              </>
            ) : (
              <>
                <NavDropdown
                  title={
                    <img
                      src={state.user.image ? state.user.image : ImgProfile}
                      alt=""
                      style={{
                        width: "50px",
                        height: "50px",
                        border: "solid orange",
                      }}
                      className="rounded-circle"
                    />
                  }>
                  <NavDropdown.Item href="/Profile">
                    <img
                      src="/images/user.svg"
                      alt=""
                      style={{ width: "20px" }}
                    />
                    <span className="ms-3 fw-bold">Profile</span>
                  </NavDropdown.Item>
                  <NavDropdown.Item href="/Payment">
                    <img
                      src="/images/bill.svg"
                      alt=""
                      style={{ width: "20px" }}
                    />
                    <span className="ms-3 fw-bold">pay</span>
                  </NavDropdown.Item>
                  <NavDropdown.Divider />
                  <NavDropdown.Item onClick={logout} href="/">
                    {" "}
                    <img
                      src="/images/logout.svg"
                      alt=""
                      style={{ width: "20px" }}
                    />
                    <span className="ms-3 fw-bold">Logout</span>
                  </NavDropdown.Item>
                </NavDropdown>
              </>
            )
          ) : (
            <>
              <Button
                variant="outline-light"
                className="lg:px-[50px] lg:py-[10px] px-[20px] py-[1px]"
                onClick={handleShowLogin}>
                Login
              </Button>
              <Button
                className="lg:px-[50px] lg:py-[10px] px-[10px] py-[1px]"
                style={{
                  color: "white",
                  backgroundColor: "#FFAF00",
                  border: "none",
                }}
                onClick={handleShowRegister}>
                Register
              </Button>
            </>
          )}
        </Nav>
      </Navbar>
      <ModalLogin
        show={showLogin}
        onHide={handleClose}
        onClick={handleShowRegister}
      />
      <ModalRegister
        show={showRegister}
        onHide={handleClose}
        onClick={handleShowLogin}
      />
    </>
  );
}

export default Navbars;
