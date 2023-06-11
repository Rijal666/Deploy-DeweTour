/** @format */
import Navbars from "../components/Navbar";
import Footer from "../components/Footer";
import { Button, Form, NavLink } from "react-bootstrap";
import { useParams, useNavigate } from "react-router-dom";
import Carousel from "react-multi-carousel";
import { API } from "../config/api";
import Swal from "sweetalert2";

import "react-multi-carousel/lib/styles.css";
import { useEffect, useState } from "react";
import { useQuery, useMutation } from "react-query";

const Detail = () => {
  document.title = "Detail | DeweTour";
  const rupiah = (number) => {
    return new Intl.NumberFormat("id-ID", {
      style: "currency",
      currency: "IDR",
    }).format(number);
  };
  const navigate = useNavigate();
  let { id } = useParams();
  // console.log(id);

  let { data: Trip } = useQuery("tripCache", async () => {
    const response = await API.get(`/trip/${id}`);
    return response.data.data;
  });

  const [count, setCount] = useState(1);
  const [dataTrans, setDataTrans] = useState();
  const [transaction, setTransaction] = useState({
    id: 0,
    name: "",
    gender: "",
    phone: "",
  });
  const { name, gender, phone } = transaction;

  const OnChangeHandler = (e) => {
    setTransaction({
      ...transaction,
      [e.target.name]: e.target.value,
    });
  };

  const Add = () => {
    if (count === 10) return;
    setCount(count + 1);
  };

  const Less = () => {
    if (count === 1) return;
    setCount(count - 1);
  };

  // const total = Trip ? Trip.price * count : 0;

  useEffect(() => {
    setDataTrans({
      qty: count,
      pay: Trip?.price * count,
    });
  }, [count]);

  // useEffect(() => {
  //   //change this to the script source you want to load, for example this is snap.js sandbox env
  //   const midtransScriptUrl = "https://app.sandbox.midtrans.com/snap/snap.js";
  //   //change this according to your client-key
  //   const myMidtransClientKey = process.env.REACT_APP_MIDTRANS_CLIENT_KEY;

  //   let scriptTag = document.createElement("script");
  //   scriptTag.src = midtransScriptUrl;
  //   // optional if you want to set script attribute
  //   // for example snap.js have data-client-key attribute
  //   scriptTag.setAttribute("data-client-key", myMidtransClientKey);

  //   document.body.appendChild(scriptTag);
  //   return () => {
  //     document.body.removeChild(scriptTag);
  //   };
  // }, []);

  const handleBuy = useMutation(async (e) => {
    try {
      e.preventDefault();

      const config = {
        headers: {
          "Content-type": "multipart/form-data",
        },
      };
      // const formData = new FormData();
      // formData.set("name", transaction.name);
      // formData.set("gender", transaction.gender);
      // formData.set("phone", transaction.phone);
      let data = {
        name: transaction?.name,
        gender: transaction?.gender,
        phone: transaction?.phone,
        counter_qty: dataTrans?.qty,
        total: dataTrans?.pay,
        status: "success",
        trip_id: Trip?.id,
      };

      console.log(data, "inininin kokojodad");

      const response = await API.post("/transaction", data, config);
      navigate("/Profile");
      Swal.fire({
        position: "center",
        icon: "success",
        title: "Add Transaction Success",
        showConfirmButton: false,
        timer: 1500,
      });
      return response.data.data;

      // const token = response.data.data.token;
      // window.snap.pay(token, {
      //   onSuccess: function (result) {
      //     /* You may add your own implementation here */
      //     console.log(result);
      //     navigate("/Profile");
      //   },
      //   onPending: function (result) {
      //     /* You may add your own implementation here */
      //     console.log(result);
      //     navigate("/Profile");
      //   },
      //   onError: function (result) {
      //     /* You may add your own implementation here */
      //     console.log(result);
      //     navigate("/Profile");
      //   },
      //   onClose: function () {
      //     /* You may add your own implementation here */
      //     alert("you closed the popup without finishing the payment");
      //   },
      // });
    } catch (error) {
      console.log("transaction failed : ", error);
    }
  });

  const responsive = {
    superLargeDesktop: {
      // the naming can be any, depends on you.
      breakpoint: { max: 4000, min: 3000 },
      items: 5,
    },
    desktop: {
      breakpoint: { max: 3000, min: 1024 },
      items: 3,
    },
    tablet: {
      breakpoint: { max: 1024, min: 464 },
      items: 2,
    },
    mobile: {
      breakpoint: { max: 464, min: 0 },
      items: 1,
    },
  };
  return (
    <>
      <Navbars />
      <div style={{ margin: "50px 130px" }}>
        <h1 className="fw-bold m-0">{Trip?.title}</h1>
        <h4 style={{ color: "#A8A8A8", fontWeight: "bold" }}>
          {Trip?.country.name}
        </h4>
        <img src={Trip?.image} className="mb-2" alt="" width="100%" />
        <Carousel responsive={responsive}>
          <div>
            <img src="/images/bekasi.png" alt="" />
          </div>
          <div>
            <img src="/images/bekasi2.png" alt="" />
          </div>
          <div>
            <img src="/images/bekasi3.png" alt="" />
          </div>
          <div>
            <img src="/images/bekasi.png" alt="" />
          </div>
        </Carousel>
        <h4 className="fw-bold my-3">Information Trip</h4>
        <div className="d-flex justify-content-between">
          <div>
            <p style={{ color: "#A8A8A8" }}>Accommodation</p>
            <div className="d-flex gap-3">
              <img src="/images/hotel.svg" alt="" width="25px" height="25px" />
              <h5 className="fw-bold">{Trip?.accomodation}</h5>
            </div>
          </div>
          <div>
            <p style={{ color: "#A8A8A8" }}>Transportation</p>
            <div className="d-flex gap-3">
              <img src="/images/plane.svg" alt="" width="25px" height="25px" />
              <h5 className="fw-bold">{Trip?.transportation}</h5>
            </div>
          </div>
          <div>
            <p style={{ color: "#A8A8A8" }}>Eat</p>
            <div className="d-flex gap-3">
              <img src="/images/meal.svg" alt="" width="25px" height="25px" />
              <h5 className="fw-bold">{Trip?.eat}</h5>
            </div>
          </div>
          <div>
            <p style={{ color: "#A8A8A8" }}>Duration</p>
            <div className="d-flex gap-3">
              <img src="/images/time.svg" alt="" width="25px" height="25px" />
              <h5 className="fw-bold">
                {Trip?.day} Day {Trip?.night} Night
              </h5>
            </div>
          </div>
          <div>
            <p style={{ color: "#A8A8A8" }}>Date Trip</p>
            <div className="d-flex gap-3">
              <img
                src="/images/calendar.svg"
                alt=""
                width="25px"
                height="25px"
              />
              <h5 className="fw-bold">{Trip?.date_trip}</h5>
            </div>
          </div>
        </div>
        <h4 className="fw-bold mt-5">Description</h4>
        <p style={{ color: "#A8A8A8" }}>{Trip?.description}</p>
        <div style={{ display: "flex", justifyContent: "space-between" }}>
          <div className="d-flex gap-2">
            <h2 style={{ color: "#FFAF00", fontWeight: "bold" }}>
              {rupiah(Trip?.price)}
            </h2>
            <h2 className="fw-bold"> / Person</h2>
          </div>
          <div>
            <img
              src="/images/Minus.png"
              alt=""
              width="40px"
              onClick={Less}
              style={{ cursor: "pointer" }}
            />
            <span className="mx-3">{count}</span>
            <img
              src="/images/Plus.png"
              alt=""
              width="40px"
              onClick={Add}
              style={{ cursor: "pointer" }}
            />
          </div>
        </div>
        <div className="d-flex justify-content-between my-5">
          <h2 className="fw-bold">Total :</h2>
          <h2 style={{ color: "#FFAF00", fontWeight: "bold" }}>
            {rupiah(count * Trip?.price)}
          </h2>
        </div>

        <div>
          <h1 className="text-center mb-5">FORMULIR PEMBELIAN TIKET</h1>

          <div
            style={{
              margin: "0 200px",
              boxShadow: "2px 2px 20px grey",
              borderRadius: "10px",
            }}
          >
            <img
              src="/images/palm.png"
              alt="#"
              width="15%"
              style={{ position: "absolute" }}
            />
            <img
              src="/images/hibiscus.png"
              alt="#"
              width="10%"
              style={{
                position: "absolute",
                left: "822px",
                borderRadius: "6px",
              }}
            />
            <div style={{ padding: " 100px" }}>
              <Form className="mt-4" onSubmit={(e) => handleBuy.mutate(e)}>
                <Form.Group
                  className="mb-3"
                  controlId="exampleForm.ControlInput1"
                >
                  <Form.Label className="fw-bold">Full Name</Form.Label>
                  <Form.Control
                    className="p-2 mb-3"
                    onChange={OnChangeHandler}
                    name="name"
                    value={name}
                    type="text"
                  />
                  <Form.Label className="fw-bold">Gender</Form.Label>
                  <Form.Control
                    className="p-2 mb-3"
                    onChange={OnChangeHandler}
                    name="gender"
                    value={gender}
                    type="text"
                  />
                  <Form.Label className="fw-bold">Phone</Form.Label>

                  <Form.Control
                    type="text"
                    onChange={OnChangeHandler}
                    name="phone"
                    value={phone}
                  />
                </Form.Group>
                <div className="d-flex justify-content-center">
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
                    BOOK NOW
                  </Button>
                </div>
              </Form>
            </div>
          </div>
        </div>
      </div>

      <Footer />
    </>
  );
};
export default Detail;
