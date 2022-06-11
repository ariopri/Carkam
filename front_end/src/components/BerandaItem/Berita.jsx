import GBerita from "./image/berita.jpg";
const Berita = () => {
  return (
    <>
      <div className="container spaciberita spacekebawahbrita">
        <center>
          <h2>Berita terakhir</h2>

          <div className="row pb-5">
            <div className="col-md-4">
              <div className="card border-0">
                <div className="card-body">
                  <img src={GBerita} alt="" width="300px" />
                  <div className="spaciberita">
                    <p className="card-text">
                    Nullam id dolor id nibh ultricies vehicula ut id elit. Cras
                    justo odio, dapibus ac facilisis in, egestas eget quam.
                    Donec id elit non mi porta gravida at eget metus.
                  </p>
                  </div>
                </div>
              </div>
            </div>
            <div className="col-md-4">
              <div className="card border-0">
                <div className="card-body">
                <img src={GBerita} alt="" width="300px" />
                  <div className="spaciberita">
                    <p className="card-text">
                    Nullam id dolor id nibh ultricies vehicula ut id elit. Cras
                    justo odio, dapibus ac facilisis in, egestas eget quam.
                    Donec id elit non mi porta gravida at eget metus.
                  </p>
                  </div>
                </div>
              </div>
            </div>
            <div className="col-md-4">
              <div className="card border-0">
                <div className="card-body">
                <img src={GBerita} alt="" width="300px" />
                  <div className="spaciberita">
                    <p className="card-text">
                    Nullam id dolor id nibh ultricies vehicula ut id elit. Cras
                    justo odio, dapibus ac facilisis in, egestas eget quam.
                    Donec id elit non mi porta gravida at eget metus.
                  </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </center>
      </div>
    </>
  );
};

export default Berita;
