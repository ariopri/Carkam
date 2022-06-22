import { Container } from "react-bootstrap";

import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";

import gambarasik from "./MasukkanreviewItem/image/iniprofil.png";
const Masukkanreview = () => {
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
                  <Form.Control type="text" placeholder="Masukan nama" />
                </Form.Group>

                <Form.Group className="mb-3">
                  <Form.Label>Alamat Email</Form.Label>
                  <Form.Control type="email" placeholder="Masukkan Email" />
                </Form.Group>
                <Form.Group className="mb-3">
                  <Form.Label>Masukan Pendapat Kamu</Form.Label>
                  <textarea class="form-control"></textarea>
                </Form.Group>
                <Button variant="primary" type="submit">
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
