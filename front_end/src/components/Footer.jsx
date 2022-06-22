import Logo from "./Login-Page/image/logo-carkam.png";

const Footer = () => {
  return (
    <>
      <div class="mt-5 pt-5 pb-5 footer bg-dark text-light">
        <div class="container">
          <div class="row">
            <div class="col-lg-6 col-xs-12 ">
              <center>
                <img src={Logo} alt="" class="pb-4" />
                <p class="col-lg-11 text-white-50">
                  Jelajahi informasi kampus - kampus di Indonesia berdasarkan
                  lokasi, biaya kuliah, jurusan serta passing grade untuk
                  membantumu memilih kampus terbaik{" "}
                </p>
              </center>
            </div>
            <div class="col-lg-2 col-xs-12 ">
              <h4 class="mt-lg-0 mt-sm-3">Email</h4>
              <p class="pr-5 text-white-50">carkam.id@gmail.com </p>
            </div>
            <div class="col-lg-2 col-xs-12 ">
              <h4 class="mt-lg-0 mt-sm-4">Telephone</h4>
              <p class="pr-5 text-white-50">+62 21 4873 392 </p>
            </div>
            <div class="col-lg-2 col-xs-12 ">
              <h4 class="mt-lg-0 mt-sm-4">Location</h4>
              <p class="pr-5 text-white-50"> South Jakarta, Indonesia</p>
              <p class="pr-5 text-white-50">
                Grand Slipi Tower Lt. 42 Jl. S. Parman Kav 22-24
              </p>
              <p class="pr-5 text-white-50"> 8 A.M - 5 P.M</p>
            </div>
          </div>
          <div class="row mt-5">
            <div class="col copyright">
              <p class="">
                <small class="text-white-50">
                  <center>© 2022 CARKAMID. All Rights Reserved.</center>
                </small>
              </p>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Footer;
