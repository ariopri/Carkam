import { Container } from 'react-bootstrap';
import Logo from './image/logo.png';
import Gambaran from './image/login.png';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
const Register = () => {
  return (
    <>
      <div className="bg-dark bg-gradient text-light mt-5">
        <div className="pt-3">
          <Container>
            <div class="row">
              <div class="col-md-6 p-5">
                <center>
                  <img src={Logo} alt="logo" className="img-fluid" />
                  <div>
                    <img src={Gambaran} alt="" />
                  </div>
                </center>
              </div>
              <div class="col-md-6">
                <div className="p-5">
                  <div className="bg-light p-4 text-dark">
                    <center>
                      <h3>
                        <b>Create Account</b>
                      </h3>
                    </center>
                    <Form className="pt-3">
                      <Form.Group className="mb-3">
                        <Form.Control type="text" placeholder="Your Name" />
                      </Form.Group>
                      <Form.Group className="mb-3">
                        <Form.Control type="email" placeholder="Your Email" />
                      </Form.Group>
                      <Form.Group className="mb-3">
                        <Form.Control type="password" placeholder="Password" />
                      </Form.Group>
                      <Form.Group className="mb-3">
                        <Form.Control type="password" placeholder="Repeat Your Password" />
                      </Form.Group>
                      <div class="d-md-flex justify-content-start align-items-center mb-4 py-2">
                        <h6 class="mb-0 me-4">Gender : </h6>
                        <div class="form-check form-check-inline mb-0 me-4">
                          <input class="form-check-input" type="radio" name="inlineRadioOptions" id="femaleGender" value="option1" />
                          <label class="form-check-label" for="femaleGender">
                            Female
                          </label>
                        </div>

                        <div class="form-check form-check-inline mb-0 me-4">
                          <input class="form-check-input" type="radio" name="inlineRadioOptions" id="maleGender" value="option2" />
                          <label class="form-check-label" for="maleGender">
                            Male
                          </label>
                        </div>
                      </div>
                      <input class="form-check-input me-3" type="checkbox" value="" id="form2Example3c" />
                      <label class="form-check-label pb-4" for="form2Example3">
                        I agree all statements in Terms of service
                      </label>
                      <Button variant="dark w-100" type="submit">
                        S U B M I T
                      </Button>
                    </Form>
                    <center>
                      <div className="p-3 text-muted ">
                        <a href="/Login" className="text-reset">
                          Sudah punya akun? Klik disini
                        </a>
                      </div>
                    </center>
                  </div>
                </div>
              </div>
            </div>
          </Container>
        </div>
      </div>
    </>
  );
};

export default Register;
