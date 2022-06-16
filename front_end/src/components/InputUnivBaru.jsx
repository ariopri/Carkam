import { Container } from "react-bootstrap";

import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import gambarasik from "./InputUnivBaruItem/image/Beranda.png";

const InputUnivBaru = () => {
  return (
    <>
      <Container>
        <div className="pt-3">
          <div className="pt-3">
            <h3>Masukkan Data Kampus Kamu</h3>
            <div class="row pt-2">
              <div class="col-md-6">
                <Form className="pt-5">
                  <Form.Group className="mb-3">
                    <Form.Label>Nama Kampus</Form.Label>
                    <Form.Control type="text" />
                  </Form.Group>
                  <Form.Group className="mb-3">
                    <Form.Label>Alamat Email Kampus</Form.Label>
                    <Form.Control type="email" />
                  </Form.Group>
                  <Form.Group className="mb-3">
                    <Form.Label>Masukan Info Singkat Kampus</Form.Label>
                    <textarea class="form-control"></textarea>
                  </Form.Group>
                  <Form.Group className="mb-3">
                    <Form.Label>Masukkan Jurusan 1</Form.Label>
                    <Form.Control type="text" />
                  </Form.Group>
                  <Form.Group className="mb-3">
                    <Form.Label>Masukkan Jurusan 2</Form.Label>
                    <Form.Control type="text" />
                  </Form.Group>
                  <Form.Group className="mb-3">
                    <Form.Label>Masukkan Jurusan 3</Form.Label>
                    <Form.Control type="text" />
                  </Form.Group>
                  <Form.Group className="mb-3">
                    <Form.Label>Masukkan Jurusan 4</Form.Label>
                    <Form.Control type="text" />
                  </Form.Group>
                  <Form.Group className="mb-3">
                    <Form.Label>Masukkan Jurusan 5</Form.Label>
                    <Form.Control type="text" />
                  </Form.Group>
                  <Button variant="primary" type="submit">
                    Simpan
                  </Button>
                </Form>
              </div>
              <div class="col-md-6">
                <center>
                  <img src={gambarasik} alt="" height="auto" />
                </center>
              </div>
            </div>
          </div>
        </div>
      </Container>
    </>
  );
};
export default InputUnivBaru;
