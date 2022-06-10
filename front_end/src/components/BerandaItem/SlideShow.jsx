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
                    <div classNameName="">
                      <center>
                        <h1>
                          Temukan kampus
                          <br />
                          Terbaik Mu
                        </h1>
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
