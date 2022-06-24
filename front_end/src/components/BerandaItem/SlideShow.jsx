import { Container } from "react-bootstrap";
import Sigambar from "./image/Beranda.png";

const SlideShow = () => {
  return (
    <>
      <div className="sibaground">
        <Container>
          <div classNameName="jumbotron spackosong">
            <center>
              <table>
                <tr>
                  <td>
                    <div className="sitext">
                      <center>
                        <h1 className="fontJudul">
                          <strong>
                            TEMUKAN KAMPUS
                            <br />
                            TERBAIKMU
                            <br />
                            <br />
                          </strong>
                        </h1>
                        <div>
                          <a href="/Kampus" className="btn btn-success">
                            Cari Kampus Disini
                          </a>
                        </div>
                        {/* <div className="input-group mb-3">
                          <input
                            type="text"
                            className="form-control siplaceholder"
                            placeholder="C A R I K A M P U S"
                            aria-describedby="button-addon2"
                          ></input>

                          <button
                            className="btn btn-success"
                            type="button"
                            id="button-addon2"
                          >
                            Cari
                          </button>
                        </div> */}
                        <div className="sisapcesubjudul">
                          <h5>
                            "We will never know the real answer, before you try"
                          </h5>
                        </div>
                      </center>
                    </div>
                  </td>
                  <img src={Sigambar} alt="" />
                  <td></td>
                </tr>
              </table>
            </center>
          </div>
        </Container>
      </div>
    </>
  );
};

export default SlideShow;
