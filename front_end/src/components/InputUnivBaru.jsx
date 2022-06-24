import { Container } from "react-bootstrap";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import gambarasik from "./InputUnivBaruItem/image/Beranda.png";
import { useState } from "react";

const InputUnivBaru = () => {
  const [inputList, setInputList] = useState([{ jurusan: "" }]);
  const handleInputChange = (e, index) => {
    const { name, value } = e.target;
    const list = [...inputList];
    list[index][name] = value;
    setInputList(list);
  };
  const handleRemoveClick = (index) => {
    const list = [...inputList];
    list.splice(index, 1);
    setInputList(list);
  };
  const handleAddClick = () => {
    setInputList([...inputList, { jurusan: "" }]);
  };
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

                  {inputList.map((x, i) => {
                    return (
                      <div className="box">
                        <Form.Group className="mb-3">
                          <Form.Label>Masukkan Jurusan </Form.Label>
                          <Form.Control
                            type="text"
                            name="jurusan"
                            value={x.jurusan}
                            onChange={(e) => handleInputChange(e, i)}
                          />
                        </Form.Group>

                        <div className="btn-box">
                          {inputList.length !== 1 && (
                            <Button
                              variant="danger"
                              onClick={() => handleRemoveClick(i)}
                            >
                              Hapus
                            </Button>
                            // <button
                            //   className="mr10"
                            //   onClick={() => handleRemoveClick(i)}
                            // >
                            //   Remove
                            // </button>
                          )}
                          <br /> <br />
                          {inputList.length - 1 === i && (
                            <Button variant="primary" onClick={handleAddClick}>
                              Tambah Jurusan
                            </Button>
                            // <button onClick={handleAddClick}>Add</button>
                          )}
                        </div>
                      </div>
                    );
                  })}
                  {/* <div style={{ marginTop: 20 }}>
                    {JSON.stringify(inputList)}
                  </div> */}

                  {/* <Form.Group className="mb-3">
                    <Form.Label>Masukkan Jurusan 1</Form.Label>
                    <Form.Control type="text" />
                  </Form.Group> */}

                  <br />
                  <Button variant="primary w-100" type="button">
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
