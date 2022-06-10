import { Container } from "react-bootstrap";
import siganteng from "./image/Animation1.png";

const Testimoni = () => {
  return (
    <>
      <div className="bg-white spaciberita">
        <Container>
          <div className="backlagi rounded-3">
            <div className="kasihjarak ">
              <center>
                <h2>Testimonial Dari Pengguna CARKAM.ID</h2>
              </center>
              <div className="kasihjarak">
                <center>
                  <div className="container">
                    <div className="row">
                      <div className="col-md-4">
                        <div className="card ">
                          <div className="card-body">
                            <img src={siganteng} alt="" />
                            <h4 className="card-title p-3">Jhonatan</h4>
                            <h6 className="text-muted card-subtitle mb-2">
                              Universitas Indonesia
                            </h6>
                            <p className="card-text p-3">
                              Lorem ipsum dolor sit amet consectetur adipisicing
                              elit. Delectus tempore ullam necessitatibus non
                              dolore quasi molestiae quae illo nisi nobis, esse,
                              possimus qui unde autem ut corrupti, itaque
                              aspernatur minus!
                            </p>
                          </div>
                        </div>
                      </div>
                      <div className="col-md-4">
                        <div className="card">
                          <div className="card-body">
                            <img src={siganteng} alt="" />
                            <h4 className="card-title p-3">Putri Aini</h4>
                            <h6 className="text-muted card-subtitle mb-2">
                              Universitas Gadjah Mada
                            </h6>
                            <p className="card-text p-3">
                              Lorem, ipsum dolor sit amet consectetur
                              adipisicing elit. Recusandae aperiam magni nam
                              provident corporis est accusantium odio quasi fuga
                              quos porro, id animi dolores laborum ipsa
                              dignissimos modi, neque quibusdam?
                            </p>
                          </div>
                        </div>
                      </div>
                      <div className="col-md-4">
                        <div className="card">
                          <div className="card-body">
                            <img src={siganteng} alt="" />
                            <h4 className="card-title p-3">Salina Hehanusa</h4>
                            <h6 className="text-muted card-subtitle mb-2">
                              Institut Pertanian Bogor
                            </h6>
                            <p className="card-text p-3">
                              Lorem ipsum dolor sit amet consectetur adipisicing
                              elit. Odit ut optio autem a magni accusamus odio
                              suscipit illum in quidem incidunt perferendis
                              error saepe molestias, commodi voluptatum facilis?
                              Harum, sed!
                            </p>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </center>
              </div>
            </div>
          </div>
        </Container>
      </div>
    </>
  );
};
export default Testimoni;
