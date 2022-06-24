import { Container } from "react-bootstrap";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import { useNavigate } from "react-router-dom";
import { useParams } from "react-router-dom";
import { useState } from "react";
import axios from "axios";

import gambarasik from "./MasukkanreviewItem/image/iniprofil.png";
const Masukkanreview = () => {
  const navigate = useNavigate();
  const { id } = useParams();
  console.log(id);

  const baseUrl = "https://reqres.in";
  const [nama, setnama] = useState("");
  const [email, setemail] = useState("");
  const [pendapat, setpendapat] = useState("");
  const inputdata = async () => {
    if (nama === "" || email === "" || pendapat === "") {
      alert("Data tidak boleh kosong");
    } else {
      const data = {
        name: nama,
        job: pendapat,
      };
      try {
        const res = await axios.post(`${baseUrl}/api/users`, data);
        console.log(res.data);
        setnama("");
        setemail("");
        setpendapat("");
        navigate("/Kampus");
      } catch (err) {
        console.log(err);
      }
    }
  };

  return (
    <>
      <Container>
        <div className="pt-3">
          <h3>Masukkan Pendapat kamu</h3>
          <div class="row pt-2">
            <div class="col-md-6">
              <Form>
                <Form.Group className="mb-3">
                  <Form.Label>Nama</Form.Label>
                  <Form.Control
                    type="text"
                    placeholder="Masukan nama"
                    onChange={(e) => {
                      setnama(e.target.value);
                    }}
                    required
                  />
                </Form.Group>

                <Form.Group className="mb-3">
                  <Form.Label>Alamat Email</Form.Label>
                  <Form.Control
                    type="email"
                    placeholder="Masukkan Email"
                    onChange={(e) => {
                      setemail(e.target.value);
                    }}
                    required
                  />
                </Form.Group>
                <Form.Group className="mb-3">
                  <Form.Label>Masukan Pendapat Kamu</Form.Label>
                  <textarea
                    class="form-control"
                    onChange={(e) => {
                      setpendapat(e.target.value);
                    }}
                    required
                  ></textarea>
                </Form.Group>
                <Button
                  variant="primary"
                  type="button"
                  onClick={() => inputdata()}
                >
                  Kirim
                </Button>
              </Form>
            </div>
            <div class="col-md-6">
              <center>
                <img src={gambarasik} alt="" />
              </center>
            </div>
          </div>
        </div>
      </Container>
    </>
  );
};
export default Masukkanreview;
