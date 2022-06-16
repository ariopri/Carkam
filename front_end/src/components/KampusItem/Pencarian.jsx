import Systemnya from "./image/FAQ-Vector.png";

const Pencarian = () => {
  return (
    <>
      <div class="container">
        <div class="row">
          <div class="col-md-6 pt-5">
            <center>
              <h3>Cari Kampus kamu disini</h3>
              <p>
                Cari kampus yang kamu inginkan lalu temukan sesuai rekomendasi
                mu
              </p>
            </center>
            <div className="input-group mb-3">
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
            </div>
            <center>
              <a href="/inpuunivbaru" className="btn btn-warning">
                <b>Tidak menemukan? Tambah disini</b>
              </a>
            </center>
          </div>
          <div class="col-md-6">
            <center>
              <img src={Systemnya} alt="" height="300px" />
            </center>
          </div>
        </div>
      </div>
    </>
  );
};

export default Pencarian;
