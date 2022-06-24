import { Container } from 'react-bootstrap';

const Isifaq = ({}) => {
  return (
    <>
      <div className="p-5">
        <Container>
          <Container>
            <h2 className="d-flex justify-content-center p-5 text-black">
              <strong>Frequently Asked Questions</strong>
            </h2>
            <div className="accordion" id="accordionExample">
              <div className="accordion-item">
                <h3 className="accordion-header" id="headingOne">
                  <button className="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne">
                    <strong>Apa itu Kampus?</strong>
                  </button>
                </h3>
                <div id="collapseOne" className="accordion-collapse collapse show" aria-labelledby="headingOne" data-bs-parent="#accordionExample">
                  <div className="accordion-body ">
                    <strong> CARKAM.ID </strong> merupakan sebuah smart education platform berbasis website untuk para siswa SMA/Calon Mahasiswa yang perlu mengenali minat, mencari kampus dan membandingkan kampus untuk mendapatkan kampus
                    pilihan.
                  </div>
                </div>
              </div>
              <div className="accordion-item">
                <h3 className="accordion-header" id="headingTwo">
                  <button className="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#collapseTwo" aria-expanded="true" aria-controls="collapseTwo">
                    <strong>Informasi apa yang bisa didapatkan?</strong>
                  </button>
                </h3>
                <div id="collapseTwo" className="accordion-collapse collapse show" aria-labelledby="headingTwo" data-bs-parent="#accordionExample">
                  <div className="accordion-body">
                    <strong>CARKAM.ID</strong> menyediakan data universitas negeri dan swasta yang terdiri dari data fakultas, program studi, alamat universitas, website universitas, passing grade, jumlah peminat, jumlah daya tampung dan
                    akreditasi.
                  </div>
                </div>
              </div>
              <div className="accordion-item">
                <h3 className="accordion-header" id="headingTwo">
                  <button className="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#collapseTwo" aria-expanded="true" aria-controls="collapseTwo">
                    <strong>Apa itu PTN dan PTS?</strong>
                  </button>
                </h3>
                <div id="collapseTwo" className="accordion-collapse collapse show" aria-labelledby="headingTwo" data-bs-parent="#accordionExample">
                  <div className="accordion-body">PTN adalah Perguruan Tinggi Negeri dan PTS adalah Perguruan Tinggi Swasta.</div>
                </div>
              </div>
              <div className="accordion-item">
                <h3 className="accordion-header" id="headingTwo">
                  <button className="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#collapseTwo" aria-expanded="true" aria-controls="collapseTwo">
                    <strong>Apa itu Akreditasi?</strong>
                  </button>
                </h3>
                <div id="collapseTwo" className="accordion-collapse collapse show" aria-labelledby="headingTwo" data-bs-parent="#accordionExample">
                  <div className="accordion-body d-flex justify-content-center">
                    Akreditasi merupakan proses evaluasi dan penilaian secara komprehensif atas komitmen perguruan tinggi terhadap mutu dan kapasitas penyelenggaraan Tridarma perguruan tinggi, untuk menentukan kelayakan program dan satuan
                    pendidikan. Akreditasi diselenggaakan oleh lembaga independen di luar dari universitas, untuk skala nasional biasanya dilakukan oleh Badan Akreditasi Nasional – Perguruan Tinggi (BAN-PT), yaitu sebuah lembaga akreditasi
                    indepeden yang menjadi rujukan nasional dan internasional.
                  </div>
                </div>
              </div>
              <div className="accordion-item">
                <h3 className="accordion-header" id="headingTwo">
                  <button className="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#collapseTwo" aria-expanded="true" aria-controls="collapseTwo">
                    <strong>Apa itu Passing Grade?</strong>
                  </button>
                </h3>
                <div id="collapseTwo" className="accordion-collapse collapse show" aria-labelledby="headingTwo" data-bs-parent="#accordionExample">
                  <div className="accordion-body">Passing Grade digunakan pada seleksi SBMPTN yang digunakan untuk mengetahui batas minimal yang harus dicapai dari hasil SBMPTN untuk memilih jurusan di suatu universitas.</div>
                </div>
              </div>
              <div className="accordion-item">
                <h3 className="accordion-header" id="headingThree">
                  <button className="accordion-button " type="button" data-bs-toggle="collapse" data-bs-target="#collapseThree" aria-expanded="true" aria-controls="collapseThree">
                    <strong>Mengapa harus membuat akun dan login di CARKAM.ID? </strong>
                  </button>
                </h3>
                <div id="collapseThree" className="accordion-collapse collapse show" aria-labelledby="headingThree" data-bs-parent="#accordionExample">
                  <div className="accordion-body">
                    Dengan membuat akun, kamu dapat menggunakan fitur kampus untuk membandingkan kampus yang diminati lalu temukan yang sesuai denganmu. Selain itu, kamu dapat menggunakan akunmu untuk memberi review dan menghapus data
                    universitas dari website kami sesuai dengan kebijakan kampus.
                  </div>
                </div>
              </div>
            </div>
          </Container>
        </Container>
      </div>
    </>
  );
};

export default Isifaq;
