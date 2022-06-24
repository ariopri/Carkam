import Systemnya from "./image/FAQ-Vector.png";
import { useState } from "react";
import ListKampus from "./ListKampus";

const Pencarian = () => {
  const [cari, setcari] = useState("");
  const kirim = cari;
  const token = localStorage.getItem("token");
  if (token === null) {
    console.log("token null");
  } else {
    console.log("token not null");
  }

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
                onChange={(e) => {
                  setcari(e.target.value);
                }}
              ></input>
              <button
                className="btn btn-success"
                type="submit"
                id="button-addon2"
              >
                Cari
              </button>
            </div>
            {token ? (
              <center>
                <a href="/Kampus/create" className="btn btn-warning">
                  <b>Tidak menemukan? Tambah disini</b>
                </a>
              </center>
            ) : (
              <center>
                <p>Kamu dapat cari melalui kata kunci yang kamu inginkan</p>
              </center>
            )}
          </div>
          <div class="col-md-6">
            <center>
              <img src={Systemnya} alt="" height="300px" />
            </center>
          </div>
        </div>
      </div>
      <ListKampus cari={kirim} />
    </>
  );
};

export default Pencarian;
