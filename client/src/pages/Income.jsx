/** @format */

import Navbar from "../components/Navbar";
import Footer from "../components/Footer";
import { Container, Table } from "react-bootstrap";
import { ModalApprove } from "../components/ModalApprove";
import { useState } from "react";
import { useQuery } from "react-query";
import { API } from "../config/api";

function Income() {
  document.title = "Income | DeweTour";

  const { data: transactions } = useQuery("transactionChace", async () => {
    const res = await API.get("/transactions");
    return res.data.data;
  });

  console.log(transactions, "ini asuuu");

  const [showApprove, setShowApprove] = useState(false);

  const handleClose = () => {
    setShowApprove(false);
  };

  const handleShowApprove = () => {
    handleClose(true);
    setShowApprove(true);
  };

  return (
    <>
      <Navbar />
      <Container>
        <div>
          <h1 className="fw-bold my-5">Income Transaction</h1>
          <Table responsive striped>
            <thead>
              <tr>
                <th>No</th>
                <th>Users</th>
                <th>Trip</th>
                <th>Status Payment</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              {transactions?.map((item, index) => {
                return (
                  <tr>
                    <td>{item.id}</td>
                    <td>{item.user.fullname}</td>
                    <td>{item.trip.country.name}</td>
                    <td>{item.status}</td>
                    <td>
                      <span onClick={handleShowApprove}>
                        <img src="/images/alat.svg" alt="" />
                      </span>
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </Table>
        </div>
      </Container>
      <ModalApprove show={showApprove} onHide={handleClose} />
      <Footer />
    </>
  );
}
export default Income;
