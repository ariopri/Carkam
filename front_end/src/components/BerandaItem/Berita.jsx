import GBerita from './image/berita-1.jpg';
import GBerita2 from './image/berita-2.jpg';
import GBerita3 from './image/berita-3.jpg';
const Berita = () => {
  return (
    <>
      <div className="container spaciberita spacekebawahbrita">
        <center>
          <h2>Berita terakhir</h2>
          <div className="row pt-5 pb-5">
            <div className="col-md-4">
              <div className="card border-0 shadow-lg">
                <div className="card-body">
                  <img src={GBerita} alt="" width="auto" height="200px" />
                  <div className="spaciberita ">
                    <p className="card-text m-4">
                      Bagi kalian yang belum berhasil lolos Seleksi Bersama Masuk Perguruan Tinggi (SBMPTN) 2022, tidak perlu khawatir. Pasalnya, masih ada sejumlah Perguruan Tinggi Negeri (PTN) yang masih membuka Seleksi Mandiri 2022.
                    </p>
                  </div>
                </div>
              </div>
            </div>
            <div className="col-md-4">
              <div className="card border-0 shadow-lg">
                <div className="card-body">
                  <img src={GBerita2} alt="" width="auto" height="200px" />
                  <div className="spaciberita">
                    <p className="card-text m-4">Baca artikel detikjabar, "Pengumuman SBMPTN 2022 Pukul 15.00 WIB, Jangan Lupa!" https://www.detik.com/jabar/berita/d-6142818/pengumuman-sbmptn-2022-pukul-1500-wib-jangan-lupa.</p>
                  </div>
                </div>
              </div>
            </div>
            <div className="col-md-4">
              <div className="card border-0 shadow-lg">
                <div className="card-body">
                  <img src={GBerita3} alt="" width="auto" height="200px" />
                  <div className="spaciberita">
                    <p className="card-text m-4">
                      Artikel ini telah tayang di Kompas.com dengan judul "Lolos SBMPTN 2022" Klik untuk baca: https://www.kompas.com/tren/read/2022/06/23/194807065/lolos-sbmptn-2022-begini-cara-unduh-sertifikat.
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
