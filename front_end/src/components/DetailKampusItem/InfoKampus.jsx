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
                  <div class="row">
                    <div class="col-md-6">
                      <b>Kampus Merdeka &nbsp;</b>
                    </div>
                    <div class="col-md-6">
                      <div className="d-flex flex-row-reverse mt-">
                        <a
                          href="/Kampus/detail/update"
                          className="btn btn-danger text-light "
                        >
                          <b>Ubah / Hapus</b>
                        </a>
                      </div>
                    </div>
                  </div>
                </Card.Header>
                <Card.Body>
                  <blockquote className="blockquote mb-0">
                    <p>
                      <b>Info singkat kampus</b>
                    </p>
                    <p>
                      Lorem ipsum dolor sit amet consectetur adipisicing elit.
                      Ullam, praesentium nam dignissimos repellendus voluptates
                      tempore.
                    </p>
                    <br />
                    <p>
                      <b>Jurusan : </b> Teknik Informatika, Teknik Sipil, Teknik
                      Mesin.
                    </p>
                    <button className="btn btn-success  w-100  ">
                      <a
                        href="/Kampus/detail/review"
                        className="btn text-light"
                      >
                        <b>R e v i e w</b>
                      </a>
                    </button>
                    <div class="container">
                      <div class="row pt-3">
                        <div class="col-md-12">Dodi</div>
                        <footer className="blockquote-footer">
                          Kampusnya bagus gaes
                        </footer>
                      </div>
                      <div class="row pt-3">
                        <div class="col-md-12">Dodi</div>
                        <footer className="blockquote-footer">
                          Kampusnya bagus gaes
                        </footer>
                      </div>
                      <div class="row pt-3">
                        <div class="col-md-12">Dodi</div>
                        <footer className="blockquote-footer">
                          Kampusnya bagus gaes
                        </footer>
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
