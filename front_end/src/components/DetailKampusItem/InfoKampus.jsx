import Logo from "./image/KM.png";
import Card from "react-bootstrap/Card";
const InfoKampus = () => {
  return (
    <>
      <div className="pt-3">
        <div class="container">
          <div class="row">
            <div class="col-md-6">
              <img src={Logo} alt="" width="400px" />
            </div>
            <div class="col-md-6 pt-5">
              <Card>
                <Card.Header className="bg-success text-light">
                  <b>Kampus Merdeka</b>
                </Card.Header>
                <Card.Body>
                  <blockquote className="blockquote mb-0">
                    <p>Info singkat kampus</p>
                    <div class="container">
                      <div class="row pt-3">
                        <div class="col-md-12">Jurusan 1</div>
                        <footer className="blockquote-footer">
                          Pendapat dan review mahasiswa
                        </footer>
                        <button className="btn btn-success  ">
                          <a href="/review" className="btn text-light">
                            <b>R e v i e w</b>
                          </a>
                        </button>
                      </div>
                      <div class="row pt-3">
                        <div class="col-md-12">Jurusan 1</div>
                        <footer className="blockquote-footer">
                          Pendapat dan review mahasiswa
                        </footer>
                        <button className="btn btn-success  ">
                          <a href="/review" className="btn text-light">
                            <b>R e v i e w</b>
                          </a>
                        </button>
                      </div>
                    </div>
                  </blockquote>
                </Card.Body>
              </Card>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default InfoKampus;
