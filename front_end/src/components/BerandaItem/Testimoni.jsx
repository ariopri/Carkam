import { Container } from 'react-bootstrap';
import Carousel from 'react-bootstrap/Carousel';

const Testimoni = () => {
  return (
    <>
      <div className="bg-light">
        <Container>
          <div class="row text-center ">
            <div class="col-md-12 ">
              <div id="carouselBasicExample" class="carousel slide carousel-dark mt-4 " data-mdb-ride="carousel">
                <h2 class="pt-2">Testimonial Dari Pengguna CARKAM.ID</h2>
                <Carousel>
                  <Carousel.Item interval={1000}>
                    <div class="mt-5 mb-5">
                      <p class="text-muted mb-2">
                        <strong class="text-success fs-4">Jessica</strong>
                      </p>
                      <p class="fs-5">Universitas Gadjah Mada</p>

                      <img src="https://mdbcdn.b-cdn.net/img/Photos/Avatars/img%20(2).webp" class="mt-4 rounded-circle img-fluid shadow-1-strong" alt="smaple image" width="100" height="100" />
                      <Carousel.Caption class="mt-4">
                        <p class="lead font-italic mx-4 mx-md-5 fst-italic">
                          " CARKAM.ID menjadi jawaban atas ratusan pertanyaan yang sering keluar dari seorang siswa SMA yang ingin melanjutkan ke perguruan tinggi. Sungguh sebuah inovasi yang sangat baik dalam memajukan pendidikan yang ada
                          di Indonesia. Disini juga ternyata ada fitur buat nyari tau info soal berbagai jurusan dari berbagai universitas, dan aku seneng banget. "
                        </p>
                      </Carousel.Caption>
                    </div>
                  </Carousel.Item>
                  <Carousel.Item interval={1000}>
                    <div class="mt-5 mb-5">
                      <p class="text-muted mb-2">
                        <strong class="text-success fs-4">Muhammad</strong>
                      </p>
                      <p class="fs-5">Universitas Indonesia</p>

                      <img src="https://mdbcdn.b-cdn.net/img/Photos/Avatars/img%20(3).webp" class="mt-4 rounded-circle img-fluid shadow-1-strong" alt="smaple image" width="100" />
                      <Carousel.Caption class="mt-4">
                        <p class="lead font-italic mx-4 mx-md-5 fst-italic">
                          " CARKAM.ID bikin aku gak ragu lagi pas mau milih kampus. Awalnya aku bingung, tapi setelah lihat website ini jadi bisa tau jurusan yang ada dan sesuai sama kita itu ke mana. Akhirnya aku cobain fitur cari kampus.
                          Dari sini aku yakin kalau aku lebih cocok kampus pilihanku tersebut dan tetap kekeh mau masuk situ. "
                        </p>
                      </Carousel.Caption>
                    </div>
                  </Carousel.Item>
                  <Carousel.Item interval={1000}>
                    <div class="mt-5 mb-5">
                      <p class="text-muted mb-2">
                        <strong class="text-success fs-4">Alma</strong>
                      </p>
                      <p class="fs-5">Institut Pertanian Bogor</p>

                      <img src="https://mdbcdn.b-cdn.net/img/Photos/Avatars/img%20(4).webp" class="mt-4 rounded-circle img-fluid shadow-1-strong" alt="smaple image" width="100" height="100" />
                      <Carousel.Caption class="mt-4">
                        <p class="lead font-italic mx-4 mx-md-5 fst-italic">
                          " Dahulu aku sebagai anak kelas 2 SMA yang tahun depan akan mengikuti seleksi masuk PTN atau PTS tentu saja aku dituntut untuk tahu apa-apa saja tentang perguruan. Pas banget nih! website ini mebyajikan informasi
                          yang dibutuhkan seputar perkuliahan. Kamu bisa mencari, menentukan atau membandingkan kampus terbaik yang kamu inginkan. "
                        </p>
                      </Carousel.Caption>
                    </div>
                  </Carousel.Item>
                </Carousel>
              </div>
            </div>
          </div>
        </Container>
      </div>
    </>
  );
};
export default Testimoni;
