/** @format */

import { Container, Form, Button } from "react-bootstrap";
import { useState, useEffect } from "react";
import { useMutation, useQu } from "react-query";
import { API } from "../config/api";
import { useNavigate } from "react-router-dom";
import Navbars from "../components/Navbar";
import ModalCountry from "../components/ModalCountry";
import Swal from "sweetalert2";
import Footer from "../components/Footer";
function AddTrip() {
  document.title = "AddTrip | DeweTour";

  const navigate = useNavigate();
  const [imageUrl, setImageUrl] = useState("/images/placeholder.webp");
  const [showCountry, setShowCountry] = useState(false);
  const [country, setCountry] = useState([]);

  const handleClose = () => {
    setShowCountry(false);
  };

  const handleAddCountry = () => {
    handleClose(true);
    setShowCountry(true);
  };

  const getCountry = async () => {
    try {
      const response = await API.get("/countries");
      setCountry(response.data.data);
    } catch (error) {
      console.log(error);
    }
  };

  const [formAddTrip, setFormAddTrip] = useState({
    id: 0,
    title: "",
    country_id: "",
    accomodation: "",
    transport: "",
    eat: "",
    day: "",
    night: "",
    date: "",
    price: "",
    quota: "",
    desc: "",
    image: "",
  });

  const handleChange = (e) => {
    setFormAddTrip({
      ...formAddTrip,
      [e.target.name]:
        e.target.type === "file" ? e.target.files : e.target.value,
    });
    // create image url for preview
    if (e.target.type === "file") {
      let url = URL.createObjectURL(e.target.files[0]);
      setImageUrl(url);
    }
  };

  const submitAddTrip = useMutation(async (e) => {
    try {
      e.preventDefault();

      const config = {
        Headers: {
          "Content-type": "multipart/form-data",
        },
      };

      //store data width FormData as object
      const formData = new FormData();
      formData.set("image", formAddTrip.image[0], formAddTrip.image[0].name);
      formData.set("title", formAddTrip.title);
      formData.set("country_id", Number(formAddTrip.country_id));
      formData.set("accomodation", formAddTrip.accomodation);
      formData.set("transportation", formAddTrip.transport);
      formData.set("eat", formAddTrip.eat);
      formData.set("day", Number(formAddTrip.day));
      formData.set("night", Number(formAddTrip.night));
      formData.set("date_trip", formAddTrip.date);
      formData.set("price", Number(formAddTrip.price));
      formData.set("quota", Number(formAddTrip.quota));
      formData.set("description", formAddTrip.desc);

      //insert trip data
      const response = await API.post("/trip", formData, config);
      console.log("add trip success : ", response);

      navigate("/IncomeTrip");
      Swal.fire({
        position: "center",
        icon: "success",
        title: "Add Trip Success",
        showConfirmButton: false,
        timer: 1500,
      });
    } catch (error) {
      Swal.fire({
        position: "center",
        icon: "error",
        title: "Add Trip Failed",
        showConfirmButton: false,
        timer: 1500,
      });
      console.log("add trip failed : ", error);
    }
  });

  useEffect(() => {
    getCountry();
  }, []);

  return (
    <>
      <ModalCountry show={showCountry} onHide={handleClose} />
      <Navbars />
      <Container>
        <div className="d-flex justify-content-between">
          <h1 className="fw-bold my-5">Add Trip</h1>
          <div className="d-flex align-items-center">
            <Button
              style={{ backgroundColor: "#FFAF00", border: "none" }}
              onClick={handleAddCountry}
            >
              ADD COUNTRY
            </Button>
          </div>
        </div>
        <div className="mx-5">
          <Form onSubmit={(e) => submitAddTrip.mutate(e)}>
            <Form.Group className="mb-5" controlId="formBasicEmail">
              <Form.Label className="fw-bold">Title Trip</Form.Label>
              <Form.Control
                type="text"
                name="title"
                value={formAddTrip.title}
                onChange={handleChange}
                style={{
                  border: "2px solid #B1B1B1",
                  backgroundColor: "#DBDBDB",
                }}
              />
            </Form.Group>

            <Form.Group className="mb-5" controlId="formBasicPassword">
              <Form.Label className="fw-bold">Country</Form.Label>
              <Form.Select
                style={{
                  border: "2px solid #B1B1B1",
                  backgroundColor: "#DBDBDB",
                }}
                onChange={handleChange}
                name="country_id"
              >
                <option>Select Country</option>
                {country.map((item, index) => {
                  return (
                    <option value={item.id} key={index}>
                      {item.name}
                    </option>
                  );
                })}
              </Form.Select>
            </Form.Group>
            <Form.Group className="mb-5" controlId="formBasicPassword">
              <Form.Label className="fw-bold">Accomodation</Form.Label>
              <Form.Control
                type="text"
                value={formAddTrip.accomodation}
                onChange={handleChange}
                name="accomodation"
                style={{
                  border: "2px solid #B1B1B1",
                  backgroundColor: "#DBDBDB",
                }}
              />
            </Form.Group>
            <Form.Group className="mb-5" controlId="formBasicPassword">
              <Form.Label className="fw-bold">transportation</Form.Label>
              <Form.Control
                type="text"
                value={formAddTrip.transport}
                name="transport"
                onChange={handleChange}
                style={{
                  border: "2px solid #B1B1B1",
                  backgroundColor: "#DBDBDB",
                }}
              />
            </Form.Group>
            <Form.Group className="mb-5" controlId="formBasicPassword">
              <Form.Label className="fw-bold">Eat</Form.Label>
              <Form.Control
                type="text"
                name="eat"
                value={formAddTrip.eat}
                onChange={handleChange}
                style={{
                  border: "2px solid #B1B1B1",
                  backgroundColor: "#DBDBDB",
                }}
              />
            </Form.Group>
            <Form.Group className="mb-5" controlId="formBasicPassword">
              <Form.Label className="fw-bold">Duration</Form.Label>
              <div className="d-flex gap-5">
                <Form.Control
                  type="text"
                  name="day"
                  value={formAddTrip.day}
                  onChange={handleChange}
                  style={{
                    border: "2px solid #B1B1B1",
                    backgroundColor: "#DBDBDB",
                  }}
                />
                <Form.Label className="fw-bold">Day</Form.Label>
                <Form.Control
                  type="text"
                  name="night"
                  value={formAddTrip.night}
                  onChange={handleChange}
                  style={{
                    border: "2px solid #B1B1B1",
                    backgroundColor: "#DBDBDB",
                  }}
                />
                <Form.Label className="fw-bold">Night</Form.Label>
              </div>
            </Form.Group>
            <Form.Group className="mb-5" controlId="formBasicPassword">
              <Form.Label className="fw-bold">Date Trip</Form.Label>
              <Form.Control
                type="text"
                name="date"
                value={formAddTrip.date}
                onChange={handleChange}
                style={{
                  border: "2px solid #B1B1B1",
                  backgroundColor: "#DBDBDB",
                }}
              />
            </Form.Group>
            <Form.Group className="mb-5" controlId="formBasicPassword">
              <Form.Label className="fw-bold">Price</Form.Label>
              <Form.Control
                type="text"
                name="price"
                value={formAddTrip.price}
                onChange={handleChange}
                style={{
                  border: "2px solid #B1B1B1",
                  backgroundColor: "#DBDBDB",
                }}
              />
            </Form.Group>
            <Form.Group className="mb-5" controlId="formBasicPassword">
              <Form.Label className="fw-bold">Quota</Form.Label>
              <Form.Control
                type="text"
                name="quota"
                value={formAddTrip.quota}
                onChange={handleChange}
                style={{
                  border: "2px solid #B1B1B1",
                  backgroundColor: "#DBDBDB",
                }}
              />
            </Form.Group>
            <Form.Group className="mb-5" controlId="formBasicPassword">
              <Form.Label className="fw-bold">Description</Form.Label>
              <Form.Control
                as="textarea"
                name="desc"
                value={formAddTrip.desc}
                onChange={handleChange}
                style={{
                  height: "100px",
                  border: "2px solid #B1B1B1",
                  backgroundColor: "#DBDBDB",
                  resize: "none",
                }}
              />
            </Form.Group>
            <Form.Label className="fw-bold">Image</Form.Label>
            <Form.Group className="mb-5 d-flex gap-5">
              <Form.Label
                style={{
                  border: "2px solid #B1B1B1",
                  width: "40%",
                  backgroundColor: "#DBDBDB",
                }}
                className="rounded"
              >
                <div className="d-flex justify-content-between align-items-center py-2 px-3">
                  <div className="d-flex align-items-center">
                    <p style={{ fontSize: "20px", color: "#FFAF00" }}>
                      Attache Here
                    </p>
                  </div>
                  <div>
                    <Form.Control
                      multiple
                      type="file"
                      name="image"
                      id="image"
                      onChange={handleChange}
                      hidden
                    />
                    <img src="/images/clip.svg" alt="" />
                  </div>
                </div>
              </Form.Label>
            </Form.Group>
            <div style={{ width: "100px" }}>
              <img src={imageUrl} alt="preview" style={{ width: "100%" }} />
            </div>
            <div className="justify-content-center d-flex">
              <Button
                style={{
                  backgroundColor: "#FFAF00",
                  border: "none",
                  padding: "5px 50px",
                }}
                type="submit"
              >
                Add Trip
              </Button>
            </div>
          </Form>
        </div>
      </Container>
      <Footer />
    </>
  );
}
export default AddTrip;
