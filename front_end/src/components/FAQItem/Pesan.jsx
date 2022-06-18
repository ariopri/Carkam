import { Container } from 'react-bootstrap';
import Back2 from './image/FAQ-Vector.png';
const Pesan = () => {
  return (
    <>
      <Container>
        <div class="container pt-5">
          <div class="row">
            <div class="col-md-6">
              <img src={Back2} alt="" />
            </div>
            <div class="col-md-6">
              <h3 class="pb-3">
                <b>Question</b>
              </h3>
              <div class="mb-3 pt-3">
                <input type="email" class="form-control" id="exampleFormControlInput1" placeholder="Email"></input>
              </div>
              <div class="mb-3">
                <input type="text" class="form-control" id="exampleFormControlInput1" placeholder="Nama Lengkap"></input>
              </div>
              <div class="mb-3">
                <textarea class="form-control" id="exampleFormControlTextarea1" rows="3" placeholder="Pertanyaan Kamu"></textarea>
              </div>
              <div class="d-flex flex-row-reverse">
                <button class="btn btn-dark btn-lg ">Kirim</button>
              </div>
            </div>
          </div>
        </div>
      </Container>
    </>
  );
};
export default Pesan;
