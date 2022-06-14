import { Container } from "react-bootstrap";
import Sigambar from "./image/iniprofil.png";

const SlideShow = () => {
  return (
    <>
      <div className="sibaground">
        <Container>
          <div className="spackosong2">
            <center>
              <table>
                <tr>
                  <td>
                    <h3>
                      <b>Our Profile</b>
                    </h3>
                    <div className="sitext">
                      <h1>
                        Kenali Minat dan
                        <br />
                        Temukan Kampus
                        <br />
                        Pilihanmu
                        <br />
                      </h1>
                      <div className="sisapcesubjudul">
                        <h5>The best companion for your brighter future</h5>
                      </div>
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
