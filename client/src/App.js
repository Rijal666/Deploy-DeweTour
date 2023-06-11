/** @format */
import Detail from "./pages/DetailTour";
import Home from "./pages/Home";
import Profile from "./pages/Profile";
import IncomeTrip from "./pages/IncomeTrip";
import { UserContext } from "./context/userContext";
import { useContext, useEffect, useState } from "react";
import Income from "./pages/Income";
import AddTrip from "./pages/AddTrip";
import { setAuthToken, API } from "./config/api";
import { Routes, Route, Navigate } from "react-router-dom";

function App() {
  const [state, dispatch] = useContext(UserContext);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    if (!isLoading) {
      if (state.isLogin === false) {
        Navigate("/");
      }
    }
  }, []);

  useEffect(() => {
    if (localStorage.token) {
      setAuthToken(localStorage.token);
      checkUser();
    } else {
      setTimeout(() => {
        setIsLoading(false);
      }, 3000);
    }
  }, []);

  const checkUser = async () => {
    try {
      const response = await API.get("/check-auth");
      let payload = response.data.data;
      payload.token = localStorage.token;
      dispatch({
        type: "USER_SUCCESS",
        payload,
      });
      setTimeout(() => {
        setIsLoading(false);
      }, 3000);
    } catch (error) {
      dispatch({
        type: "AUTH_ERROR",
      });
      setTimeout(() => {
        setIsLoading(false);
      }, 3000);
      return;
    }
  };

  return (
    <>
      {isLoading ? (
        <div
          className="d-flex justify-content-center align-items-center"
          style={{ width: "100vw", height: "100vh" }}
        >
          <img
            src="/images/icon.png"
            alt="Dewetour"
            className="animate__heartBeat"
          />
        </div>
      ) : (
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/Detail/:id" element={<Detail />} />
          <Route path="/Profile" element={<Profile />} />
          <Route path="/AddTrip" element={<AddTrip />} />
          <Route path="/IncomeTrip" element={<IncomeTrip />} />
          <Route path="/Income" element={<Income />} />
        </Routes>
      )}
    </>
  );
}
export default App;
